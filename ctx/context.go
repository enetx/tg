package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
	"github.com/enetx/tg/internal/tgfile"
	"github.com/enetx/tg/keyboard"
)

// Context represents a Telegram Bot API context with convenient methods for sending messages and media.
// It provides a high-level interface over the raw gotgbot API with fluent method chaining.
//
// The Context contains all necessary information about the current update including:
//   - Bot instance for making API calls
//   - Effective chat, message, user, and sender information
//   - Raw callback query data if applicable
//   - Original update and gotgbot context for advanced usage
type Context struct {
	Bot              core.BotAPI            // The bot instance for making API calls
	Callback         *gotgbot.CallbackQuery // Callback query data if the update is a callback query
	EffectiveChat    *gotgbot.Chat          // The chat where the update occurred
	EffectiveMessage *gotgbot.Message       // The message associated with the update
	EffectiveSender  *gotgbot.Sender        // The sender of the update
	EffectiveUser    *gotgbot.User          // The user who sent the update
	Update           *gotgbot.Update        // The original update object
	Raw              *ext.Context           // The raw gotgbot context for advanced usage
}

// New creates a new Context instance from a bot and raw gotgbot context.
// This is typically called automatically by the framework when handling updates.
func New(bot core.BotAPI, raw *ext.Context) *Context {
	return &Context{
		Bot:              bot,
		Callback:         raw.CallbackQuery,
		EffectiveChat:    raw.EffectiveChat,
		EffectiveMessage: raw.EffectiveMessage,
		EffectiveSender:  raw.EffectiveSender,
		EffectiveUser:    raw.EffectiveUser,
		Update:           raw.Update,
		Raw:              raw,
	}
}

// Ban creates a new Ban request to ban a user from the chat.
func (ctx *Context) Ban(userID int64) *Ban {
	return &Ban{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.BanChatMemberOpts),
	}
}

// Unban creates a new Unban request to unban a user from the chat.
func (ctx *Context) Unban(userID int64) *Unban {
	return &Unban{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.UnbanChatMemberOpts),
	}
}

// Restrict creates a new Restrict request to restrict a user's permissions in the chat.
func (ctx *Context) Restrict(userID int64) *Restrict {
	return &Restrict{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.RestrictChatMemberOpts),
	}
}

// Promote creates a new Promote request to promote a user to administrator in the chat.
func (ctx *Context) Promote(userID int64) *Promote {
	return &Promote{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.PromoteChatMemberOpts),
	}
}

// SendPoll creates a new Poll request with the specified question.
func (ctx *Context) SendPoll(question String) *SendPoll {
	return &SendPoll{
		ctx:      ctx,
		question: question,
		opts:     new(gotgbot.SendPollOpts),
	}
}

// Reply creates a new Reply request that replies to the current message.
func (ctx *Context) Reply(text String) *Reply {
	return &Reply{
		ctx:  ctx,
		text: text,
		opts: new(gotgbot.SendMessageOpts),
	}
}

// SendMessage creates a new SendMessage request to send a text message.
func (ctx *Context) SendMessage(text String) *SendMessage {
	return &SendMessage{
		ctx:  ctx,
		text: text,
		opts: new(gotgbot.SendMessageOpts),
	}
}

// SendDocument creates a new SendDocument request to send a document file.
func (ctx *Context) SendDocument(filename String) *SendDocument {
	d := &SendDocument{
		ctx:  ctx,
		opts: new(gotgbot.SendDocumentOpts),
	}

	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		d.err = result.Err()
		return d
	}

	d.doc = result.Ok().Doc
	d.file = result.Ok().File

	return d
}

// SendAudio creates a new SendAudio request to send an audio file.
func (ctx *Context) SendAudio(filename String) *SendAudio {
	a := &SendAudio{
		ctx:  ctx,
		opts: new(gotgbot.SendAudioOpts),
	}

	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		a.err = result.Err()
		return a
	}

	a.doc = result.Ok().Doc
	a.file = result.Ok().File

	return a
}

// SendPhoto creates a new SendPhoto request to send a photo.
func (ctx *Context) SendPhoto(filename String) *SendPhoto {
	p := &SendPhoto{
		ctx:  ctx,
		opts: new(gotgbot.SendPhotoOpts),
	}

	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		p.err = result.Err()
		return p
	}

	p.doc = result.Ok().Doc
	p.file = result.Ok().File

	return p
}

// SendVideo creates a new SendVideo request to send a video file.
func (ctx *Context) SendVideo(filename String) *SendVideo {
	v := &SendVideo{
		ctx:  ctx,
		opts: new(gotgbot.SendVideoOpts),
	}

	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		v.err = result.Err()
		return v
	}

	v.doc = result.Ok().Doc
	v.file = result.Ok().File

	return v
}

