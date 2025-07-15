package core

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type BotAPI interface {
	Dispatcher() *ext.Dispatcher
	Updater() *ext.Updater
	Raw() *gotgbot.Bot
}
