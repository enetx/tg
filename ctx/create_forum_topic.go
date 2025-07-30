package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// CreateForumTopic represents a request to create a forum topic.
type CreateForumTopic struct {
	ctx    *Context
	name   g.String
	opts   *gotgbot.CreateForumTopicOpts
	chatID g.Option[int64]
}

// IconColor sets the color of the topic icon in RGB format.
func (cf *CreateForumTopic) IconColor(color int64) *CreateForumTopic {
	cf.opts.IconColor = color
	return cf
}

// IconCustomEmojiID sets the unique identifier of the custom emoji.
func (cf *CreateForumTopic) IconCustomEmojiID(emojiID g.String) *CreateForumTopic {
	cf.opts.IconCustomEmojiId = emojiID.Std()
	return cf
}

// Timeout sets a custom timeout for this request.
func (cf *CreateForumTopic) Timeout(duration time.Duration) *CreateForumTopic {
	if cf.opts.RequestOpts == nil {
		cf.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cf.opts.RequestOpts.Timeout = duration

	return cf
}

// APIURL sets a custom API URL for this request.
func (cf *CreateForumTopic) APIURL(url g.String) *CreateForumTopic {
	if cf.opts.RequestOpts == nil {
		cf.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cf.opts.RequestOpts.APIURL = url.Std()

	return cf
}

// ChatID sets the target chat ID for this request.
func (cf *CreateForumTopic) ChatID(id int64) *CreateForumTopic {
	cf.chatID = g.Some(id)
	return cf
}

// Send executes the CreateForumTopic request.
func (cf *CreateForumTopic) Send() g.Result[*gotgbot.ForumTopic] {
	chatID := cf.chatID.UnwrapOr(cf.ctx.EffectiveChat.Id)
	return g.ResultOf(cf.ctx.Bot.Raw().CreateForumTopic(chatID, cf.name.Std(), cf.opts))
}
