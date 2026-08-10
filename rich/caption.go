package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/richtext"
)

// Caption provides a builder for captions attached to media and container blocks.
type Caption struct {
	caption *gotgbot.RichBlockCaption
}

// NewCaption creates a new Caption builder with the required text.
func NewCaption(text *richtext.Text) *Caption {
	return &Caption{
		caption: &gotgbot.RichBlockCaption{
			Text: text.Std(),
		},
	}
}

// Text sets the caption text.
func (c *Caption) Text(text *richtext.Text) *Caption {
	c.caption.Text = text.Std()
	return c
}

// Credit sets the caption credit, usually the author or the source.
func (c *Caption) Credit(credit *richtext.Text) *Caption {
	c.caption.Credit = credit.Std()
	return c
}

// Std returns the underlying gotgbot.RichBlockCaption.
func (c *Caption) Std() *gotgbot.RichBlockCaption {
	if c == nil {
		return nil
	}

	return c.caption
}
