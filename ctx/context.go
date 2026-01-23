package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/core"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/inline"
	"github.com/enetx/tg/input"
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

// BanChatMember creates a new BanChatMember request to ban a user from the chat.
func (ctx *Context) BanChatMember(userID int64) *BanChatMember {
	return &BanChatMember{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.BanChatMemberOpts),
	}
}

// BanChatSenderChat creates a new BanChatSenderChat request to ban a sender chat.
func (ctx *Context) BanChatSenderChat(senderChatID int64) *BanChatSenderChat {
	return &BanChatSenderChat{
		ctx:          ctx,
		senderChatID: senderChatID,
		opts:         new(gotgbot.BanChatSenderChatOpts),
	}
}

// UnbanChatSenderChat creates a new UnbanChatSenderChat request to unban a sender chat.
func (ctx *Context) UnbanChatSenderChat(senderChatID int64) *UnbanChatSenderChat {
	return &UnbanChatSenderChat{
		ctx:          ctx,
		senderChatID: senderChatID,
		opts:         new(gotgbot.UnbanChatSenderChatOpts),
	}
}

// UnbanChatMember creates a new UnbanChatMember request to unban a user from the chat.
func (ctx *Context) UnbanChatMember(userID int64) *UnbanChatMember {
	return &UnbanChatMember{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.UnbanChatMemberOpts),
	}
}

// RestrictChatMember creates a new RestrictChatMember request to restrict a user's permissions in the chat.
func (ctx *Context) RestrictChatMember(userID int64) *RestrictChatMember {
	return &RestrictChatMember{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.RestrictChatMemberOpts),
	}
}

// PromoteChatMember creates a new PromoteChatMember request to promote a user to administrator in the chat.
func (ctx *Context) PromoteChatMember(userID int64) *PromoteChatMember {
	return &PromoteChatMember{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.PromoteChatMemberOpts),
	}
}

// SendPoll creates a new Poll request with the specified question.
func (ctx *Context) SendPoll(question g.String) *SendPoll {
	return &SendPoll{
		ctx:      ctx,
		question: question,
		opts:     new(gotgbot.SendPollOpts),
	}
}

// GiftPremiumSubscription creates a new GiftPremiumSubscription request.
func (ctx *Context) GiftPremiumSubscription(userID, monthCount, starCount int64) *GiftPremiumSubscription {
	return &GiftPremiumSubscription{
		ctx:        ctx,
		userID:     userID,
		monthCount: monthCount,
		starCount:  starCount,
		opts:       new(gotgbot.GiftPremiumSubscriptionOpts),
	}
}

// Reply creates a new Reply request that replies to the current message.
func (ctx *Context) Reply(text g.String) *Reply {
	return &Reply{
		ctx:  ctx,
		text: text,
		opts: new(gotgbot.SendMessageOpts),
	}
}

// EditMessageLiveLocation creates a new EditMessageLiveLocation request.
func (ctx *Context) EditMessageLiveLocation(latitude, longitude float64) *EditMessageLiveLocation {
	return &EditMessageLiveLocation{
		ctx:       ctx,
		latitude:  latitude,
		longitude: longitude,
		opts:      new(gotgbot.EditMessageLiveLocationOpts),
	}
}

// StopMessageLiveLocation creates a new StopMessageLiveLocation request.
func (ctx *Context) StopMessageLiveLocation() *StopMessageLiveLocation {
	return &StopMessageLiveLocation{
		ctx:  ctx,
		opts: new(gotgbot.StopMessageLiveLocationOpts),
	}
}

// GetUserGifts creates a new GetUserGifts request to get information about a user gifts.
func (ctx *Context) GetUserGifts(userID int64) *GetUserGifts {
	return &GetUserGifts{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.GetUserGiftsOpts),
	}
}

// GetUserChatBoosts creates a new GetUserChatBoosts request.
func (ctx *Context) GetUserChatBoosts(userID int64) *GetUserChatBoosts {
	return &GetUserChatBoosts{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.GetUserChatBoostsOpts),
	}
}

// SavePreparedInlineMessage creates a new SavePreparedInlineMessage request.
func (ctx *Context) SavePreparedInlineMessage(userID int64, result inline.QueryResult) *SavePreparedInlineMessage {
	return &SavePreparedInlineMessage{
		ctx:    ctx,
		userID: userID,
		result: result,
		opts:   new(gotgbot.SavePreparedInlineMessageOpts),
	}
}

// CreateChatSubscriptionInviteLink creates a new CreateChatSubscriptionInviteLink request.
// subscriptionPeriod must be 2592000 (30 days), subscriptionPrice is 1-10000 stars.
func (ctx *Context) CreateChatSubscriptionInviteLink(period, price int64) *CreateChatSubscriptionInviteLink {
	return &CreateChatSubscriptionInviteLink{
		ctx:                ctx,
		subscriptionPeriod: period,
		subscriptionPrice:  price,
		opts:               new(gotgbot.CreateChatSubscriptionInviteLinkOpts),
	}
}

// EditChatSubscriptionInviteLink creates a new EditChatSubscriptionInviteLink request.
func (ctx *Context) EditChatSubscriptionInviteLink(inviteLink g.String) *EditChatSubscriptionInviteLink {
	return &EditChatSubscriptionInviteLink{
		ctx:        ctx,
		inviteLink: inviteLink,
		opts:       new(gotgbot.EditChatSubscriptionInviteLinkOpts),
	}
}

// StopPoll creates a new StopPoll request.
func (ctx *Context) StopPoll() *StopPoll {
	return &StopPoll{
		ctx:  ctx,
		opts: new(gotgbot.StopPollOpts),
	}
}

