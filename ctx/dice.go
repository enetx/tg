package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/types/effects"
)

type Dice struct {
	ctx         *Context
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	opts        *gotgbot.SendDiceOpts
}

// After schedules the dice to be sent after the specified duration.
func (d *Dice) After(duration time.Duration) *Dice {
	d.after = Some(duration)
	return d
}

// DeleteAfter schedules the dice message to be deleted after the specified duration.
func (d *Dice) DeleteAfter(duration time.Duration) *Dice {
	d.deleteAfter = Some(duration)
	return d
}

// Emoji sets a custom emoji for the dice.
func (d *Dice) Emoji(e String) *Dice {
	d.opts.Emoji = e.Std()
	return d
}

// Dart sets the dice emoji to dart.
func (d *Dice) Dart() *Dice {
	d.opts.Emoji = "🎯"
	return d
}

// Slot sets the dice emoji to slot machine.
func (d *Dice) Slot() *Dice {
	d.opts.Emoji = "🎰"
	return d
}

// Ball sets the dice emoji to basketball.
func (d *Dice) Ball() *Dice {
	d.opts.Emoji = "🏀"
	return d
}

// Soccer sets the dice emoji to soccer ball.
func (d *Dice) Soccer() *Dice {
	d.opts.Emoji = "⚽"
	return d
}

// Bowling sets the dice emoji to bowling.
func (d *Dice) Bowling() *Dice {
	d.opts.Emoji = "🎳"
	return d
}

// Silent disables notification for the dice message.
func (d *Dice) Silent() *Dice {
	d.opts.DisableNotification = true
	return d
}

// Thread sets the message thread ID for the dice message.
func (d *Dice) Thread(id int64) *Dice {
	d.opts.MessageThreadId = id
	return d
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (d *Dice) AllowPaidBroadcast() *Dice {
	d.opts.AllowPaidBroadcast = true
	return d
}

// Effect sets a message effect for the dice message.
func (d *Dice) Effect(effect effects.EffectType) *Dice {
	d.opts.MessageEffectId = effect.String()
	return d
}

// ReplyTo sets the message ID to reply to.
func (d *Dice) ReplyTo(id int64) *Dice {
	d.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: id}
	return d
}

// Markup sets the reply markup keyboard for the dice message.
func (d *Dice) Markup(kb keyboard.KeyboardBuilder) *Dice {
	d.opts.ReplyMarkup = kb.Markup()
	return d
}

// Business sets the business connection ID for the dice message.
func (d *Dice) Business(id String) *Dice {
	d.opts.BusinessConnectionId = id.Std()
	return d
}

// Protect enables content protection for the dice message.
func (d *Dice) Protect() *Dice {
	d.opts.ProtectContent = true
	return d
}

// To sets the target chat ID for the dice message.
func (d *Dice) To(chatID int64) *Dice {
	d.chatID = Some(chatID)
	return d
}

// Send sends the dice message to Telegram and returns the result.
func (d *Dice) Send() Result[*gotgbot.Message] {
	return d.ctx.timers(d.after, d.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := d.chatID.UnwrapOr(d.ctx.EffectiveChat.Id)
		return ResultOf(d.ctx.Bot.Raw().SendDice(chatID, d.opts))
	})
}
