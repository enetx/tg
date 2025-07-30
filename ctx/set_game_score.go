package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetGameScore represents a request to set the score for a game.
type SetGameScore struct {
	ctx             *Context
	userID          int64
	score           int64
	opts            *gotgbot.SetGameScoreOpts
	chatID          g.Option[int64]
	messageID       g.Option[int64]
	inlineMessageID g.Option[g.String]
}

// UserID sets the user ID for the score.
func (sgs *SetGameScore) UserID(userID int64) *SetGameScore {
	sgs.userID = userID
	return sgs
}

// Score sets the new score value.
func (sgs *SetGameScore) Score(score int64) *SetGameScore {
	sgs.score = score
	return sgs
}

// Force forces the score update even if it's lower than current.
func (sgs *SetGameScore) Force() *SetGameScore {
	sgs.opts.Force = true
	return sgs
}

// DisableEditMessage prevents editing the game message.
func (sgs *SetGameScore) DisableEditMessage() *SetGameScore {
	sgs.opts.DisableEditMessage = true
	return sgs
}

// ChatID sets the chat ID where the game message is located.
func (sgs *SetGameScore) ChatID(chatID int64) *SetGameScore {
	sgs.chatID = g.Some(chatID)
	return sgs
}

// MessageID sets the message ID of the game message.
func (sgs *SetGameScore) MessageID(messageID int64) *SetGameScore {
	sgs.messageID = g.Some(messageID)
	return sgs
}

// InlineMessageID sets the inline message ID for inline games.
func (sgs *SetGameScore) InlineMessageID(inlineMessageID g.String) *SetGameScore {
	sgs.inlineMessageID = g.Some(inlineMessageID)
	return sgs
}

// Timeout sets a custom timeout for this request.
func (sgs *SetGameScore) Timeout(duration time.Duration) *SetGameScore {
	if sgs.opts.RequestOpts == nil {
		sgs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sgs.opts.RequestOpts.Timeout = duration

	return sgs
}

// APIURL sets a custom API URL for this request.
func (sgs *SetGameScore) APIURL(url g.String) *SetGameScore {
	if sgs.opts.RequestOpts == nil {
		sgs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sgs.opts.RequestOpts.APIURL = url.Std()

	return sgs
}

// Send sets the game score and returns the result.
func (sgs *SetGameScore) Send() g.Result[*gotgbot.Message] {
	if sgs.score < 0 {
		return g.Err[*gotgbot.Message](g.Errorf("score cannot be negative: {}", sgs.score))
	}

	sgs.opts.ChatId = sgs.chatID.UnwrapOr(sgs.ctx.EffectiveChat.Id)
	sgs.opts.MessageId = sgs.messageID.UnwrapOr(sgs.ctx.EffectiveMessage.MessageId)

	if sgs.inlineMessageID.IsSome() {
		sgs.opts.InlineMessageId = sgs.inlineMessageID.Some().Std()
	}

	msg, _, err := sgs.ctx.Bot.Raw().SetGameScore(sgs.userID, sgs.score, sgs.opts)
	return g.ResultOf(msg, err)
}
