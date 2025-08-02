package business_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestNewAccount(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")

	account := business.NewAccount(bot, connectionID)

	if account == nil {
		t.Error("Expected Account to be created")
	}
}

func TestAccount_SetName(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)

	firstName := g.String("John")
	result := account.SetName(firstName)

	if result == nil {
		t.Error("Expected SetName builder to be created")
	}
}

func TestAccount_SetUsername(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)

	username := g.String("john_business")
	result := account.SetUsername(username)

	if result == nil {
		t.Error("Expected SetUsername builder to be created")
	}
}

func TestAccount_SetBio(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)

	bio := g.String("Business owner")
	result := account.SetBio(bio)

	if result == nil {
		t.Error("Expected SetBio builder to be created")
	}
}

func TestAccount_SetPhoto(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)

	photo := g.String("photo.jpg")
	result := account.SetPhoto(photo)

	if result == nil {
		t.Error("Expected SetPhoto builder to be created")
	}
}

func TestAccount_Balance(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)

	result := account.Balance()

	if result == nil {
		t.Error("Expected Balance handler to be created")
	}
}

func TestAccount_Message(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)

	result := account.Message()

	if result == nil {
		t.Error("Expected Message handler to be created")
	}
}
