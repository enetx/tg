package tg

import . "github.com/enetx/g"

type Flow struct {
	bot      *Bot
	steps    Map[String, func(*Context) error]
	fallback Handler
}

func NewFlow(bot *Bot) *Flow {
	return &Flow{
		bot:   bot,
		steps: NewMap[String, func(*Context) error](),
	}
}

func (f *Flow) Entry(cmd, state String) *Flow {
	f.bot.Command(cmd, func(ctx *Context) error {
		ctx.State.flow = f
		return ctx.State.Jump(state)
	})

	return f
}

func (f *Flow) Step(state String, handler Handler) *Flow {
	f.steps.Set(state, handler)
	return f
}

func (f *Flow) Fallback(fn Handler) *Flow {
	f.fallback = fn
	return f
}

func (f *Flow) Register(message ...String) {
	handler := func(ctx *Context) error {
		ctx.State.flow = f
		state := ctx.State.Get().UnwrapOrDefault()
		return f.call(state, ctx, Slice[String](message).Get(0).UnwrapOr("Please type /start to begin."))
	}

	f.bot.On.Any(handler)
}

func (f *Flow) call(state String, ctx *Context, message String) error {
	if h := f.steps.Get(state); h.IsSome() {
		return h.Some()(ctx)
	}

	if f.fallback != nil {
		return f.fallback(ctx)
	}

	ctx.State.Clear()

	return ctx.Message(message).Send().Err()
}
