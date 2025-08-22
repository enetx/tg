package reply

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
)

// Parameters provides a builder for creating ReplyParameters.
type Parameters struct {
	params *gotgbot.ReplyParameters
}

// New creates a new Parameters builder for the specified message ID.
func New(messageID int64) *Parameters {
	return &Parameters{
		params: &gotgbot.ReplyParameters{
			MessageId: messageID,
		},
	}
}

// ChatID sets the chat ID if replying to a message from a different chat.
func (p *Parameters) ChatID(chatID int64) *Parameters {
	p.params.ChatId = chatID
	return p
}

// AllowSendingWithoutReply allows sending even if the replied message is not found.
func (p *Parameters) AllowSendingWithoutReply() *Parameters {
	p.params.AllowSendingWithoutReply = true
	return p
}

// Quote sets the quoted part of the message to be replied to.
func (p *Parameters) Quote(quote g.String) *Parameters {
	p.params.Quote = quote.Std()
	return p
}

// QuoteHTML sets the quote parse mode to HTML.
func (p *Parameters) QuoteHTML() *Parameters {
	p.params.QuoteParseMode = "HTML"
	return p
}

// QuoteMarkdown sets the quote parse mode to MarkdownV2.
func (p *Parameters) QuoteMarkdown() *Parameters {
	p.params.QuoteParseMode = "MarkdownV2"
	return p
}

// QuoteEntities sets custom entities for the quoted text.
func (p *Parameters) QuoteEntities(e *entities.Entities) *Parameters {
	p.params.QuoteEntities = e.Std()
	return p
}

// QuotePosition sets the position of the quote in the original message.
func (p *Parameters) QuotePosition(position int64) *Parameters {
	p.params.QuotePosition = position
	return p
}

// ChecklistTask sets the specific checklist task ID to reply to.
func (p *Parameters) ChecklistTask(taskID int64) *Parameters {
	p.params.ChecklistTaskId = taskID
	return p
}

// Build returns the ReplyParameters for use with the Telegram API.
func (p *Parameters) Build() *gotgbot.ReplyParameters {
	return p.params
}

// Std returns the ReplyParameters for use with the Telegram API.
func (p *Parameters) Std() *gotgbot.ReplyParameters {
	return p.params
}
