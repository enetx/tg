package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/suggested"
	"github.com/enetx/tg/types/effects"
)

type SendMessage struct {
	ctx         *Context
	text        g.String
	chatID      g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
	opts        *gotgbot.SendMessageOpts
}

// Entities sets special entities in the message text using Entities builder.
func (sm *SendMessage) Entities(e *entities.Entities) *SendMessage {
	sm.opts.Entities = e.Std()
	return sm
}

// After schedules the message to be sent after the specified duration.
func (sm *SendMessage) After(duration time.Duration) *SendMessage {
	sm.after = g.Some(duration)
	return sm
}

// DeleteAfter schedules the message to be deleted after the specified duration.
func (sm *SendMessage) DeleteAfter(duration time.Duration) *SendMessage {
	sm.deleteAfter = g.Some(duration)
	return sm
}

// To sets the target chat ID for the message.
func (sm *SendMessage) To(chatID int64) *SendMessage {
	sm.chatID = g.Some(chatID)
	return sm
}

// HTML sets the message parse mode to HTML.
func (sm *SendMessage) HTML() *SendMessage {
	sm.opts.ParseMode = "HTML"
	return sm
}

// Markdown sets the message parse mode to MarkdownV2.
func (sm *SendMessage) Markdown() *SendMessage {
	sm.opts.ParseMode = "MarkdownV2"
	return sm
}

// Silent disables notification for the message.
func (sm *SendMessage) Silent() *SendMessage {
	sm.opts.DisableNotification = true
	return sm
}

// Effect sets a message effect for the message.
func (sm *SendMessage) Effect(effect effects.EffectType) *SendMessage {
	sm.opts.MessageEffectId = effect.String()
	return sm
}

// Reply sets reply parameters using the reply builder.
func (sm *SendMessage) Reply(params *reply.Parameters) *SendMessage {
	if params != nil {
		sm.opts.ReplyParameters = params.Std()
	}
	return sm
}

// Markup sets the reply markup keyboard for the message.
func (sm *SendMessage) Markup(kb keyboard.Keyboard) *SendMessage {
	sm.opts.ReplyMarkup = kb.Markup()
	return sm
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (sm *SendMessage) AllowPaidBroadcast() *SendMessage {
	sm.opts.AllowPaidBroadcast = true
	return sm
}

// Thread sets the message thread ID for the message.
func (sm *SendMessage) Thread(id int64) *SendMessage {
	sm.opts.MessageThreadId = id
	return sm
}

// ForceReply forces users to reply to the message.
func (sm *SendMessage) ForceReply() *SendMessage {
	sm.opts.ReplyMarkup = gotgbot.ForceReply{ForceReply: true}
	return sm
}

// RemoveKeyboard removes the custom keyboard.
func (sm *SendMessage) RemoveKeyboard() *SendMessage {
	sm.opts.ReplyMarkup = gotgbot.ReplyKeyboardRemove{RemoveKeyboard: true}
	return sm
}

// Preview sets link preview options for the message.
func (sm *SendMessage) Preview(p *preview.Preview) *SendMessage {
	sm.opts.LinkPreviewOptions = p.Std()
	return sm
}

// Business sets the business connection ID for the message.
func (sm *SendMessage) Business(id g.String) *SendMessage {
	sm.opts.BusinessConnectionId = id.Std()
	return sm
}

// SuggestedPost sets suggested post parameters for direct messages chats.
func (sm *SendMessage) SuggestedPost(params *suggested.PostParameters) *SendMessage {
	if params != nil {
		sm.opts.SuggestedPostParameters = params.Std()
	}
	return sm
}

// Protect enables content protection for the message.
func (sm *SendMessage) Protect() *SendMessage {
	sm.opts.ProtectContent = true
	return sm
}

// Timeout sets a custom timeout for this request.
func (sm *SendMessage) Timeout(duration time.Duration) *SendMessage {
	if sm.opts.RequestOpts == nil {
		sm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sm.opts.RequestOpts.Timeout = duration

	return sm
}

// APIURL sets a custom API URL for this request.
func (sm *SendMessage) APIURL(url g.String) *SendMessage {
	if sm.opts.RequestOpts == nil {
		sm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sm.opts.RequestOpts.APIURL = url.Std()

	return sm
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (sm *SendMessage) DirectMessagesTopic(topicID int64) *SendMessage {
	sm.opts.DirectMessagesTopicId = topicID
	return sm
}

// Send sends the message to Telegram and returns the result.
func (sm *SendMessage) Send() g.Result[*gotgbot.Message] {
	return sm.ctx.timers(sm.after, sm.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := sm.chatID.UnwrapOr(sm.ctx.EffectiveChat.Id)
		return g.ResultOf(sm.ctx.Bot.Raw().SendMessage(chatID, sm.text.Std(), sm.opts))
	})
}