// SendVoice creates a new SendVoice request to send a voice message.
func (ctx *Context) SendVoice(filename String) *SendVoice {
	v := &SendVoice{
		ctx:  ctx,
		opts: new(gotgbot.SendVoiceOpts),
	}

	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		v.err = result.Err()
		return v
	}

	v.doc = result.Ok().Doc
	v.file = result.Ok().File

	return v
}

// SendVideoNote creates a new SendVideoNote request to send a video note.
func (ctx *Context) SendVideoNote(filename String) *SendVideoNote {
	vn := &SendVideoNote{
		ctx:  ctx,
		opts: new(gotgbot.SendVideoNoteOpts),
	}

	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		vn.err = result.Err()
		return vn
	}

	vn.doc = result.Ok().Doc
	vn.file = result.Ok().File

	return vn
}

// SendAnimation creates a new SendAnimation request to send an animated GIF or video.
func (ctx *Context) SendAnimation(filename String) *SendAnimation {
	a := &SendAnimation{
		ctx:  ctx,
		opts: new(gotgbot.SendAnimationOpts),
	}

	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		a.err = result.Err()
		return a
	}

	a.doc = result.Ok().Doc
	a.file = result.Ok().File

	return a
}

// SendSticker creates a new SendSticker request to send a sticker.
func (ctx *Context) SendSticker(filename String) *SendSticker {
	s := &SendSticker{
		ctx:  ctx,
		opts: new(gotgbot.SendStickerOpts),
	}

	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		s.err = result.Err()
		return s
	}

	s.doc = result.Ok().Doc
	s.file = result.Ok().File

	return s
}

// SendLocation creates a new SendLocation request to send a location.
func (ctx *Context) SendLocation(latitude, longitude float64) *SendLocation {
	return &SendLocation{
		ctx:       ctx,
		latitude:  latitude,
		longitude: longitude,
		opts:      new(gotgbot.SendLocationOpts),
	}
}

// SendVenue creates a new SendVenue request to send a venue location.
func (ctx *Context) SendVenue(latitude, longitude float64, title, address String) *SendVenue {
	return &SendVenue{
		ctx:       ctx,
		latitude:  latitude,
		longitude: longitude,
		title:     title,
		address:   address,
		opts:      new(gotgbot.SendVenueOpts),
	}
}

// SendContact creates a new SendContact request to send a contact.
func (ctx *Context) SendContact(phoneNumber, firstName String) *SendContact {
	return &SendContact{
		ctx:         ctx,
		phoneNumber: phoneNumber,
		firstName:   firstName,
		opts:        new(gotgbot.SendContactOpts),
	}
}

// Forward creates a new Forward request to forward a message.
func (ctx *Context) Forward(fromChatID, messageID int64) *Forward {
	return &Forward{
		ctx:        ctx,
		fromChatID: fromChatID,
		messageID:  messageID,
		opts:       new(gotgbot.ForwardMessageOpts),
	}
}

// Copy creates a new Copy request to copy a message.
func (ctx *Context) Copy(fromChatID, messageID int64) *Copy {
	return &Copy{
		ctx:        ctx,
		fromChatID: fromChatID,
		messageID:  messageID,
		opts:       new(gotgbot.CopyMessageOpts),
	}
}

// EditMarkup creates a new EditMarkup request to edit a message's reply markup.
func (ctx *Context) EditMarkup(kb keyboard.KeyboardBuilder) *EditMarkup {
	return &EditMarkup{
		ctx:  ctx,
		kb:   kb,
		opts: new(gotgbot.EditMessageReplyMarkupOpts),
	}
}

// EditText creates a new EditText request to edit a message's text.
func (ctx *Context) EditText(text String) *EditText {
	return &EditText{
		ctx:  ctx,
		text: text,
		opts: new(gotgbot.EditMessageTextOpts),
	}
}

// Answer creates a new Answer request to answer a callback query.
func (ctx *Context) Answer(text String) *Answer {
	return &Answer{
		ctx:  ctx,
		text: text,
		opts: new(gotgbot.AnswerCallbackQueryOpts),
	}
}

// SendDice creates a new SendDice request to send a dice animation.
func (ctx *Context) SendDice() *SendDice {
	return &SendDice{
		ctx:  ctx,
		opts: new(gotgbot.SendDiceOpts),
	}
}

// Invoice creates a new Invoice request to send an invoice for payment.
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