// SendMessage creates a new SendMessage request to send a text message.
func (ctx *Context) SendMessage(text g.String) *SendMessage {
	return &SendMessage{
		ctx:  ctx,
		text: text,
		opts: new(gotgbot.SendMessageOpts),
	}
}

// SendDocument creates a new SendDocument request to send a document file.
func (ctx *Context) SendDocument(filename g.String) *SendDocument {
	d := &SendDocument{
		ctx:  ctx,
		opts: new(gotgbot.SendDocumentOpts),
	}

	result := file.Input(filename)
	if result.IsErr() {
		d.err = result.Err()
		return d
	}

	d.doc = result.Ok().Doc
	d.file = result.Ok().File

	return d
}

// SendAudio creates a new SendAudio request to send an audio file.
func (ctx *Context) SendAudio(filename g.String) *SendAudio {
	a := &SendAudio{
		ctx:  ctx,
		opts: new(gotgbot.SendAudioOpts),
	}

	result := file.Input(filename)
	if result.IsErr() {
		a.err = result.Err()
		return a
	}

	a.doc = result.Ok().Doc
	a.file = result.Ok().File

	return a
}

// SendPhoto creates a new SendPhoto request to send a photo.
func (ctx *Context) SendPhoto(filename g.String) *SendPhoto {
	p := &SendPhoto{
		ctx:  ctx,
		opts: new(gotgbot.SendPhotoOpts),
	}

	result := file.Input(filename)
	if result.IsErr() {
		p.err = result.Err()
		return p
	}

	p.doc = result.Ok().Doc
	p.file = result.Ok().File

	return p
}

// SendVideo creates a new SendVideo request to send a video file.
func (ctx *Context) SendVideo(filename g.String) *SendVideo {
	v := &SendVideo{
		ctx:  ctx,
		opts: new(gotgbot.SendVideoOpts),
	}

	result := file.Input(filename)
	if result.IsErr() {
		v.err = result.Err()
		return v
	}

	v.doc = result.Ok().Doc
	v.file = result.Ok().File

	return v
}

// SendVoice creates a new SendVoice request to send a voice message.
func (ctx *Context) SendVoice(filename g.String) *SendVoice {
	v := &SendVoice{
		ctx:  ctx,
		opts: new(gotgbot.SendVoiceOpts),
	}

	result := file.Input(filename)
	if result.IsErr() {
		v.err = result.Err()
		return v
	}

	v.doc = result.Ok().Doc
	v.file = result.Ok().File

	return v
}

// SendVideoNote creates a new SendVideoNote request to send a video note.
func (ctx *Context) SendVideoNote(filename g.String) *SendVideoNote {
	vn := &SendVideoNote{
		ctx:  ctx,
		opts: new(gotgbot.SendVideoNoteOpts),
	}

	result := file.Input(filename)
	if result.IsErr() {
		vn.err = result.Err()
		return vn
	}

	vn.doc = result.Ok().Doc
	vn.file = result.Ok().File

	return vn
}

// SendAnimation creates a new SendAnimation request to send an animated GIF or video.
func (ctx *Context) SendAnimation(filename g.String) *SendAnimation {
	a := &SendAnimation{
		ctx:  ctx,
		opts: new(gotgbot.SendAnimationOpts),
	}

	result := file.Input(filename)
	if result.IsErr() {
		a.err = result.Err()
		return a
	}

	a.doc = result.Ok().Doc
	a.file = result.Ok().File

	return a
}

