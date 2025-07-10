package tg

import . "github.com/enetx/g"

type State struct {
	ctx     *Context
	current *MapSafe[int64, String]
	data    *MapSafe[int64, *MapSafe[String, any]]
	flow    *Flow
}

func (s *State) Get() Option[String] {
	return s.current.Get(s.key())
}

func (s *State) Set(state String) {
	s.current.Set(s.key(), state)
}

func (s *State) Clear() {
	s.current.Delete(s.key())
	s.data.Delete(s.key())
}

func (s *State) Data() *MapSafe[String, any] {
	entry := s.data.Entry(s.key())
	if data := entry.OrSet(NewMapSafe[String, any]()); data.IsSome() {
		return data.Some()
	}

	return entry.Get().Some()
}

func (s *State) key() int64 {
	if s.ctx.EffectiveUser != nil && s.ctx.EffectiveUser.Id != 0 {
		return s.ctx.EffectiveUser.Id
	}

	if s.ctx.EffectiveChat != nil {
		return s.ctx.EffectiveChat.Id
	}

	return 0
}

func (s *State) Jump(state String) error {
	if s.flow == nil {
		return Errorf("cannot Jump: no flow registered for key {}", s.key())
	}

	s.Set(state)

	if h := s.flow.steps.Get(state); h.IsSome() {
		return h.Some()(s.ctx)
	}

	s.Clear()
	return Errorf("flow: unknown state '{}'", state)
}
