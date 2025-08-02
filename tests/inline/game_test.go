package inline_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/inline"
)

func TestNewGame(t *testing.T) {
	game := inline.NewGame(testID, g.String("mygame"))

	if game == nil {
		t.Error("Expected Game to be created")
	}

	built := game.Build()
	if built == nil {
		t.Error("Expected Game to build correctly")
	}

	if result, ok := built.(gotgbot.InlineQueryResultGame); ok {
		if result.GetType() != "game" {
			t.Error("Expected type to be 'game'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultGame")
	}
}

func TestGame_Markup(t *testing.T) {
	game := inline.NewGame(testID, g.String("mygame"))
	result := game.Markup(createTestKeyboard())

	if result == nil {
		t.Error("Expected Markup method to return Game")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultGame); ok {
		if v.ReplyMarkup == nil {
			t.Error("Expected ReplyMarkup to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultGame")
	}
}

func TestGame_MethodChaining(t *testing.T) {
	result := inline.NewGame(testID, g.String("mygame")).
		Markup(createTestKeyboard())

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Game to build correctly")
	}

	if _, ok := built.(gotgbot.InlineQueryResultGame); !ok {
		t.Error("Expected result to be InlineQueryResultGame")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
