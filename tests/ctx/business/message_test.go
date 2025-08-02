package business_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestMessage_Read(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)
	message := account.Message()

	chatID := int64(456)
	messageID := int64(123)

	result := message.Read(chatID, messageID)

	if result == nil {
		t.Error("Expected ReadMessage builder to be created")
	}
}

func TestMessage_Delete(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)
	message := account.Message()

	messageIDs := g.Slice[int64]{}
	messageIDs.Push(123)
	result := message.Delete(messageIDs)

	if result == nil {
		t.Error("Expected DeleteMessage builder to be created")
	}
}
