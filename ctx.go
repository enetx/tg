package tg

import (
	"errors"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type Context struct {
	Bot              *Bot
	Callback         *gotgbot.CallbackQuery
	EffectiveChat    *gotgbot.Chat
	EffectiveMessage *gotgbot.Message
	EffectiveSender  *gotgbot.Sender
	EffectiveUser    *gotgbot.User
	State            *State
	Update           *gotgbot.Update
	Raw              *ext.Context
}

func newCtx(bot *Bot, ctx *ext.Context) *Context {
	context := &Context{
		Bot:              bot,
		EffectiveChat:    ctx.EffectiveChat,
		EffectiveMessage: ctx.EffectiveMessage,
		EffectiveUser:    ctx.EffectiveUser,
		EffectiveSender:  ctx.EffectiveSender,
		Update:           ctx.Update,
		Callback:         ctx.Update.CallbackQuery,
		Raw:              ctx,
	}

	context.State = &State{
		ctx:     context,
		current: bot.states,
		data:    bot.stateData,
	}

	return context
}

func (ctx *Context) Poll(question String) *Poll {
	return &Poll{
		ctx:      ctx,
		question: question,
		opts:     new(gotgbot.SendPollOpts),
	}
}

func (ctx *Context) Reply(text String) *Reply {
	return &Reply{
		ctx:  ctx,
		text: text,
		opts: new(gotgbot.SendMessageOpts),
	}
}

func (ctx *Context) Message(text String) *Message {
	return &Message{
		ctx:  ctx,
		text: text,
		opts: new(gotgbot.SendMessageOpts),
	}
}

func (ctx *Context) Document(filename String) *Document {
	d := &Document{
		ctx:  ctx,
		opts: new(gotgbot.SendDocumentOpts),
	}

	if filename.Empty() {
		d.err = errors.New("filename is empty")
		return d
	}

	switch {
	case filename.StartsWithAny("http://", "https://"):
		d.doc = gotgbot.InputFileByURL(filename.Std())
	case filename.StartsWith(FileIDPrefix):
		d.doc = gotgbot.InputFileByID(filename.StripPrefix(FileIDPrefix).Std())
	default:
		file := NewFile(filename).Open()
		if file.IsErr() {
			d.err = file.Err()
			return d
		}

		d.file = file.Ok()
		d.doc = gotgbot.InputFileByReader(file.Ok().Name().Std(), file.Ok().Std())
	}

	return d
}

func (ctx *Context) Audio(filename String) *Audio {
	a := &Audio{
		ctx:  ctx,
		opts: new(gotgbot.SendAudioOpts),
	}

	if filename.Empty() {
		a.err = errors.New("filename is empty")
		return a
	}

	switch {
	case filename.StartsWithAny("http://", "https://"):
		a.doc = gotgbot.InputFileByURL(filename.Std())
	case filename.StartsWith(FileIDPrefix):
		a.doc = gotgbot.InputFileByID(filename.StripPrefix(FileIDPrefix).Std())
	default:
		file := NewFile(filename).Open()
		if file.IsErr() {
			a.err = file.Err()
			return a
		}

		a.file = file.Ok()
		a.doc = gotgbot.InputFileByReader(file.Ok().Name().Std(), file.Ok().Std())
	}

	return a
}

func (ctx *Context) Photo(filename String) *Photo {
	p := &Photo{
		ctx:  ctx,
		opts: new(gotgbot.SendPhotoOpts),
	}

	if filename.Empty() {
		p.err = errors.New("filename is empty")
		return p
	}

	switch {
	case filename.StartsWithAny("http://", "https://"):
		p.doc = gotgbot.InputFileByURL(filename.Std())
	case filename.StartsWith(FileIDPrefix):
		p.doc = gotgbot.InputFileByID(filename.StripPrefix(FileIDPrefix).Std())
	default:
		file := NewFile(filename).Open()
		if file.IsErr() {
			p.err = file.Err()
			return p
		}

		p.file = file.Ok()
		p.doc = gotgbot.InputFileByReader(file.Ok().Name().Std(), file.Ok().Std())
	}

	return p
}

func (ctx *Context) Video(filename String) *Video {
	v := &Video{
		ctx:  ctx,
		opts: new(gotgbot.SendVideoOpts),
	}

	if filename.Empty() {
		v.err = errors.New("filename is empty")
		return v
	}

	switch {
	case filename.StartsWithAny("http://", "https://"):
		v.doc = gotgbot.InputFileByURL(filename.Std())
	case filename.StartsWith(FileIDPrefix):
		v.doc = gotgbot.InputFileByID(filename.StripPrefix(FileIDPrefix).Std())
	default:
		file := NewFile(filename).Open()
		if file.IsErr() {
			v.err = file.Err()
			return v
		}

		v.file = file.Ok()
		v.doc = gotgbot.InputFileByReader(file.Ok().Name().Std(), file.Ok().Std())
	}

	return v
}

func (ctx *Context) EditMarkup(kb keyboard.KeyboardBuilder) *EditMarkup {
	return &EditMarkup{
		ctx:  ctx,
		kb:   kb,
		opts: new(gotgbot.EditMessageReplyMarkupOpts),
	}
}

func (ctx *Context) EditText(text String) *EditText {
	return &EditText{
		ctx:  ctx,
		text: text,
		opts: new(gotgbot.EditMessageTextOpts),
	}
}

func (ctx *Context) Answer(text String) *Answer {
	return &Answer{
		ctx:  ctx,
		text: text,
		opts: new(gotgbot.AnswerCallbackQueryOpts),
	}
}

func (ctx *Context) Dice() *Dice {
	return &Dice{
		ctx:  ctx,
		opts: new(gotgbot.SendDiceOpts),
	}
}

func (ctx *Context) Invoice(title, desc, payload, currency String) *Invoice {
	return &Invoice{
		ctx:      ctx,
		title:    title,
		desc:     desc,
		payload:  payload,
		currency: currency,
		prices:   NewSlice[gotgbot.LabeledPrice](),
		opts:     new(gotgbot.SendInvoiceOpts),
	}
}

func (ctx *Context) PreCheckout() *PreCheckout {
	return &PreCheckout{
		ctx:  ctx,
		opts: new(gotgbot.AnswerPreCheckoutQueryOpts),
	}
}

func (ctx *Context) RefundStarPayment(chargeID String) *RefundStarPayment {
	return &RefundStarPayment{
		ctx:      ctx,
		chargeID: chargeID,
		opts:     new(gotgbot.RefundStarPaymentOpts),
	}
}

func (ctx *Context) IsAdmin() Result[bool] {
	member, err := ctx.Bot.Raw.GetChatMember(ctx.EffectiveChat.Id, ctx.EffectiveUser.Id, nil)
	if err != nil {
		return Err[bool](nil)
	}

	return Ok(member.GetStatus() == "administrator" || member.GetStatus() == "creator")
}

func (ctx *Context) Args() Slice[String] {
	return String(ctx.EffectiveMessage.Text).Fields().Skip(1).Collect()
}

func (ctx *Context) Delete() *Delete {
	return &Delete{
		ctx:  ctx,
		opts: new(gotgbot.DeleteMessageOpts),
	}
}

func (ctx *Context) timers(
	after Option[time.Duration],
	deleteAfter Option[time.Duration],
	send func() Result[*gotgbot.Message],
) Result[*gotgbot.Message] {
	if after.IsSome() {
		delay := after.Some()

		go func() {
			<-time.After(delay)
			msg := send()
			if msg.IsOk() && deleteAfter.IsSome() {
				ctx.Delete().MessageID(msg.Ok().MessageId).After(deleteAfter.Some()).Send()
			}
		}()

		return Ok[*gotgbot.Message](nil)
	}

	msg := send()

	if msg.IsOk() && deleteAfter.IsSome() {
		ctx.Delete().MessageID(msg.Ok().MessageId).After(deleteAfter.Some()).Send()
	}

	return msg
}
