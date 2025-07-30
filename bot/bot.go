package bot

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/core"
	"github.com/enetx/tg/handlers"
)

// Bot represents a Telegram bot instance with all necessary components for handling updates,
// managing middleware, and interacting with the Telegram Bot API.
type Bot struct {
	token       g.String                  // Bot token for API authentication
	dispatcher  *ext.Dispatcher           // Event dispatcher for handling updates
	updater     *ext.Updater              // Updater for receiving updates
	middlewares g.Slice[handlers.Handler] // Global middleware stack
	On          *handlers.Handlers        // Event handlers for different update types
	raw         *gotgbot.Bot              // Raw gotgbot instance for direct API access
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
		token: g.String(token),
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
func (b *Bot) Command(cmd g.String, fn handlers.Handler) *handlers.Command {
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
func (b *Bot) Webhook() *SetWebhook {
	return &SetWebhook{
		bot: b,
		opt: new(gotgbot.SetWebhookOpts),
	}
}

// HandleWebhook processes a webhook update from raw JSON data.
func (b *Bot) HandleWebhook(data []byte) error {
	var update gotgbot.Update
	if err := json.Unmarshal(data, &update); err != nil {
		return g.Errorf("failed to unmarshal update: {}", err)
	}

	return b.dispatcher.ProcessUpdate(b.Raw(), &update, nil)
}

// Use adds a global middleware to the bot.
func (b *Bot) Use(middleware handlers.Handler) *Bot {
	b.middlewares.Push(middleware)
	return b
}

// Middlewares returns the current global middleware stack.
func (b *Bot) Middlewares() g.Slice[handlers.Handler] {
	return b.middlewares
}

// GetMyDescription creates a new GetMyDescription request to get the bot's description.
func (b *Bot) GetMyDescription() *GetMyDescription {
	return &GetMyDescription{
		bot:  b,
		opts: new(gotgbot.GetMyDescriptionOpts),
	}
}

// GetMyShortDescription creates a new GetMyShortDescription request to get the bot's short description.
func (b *Bot) GetMyShortDescription() *GetMyShortDescription {
	return &GetMyShortDescription{
		bot:  b,
		opts: new(gotgbot.GetMyShortDescriptionOpts),
	}
}

// SetMyDescription creates a new SetMyDescription request to set the bot's description.
func (b *Bot) SetMyDescription() *SetMyDescription {
	return &SetMyDescription{
		bot:  b,
		opts: new(gotgbot.SetMyDescriptionOpts),
	}
}

// SetMyShortDescription creates a new SetMyShortDescription request to set the bot's short description.
func (b *Bot) SetMyShortDescription() *SetMyShortDescription {
	return &SetMyShortDescription{
		bot:  b,
		opts: new(gotgbot.SetMyShortDescriptionOpts),
	}
}

// SetMyName creates a new SetMyName request to set the bot's name.
func (b *Bot) SetMyName() *SetMyName {
	return &SetMyName{
		bot:  b,
		opts: new(gotgbot.SetMyNameOpts),
	}
}

// GetMyName creates a new GetMyName request to get the bot's name.
func (b *Bot) GetMyName() *GetMyName {
	return &GetMyName{
		bot:  b,
		opts: new(gotgbot.GetMyNameOpts),
	}
}

// SetMyCommands sets the list of bot commands.
func (b *Bot) SetMyCommands() *SetMyCommands {
	return &SetMyCommands{
		bot:  b,
		opts: new(gotgbot.SetMyCommandsOpts),
	}
}

// GetMyCommands gets the current list of bot commands.
func (b *Bot) GetMyCommands() *GetMyCommands {
	return &GetMyCommands{
		bot:  b,
		opts: new(gotgbot.GetMyCommandsOpts),
	}
}

// DeleteMyCommands deletes the list of bot commands.
func (b *Bot) DeleteMyCommands() *DeleteMyCommands {
	return &DeleteMyCommands{
		bot:  b,
		opts: new(gotgbot.DeleteMyCommandsOpts),
	}
}

// LogOut creates a new LogOut request to log out from the cloud Bot API server.
func (b *Bot) LogOut() *LogOut {
	return &LogOut{
		bot:  b,
		opts: new(gotgbot.LogOutOpts),
	}
}

// Close creates a new Close request to close the bot instance.
func (b *Bot) Close() *Close {
	return &Close{
		bot:  b,
		opts: new(gotgbot.CloseOpts),
	}
}

// SetMyDefaultAdministratorRights creates a new SetMyDefaultAdministratorRights request.
func (b *Bot) SetMyDefaultAdministratorRights() *SetMyDefaultAdministratorRights {
	return &SetMyDefaultAdministratorRights{
		bot:  b,
		opts: new(gotgbot.SetMyDefaultAdministratorRightsOpts),
	}
}

// GetMyDefaultAdministratorRights creates a new GetMyDefaultAdministratorRights request.
func (b *Bot) GetMyDefaultAdministratorRights() *GetMyDefaultAdministratorRights {
	return &GetMyDefaultAdministratorRights{
		bot:  b,
		opts: new(gotgbot.GetMyDefaultAdministratorRightsOpts),
	}
}

// GetWebhookInfo creates a new GetWebhookInfo request.
func (b *Bot) GetWebhookInfo() *GetWebhookInfo {
	return &GetWebhookInfo{
		bot:  b,
		opts: new(gotgbot.GetWebhookInfoOpts),
	}
}

// DeleteWebhook creates a new DeleteWebhook request.
func (b *Bot) DeleteWebhook() *DeleteWebhook {
	return &DeleteWebhook{
		bot:  b,
		opts: new(gotgbot.DeleteWebhookOpts),
	}
}

// GetMe creates a new GetMe request.
func (b *Bot) GetMe() *GetMe {
	return &GetMe{
		bot:  b,
		opts: new(gotgbot.GetMeOpts),
	}
}