// SendSticker creates a new SendSticker request to send a sticker.
func (ctx *Context) SendSticker(filename g.String) *SendSticker {
	s := &SendSticker{
		ctx:  ctx,
		opts: new(gotgbot.SendStickerOpts),
	}

	result := file.Input(filename)
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
func (ctx *Context) SendVenue(latitude, longitude float64, title, address g.String) *SendVenue {
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
func (ctx *Context) SendContact(phoneNumber, firstName g.String) *SendContact {
	return &SendContact{
		ctx:         ctx,
		phoneNumber: phoneNumber,
		firstName:   firstName,
		opts:        new(gotgbot.SendContactOpts),
	}
}

// SendMessageDraft creates a new SendMessageDraft request to send a message draft.
func (ctx *Context) SendMessageDraft(draftID int64, text g.String) *SendMessageDraft {
	return &SendMessageDraft{
		ctx:     ctx,
		draftID: draftID,
		text:    text,
		opts:    new(gotgbot.SendMessageDraftOpts),
	}
}

// ForwardMessage creates a new ForwardMessage request to forward a message.
func (ctx *Context) ForwardMessage(fromChatID, messageID int64) *ForwardMessage {
	return &ForwardMessage{
		ctx:        ctx,
		fromChatID: fromChatID,
		messageID:  messageID,
		opts:       new(gotgbot.ForwardMessageOpts),
	}
}

// CopyMessage creates a new CopyMessage request to copy a message.
func (ctx *Context) CopyMessage(fromChatID, messageID int64) *CopyMessage {
	return &CopyMessage{
		ctx:        ctx,
		fromChatID: fromChatID,
		messageID:  messageID,
		opts:       new(gotgbot.CopyMessageOpts),
	}
}

// EditMessageReplyMarkup creates a new EditMessageReplyMarkup request to edit a message's reply markup.
func (ctx *Context) EditMessageReplyMarkup(kb keyboard.Keyboard) *EditMessageReplyMarkup {
	return &EditMessageReplyMarkup{
		ctx:  ctx,
		kb:   kb,
		opts: new(gotgbot.EditMessageReplyMarkupOpts),
	}
}

// EditMessageCaption creates a new EditMessageCaption request to edit a message's caption.
func (ctx *Context) EditMessageCaption(caption g.String) *EditMessageCaption {
	return &EditMessageCaption{
		ctx:  ctx,
		opts: &gotgbot.EditMessageCaptionOpts{Caption: caption.Std()},
	}
}

// EditMessageMedia creates a new EditMessageMedia request to edit message media.
func (ctx *Context) EditMessageMedia(media input.Media) *EditMessageMedia {
	return &EditMessageMedia{
		ctx:   ctx,
		media: media,
		opts:  new(gotgbot.EditMessageMediaOpts),
	}
}

// EditMessageText creates a new EditMessageText request to edit a message's text.
func (ctx *Context) EditMessageText(text g.String) *EditMessageText {
	return &EditMessageText{
		ctx:  ctx,
		text: text,
		opts: new(gotgbot.EditMessageTextOpts),
	}
}

// AnswerCallbackQuery creates a new AnswerCallbackQuery request to answer a callback query.
func (ctx *Context) AnswerCallbackQuery(text g.String) *AnswerCallbackQuery {
	return &AnswerCallbackQuery{
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

// SendInvoice creates a new SendInvoice request to send an invoice for payment.
func (ctx *Context) SendInvoice(title, desc, payload, currency g.String) *SendInvoice {
	return &SendInvoice{
		ctx:      ctx,
		title:    title,
		desc:     desc,
		payload:  payload,
		currency: currency,
		prices:   g.NewSlice[gotgbot.LabeledPrice](),
		opts:     new(gotgbot.SendInvoiceOpts),
	}
}

// AnswerPreCheckoutQuery creates a new AnswerPreCheckoutQuery request to answer a pre-checkout query.
func (ctx *Context) AnswerPreCheckoutQuery() *AnswerPreCheckoutQuery {
	return &AnswerPreCheckoutQuery{
		ctx:  ctx,
		opts: new(gotgbot.AnswerPreCheckoutQueryOpts),
	}
}

// AnswerShippingQuery creates a new AnswerShippingQuery request to answer a shipping query.
func (ctx *Context) AnswerShippingQuery() *AnswerShippingQuery {
	return &AnswerShippingQuery{
		ctx:     ctx,
		options: g.NewSlice[gotgbot.ShippingOption](),
		opts:    new(gotgbot.AnswerShippingQueryOpts),
	}
}

// RefundStarPayment creates a new RefundStarPayment request to refund a star payment.
func (ctx *Context) RefundStarPayment(chargeID g.String) *RefundStarPayment {
	return &RefundStarPayment{
		ctx:      ctx,
		chargeID: chargeID,
		opts:     new(gotgbot.RefundStarPaymentOpts),
	}
}

// Gift creates a Gift request builder.
func (ctx *Context) SendGift(giftID g.String) *SendGift {
	return &SendGift{
		ctx:    ctx,
		giftID: giftID,
		opts:   new(gotgbot.SendGiftOpts),
	}
}

// ConvertGiftToStars creates a ConvertGiftToStars request builder.
func (ctx *Context) ConvertGiftToStars(businessConnectionID, ownedGiftID g.String) *ConvertGiftToStars {
	return &ConvertGiftToStars{
		ctx:                  ctx,
		businessConnectionID: businessConnectionID,
		ownedGiftID:          ownedGiftID,
		opts:                 new(gotgbot.ConvertGiftToStarsOpts),
	}
}

// CreateInvoiceLink creates a new CreateInvoiceLink request to create an invoice link.
func (ctx *Context) CreateInvoiceLink(title, desc, payload, currency g.String) *CreateInvoiceLink {
	return &CreateInvoiceLink{
		ctx:      ctx,
		title:    title,
		desc:     desc,
		payload:  payload,
		currency: currency,
		opts:     new(gotgbot.CreateInvoiceLinkOpts),
	}
}

// TransferGift creates a TransferGift request builder.
func (ctx *Context) TransferGift(businessConnectionID, ownedGiftID g.String, newOwnerChatID int64) *TransferGift {
	return &TransferGift{
		ctx:                  ctx,
		businessConnectionID: businessConnectionID,
		ownedGiftID:          ownedGiftID,
		newOwnerChatID:       newOwnerChatID,
		opts:                 new(gotgbot.TransferGiftOpts),
	}
}

// UpgradeGift creates an UpgradeGift request builder.
func (ctx *Context) UpgradeGift(businessConnectionID, ownedGiftID g.String) *UpgradeGift {
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

// GetForumTopicIconStickers gets custom emoji stickers that can be used as forum topic icons.
func (ctx *Context) GetForumTopicIconStickers() *GetForumTopicIconStickers {
	return &GetForumTopicIconStickers{
		ctx:  ctx,
		opts: new(gotgbot.GetForumTopicIconStickersOpts),
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

// EditUserStarSubscription creates an EditUserStarSubscription request builder.
func (ctx *Context) EditUserStarSubscription(
	userID int64,
	telegramPaymentChargeID g.String,
	isCanceled bool,
) *EditUserStarSubscription {
	return &EditUserStarSubscription{
		ctx:                     ctx,
		userID:                  userID,
		telegramPaymentChargeID: telegramPaymentChargeID,
		isCanceled:              isCanceled,
		opts:                    new(gotgbot.EditUserStarSubscriptionOpts),
	}
}

// PostStory creates a new PostStory request for posting a photo story.
func (ctx *Context) PostStory(businessConnectionID g.String, content input.StoryContent) *PostStory {
	return &PostStory{
		ctx:                  ctx,
		businessConnectionID: businessConnectionID,
		activePeriod:         86400, // Default 24 hours
		opts:                 new(gotgbot.PostStoryOpts),
		content:              content,
	}
}

// EditPhotoStory creates a new EditStory request for editing a photo story.
func (ctx *Context) EditStory(businessConnectionID g.String, storyID int64, content input.StoryContent) *EditStory {
	return &EditStory{
		ctx:                  ctx,
		businessConnectionID: businessConnectionID,
		storyID:              storyID,
		opts:                 new(gotgbot.EditStoryOpts),
		content:              content,
	}
}

// RepostStory creates a new RepostStory request to repost a story from another business account.
func (ctx *Context) RepostStory(businessConnectionID g.String, fromChatID, fromStoryID int64) *RepostStory {
	return &RepostStory{
		ctx:                  ctx,
		businessConnectionID: businessConnectionID,
		fromChatID:           fromChatID,
		fromStoryID:          fromStoryID,
		opts:                 new(gotgbot.RepostStoryOpts),
	}
}

// DeleteStory creates a new DeleteStory request for the specified business connection and story.
func (ctx *Context) DeleteStory(businessConnectionID g.String, storyID int64) *DeleteStory {
	return &DeleteStory{
		ctx:                  ctx,
		businessConnectionID: businessConnectionID,
		storyID:              storyID,
		opts:                 new(gotgbot.DeleteStoryOpts),
	}
}

// IsAdmin checks if the effective user is an administrator in the current chat.
func (ctx *Context) IsAdmin() g.Result[bool] {
	member, err := ctx.Bot.Raw().GetChatMember(ctx.EffectiveChat.Id, ctx.EffectiveUser.Id, nil)
	if err != nil {
		return g.Err[bool](err)
	}

	return g.Ok(member.GetStatus() == "administrator" || member.GetStatus() == "creator")
}

// Args returns command arguments from the message text, excluding the command itself.
func (ctx *Context) Args() g.Slice[g.String] {
	return g.String(ctx.EffectiveMessage.Text).Fields().Skip(1).Collect()
}

// MediaGroup creates a new MediaGroup request to send multiple media as an album.
func (ctx *Context) MediaGroup() *MediaGroup {
	return &MediaGroup{
		ctx:   ctx,
		media: g.NewSlice[input.Media](),
		opts:  new(gotgbot.SendMediaGroupOpts),
	}
}

// SendChatAction creates a new SendChatAction request to send a chat action.
func (ctx *Context) SendChatAction() *SendChatAction {
	return &SendChatAction{
		ctx:  ctx,
		opts: new(gotgbot.SendChatActionOpts),
	}
}

// DeleteMessage creates a new DeleteMessage request to delete a message.
func (ctx *Context) DeleteMessage() *DeleteMessage {
	return &DeleteMessage{
		ctx:  ctx,
		opts: new(gotgbot.DeleteMessageOpts),
	}
}

// DeleteMessages creates a new DeleteMessages request to delete multiple messages.
func (ctx *Context) DeleteMessages() *DeleteMessages {
	return &DeleteMessages{
		ctx:        ctx,
		messageIDs: g.NewSlice[int64](),
		opts:       new(gotgbot.DeleteMessagesOpts),
	}
}

// ForwardMessages creates a new ForwardMessages request to forward multiple messages.
func (ctx *Context) ForwardMessages() *ForwardMessages {
	return &ForwardMessages{
		ctx:        ctx,
		messageIDs: g.NewSlice[int64](),
		opts:       new(gotgbot.ForwardMessagesOpts),
	}
}

// CopyMessages creates a new CopyMessages request to copy multiple messages.
func (ctx *Context) CopyMessages() *CopyMessages {
	return &CopyMessages{
		ctx:        ctx,
		messageIDs: g.NewSlice[int64](),
		opts:       new(gotgbot.CopyMessagesOpts),
	}
}

// SendPaidMedia creates a new SendPaidMedia request to send paid media content.
func (ctx *Context) SendPaidMedia(starCount int64) *SendPaidMedia {
	return &SendPaidMedia{
		ctx:       ctx,
		starCount: starCount,
		media:     g.NewSlice[input.PaidMedia](),
		opts:      new(gotgbot.SendPaidMediaOpts),
	}
}

// CreateForumTopic creates a new CreateForumTopic request.
func (ctx *Context) CreateForumTopic(name g.String) *CreateForumTopic {
	return &CreateForumTopic{
		ctx:  ctx,
		name: name,
		opts: new(gotgbot.CreateForumTopicOpts),
	}
}

// EditForumTopic creates a new EditForumTopic request.
func (ctx *Context) EditForumTopic(messageThreadID int64) *EditForumTopic {
	return &EditForumTopic{
		ctx:             ctx,
		messageThreadID: messageThreadID,
		opts:            new(gotgbot.EditForumTopicOpts),
	}
}

// CloseForumTopic creates a new CloseForumTopic request.
func (ctx *Context) CloseForumTopic(messageThreadID int64) *CloseForumTopic {
	return &CloseForumTopic{
		ctx:             ctx,
		messageThreadID: messageThreadID,
		opts:            new(gotgbot.CloseForumTopicOpts),
	}
}

// ReopenForumTopic creates a new ReopenForumTopic request.
func (ctx *Context) ReopenForumTopic(messageThreadID int64) *ReopenForumTopic {
	return &ReopenForumTopic{
		ctx:             ctx,
		messageThreadID: messageThreadID,
		opts:            new(gotgbot.ReopenForumTopicOpts),
	}
}

// DeleteForumTopic creates a new DeleteForumTopic request.
func (ctx *Context) DeleteForumTopic(messageThreadID int64) *DeleteForumTopic {
	return &DeleteForumTopic{
		ctx:             ctx,
		messageThreadID: messageThreadID,
		opts:            new(gotgbot.DeleteForumTopicOpts),
	}
}

// EditGeneralForumTopic creates a new EditGeneralForumTopic request.
func (ctx *Context) EditGeneralForumTopic(name g.String) *EditGeneralForumTopic {
	return &EditGeneralForumTopic{
		ctx:  ctx,
		name: name,
		opts: new(gotgbot.EditGeneralForumTopicOpts),
	}
}

// CloseGeneralForumTopic creates a new CloseGeneralForumTopic request.
func (ctx *Context) CloseGeneralForumTopic() *CloseGeneralForumTopic {
	return &CloseGeneralForumTopic{
		ctx:  ctx,
		opts: new(gotgbot.CloseGeneralForumTopicOpts),
	}
}

// HideGeneralForumTopic hides the general forum topic.
func (ctx *Context) HideGeneralForumTopic() *HideGeneralForumTopic {
	return &HideGeneralForumTopic{
		ctx:  ctx,
		opts: new(gotgbot.HideGeneralForumTopicOpts),
	}
}

// UnhideGeneralForumTopic unhides the general forum topic.
func (ctx *Context) UnhideGeneralForumTopic() *UnhideGeneralForumTopic {
	return &UnhideGeneralForumTopic{
		ctx:  ctx,
		opts: new(gotgbot.UnhideGeneralForumTopicOpts),
	}
}

// ReopenGeneralForumTopic reopens the general forum topic.
func (ctx *Context) ReopenGeneralForumTopic() *ReopenGeneralForumTopic {
	return &ReopenGeneralForumTopic{
		ctx:  ctx,
		opts: new(gotgbot.ReopenGeneralForumTopicOpts),
	}
}

// SetChatTitle creates a new SetChatTitle request.
func (ctx *Context) SetChatTitle(title g.String) *SetChatTitle {
	return &SetChatTitle{
		ctx:   ctx,
		title: title,
		opts:  new(gotgbot.SetChatTitleOpts),
	}
}

// SetChatDescription creates a new SetChatDescription request.
func (ctx *Context) SetChatDescription(description g.String) *SetChatDescription {
	return &SetChatDescription{
		ctx:  ctx,
		opts: &gotgbot.SetChatDescriptionOpts{Description: description.Std()},
	}
}

// SetChatPhoto creates a new SetChatPhoto request.
func (ctx *Context) SetChatPhoto(filename g.String) *SetChatPhoto {
	p := &SetChatPhoto{
		ctx:  ctx,
		opts: new(gotgbot.SetChatPhotoOpts),
	}

	result := file.Input(filename)
	if result.IsErr() {
		p.err = result.Err()
		return p
	}

	p.doc = result.Ok().Doc.(gotgbot.InputFile)
	p.file = result.Ok().File

	return p
}

// DeleteChatPhoto creates a new DeleteChatPhoto request.
func (ctx *Context) DeleteChatPhoto() *DeleteChatPhoto {
	return &DeleteChatPhoto{
		ctx:  ctx,
		opts: new(gotgbot.DeleteChatPhotoOpts),
	}
}

// SetChatPermissions creates a new SetChatPermissions request.
func (ctx *Context) SetChatPermissions() *SetChatPermissions {
	return &SetChatPermissions{
		ctx:  ctx,
		opts: new(gotgbot.SetChatPermissionsOpts),
	}
}

// SetChatAdministratorCustomTitle creates a new SetChatAdministratorCustomTitle request.
func (ctx *Context) SetChatAdministratorCustomTitle(
	userID int64,
	customTitle g.String,
) *SetChatAdministratorCustomTitle {
	return &SetChatAdministratorCustomTitle{
		ctx:         ctx,
		userID:      userID,
		customTitle: customTitle,
		opts:        new(gotgbot.SetChatAdministratorCustomTitleOpts),
	}
}

// PinChatMessage creates a new PinChatMessage request.
func (ctx *Context) PinChatMessage(messageID int64) *PinChatMessage {
	return &PinChatMessage{
		ctx:       ctx,
		messageID: messageID,
		opts:      new(gotgbot.PinChatMessageOpts),
	}
}

// UnpinChatMessage creates a new UnpinChatMessage request.
func (ctx *Context) UnpinChatMessage() *UnpinChatMessage {
	return &UnpinChatMessage{
		ctx:  ctx,
		opts: new(gotgbot.UnpinChatMessageOpts),
	}
}

// UnpinAllChatMessages creates a new UnpinAllChatMessages request.
func (ctx *Context) UnpinAllChatMessages() *UnpinAllChatMessages {
	return &UnpinAllChatMessages{
		ctx:  ctx,
		opts: new(gotgbot.UnpinAllChatMessagesOpts),
	}
}

// UnpinAllForumTopicMessages unpins all messages in a forum topic.
func (ctx *Context) UnpinAllForumTopicMessages(messageThreadID int64) *UnpinAllForumTopicMessages {
	return &UnpinAllForumTopicMessages{
		ctx:             ctx,
		messageThreadID: messageThreadID,
		opts:            new(gotgbot.UnpinAllForumTopicMessagesOpts),
	}
}

// UnpinAllGeneralForumTopicMessages unpins all messages in the general forum topic.
func (ctx *Context) UnpinAllGeneralForumTopicMessages() *UnpinAllGeneralForumTopicMessages {
	return &UnpinAllGeneralForumTopicMessages{
		ctx:  ctx,
		opts: new(gotgbot.UnpinAllGeneralForumTopicMessagesOpts),
	}
}

// GetChat creates a new GetChat request to get chat information.
func (ctx *Context) GetChat() *GetChat {
	return &GetChat{
		ctx:  ctx,
		opts: new(gotgbot.GetChatOpts),
	}
}

// GetChatAdministrators creates a new GetChatAdministrators request.
func (ctx *Context) GetChatAdministrators() *GetChatAdministrators {
	return &GetChatAdministrators{
		ctx:  ctx,
		opts: new(gotgbot.GetChatAdministratorsOpts),
	}
}

// GetChatGifts creates a new GetChatGifts request to get information about a chat gifts.
func (ctx *Context) GetChatGifts() *GetChatGifts {
	return &GetChatGifts{
		ctx:  ctx,
		opts: new(gotgbot.GetChatGiftsOpts),
	}
}

// GetChatMember creates a new GetChatMember request to get information about a chat member.
func (ctx *Context) GetChatMember(userID int64) *GetChatMember {
	return &GetChatMember{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.GetChatMemberOpts),
	}
}

// GetChatMemberCount creates a new GetChatMemberCount request.
func (ctx *Context) GetChatMemberCount() *GetChatMemberCount {
	return &GetChatMemberCount{
		ctx:  ctx,
		opts: new(gotgbot.GetChatMemberCountOpts),
	}
}

// CreateNewStickerSet creates a new sticker set.
func (ctx *Context) CreateNewStickerSet(userID int64, name, title g.String) *CreateNewStickerSet {
	return &CreateNewStickerSet{
		ctx:    ctx,
		userID: userID,
		name:   name,
		title:  title,
		opts:   new(gotgbot.CreateNewStickerSetOpts),
	}
}

// AddStickerToSet adds a sticker to an existing sticker set.
func (ctx *Context) AddStickerToSet(userID int64, name g.String) *AddStickerToSet {
	return &AddStickerToSet{
		ctx:    ctx,
		userID: userID,
		name:   name,
		opts:   new(gotgbot.AddStickerToSetOpts),
	}
}

// GetStickerSet gets sticker set information by name.
func (ctx *Context) GetStickerSet(name g.String) *GetStickerSet {
	return &GetStickerSet{
		ctx:  ctx,
		name: name,
		opts: new(gotgbot.GetStickerSetOpts),
	}
}

// DeleteStickerSet deletes a sticker set.
func (ctx *Context) DeleteStickerSet(name g.String) *DeleteStickerSet {
	return &DeleteStickerSet{
		ctx:  ctx,
		name: name,
		opts: new(gotgbot.DeleteStickerSetOpts),
	}
}

// DeleteStickerFromSet deletes a sticker from a set.
func (ctx *Context) DeleteStickerFromSet(sticker file.InputFile) *DeleteStickerFromSet {
	return &DeleteStickerFromSet{
		ctx:     ctx,
		sticker: sticker.Doc,
		opts:    new(gotgbot.DeleteStickerFromSetOpts),
	}
}

// SetStickerPositionInSet sets the position of a sticker in a set.
func (ctx *Context) SetStickerPositionInSet(sticker file.InputFile, position int64) *SetStickerPositionInSet {
	return &SetStickerPositionInSet{
		ctx:      ctx,
		sticker:  sticker.Doc,
		position: position,
		opts:     new(gotgbot.SetStickerPositionInSetOpts),
	}
}

// ReplaceStickerInSet replaces a sticker in a sticker set.
func (ctx *Context) ReplaceStickerInSet(
	userID int64,
	name, oldSticker g.String,
	sticker gotgbot.InputSticker,
) *ReplaceStickerInSet {
	return &ReplaceStickerInSet{
		ctx:        ctx,
		userID:     userID,
		name:       name,
		oldSticker: oldSticker,
		sticker:    sticker,
		opts:       new(gotgbot.ReplaceStickerInSetOpts),
	}
}

// SetCustomEmojiStickerSetThumbnail sets the thumbnail of a custom emoji sticker set.
func (ctx *Context) SetCustomEmojiStickerSetThumbnail(name g.String) *SetCustomEmojiStickerSetThumbnail {
	return &SetCustomEmojiStickerSetThumbnail{
		ctx:  ctx,
		name: name,
		opts: new(gotgbot.SetCustomEmojiStickerSetThumbnailOpts),
	}
}

// SetStickerSetTitle sets the title of a sticker set.
func (ctx *Context) SetStickerSetTitle(name, title g.String) *SetStickerSetTitle {
	return &SetStickerSetTitle{
		ctx:   ctx,
		name:  name,
		title: title,
		opts:  new(gotgbot.SetStickerSetTitleOpts),
	}
}

// SetUserEmojiStatus sets the emoji status for a user.
func (ctx *Context) SetUserEmojiStatus(userID int64) *SetUserEmojiStatus {
	return &SetUserEmojiStatus{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.SetUserEmojiStatusOpts),
	}
}

// SetStickerEmojiList sets the emoji list for a sticker.
func (ctx *Context) SetStickerEmojiList(sticker file.InputFile) *SetStickerEmojiList {
	return &SetStickerEmojiList{
		ctx:     ctx,
		sticker: sticker.Doc,
		opts:    new(gotgbot.SetStickerEmojiListOpts),
	}
}

// SetStickerKeywords sets keywords for a sticker.
func (ctx *Context) SetStickerKeywords(sticker file.InputFile) *SetStickerKeywords {
	return &SetStickerKeywords{
		ctx:     ctx,
		sticker: sticker.Doc,
		opts:    new(gotgbot.SetStickerKeywordsOpts),
	}
}

// SetStickerMaskPosition sets the mask position for a sticker.
func (ctx *Context) SetStickerMaskPosition(sticker file.InputFile) *SetStickerMaskPosition {
	return &SetStickerMaskPosition{
		ctx:     ctx,
		sticker: sticker.Doc,
		opts:    new(gotgbot.SetStickerMaskPositionOpts),
	}
}

// SetStickerSetThumbnail sets the thumbnail for a sticker set.
func (ctx *Context) SetStickerSetThumbnail(name g.String, userID int64) *SetStickerSetThumbnail {
	return &SetStickerSetThumbnail{
		ctx:    ctx,
		name:   name,
		userID: userID,
		opts:   new(gotgbot.SetStickerSetThumbnailOpts),
	}
}

// UploadStickerFile uploads a sticker file.
func (ctx *Context) UploadStickerFile(userID int64, format g.String) *UploadStickerFile {
	return &UploadStickerFile{
		ctx:    ctx,
		userID: userID,
		format: format,
		opts:   new(gotgbot.UploadStickerFileOpts),
	}
}

// GetCustomEmojiStickers gets custom emoji stickers by IDs.
func (ctx *Context) GetCustomEmojiStickers(ids g.Slice[g.String]) *GetCustomEmojiStickers {
	return &GetCustomEmojiStickers{
		ctx:            ctx,
		customEmojiIDs: ids,
		opts:           new(gotgbot.GetCustomEmojiStickersOpts),
	}
}

// SetChatMenuButton sets the menu button of a chat.
func (ctx *Context) SetChatMenuButton() *SetChatMenuButton {
	return &SetChatMenuButton{
		ctx:  ctx,
		opts: new(gotgbot.SetChatMenuButtonOpts),
	}
}

// GetChatMenuButton gets the menu button of a chat.
func (ctx *Context) GetChatMenuButton() *GetChatMenuButton {
	return &GetChatMenuButton{
		ctx:  ctx,
		opts: new(gotgbot.GetChatMenuButtonOpts),
	}
}

// VerifyUser verifies a user.
func (ctx *Context) VerifyUser(userID int64) *VerifyUser {
	return &VerifyUser{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.VerifyUserOpts),
	}
}

// VerifyChat verifies a chat.
func (ctx *Context) VerifyChat(chatID int64) *VerifyChat {
	return &VerifyChat{
		ctx:    ctx,
		chatID: chatID,
		opts:   new(gotgbot.VerifyChatOpts),
	}
}

// RemoveUserVerification removes verification from a user.
func (ctx *Context) RemoveUserVerification(userID int64) *RemoveUserVerification {
	return &RemoveUserVerification{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.RemoveUserVerificationOpts),
	}
}

// RemoveChatVerification removes verification from a chat.
func (ctx *Context) RemoveChatVerification(chatID int64) *RemoveChatVerification {
	return &RemoveChatVerification{
		ctx:    ctx,
		chatID: chatID,
		opts:   new(gotgbot.RemoveChatVerificationOpts),
	}
}

// SendChecklist sends a checklist message.
func (ctx *Context) SendChecklist(businessConnectionID, title g.String) *SendChecklist {
	return &SendChecklist{
		ctx:                  ctx,
		checklist:            gotgbot.InputChecklist{Title: title.Std()},
		businessConnectionID: businessConnectionID,
		opts:                 new(gotgbot.SendChecklistOpts),
		taskIDCounter:        0,
	}
}

// EditMessageChecklist edits a checklist message.
func (ctx *Context) EditMessageChecklist(businessConnectionID g.String) *EditMessageChecklist {
	return &EditMessageChecklist{
		ctx:                  ctx,
		businessConnectionID: businessConnectionID,
		opts:                 new(gotgbot.EditMessageChecklistOpts),
		taskIDCounter:        0,
	}
}

// SendGame creates a new SendGame request to send a game.
func (ctx *Context) SendGame(gameShortName g.String) *SendGame {
	return &SendGame{
		ctx:           ctx,
		gameShortName: gameShortName,
		opts:          new(gotgbot.SendGameOpts),
	}
}

// SetGameScore sets the score for a game.
func (ctx *Context) SetGameScore(userID, score int64) *SetGameScore {
	return &SetGameScore{
		ctx:    ctx,
		userID: userID,
		score:  score,
		opts:   new(gotgbot.SetGameScoreOpts),
	}
}

// GetGameHighScores gets high scores for a game.
func (ctx *Context) GetGameHighScores(userID int64) *GetGameHighScores {
	return &GetGameHighScores{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.GetGameHighScoresOpts),
	}
}

// CreateChatInviteLink creates a new invite link for a chat.
func (ctx *Context) CreateChatInviteLink() *CreateChatInviteLink {
	return &CreateChatInviteLink{
		ctx:  ctx,
		opts: new(gotgbot.CreateChatInviteLinkOpts),
	}
}

// EditChatInviteLink edits an existing chat invite link.
func (ctx *Context) EditChatInviteLink(inviteLink g.String) *EditChatInviteLink {
	return &EditChatInviteLink{
		ctx:        ctx,
		inviteLink: inviteLink,
		opts:       new(gotgbot.EditChatInviteLinkOpts),
	}
}

// RevokeChatInviteLink revokes a chat invite link.
func (ctx *Context) RevokeChatInviteLink(inviteLink g.String) *RevokeChatInviteLink {
	return &RevokeChatInviteLink{
		ctx:        ctx,
		inviteLink: inviteLink,
		opts:       new(gotgbot.RevokeChatInviteLinkOpts),
	}
}

// ApproveChatJoinRequest approves a chat join request.
func (ctx *Context) ApproveChatJoinRequest(userID int64) *ApproveChatJoinRequest {
	return &ApproveChatJoinRequest{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.ApproveChatJoinRequestOpts),
	}
}

// DeclineChatJoinRequest declines a chat join request.
func (ctx *Context) DeclineChatJoinRequest(userID int64) *DeclineChatJoinRequest {
	return &DeclineChatJoinRequest{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.DeclineChatJoinRequestOpts),
	}
}

