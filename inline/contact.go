package inline

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

// Contact represents an inline query result contact builder.
type Contact struct {
	inline *gotgbot.InlineQueryResultContact
}

// NewContact creates a new Contact builder with the required fields.
func NewContact(id, phoneNumber, firstName g.String) *Contact {
	return &Contact{
		inline: &gotgbot.InlineQueryResultContact{
			Id:          id.Std(),
			PhoneNumber: phoneNumber.Std(),
			FirstName:   firstName.Std(),
		},
	}
}

// LastName sets the contact's last name.
func (c *Contact) LastName(lastName g.String) *Contact {
	c.inline.LastName = lastName.Std()
	return c
}

// VCard sets additional data about the contact in the form of a vCard.
func (c *Contact) VCard(vcard g.String) *Contact {
	c.inline.Vcard = vcard.Std()
	return c
}

// Markup sets the inline keyboard attached to the message.
func (c *Contact) Markup(kb keyboard.Keyboard) *Contact {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			c.inline.ReplyMarkup = &ikm
		}
	}

	return c
}

// ThumbnailURL sets the URL of the thumbnail for the result.
func (c *Contact) ThumbnailURL(url g.String) *Contact {
	c.inline.ThumbnailUrl = url.Std()
	return c
}

// ThumbnailSize sets the thumbnail width and height.
func (c *Contact) ThumbnailSize(width, height int64) *Contact {
	c.inline.ThumbnailWidth = width
	c.inline.ThumbnailHeight = height

	return c
}

// InputMessageContent sets the content of the message to be sent instead of the contact.
func (c *Contact) InputMessageContent(message input.MessageContent) *Contact {
	c.inline.InputMessageContent = message.Build()
	return c
}

// Build creates the gotgbot.InlineQueryResultContact.
func (c *Contact) Build() gotgbot.InlineQueryResult {
	return *c.inline
}
