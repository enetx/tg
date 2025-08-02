package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
)

func TestContact(t *testing.T) {
	contact := input.Contact(testPhoneNumber, testFirstName)
	if contact == nil {
		t.Error("Expected MessageContact to be created")
	}
	if !assertMessageContent(contact) {
		t.Error("MessageContact should implement MessageContent correctly")
	}
}

func TestContact_LastName(t *testing.T) {
	contact := input.Contact(testPhoneNumber, testFirstName)
	result := contact.LastName(testLastName)
	if result == nil {
		t.Error("Expected LastName method to return MessageContact")
	}
	if result != contact {
		t.Error("Expected LastName to return same MessageContact instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputContactMessageContent); ok {
		if v.LastName != testLastName.Std() {
			t.Error("Expected LastName to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputContactMessageContent")
	}
}

func TestContact_Vcard(t *testing.T) {
	contact := input.Contact(testPhoneNumber, testFirstName)
	vcard := g.String("BEGIN:VCARD\nVERSION:3.0\nFN:John Doe\nEND:VCARD")
	result := contact.Vcard(vcard)
	if result == nil {
		t.Error("Expected Vcard method to return MessageContact")
	}
	if result != contact {
		t.Error("Expected Vcard to return same MessageContact instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputContactMessageContent); ok {
		if v.Vcard != vcard.Std() {
			t.Error("Expected Vcard to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputContactMessageContent")
	}
}

func TestContact_Build(t *testing.T) {
	contact := input.Contact(testPhoneNumber, testFirstName)
	built := contact.Build()

	if v, ok := built.(gotgbot.InputContactMessageContent); ok {
		if v.PhoneNumber != testPhoneNumber.Std() {
			t.Errorf("Expected PhoneNumber to be %s, got %s", testPhoneNumber.Std(), v.PhoneNumber)
		}
		if v.FirstName != testFirstName.Std() {
			t.Errorf("Expected FirstName to be %s, got %s", testFirstName.Std(), v.FirstName)
		}
	} else {
		t.Error("Expected result to be InputContactMessageContent")
	}
}

func TestContact_MethodChaining(t *testing.T) {
	vcard := g.String("BEGIN:VCARD\nVERSION:3.0\nFN:John Doe\nEND:VCARD")
	result := input.Contact(testPhoneNumber, testFirstName).
		LastName(testLastName).
		Vcard(vcard)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Contact to build correctly")
	}

	if _, ok := built.(gotgbot.InputContactMessageContent); !ok {
		t.Error("Expected result to be InputContactMessageContent")
	}

	if !assertMessageContent(result) {
		t.Error("Expected result to implement MessageContent interface")
	}
}