// LeaveChat leaves a chat.
func (ctx *Context) LeaveChat() *LeaveChat {
	return &LeaveChat{
		ctx:  ctx,
		opts: new(gotgbot.LeaveChatOpts),
	}
}

// GetFile gets information about a file.
func (ctx *Context) GetFile(fileID g.String) *GetFile {
	return &GetFile{
		ctx:    ctx,
		fileID: fileID,
		opts:   new(gotgbot.GetFileOpts),
	}
}

// SetChatStickerSet sets a chat's sticker set.
func (ctx *Context) SetChatStickerSet(stickerSetName g.String) *SetChatStickerSet {
	return &SetChatStickerSet{
		ctx:            ctx,
		stickerSetName: stickerSetName,
		opts:           new(gotgbot.SetChatStickerSetOpts),
	}
}

// DeleteChatStickerSet deletes a chat's sticker set.
func (ctx *Context) DeleteChatStickerSet() *DeleteChatStickerSet {
	return &DeleteChatStickerSet{
		ctx:  ctx,
		opts: new(gotgbot.DeleteChatStickerSetOpts),
	}
}

// AnswerWebAppQuery answers a web app query.
func (ctx *Context) AnswerWebAppQuery(webAppQueryID g.String, result inline.QueryResult) *AnswerWebAppQuery {
	return &AnswerWebAppQuery{
		ctx:           ctx,
		webAppQueryID: webAppQueryID,
		result:        result,
		opts:          new(gotgbot.AnswerWebAppQueryOpts),
	}
}

