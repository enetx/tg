package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetGameHighScores represents a request to get high scores for a game.
type GetGameHighScores struct {
	ctx             *Context
	userID          int64
	opts            *gotgbot.GetGameHighScoresOpts
	chatID          g.Option[int64]
	messageID       g.Option[int64]
	inlineMessageID g.Option[g.String]
}

// UserID sets the user ID to get scores for.
func (gghs *GetGameHighScores) UserID(userID int64) *GetGameHighScores {
	gghs.userID = userID
	return gghs
}

// ChatID sets the chat ID where the game message is located.
func (gghs *GetGameHighScores) ChatID(chatID int64) *GetGameHighScores {
	gghs.chatID = g.Some(chatID)
	return gghs
}

// MessageID sets the message ID of the game message.
func (gghs *GetGameHighScores) MessageID(messageID int64) *GetGameHighScores {
	gghs.messageID = g.Some(messageID)
	return gghs
}

// InlineMessageID sets the inline message ID for inline games.
func (gghs *GetGameHighScores) InlineMessageID(inlineMessageID g.String) *GetGameHighScores {
	gghs.inlineMessageID = g.Some(inlineMessageID)
	return gghs
}

// Timeout sets a custom timeout for this request.
func (gghs *GetGameHighScores) Timeout(duration time.Duration) *GetGameHighScores {
	if gghs.opts.RequestOpts == nil {
		gghs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gghs.opts.RequestOpts.Timeout = duration

	return gghs
}

// APIURL sets a custom API URL for this request.
func (gghs *GetGameHighScores) APIURL(url g.String) *GetGameHighScores {
	if gghs.opts.RequestOpts == nil {
		gghs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gghs.opts.RequestOpts.APIURL = url.Std()

	return gghs
}

// Send gets the game high scores and returns the result.
func (gghs *GetGameHighScores) Send() g.Result[g.Slice[gotgbot.GameHighScore]] {
	gghs.opts.ChatId = gghs.chatID.UnwrapOr(gghs.ctx.EffectiveChat.Id)
	gghs.opts.MessageId = gghs.messageID.UnwrapOr(gghs.ctx.EffectiveMessage.MessageId)

	if gghs.inlineMessageID.IsSome() {
		gghs.opts.InlineMessageId = gghs.inlineMessageID.Some().Std()
	}

	scores, err := gghs.ctx.Bot.Raw().GetGameHighScores(gghs.userID, gghs.opts)
	return g.ResultOf[g.Slice[gotgbot.GameHighScore]](scores, err)
}
