package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/types/effects"
)

type SendGame struct {
	ctx           *Context
	gameShortName g.String
	opts          *gotgbot.SendGameOpts
	chatID        g.Option[int64]
	after         g.Option[time.Duration]
	deleteAfter   g.Option[time.Duration]
}

// After schedules the game to be sent after the specified duration.
func (sg *SendGame) After(duration time.Duration) *SendGame {
	sg.after = g.Some(duration)
	return sg
}

// DeleteAfter schedules the game message to be deleted after the specified duration.
func (sg *SendGame) DeleteAfter(duration time.Duration) *SendGame {
	sg.deleteAfter = g.Some(duration)
	return sg
}

// Silent disables notification for the game message.
func (sg *SendGame) Silent() *SendGame {
	sg.opts.DisableNotification = true
	return sg
}

// Protect enables content protection for the game message.
func (sg *SendGame) Protect() *SendGame {
	sg.opts.ProtectContent = true
	return sg
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (sg *SendGame) AllowPaidBroadcast() *SendGame {
	sg.opts.AllowPaidBroadcast = true
	return sg
}

// Thread sets the message thread ID for the game message.
func (sg *SendGame) Thread(id int64) *SendGame {
	sg.opts.MessageThreadId = id
	return sg
}

// Effect sets a message effect for the game message.
func (sg *SendGame) Effect(effect effects.EffectType) *SendGame {
	sg.opts.MessageEffectId = effect.String()
	return sg
}

// Reply sets reply parameters using the reply builder.
func (sg *SendGame) Reply(params *reply.Parameters) *SendGame {
	sg.opts.ReplyParameters = params.Std()
	return sg
}

// Markup sets the reply markup keyboard for the game message.
func (sg *SendGame) Markup(kb keyboard.Keyboard) *SendGame {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		sg.opts.ReplyMarkup = markup
	}

	return sg
}

// Business sets the business connection ID for the game message.
func (sg *SendGame) Business(id g.String) *SendGame {
	sg.opts.BusinessConnectionId = id.Std()
	return sg
}

// To sets the target chat ID for the game message.
func (sg *SendGame) To(chatID int64) *SendGame {
	sg.chatID = g.Some(chatID)
	return sg
}

// Timeout sets a custom timeout for this request.
func (sg *SendGame) Timeout(duration time.Duration) *SendGame {
	if sg.opts.RequestOpts == nil {
		sg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sg.opts.RequestOpts.Timeout = duration

	return sg
}

// APIURL sets a custom API URL for this request.
func (sg *SendGame) APIURL(url g.String) *SendGame {
	if sg.opts.RequestOpts == nil {
		sg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sg.opts.RequestOpts.APIURL = url.Std()

	return sg
}

// Send sends the game message to Telegram and returns the result.
func (sg *SendGame) Send() g.Result[*gotgbot.Message] {
	return sg.ctx.timers(sg.after, sg.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := sg.chatID.UnwrapOr(sg.ctx.EffectiveChat.Id)
		return g.ResultOf(sg.ctx.Bot.Raw().SendGame(chatID, sg.gameShortName.Std(), sg.opts))
	})
}