// SetMessageReaction creates a new SetMessageReaction request to set reactions on a message.
func (ctx *Context) SetMessageReaction(messageID int64) *SetMessageReaction {
	return &SetMessageReaction{
		ctx:       ctx,
		messageID: messageID,
		opts:      new(gotgbot.SetMessageReactionOpts),
	}
}

// AnswerInlineQuery answers an inline query.
func (ctx *Context) AnswerInlineQuery(inlineQueryID g.String) *AnswerInlineQuery {
	return &AnswerInlineQuery{
		ctx:           ctx,
		inlineQueryID: inlineQueryID,
		results:       g.NewSlice[gotgbot.InlineQueryResult](),
		opts:          new(gotgbot.AnswerInlineQueryOpts),
	}
}

// GetUserProfilePhotos gets user profile photos.
func (ctx *Context) GetUserProfilePhotos(userID int64) *GetUserProfilePhotos {
	return &GetUserProfilePhotos{
		ctx:    ctx,
		userID: userID,
		opts:   new(gotgbot.GetUserProfilePhotosOpts),
	}
}

// ExportChatInviteLink exports a basic chat invite link.
func (ctx *Context) ExportChatInviteLink() *ExportChatInviteLink {
	return &ExportChatInviteLink{
		ctx:  ctx,
		opts: new(gotgbot.ExportChatInviteLinkOpts),
	}
}

