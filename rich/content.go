package rich

import "github.com/PaulSonOfLars/gotgbot/v2"

// Content represents a rich message used as the content of an inline query result.
type Content struct {
	content *gotgbot.InputRichMessageContent
}

// NewContent creates a new Content builder wrapping the given rich message.
func NewContent(message *Message) *Content {
	content := new(gotgbot.InputRichMessageContent)
	if message != nil {
		content.RichMessage = message.Std()
	}

	return &Content{content: content}
}

// Message sets the rich message to be sent.
func (c *Content) Message(message *Message) *Content {
	if message != nil {
		c.content.RichMessage = message.Std()
	}

	return c
}

// Build creates the gotgbot.InputMessageContent.
func (c *Content) Build() gotgbot.InputMessageContent {
	return *c.content
}
