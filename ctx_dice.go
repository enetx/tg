package tg

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

func (d *Dice) After(duration time.Duration) *Dice {
	d.after = Some(duration)
	return d
}

func (d *Dice) DeleteAfter(duration time.Duration) *Dice {
	d.deleteAfter = Some(duration)
	return d
}

func (d *Dice) Emoji(e String) *Dice {
	d.opts.Emoji = e.Std()
	return d
}

func (d *Dice) Dart() *Dice {
	d.opts.Emoji = "🎯"
	return d
}

func (d *Dice) Slot() *Dice {
	d.opts.Emoji = "🎰"
	return d
}

func (d *Dice) Ball() *Dice {
	d.opts.Emoji = "🏀"
	return d
}

func (d *Dice) Soccer() *Dice {
	d.opts.Emoji = "⚽"
	return d
}

func (d *Dice) Bowling() *Dice {
	d.opts.Emoji = "🎳"
	return d
}

func (d *Dice) Silent() *Dice {
	d.opts.DisableNotification = true
	return d
}

func (d *Dice) Thread(id int64) *Dice {
	d.opts.MessageThreadId = id
	return d
}

func (d *Dice) AllowPaidBroadcast() *Dice {
	d.opts.AllowPaidBroadcast = true
	return d
}

func (d *Dice) Effect(effect effects.EffectType) *Dice {
	d.opts.MessageEffectId = effect.String()
	return d
}

func (d *Dice) ReplyTo(id int64) *Dice {
	d.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: id}
	return d
}

func (d *Dice) Markup(kb keyboard.KeyboardBuilder) *Dice {
	d.opts.ReplyMarkup = kb.Markup()
	return d
}

func (d *Dice) To(chatID int64) *Dice {
	d.chatID = Some(chatID)
	return d
}

func (d *Dice) Send() Result[*gotgbot.Message] {
	return d.ctx.timers(d.after, d.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := d.chatID.UnwrapOr(d.ctx.EffectiveChat.Id)
		return ResultOf(d.ctx.Bot.Raw.SendDice(chatID, d.opts))
	})
}
