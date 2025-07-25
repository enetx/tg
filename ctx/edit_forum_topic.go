package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/g/ref"
)

// EditForumTopic represents a request to edit a forum topic.
type EditForumTopic struct {
	ctx             *Context
	messageThreadID int64
	opts            *gotgbot.EditForumTopicOpts
	chatID          Option[int64]
}

// Name sets the new name of the topic.
func (eft *EditForumTopic) Name(name String) *EditForumTopic {
	eft.opts.Name = name.Std()
	return eft
}

// IconCustomEmojiID sets the new custom emoji identifier.
func (eft *EditForumTopic) IconCustomEmojiID(emojiID String) *EditForumTopic {
	eft.opts.IconCustomEmojiId = ref.Of(emojiID.Std())
	return eft
}

// Timeout sets a custom timeout for this request.
func (eft *EditForumTopic) Timeout(duration time.Duration) *EditForumTopic {
	if eft.opts.RequestOpts == nil {
		eft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	eft.opts.RequestOpts.Timeout = duration

	return eft
}

// APIURL sets a custom API URL for this request.
func (eft *EditForumTopic) APIURL(url String) *EditForumTopic {
	if eft.opts.RequestOpts == nil {
		eft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	eft.opts.RequestOpts.APIURL = url.Std()

	return eft
}

// ChatID sets the target chat ID for this request.
func (eft *EditForumTopic) ChatID(id int64) *EditForumTopic {
	eft.chatID = Some(id)
	return eft
}

// Send executes the EditForumTopic request.
func (eft *EditForumTopic) Send() Result[bool] {
	chatID := eft.chatID.UnwrapOr(eft.ctx.EffectiveChat.Id)
	return ResultOf(eft.ctx.Bot.Raw().EditForumTopic(chatID, eft.messageThreadID, eft.opts))
}
