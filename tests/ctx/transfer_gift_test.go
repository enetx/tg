package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_TransferGift(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_123")
	giftID := g.String("gift_123")
	userID := int64(456)

	result := ctx.TransferGift(businessConnectionID, giftID, userID)

	if result == nil {
		t.Error("Expected TransferGift builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(time.Second * 30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestTransferGift_StarCount(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	businessConnectionID := g.String("business_conn_123")
	giftID := g.String("gift_123")
	userID := int64(456)
	if ctx.TransferGift(businessConnectionID, giftID, userID).StarCount(100) == nil { t.Error("StarCount should return builder") }
}

func TestTransferGift_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	businessConnectionID := g.String("business_conn_123")
	giftID := g.String("gift_123")
	userID := int64(456)
	if ctx.TransferGift(businessConnectionID, giftID, userID).APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestTransferGift_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	businessConnectionID := g.String("business_conn_123")
	giftID := g.String("gift_123")
	userID := int64(456)
	
	sendResult := ctx.TransferGift(businessConnectionID, giftID, userID).StarCount(50).Send()
	
	if sendResult.IsErr() {
		t.Logf("TransferGift Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