// SetPassportDataErrors sets passport data errors for the specified user.
func (ctx *Context) SetPassportDataErrors(userID int64) *SetPassportDataErrors {
	return &SetPassportDataErrors{
		ctx:    ctx,
		userID: userID,
		errors: g.NewSlice[gotgbot.PassportElementError](),
		opts:   new(gotgbot.SetPassportDataErrorsOpts),
	}
}

// ApproveSuggestedPost approves a suggested post in a direct messages chat.
func (ctx *Context) ApproveSuggestedPost() *ApproveSuggestedPost {
	return &ApproveSuggestedPost{
		ctx:  ctx,
		opts: new(gotgbot.ApproveSuggestedPostOpts),
	}
}

// DeclineSuggestedPost declines a suggested post in a direct messages chat.
func (ctx *Context) DeclineSuggestedPost() *DeclineSuggestedPost {
	return &DeclineSuggestedPost{
		ctx:  ctx,
		opts: new(gotgbot.DeclineSuggestedPostOpts),
	}
}

func (ctx *Context) timers(
	after g.Option[time.Duration],
	deleteAfter g.Option[time.Duration],
	send func() g.Result[*gotgbot.Message],
) g.Result[*gotgbot.Message] {
	if after.IsSome() {
		go func() {
			<-time.After(after.Some())
			msg := send()
			if msg.IsOk() && deleteAfter.IsSome() {
				ctx.DeleteMessage().MessageID(msg.Ok().MessageId).After(deleteAfter.Some()).Send()
			}
		}()

		return g.Ok[*gotgbot.Message](nil)
	}

	msg := send()

	if msg.IsOk() && deleteAfter.IsSome() {
		ctx.DeleteMessage().MessageID(msg.Ok().MessageId).After(deleteAfter.Some()).Send()
	}

	return msg
}
