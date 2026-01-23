package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/g/ref"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/suggested"
	"github.com/enetx/tg/types/effects"
)

type CopyMessage struct {
	ctx         *Context
	fromChatID  int64
	messageID   int64
	opts        *gotgbot.CopyMessageOpts
	toChatID    g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
}

// CaptionEntities sets custom entities for the copied message caption.
func (c *CopyMessage) CaptionEntities(e *entities.Entities) *CopyMessage {
	c.opts.CaptionEntities = e.Std()
	return c
}

// After schedules the copy to be sent after the specified duration.
func (c *CopyMessage) After(duration time.Duration) *CopyMessage {
	c.after = g.Some(duration)
	return c
}

// DeleteAfter schedules the copied message to be deleted after the specified duration.
func (c *CopyMessage) DeleteAfter(duration time.Duration) *CopyMessage {
	c.deleteAfter = g.Some(duration)
	return c
}

// Caption sets a new caption for the copied message.
func (c *CopyMessage) Caption(caption g.String) *CopyMessage {
	c.opts.Caption = ref.Of(caption.Std())
	return c
}

// HTML sets the caption parse mode to HTML.
func (c *CopyMessage) HTML() *CopyMessage {
	c.opts.ParseMode = "HTML"
	return c
}

// Markdown sets the caption parse mode to MarkdownV2.
func (c *CopyMessage) Markdown() *CopyMessage {
	c.opts.ParseMode = "MarkdownV2"
	return c
}

// Silent disables notification for the copied message.
func (c *CopyMessage) Silent() *CopyMessage {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the copied message.
func (c *CopyMessage) Protect() *CopyMessage {
	c.opts.ProtectContent = true
	return c
}

// Markup sets the reply markup keyboard for the copied message.
func (c *CopyMessage) Markup(kb keyboard.Keyboard) *CopyMessage {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// Thread sets the message thread ID for forum supergroups.
func (c *CopyMessage) Thread(id int64) *CopyMessage {
	c.opts.MessageThreadId = id
	return c
}

// VideoStartAt sets the start timestamp for copied video.
func (c *CopyMessage) VideoStartAt(offset time.Duration) *CopyMessage {
	c.opts.VideoStartTimestamp = int64(offset.Seconds())
	return c
}

// ShowCaptionAbove displays the caption above the media instead of below.
func (c *CopyMessage) ShowCaptionAbove() *CopyMessage {
	c.opts.ShowCaptionAboveMedia = true
	return c
}

// AllowPaidBroadcast allows paid broadcast for high-speed delivery.
func (c *CopyMessage) AllowPaidBroadcast() *CopyMessage {
	c.opts.AllowPaidBroadcast = true
	return c
}

// Reply sets reply parameters using the reply builder.
func (c *CopyMessage) Reply(params *reply.Parameters) *CopyMessage {
	if params != nil {
		c.opts.ReplyParameters = params.Std()
	}
	return c
}

// Effect sets a message effect for the copied message.
func (c *CopyMessage) Effect(effect effects.EffectType) *CopyMessage {
	c.opts.MessageEffectId = effect.String()
	return c
}

// Timeout sets a custom timeout for this request.
func (c *CopyMessage) Timeout(duration time.Duration) *CopyMessage {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *CopyMessage) APIURL(url g.String) *CopyMessage {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// SuggestedPost sets suggested post parameters for direct messages chats.
func (c *CopyMessage) SuggestedPost(params *suggested.PostParameters) *CopyMessage {
	if params != nil {
		c.opts.SuggestedPostParameters = params.Std()
	}
	return c
}

// To sets the target chat ID for the copied message.
func (c *CopyMessage) To(chatID int64) *CopyMessage {
	c.toChatID = g.Some(chatID)
	return c
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (c *CopyMessage) DirectMessagesTopic(topicID int64) *CopyMessage {
	c.opts.DirectMessagesTopicId = topicID
	return c
}

// Send copies the message to the target chat and returns the result.
func (c *CopyMessage) Send() g.Result[*gotgbot.MessageId] {
	if c.after.IsSome() {
		go func() {
			<-time.After(c.after.Some())
			chatID := c.toChatID.UnwrapOr(c.ctx.EffectiveChat.Id)
			msgID, err := c.ctx.Bot.Raw().CopyMessage(chatID, c.fromChatID, c.messageID, c.opts)
			if err == nil && msgID != nil && c.deleteAfter.IsSome() {
				c.ctx.DeleteMessage().MessageID(msgID.MessageId).ChatID(chatID).After(c.deleteAfter.Some()).Send()
			}
		}()

		return g.Ok[*gotgbot.MessageId](nil)
	}

	chatID := c.toChatID.UnwrapOr(c.ctx.EffectiveChat.Id)
	result := g.ResultOf(c.ctx.Bot.Raw().CopyMessage(chatID, c.fromChatID, c.messageID, c.opts))

	if result.IsOk() && c.deleteAfter.IsSome() {
		c.ctx.DeleteMessage().MessageID(result.Ok().MessageId).ChatID(chatID).After(c.deleteAfter.Some()).Send()
	}

	return result
}
