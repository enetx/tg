package tg

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	. "github.com/enetx/g"
)

type Bot struct {
	token       string
	dispatcher  *ext.Dispatcher
	updater     *ext.Updater
	states      *MapSafe[int64, String]
	stateData   *MapSafe[int64, *MapSafe[String, any]]
	middlewares Slice[Handler]
	On          *Handlers
	Raw         *gotgbot.Bot
}

func NewBot[T ~string](token T) *Bot {
	bot, err := gotgbot.NewBot(string(token), nil)
	if err != nil {
		panic("failed to create bot: " + err.Error())
	}

	b := &Bot{
		Raw:        bot,
		token:      string(token),
		dispatcher: ext.NewDispatcher(nil),
		states:     NewMapSafe[int64, String](),
		stateData:  NewMapSafe[int64, *MapSafe[String, any]](),
	}

	b.updater = ext.NewUpdater(b.dispatcher, nil)
	b.On = newHandlers(b)

	return b
}

func (b *Bot) Command(cmd String, fn Handler) *Command {
	c := &Command{
		bot:          b,
		command:      cmd.Lower(),
		handler:      fn,
		name:         "command_" + cmd.Lower(),
		triggers:     []rune{'/'},
		allowEdited:  false,
		allowChannel: false,
	}

	c.Register()

	return c
}

// func (b *Bot) Handlers(fns ...HandlerFunc) *Bot {
// 	for _, fn := range fns {
// 		b.Any(fn)
// 	}
//
// 	return b
// }

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

func (b *Bot) Dispatcher() *ext.Dispatcher {
	return b.dispatcher
}

func (b *Bot) Updater() *ext.Updater {
	return b.updater
}

func (b *Bot) HandleWebhook(data []byte) error {
	var update gotgbot.Update
	if err := json.Unmarshal(data, &update); err != nil {
		return Errorf("failed to unmarshal update: {}", err)
	}

	return b.dispatcher.ProcessUpdate(b.Raw, &update, nil)
}

func (b *Bot) Use(middleware Handler) *Bot {
	b.middlewares.Push(middleware)
	return b
}

func (b *Bot) handleChatJoinRequest(f filters.ChatJoinRequest, fn Handler) {
	b.dispatcher.AddHandler(handlers.NewChatJoinRequest(f, wrap(b, fn)))
}

func (b *Bot) handleChatMember(f filters.ChatMember, fn Handler) {
	b.dispatcher.AddHandler(handlers.NewChatMember(f, wrap(b, fn)))
}

func (b *Bot) handleMyChatMember(f filters.ChatMember, fn Handler) {
	b.dispatcher.AddHandler(handlers.NewMyChatMember(f, wrap(b, fn)))
}

func (b *Bot) handleChosenInlineResult(f filters.ChosenInlineResult, fn Handler) {
	b.dispatcher.AddHandler(handlers.NewChosenInlineResult(f, wrap(b, fn)))
}

func (b *Bot) handleInlineQuery(f filters.InlineQuery, fn Handler) {
	b.dispatcher.AddHandler(handlers.NewInlineQuery(f, wrap(b, fn)))
}

func (b *Bot) handleMessage(f filters.Message, fn Handler) *MessageHandler {
	m := &MessageHandler{
		bot:     b,
		filter:  f,
		handler: fn,
		name:    fmt.Sprintf("message_%x_%x", reflect.ValueOf(f).Pointer(), reflect.ValueOf(fn).Pointer()),
	}

	return m.Register()
}

func (b *Bot) handleCallback(f filters.CallbackQuery, fn Handler) *CallbackHandler {
	h := &CallbackHandler{
		bot:     b,
		filter:  f,
		handler: fn,
		name:    fmt.Sprintf("callback_%x_%x", reflect.ValueOf(f).Pointer(), reflect.ValueOf(fn).Pointer()),
	}

	return h.Register()
}

func (b *Bot) handlePoll(f filters.Poll, fn Handler) {
	b.dispatcher.AddHandler(handlers.NewPoll(f, wrap(b, fn)))
}

func (b *Bot) handlePollAnswer(f filters.PollAnswer, fn Handler) {
	b.dispatcher.AddHandler(handlers.NewPollAnswer(f, wrap(b, fn)))
}

func (b *Bot) handlePreCheckoutQuery(f filters.PreCheckoutQuery, fn Handler) {
	b.dispatcher.AddHandler(handlers.NewPreCheckoutQuery(f, wrap(b, fn)))
}

func (b *Bot) handleShippingQuery(f filters.ShippingQuery, fn Handler) {
	b.dispatcher.AddHandler(handlers.NewShippingQuery(f, wrap(b, fn)))
}

func (b *Bot) handleReaction(f filters.Reaction, fn Handler) {
	b.dispatcher.AddHandler(handlers.NewReaction(f, wrap(b, fn)))
}

func (b *Bot) handlePurchasedPaidMedia(f filters.PurchasedPaidMedia, fn Handler) {
	b.dispatcher.AddHandler(handlers.NewPurchasedPaidMedia(f, wrap(b, fn)))
}
