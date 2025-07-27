package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// MessageContact represents an input contact message content builder.
type MessageContact struct {
	input *gotgbot.InputContactMessageContent
}

// Contact creates a new MessageContact builder with the required fields.
func Contact(phoneNumber, firstName String) *MessageContact {
	return &MessageContact{
		input: &gotgbot.InputContactMessageContent{
			PhoneNumber: phoneNumber.Std(),
			FirstName:   firstName.Std(),
		},
	}
}

// LastName sets the contact's last name.
func (mc *MessageContact) LastName(lastName String) *MessageContact {
	mc.input.LastName = lastName.Std()
	return mc
}

// Vcard sets the additional data about the contact in the form of a vCard.
func (mc *MessageContact) Vcard(vcard String) *MessageContact {
	mc.input.Vcard = vcard.Std()
	return mc
}

// Build creates the gotgbot.InputContactMessageContent.
func (mc *MessageContact) Build() gotgbot.InputMessageContent {
	return *mc.input
}
