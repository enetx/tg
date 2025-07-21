package bot

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
	"github.com/enetx/tg/handlers"
)

// Bot represents a Telegram bot instance with all necessary components for handling updates,
// managing middleware, and interacting with the Telegram Bot API.
type Bot struct {
	token       String                  // Bot token for API authentication
	dispatcher  *ext.Dispatcher         // Event dispatcher for handling updates
	updater     *ext.Updater            // Updater for receiving updates
	middlewares Slice[handlers.Handler] // Global middleware stack
	On          *handlers.Handlers      // Event handlers for different update types
	raw         *gotgbot.Bot            // Raw gotgbot instance for direct API access
}

var _ core.BotAPI = (*Bot)(nil)

// Dispatcher returns the bot's event dispatcher for advanced usage.
func (b *Bot) Dispatcher() *ext.Dispatcher {
	return b.dispatcher
}

// Updater returns the bot's updater instance for managing update polling/webhook.
func (b *Bot) Updater() *ext.Updater {
	return b.updater
}

// Raw returns the underlying gotgbot.Bot instance for direct API access.
func (b *Bot) Raw() *gotgbot.Bot {
	return b.raw
}

// New creates a new BotBuilder instance with the provided token.
func New[T ~string](token T) *BotBuilder {
	return &BotBuilder{
		token: String(token),
		opts: &gotgbot.BotOpts{
			BotClient: &gotgbot.BaseBotClient{
				Client: http.Client{},
				DefaultRequestOpts: &gotgbot.RequestOpts{
					Timeout: 10 * time.Second,
				},
			},
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: 10 * time.Second,
			},
		},
	}
}

// Command registers a command handler for the specified command.
func (b *Bot) Command(cmd String, fn handlers.Handler) *handlers.Command {
	c := handlers.NewCommand(b, cmd, fn)
	c.Register()

	return c
}

// Polling returns a Polling instance for receiving updates via long polling.
func (b *Bot) Polling() *Polling {
	return &Polling{
		bot:  b,
		opts: &ext.PollingOpts{GetUpdatesOpts: new(gotgbot.GetUpdatesOpts)},
	}
}

// Webhook returns a Webhook instance for receiving updates via webhook.
func (b *Bot) Webhook() *Webhook {
	return &Webhook{
		bot: b,
		opt: new(gotgbot.SetWebhookOpts),
	}
}

// HandleWebhook processes a webhook update from raw JSON data.
func (b *Bot) HandleWebhook(data []byte) error {
	var update gotgbot.Update
	if err := json.Unmarshal(data, &update); err != nil {
		return Errorf("failed to unmarshal update: {}", err)
	}

	return b.dispatcher.ProcessUpdate(b.Raw(), &update, nil)
}

// Use adds a global middleware to the bot.
func (b *Bot) Use(middleware handlers.Handler) *Bot {
	b.middlewares.Push(middleware)
	return b
}

// Middlewares returns the current global middleware stack.
func (b *Bot) Middlewares() Slice[handlers.Handler] {
	return b.middlewares
}
