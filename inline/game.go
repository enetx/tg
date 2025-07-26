package inline

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

// Game represents an inline query result game builder.
type Game struct {
	inline *gotgbot.InlineQueryResultGame
}

// NewGame creates a new Game builder with the required fields.
func NewGame(id, gameShortName g.String) *Game {
	return &Game{
		inline: &gotgbot.InlineQueryResultGame{
			Id:            id.Std(),
			GameShortName: gameShortName.Std(),
		},
	}
}

// Markup sets the inline keyboard attached to the message.
func (gm *Game) Markup(kb keyboard.Keyboard) *Game {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			gm.inline.ReplyMarkup = &ikm
		}
	}

	return gm
}

// Build creates the gotgbot.InlineQueryResultGame.
func (gm *Game) Build() gotgbot.InlineQueryResult {
	return *gm.inline
}
