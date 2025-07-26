package content

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/preview"
)

// TextBuilder represents an input text message content builder.
type TextBuilder struct {
	input *gotgbot.InputTextMessageContent
}

// NewTextContent creates a new TextBuilder.
func NewTextContent(text String) *TextBuilder {
	return &TextBuilder{
		input: &gotgbot.InputTextMessageContent{MessageText: text.Std()},
	}
}

// HTML sets parse mode to HTML.
func (t *TextBuilder) HTML() *TextBuilder {
	t.input.ParseMode = "HTML"
	return t
}

// Markdown sets parse mode to Markdown2.
func (t *TextBuilder) Markdown() *TextBuilder {
	t.input.ParseMode = "MarkdownV2"
	return t
}

// Entities sets the message entities for the text.
func (t *TextBuilder) Entities(e entities.Entities) *TextBuilder {
	t.input.Entities = e.Std()
	return t
}

// Preview sets custom link preview options.
func (t *TextBuilder) Preview(p *preview.Preview) *TextBuilder {
	t.input.LinkPreviewOptions = p.Std()
	return t
}

// Build creates the gotgbot.InputTextMessageContent.
func (t *TextBuilder) Build() gotgbot.InputMessageContent {
	return *t.input
}
