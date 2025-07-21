package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/types/effects"
)

type Game struct {
	ctx           *Context
	gameShortName String
	opts          *gotgbot.SendGameOpts
	chatID        Option[int64]
	after         Option[time.Duration]
	deleteAfter   Option[time.Duration]
}

// After schedules the game to be sent after the specified duration.
func (g *Game) After(duration time.Duration) *Game {
	g.after = Some(duration)
	return g
}

// DeleteAfter schedules the game message to be deleted after the specified duration.
func (g *Game) DeleteAfter(duration time.Duration) *Game {
	g.deleteAfter = Some(duration)
	return g
}

// Silent disables notification for the game message.
func (g *Game) Silent() *Game {
	g.opts.DisableNotification = true
	return g
}

// Protect enables content protection for the game message.
func (g *Game) Protect() *Game {
	g.opts.ProtectContent = true
	return g
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (g *Game) AllowPaidBroadcast() *Game {
	g.opts.AllowPaidBroadcast = true
	return g
}

// Thread sets the message thread ID for the game message.
func (g *Game) Thread(id int64) *Game {
	g.opts.MessageThreadId = id
	return g
}

// Effect sets a message effect for the game message.
func (g *Game) Effect(effect effects.EffectType) *Game {
	g.opts.MessageEffectId = effect.String()
	return g
}

// ReplyTo sets the message ID to reply to.
func (g *Game) ReplyTo(messageID int64) *Game {
	g.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return g
}

// Markup sets the reply markup keyboard for the game message.
func (g *Game) Markup(kb keyboard.KeyboardBuilder) *Game {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		g.opts.ReplyMarkup = markup
	}

	return g
}

// Business sets the business connection ID for the game message.
func (g *Game) Business(id String) *Game {
	g.opts.BusinessConnectionId = id.Std()
	return g
}

// To sets the target chat ID for the game message.
func (g *Game) To(chatID int64) *Game {
	g.chatID = Some(chatID)
	return g
}

// Send sends the game message to Telegram and returns the result.
func (g *Game) Send() Result[*gotgbot.Message] {
	return g.ctx.timers(g.after, g.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := g.chatID.UnwrapOr(g.ctx.EffectiveChat.Id)
		return ResultOf(g.ctx.Bot.Raw().SendGame(chatID, g.gameShortName.Std(), g.opts))
	})
}
