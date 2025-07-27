package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/preview"
)

// MessageText represents an input text message content builder.
type MessageText struct {
	input *gotgbot.InputTextMessageContent
}

// Text creates a new MessageText builder with the required fields.
func Text(messageText String) *MessageText {
	return &MessageText{
		input: &gotgbot.InputTextMessageContent{
			MessageText: messageText.Std(),
		},
	}
}

// HTML sets parse mode to HTML.
func (mt *MessageText) HTML() *MessageText {
	mt.input.ParseMode = "HTML"
	return mt
}

// Markdown sets parse mode to MarkdownV2.
func (mt *MessageText) Markdown() *MessageText {
	mt.input.ParseMode = "MarkdownV2"
	return mt
}

// Entities sets the message entities for the text.
func (mt *MessageText) Entities(e entities.Entities) *MessageText {
	mt.input.Entities = e.Std()
	return mt
}

// Preview sets custom link preview options.
func (mt *MessageText) Preview(p *preview.Preview) *MessageText {
	mt.input.LinkPreviewOptions = p.Std()
	return mt
}

// Build creates the gotgbot.InputTextMessageContent.
func (mt *MessageText) Build() gotgbot.InputMessageContent {
	return *mt.input
}
