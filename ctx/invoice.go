package ctx

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type Invoice struct {
	ctx      *Context
	title    String
	desc     String
	payload  String
	currency String
	prices   Slice[gotgbot.LabeledPrice]
	chatID   Option[int64]
	opts     *gotgbot.SendInvoiceOpts
}

func (i *Invoice) To(chatID int64) *Invoice {
	i.chatID = Some(chatID)
	return i
}

func (i *Invoice) Price(label String, amount int) *Invoice {
	i.prices.Push(gotgbot.LabeledPrice{Label: label.Std(), Amount: int64(amount)})
	return i
}

func (i *Invoice) Thread(id int64) *Invoice {
	i.opts.MessageThreadId = id
	return i
}

func (i *Invoice) ProviderToken(token String) *Invoice {
	i.opts.ProviderToken = token.Std()
	return i
}

func (i *Invoice) MaxTip(amount int64) *Invoice {
	i.opts.MaxTipAmount = amount
	return i
}

func (i *Invoice) SuggestedTips(tips ...int64) *Invoice {
	i.opts.SuggestedTipAmounts = tips
	return i
}

func (i *Invoice) StartParameter(param String) *Invoice {
	i.opts.StartParameter = param.Std()
	return i
}

func (i *Invoice) ProviderData(data String) *Invoice {
	i.opts.ProviderData = data.Std()
	return i
}

func (i *Invoice) Photo(url String, size, width, height int64) *Invoice {
	i.opts.PhotoUrl = url.Std()
	i.opts.PhotoSize = size
	i.opts.PhotoWidth = width
	i.opts.PhotoHeight = height

	return i
}

func (i *Invoice) NeedName() *Invoice {
	i.opts.NeedName = true
	return i
}

func (i *Invoice) NeedPhone() *Invoice {
	i.opts.NeedPhoneNumber = true
	return i
}

func (i *Invoice) NeedEmail() *Invoice {
	i.opts.NeedEmail = true
	return i
}

func (i *Invoice) NeedShipping() *Invoice {
	i.opts.NeedShippingAddress = true
	return i
}

func (i *Invoice) SendPhone() *Invoice {
	i.opts.SendPhoneNumberToProvider = true
	return i
}

func (i *Invoice) SendEmail() *Invoice {
	i.opts.SendEmailToProvider = true
	return i
}

func (i *Invoice) Flexible() *Invoice {
	i.opts.IsFlexible = true
	return i
}

func (i *Invoice) Silent() *Invoice {
	i.opts.DisableNotification = true
	return i
}

func (i *Invoice) Protect() *Invoice {
	i.opts.ProtectContent = true
	return i
}

func (i *Invoice) AllowPaidBroadcast() *Invoice {
	i.opts.AllowPaidBroadcast = true
	return i
}

func (i *Invoice) Effect(effect string) *Invoice {
	i.opts.MessageEffectId = effect
	return i
}

func (i *Invoice) ReplyTo(messageID int64) *Invoice {
	i.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return i
}

func (i *Invoice) Markup(kb keyboard.KeyboardBuilder) *Invoice {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		i.opts.ReplyMarkup = markup
	}

	return i
}

func (i *Invoice) Send() Result[*gotgbot.Message] {
	return ResultOf(i.ctx.Bot.Raw().SendInvoice(
		i.chatID.UnwrapOr(i.ctx.EffectiveChat.Id),
		i.title.Std(),
		i.desc.Std(),
		i.payload.Std(),
		i.currency.Std(),
		i.prices,
		i.opts,
	))
}
