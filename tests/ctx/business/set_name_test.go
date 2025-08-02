package business_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestSetName(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)

	firstName := g.String("John")
	result := account.SetName(firstName)

	if result == nil {
		t.Error("Expected SetName builder to be created")
	}

	// Test method chaining
	withLastName := result.LastName(g.String("Doe"))
	if withLastName == nil {
		t.Error("Expected LastName method to return builder")
	}
}
