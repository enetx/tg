package business

import "github.com/PaulSonOfLars/gotgbot/v2"

// Bot defines the minimal interface required to perform business account operations.
type Bot interface {
	Raw() *gotgbot.Bot
}
