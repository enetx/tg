package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/types/passport"
)

func TestContext_SetPassportDataErrors(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)

	result := ctx.SetPassportDataErrors(userID)

	if result == nil {
		t.Error("Expected SetPassportDataErrors builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(time.Second * 30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestSetPassportDataErrors_Errors(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)

	error1 := passport.NewDataFieldError(passport.Passport, "first_name", g.String("invalid_value"), g.String("Invalid first name"))
	error2 := passport.NewFrontSideError(passport.Passport, g.String("hash123"), g.String("blurry_image"))

	if ctx.SetPassportDataErrors(userID).Errors(error1, error2) == nil {
		t.Error("Errors should return builder")
	}
}

func TestSetPassportDataErrors_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	if ctx.SetPassportDataErrors(userID).APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL should return builder")
	}
}

func TestSetPassportDataErrors_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)

	error1 := passport.NewDataFieldError(passport.Passport, "first_name", g.String("invalid_value"), g.String("Invalid first name"))
	sendResult := ctx.SetPassportDataErrors(userID).Errors(error1).Send()

	if sendResult.IsErr() {
		t.Logf("SetPassportDataErrors Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
