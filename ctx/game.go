package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/types/effects"
)

type SendGame struct {
	ctx           *Context
	gameShortName String
	opts          *gotgbot.SendGameOpts
	chatID        Option[int64]
	after         Option[time.Duration]
	deleteAfter   Option[time.Duration]
}

// After schedules the game to be sent after the specified duration.
func (sg *SendGame) After(duration time.Duration) *SendGame {
	sg.after = Some(duration)
	return sg
}

// DeleteAfter schedules the game message to be deleted after the specified duration.
func (sg *SendGame) DeleteAfter(duration time.Duration) *SendGame {
	sg.deleteAfter = Some(duration)
	return sg
}

// Silent disables notification for the game message.
func (sg *SendGame) Silent() *SendGame {
	sg.opts.DisableNotification = true
	return sg
}

// Protect enables content protection for the game message.
func (sg *SendGame) Protect() *SendGame {
	sg.opts.ProtectContent = true
	return sg
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (sg *SendGame) AllowPaidBroadcast() *SendGame {
	sg.opts.AllowPaidBroadcast = true
	return sg
}

// Thread sets the message thread ID for the game message.
func (sg *SendGame) Thread(id int64) *SendGame {
	sg.opts.MessageThreadId = id
	return sg
}

// Effect sets a message effect for the game message.
func (sg *SendGame) Effect(effect effects.EffectType) *SendGame {
	sg.opts.MessageEffectId = effect.String()
	return sg
}

// ReplyTo sets the message ID to reply to.
func (sg *SendGame) ReplyTo(messageID int64) *SendGame {
	sg.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return sg
}

// Markup sets the reply markup keyboard for the game message.
func (sg *SendGame) Markup(kb keyboard.KeyboardBuilder) *SendGame {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		sg.opts.ReplyMarkup = markup
	}

	return sg
}

// Business sets the business connection ID for the game message.
func (sg *SendGame) Business(id String) *SendGame {
	sg.opts.BusinessConnectionId = id.Std()
	return sg
}

// To sets the target chat ID for the game message.
func (sg *SendGame) To(chatID int64) *SendGame {
	sg.chatID = Some(chatID)
	return sg
}

// Timeout sets a custom timeout for this request.
func (sg *SendGame) Timeout(duration time.Duration) *SendGame {
	if sg.opts.RequestOpts == nil {
		sg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sg.opts.RequestOpts.Timeout = duration

	return sg
}

// APIURL sets a custom API URL for this request.
func (sg *SendGame) APIURL(url String) *SendGame {
	if sg.opts.RequestOpts == nil {
		sg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sg.opts.RequestOpts.APIURL = url.Std()

	return sg
}

// Send sends the game message to Telegram and returns the result.
func (sg *SendGame) Send() Result[*gotgbot.Message] {
	return sg.ctx.timers(sg.after, sg.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := sg.chatID.UnwrapOr(sg.ctx.EffectiveChat.Id)
		return ResultOf(sg.ctx.Bot.Raw().SendGame(chatID, sg.gameShortName.Std(), sg.opts))
	})
}

// SetGameScore represents a request to set the score for a game.
type SetGameScore struct {
	ctx             *Context
	userID          int64
	score           int64
	opts            *gotgbot.SetGameScoreOpts
	chatID          Option[int64]
	messageID       Option[int64]
	inlineMessageID Option[String]
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
	sgs.chatID = Some(chatID)
	return sgs
}

// MessageID sets the message ID of the game message.
func (sgs *SetGameScore) MessageID(messageID int64) *SetGameScore {
	sgs.messageID = Some(messageID)
	return sgs
}

// InlineMessageID sets the inline message ID for inline games.
func (sgs *SetGameScore) InlineMessageID(inlineMessageID String) *SetGameScore {
	sgs.inlineMessageID = Some(inlineMessageID)
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
func (sgs *SetGameScore) APIURL(url String) *SetGameScore {
	if sgs.opts.RequestOpts == nil {
		sgs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sgs.opts.RequestOpts.APIURL = url.Std()

	return sgs
}

// Send sets the game score and returns the result.
func (sgs *SetGameScore) Send() Result[*gotgbot.Message] {
	if sgs.score < 0 {
		return Err[*gotgbot.Message](Errorf("score cannot be negative: {}", sgs.score))
	}

	sgs.opts.ChatId = sgs.chatID.UnwrapOr(sgs.ctx.EffectiveChat.Id)
	sgs.opts.MessageId = sgs.messageID.UnwrapOr(sgs.ctx.EffectiveMessage.MessageId)

	if sgs.inlineMessageID.IsSome() {
		sgs.opts.InlineMessageId = sgs.inlineMessageID.Some().Std()
	}

	msg, _, err := sgs.ctx.Bot.Raw().SetGameScore(sgs.userID, sgs.score, sgs.opts)
	return ResultOf(msg, err)
}

// GetGameHighScores represents a request to get high scores for a game.
type GetGameHighScores struct {
	ctx             *Context
	userID          int64
	opts            *gotgbot.GetGameHighScoresOpts
	chatID          Option[int64]
	messageID       Option[int64]
	inlineMessageID Option[String]
}

// UserID sets the user ID to get scores for.
func (gghs *GetGameHighScores) UserID(userID int64) *GetGameHighScores {
	gghs.userID = userID
	return gghs
}

// ChatID sets the chat ID where the game message is located.
func (gghs *GetGameHighScores) ChatID(chatID int64) *GetGameHighScores {
	gghs.chatID = Some(chatID)
	return gghs
}

// MessageID sets the message ID of the game message.
func (gghs *GetGameHighScores) MessageID(messageID int64) *GetGameHighScores {
	gghs.messageID = Some(messageID)
	return gghs
}

// InlineMessageID sets the inline message ID for inline games.
func (gghs *GetGameHighScores) InlineMessageID(inlineMessageID String) *GetGameHighScores {
	gghs.inlineMessageID = Some(inlineMessageID)
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
func (gghs *GetGameHighScores) APIURL(url String) *GetGameHighScores {
	if gghs.opts.RequestOpts == nil {
		gghs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gghs.opts.RequestOpts.APIURL = url.Std()

	return gghs
}

// Send gets the game high scores and returns the result.
func (gghs *GetGameHighScores) Send() Result[Slice[gotgbot.GameHighScore]] {
	gghs.opts.ChatId = gghs.chatID.UnwrapOr(gghs.ctx.EffectiveChat.Id)
	gghs.opts.MessageId = gghs.messageID.UnwrapOr(gghs.ctx.EffectiveMessage.MessageId)

	if gghs.inlineMessageID.IsSome() {
		gghs.opts.InlineMessageId = gghs.inlineMessageID.Some().Std()
	}

	scores, err := gghs.ctx.Bot.Raw().GetGameHighScores(gghs.userID, gghs.opts)
	return ResultOf[Slice[gotgbot.GameHighScore]](scores, err)
}
