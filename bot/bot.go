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

type Bot struct {
	token       String
	dispatcher  *ext.Dispatcher
	updater     *ext.Updater
	middlewares Slice[handlers.Handler]
	On          *handlers.Handlers
	raw         *gotgbot.Bot
}

var _ core.BotAPI = (*Bot)(nil)

func (b *Bot) Dispatcher() *ext.Dispatcher {
	return b.dispatcher
}

func (b *Bot) Updater() *ext.Updater {
	return b.updater
}

func (b *Bot) Raw() *gotgbot.Bot {
	return b.raw
}

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

func (b *Bot) Command(cmd String, fn handlers.Handler) *handlers.Command {
	c := handlers.NewCommand(b, cmd, fn)
	c.Register()

	return c
}

func (b *Bot) Polling() *Polling {
	return &Polling{
		bot:  b,
		opts: &ext.PollingOpts{GetUpdatesOpts: new(gotgbot.GetUpdatesOpts)},
	}
}

func (b *Bot) Webhook() *Webhook {
	return &Webhook{
		bot: b,
		opt: new(gotgbot.SetWebhookOpts),
	}
}

func (b *Bot) HandleWebhook(data []byte) error {
	var update gotgbot.Update
	if err := json.Unmarshal(data, &update); err != nil {
		return Errorf("failed to unmarshal update: {}", err)
	}

	return b.dispatcher.ProcessUpdate(b.Raw(), &update, nil)
}

func (b *Bot) Use(middleware handlers.Handler) *Bot {
	b.middlewares.Push(middleware)
	return b
}

func (b *Bot) Middlewares() Slice[handlers.Handler] {
	return b.middlewares
}
