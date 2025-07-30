package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetMessageReaction represents a request to set reactions on a message.
type SetMessageReaction struct {
	ctx       *Context
	messageID int64
	reactions g.Slice[gotgbot.ReactionType]
	opts      *gotgbot.SetMessageReactionOpts
	chatID    g.Option[int64]
}

// ChatID sets the chat ID where the message is located.
func (smr *SetMessageReaction) ChatID(chatID int64) *SetMessageReaction {
	smr.chatID = g.Some(chatID)
	return smr
}

// Reaction adds a reaction to the message. Can be called multiple times.
func (smr *SetMessageReaction) Reaction(emoji g.String) *SetMessageReaction {
	reaction := gotgbot.ReactionTypeEmoji{Emoji: emoji.Std()}
	smr.reactions.Push(reaction)

	return smr
}

// CustomEmoji adds a custom emoji reaction to the message.
func (smr *SetMessageReaction) CustomEmoji(customEmojiID g.String) *SetMessageReaction {
	reaction := gotgbot.ReactionTypeCustomEmoji{CustomEmojiId: customEmojiID.Std()}
	smr.reactions.Push(reaction)

	return smr
}

// Big makes the reaction animation bigger.
func (smr *SetMessageReaction) Big() *SetMessageReaction {
	smr.opts.IsBig = true
	return smr
}

// RemoveReactions removes all reactions from the message.
func (smr *SetMessageReaction) RemoveReactions() *SetMessageReaction {
	smr.reactions = []gotgbot.ReactionType{}
	return smr
}

// Timeout sets a custom timeout for this request.
func (smr *SetMessageReaction) Timeout(duration time.Duration) *SetMessageReaction {
	if smr.opts.RequestOpts == nil {
		smr.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smr.opts.RequestOpts.Timeout = duration

	return smr
}

// APIURL sets a custom API URL for this request.
func (smr *SetMessageReaction) APIURL(url g.String) *SetMessageReaction {
	if smr.opts.RequestOpts == nil {
		smr.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smr.opts.RequestOpts.APIURL = url.Std()

	return smr
}

// Send sets the message reactions and returns the result.
func (smr *SetMessageReaction) Send() g.Result[bool] {
	chatID := smr.chatID.UnwrapOr(smr.ctx.EffectiveChat.Id)
	smr.opts.Reaction = smr.reactions

	return g.ResultOf(smr.ctx.Bot.Raw().SetMessageReaction(chatID, smr.messageID, smr.opts))
}
