package content

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// ContactBuilder represents an input contact message content builder.
type ContactBuilder struct {
	input *gotgbot.InputContactMessageContent
}

// NewContactBuilder creates a new ContactBuilder builder.
func NewContactBuilder(phoneNumber, firstName String) *ContactBuilder {
	return &ContactBuilder{
		input: &gotgbot.InputContactMessageContent{
			PhoneNumber: phoneNumber.Std(),
			FirstName:   firstName.Std(),
		},
	}
}

// LastName sets the contact's last name.
func (c *ContactBuilder) LastName(lastName String) *ContactBuilder {
	c.input.LastName = lastName.Std()
	return c
}

// VCard sets additional data about the contact in the form of a vCard.
func (c *ContactBuilder) VCard(vcard String) *ContactBuilder {
	c.input.Vcard = vcard.Std()
	return c
}

// Build creates the gotgbot.InputContactMessageContent.
func (c *ContactBuilder) Build() gotgbot.InputMessageContent {
	return *c.input
}