// PreCheckout creates a new PreCheckout request to answer a pre-checkout query.
func (ctx *Context) PreCheckout() *PreCheckout {
	return &PreCheckout{
		ctx:  ctx,
		opts: new(gotgbot.AnswerPreCheckoutQueryOpts),
	}
}

// RefundStarPayment creates a new RefundStarPayment request to refund a star payment.
func (ctx *Context) RefundStarPayment(chargeID String) *RefundStarPayment {
	return &RefundStarPayment{
		ctx:      ctx,
		chargeID: chargeID,
		opts:     new(gotgbot.RefundStarPaymentOpts),
	}
}

// Gift creates a Gift request builder.
func (ctx *Context) SendGift(giftID String) *SendGift {
	return &SendGift{
		ctx:    ctx,
		giftID: giftID,
		opts:   new(gotgbot.SendGiftOpts),
	}
}

// ConvertGiftToStars creates a ConvertGiftToStars request builder.
func (ctx *Context) ConvertGiftToStars(businessConnectionID, ownedGiftID String) *ConvertGiftToStars {
	return &ConvertGiftToStars{
		ctx:                  ctx,
		businessConnectionID: businessConnectionID,
		ownedGiftID:          ownedGiftID,
		opts:                 new(gotgbot.ConvertGiftToStarsOpts),
	}
}

// TransferGift creates a TransferGift request builder.
func (ctx *Context) TransferGift(businessConnectionID, ownedGiftID String, newOwnerChatID int64) *TransferGift {
	return &TransferGift{
		ctx:                  ctx,
		businessConnectionID: businessConnectionID,
		ownedGiftID:          ownedGiftID,
		newOwnerChatID:       newOwnerChatID,
		opts:                 new(gotgbot.TransferGiftOpts),
	}
}

// UpgradeGift creates an UpgradeGift request builder.
func (ctx *Context) UpgradeGift(businessConnectionID, ownedGiftID String) *UpgradeGift {
	return &UpgradeGift{
		ctx:                  ctx,
		businessConnectionID: businessConnectionID,
		ownedGiftID:          ownedGiftID,
		opts:                 new(gotgbot.UpgradeGiftOpts),
	}
}

// GetAvailableGifts creates a GetAvailableGifts request builder.
func (ctx *Context) GetAvailableGifts() *GetAvailableGifts {
	return &GetAvailableGifts{
		ctx:  ctx,
		opts: new(gotgbot.GetAvailableGiftsOpts),
	}
}

// GetMyStarBalance creates a GetMyStarBalance request builder.
func (ctx *Context) GetMyStarBalance() *GetMyStarBalance {
	return &GetMyStarBalance{
		ctx:  ctx,
		opts: new(gotgbot.GetMyStarBalanceOpts),
	}
}

// GetStarTransactions creates a GetStarTransactions request builder.
func (ctx *Context) GetStarTransactions() *GetStarTransactions {
	return &GetStarTransactions{
		ctx:  ctx,
		opts: new(gotgbot.GetStarTransactionsOpts),
	}
}

// IsAdmin checks if the effective user is an administrator in the current chat.
func (ctx *Context) IsAdmin() Result[bool] {
	member, err := ctx.Bot.Raw().GetChatMember(ctx.EffectiveChat.Id, ctx.EffectiveUser.Id, nil)
	if err != nil {
		return Err[bool](nil)
	}

	return Ok(member.GetStatus() == "administrator" || member.GetStatus() == "creator")
}

// Args returns command arguments from the message text, excluding the command itself.
func (ctx *Context) Args() Slice[String] {
	return String(ctx.EffectiveMessage.Text).Fields().Skip(1).Collect()
}

// MediaGroup creates a new MediaGroup request to send multiple media as an album.
func (ctx *Context) MediaGroup() *MediaGroup {
	return &MediaGroup{
		ctx:   ctx,
		media: NewSlice[gotgbot.InputMedia](),
		files: NewSlice[*File](),
		opts:  new(gotgbot.SendMediaGroupOpts),
	}
}

// Game creates a new Game request to send a game.
func (ctx *Context) Game(gameShortName String) *Game {
	return &Game{
		ctx:           ctx,
		gameShortName: gameShortName,
		opts:          new(gotgbot.SendGameOpts),
	}
}

// ChatAction creates a new ChatAction request to send a chat action.
func (ctx *Context) ChatAction() *ChatAction {
	return &ChatAction{
		ctx:  ctx,
		opts: new(gotgbot.SendChatActionOpts),
	}
}

// Delete creates a new Delete request to delete a message.
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
		go func() {
			<-time.After(after.Some())
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
