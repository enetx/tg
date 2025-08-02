package business_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestBalance_GetStarBalance(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	result := balance.GetStarBalance()

	if result == nil {
		t.Error("Expected GetStarBalance builder to be created")
	}
}

func TestBalance_TransferStars(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	amount := int64(100)
	result := balance.TransferStars(amount)

	if result == nil {
		t.Error("Expected TransferStars builder to be created")
	}
}
