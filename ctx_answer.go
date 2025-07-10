package tg

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type Answer struct {
	ctx  *Context
	text String
	opts *gotgbot.AnswerCallbackQueryOpts
}

func (a *Answer) URL(url String) *Answer {
	a.opts.Url = url.Std()
	return a
}

func (a *Answer) Alert() *Answer {
	a.opts.ShowAlert = true
	return a
}

func (a *Answer) CacheTime(seconds int64) *Answer {
	a.opts.CacheTime = seconds
	return a
}

func (a *Answer) Timeout(duration time.Duration) *Answer {
	a.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return a
}

func (a *Answer) Send() Result[bool] {
	a.opts.Text = a.text.Std()
	return ResultOf(a.ctx.Update.CallbackQuery.Answer(a.ctx.Bot.Raw, a.opts))
}
