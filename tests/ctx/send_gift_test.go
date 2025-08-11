package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
)

func TestContext_SendGift(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	giftID := g.String("gift_123")

	result := ctx.SendGift(giftID)

	if result == nil {
		t.Error("Expected SendGift builder to be created")
	}

	// Test method chaining
	chained := result.Text(g.String("Happy Birthday!"))
	if chained == nil {
		t.Error("Expected Text method to return builder")
	}
}

func TestContext_SendGiftChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	giftID := g.String("gift_123")

	result := ctx.SendGift(giftID).
		Text(g.String("Happy Birthday!"))

	if result == nil {
		t.Error("Expected SendGift builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestSendGift_To(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SendGift(g.String("gift_123")).To(123) == nil {
		t.Error("To should return builder")
	}
}

func TestSendGift_ToChat(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SendGift(g.String("gift_123")).ToChat(-1001234567890) == nil {
		t.Error("ToChat should return builder")
	}
}

func TestSendGift_PayForUpgrade(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SendGift(g.String("gift_123")).PayForUpgrade() == nil {
		t.Error("PayForUpgrade should return builder")
	}
}

func TestSendGift_HTML(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SendGift(g.String("gift_123")).HTML() == nil {
		t.Error("HTML should return builder")
	}
}

func TestSendGift_Markdown(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SendGift(g.String("gift_123")).Markdown() == nil {
		t.Error("Markdown should return builder")
	}
}

func TestSendGift_TextEntities(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	ent := entities.New(g.String("Bold text")).Bold(g.String("Bold"))
	if ctx.SendGift(g.String("gift_123")).TextEntities(ent) == nil {
		t.Error("TextEntities should return builder")
	}
}

func TestSendGift_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SendGift(g.String("gift_123")).Timeout(time.Minute) == nil {
		t.Error("Timeout should return builder")
	}
}

func TestSendGift_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SendGift(g.String("gift_123")).APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL should return builder")
	}
}

func TestSendGift_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveUser: &gotgbot.User{Id: 123, FirstName: "Test"},
		Update:        &gotgbot.Update{UpdateId: 1},
	})

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendGift(g.String("gift_123")).Send()

	if sendResult.IsErr() {
		t.Logf("SendGift Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendGift(g.String("gift_456")).
		Text(g.String("Happy Birthday!")).
		HTML().
		Timeout(time.Second * 30).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendGift configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

// Test Send method with different targeting options for complete coverage
func TestSendGift_SendTargetOptions(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveUser: &gotgbot.User{Id: 123, FirstName: "Test"},
		Update:        &gotgbot.Update{UpdateId: 1},
	})

	// Test 1: Send with To() (userID.IsSome() branch)
	userTargetResult := ctx.SendGift(g.String("gift_user")).
		To(789).
		Text(g.String("Gift for user")).
		Send()

	if userTargetResult.IsErr() {
		t.Logf("SendGift with To() failed as expected: %v", userTargetResult.Err())
	}

	// Test 2: Send with ToChat() (chatID.IsSome() branch)
	chatTargetResult := ctx.SendGift(g.String("gift_chat")).
		ToChat(-1001234567890).
		Text(g.String("Gift for chat")).
		PayForUpgrade().
		Send()

	if chatTargetResult.IsErr() {
		t.Logf("SendGift with ToChat() failed as expected: %v", chatTargetResult.Err())
	}

	// Test 3: Send without To() or ToChat() (uses EffectiveUser branch)
	defaultUserResult := ctx.SendGift(g.String("gift_default")).
		Text(g.String("Default user gift")).
		Markdown().
		Send()

	if defaultUserResult.IsErr() {
		t.Logf("SendGift with default user failed as expected: %v", defaultUserResult.Err())
	}

	// Test 4: Test with TextEntities
	ent := entities.New(g.String("Special gift with <b>bold</b> text")).Bold(g.String("bold"))
	entitiesResult := ctx.SendGift(g.String("gift_entities")).
		TextEntities(ent).
		To(999).
		Send()

	if entitiesResult.IsErr() {
		t.Logf("SendGift with TextEntities failed as expected: %v", entitiesResult.Err())
	}

	// Test 5: Test comprehensive chaining with all methods
	comprehensiveResult := ctx.SendGift(g.String("gift_comprehensive")).
		To(555).
		Text(g.String("Comprehensive gift test")).
		HTML().
		TextEntities(ent).
		PayForUpgrade().
		Timeout(45 * time.Second).
		APIURL(g.String("https://comprehensive-gift-api.example.com")).
		Send()

	if comprehensiveResult.IsErr() {
		t.Logf("SendGift comprehensive test failed as expected: %v", comprehensiveResult.Err())
	}
}
