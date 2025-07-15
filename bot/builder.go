package bot

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	. "github.com/enetx/g"
	"github.com/enetx/tg/handlers"
)

type BotBuilder struct {
	token String
	opts  *gotgbot.BotOpts
}

func (b *BotBuilder) APIURL(url String) *BotBuilder {
	if base, ok := b.opts.BotClient.(*gotgbot.BaseBotClient); ok {
		base.DefaultRequestOpts.APIURL = url.Std()
	}

	return b
}

func (b *BotBuilder) UseTestEnvironment() *BotBuilder {
	if base, ok := b.opts.BotClient.(*gotgbot.BaseBotClient); ok {
		base.UseTestEnvironment = true
	}

	return b
}

func (b *BotBuilder) DisableTokenCheck() *BotBuilder {
	b.opts.DisableTokenCheck = true
	return b
}

func (b *BotBuilder) UseClient(c *http.Client) *BotBuilder {
	b.opts.BotClient = &gotgbot.BaseBotClient{Client: *c}
	return b
}

func (b *BotBuilder) Timeout(d time.Duration) *BotBuilder {
	b.opts.RequestOpts.Timeout = d
	return b
}

func (b *BotBuilder) Build() Result[*Bot] {
	raw := &gotgbot.Bot{
		Token:     b.token.Std(),
		BotClient: b.opts.BotClient,
	}

	if !b.opts.DisableTokenCheck {
		user, err := raw.GetMe(&gotgbot.GetMeOpts{RequestOpts: b.opts.RequestOpts})
		if err != nil {
			return Err[*Bot](fmt.Errorf("failed to check bot token: %w", err))
		}

		raw.User = *user
	} else {
		split := b.token.Split(":").Collect()
		if split.Len().Ne(2) {
			return Err[*Bot](fmt.Errorf("invalid token format: expected '123456:ABCDEF', got '%s'", b.token))
		}

		id := split[0].ToInt()
		if id.IsErr() {
			return Err[*Bot](fmt.Errorf("failed to parse bot ID from token: %w", id.Err()))
		}

		raw.User = gotgbot.User{
			Id:        id.Ok().Int64(),
			IsBot:     true,
			FirstName: "<unknown>",
			Username:  "<unknown>",
		}
	}

	bot := &Bot{
		token:      b.token,
		raw:        raw,
		dispatcher: ext.NewDispatcher(nil),
	}

	bot.updater = ext.NewUpdater(bot.dispatcher, nil)
	bot.On = handlers.NewHandlers(bot)

	return Ok(bot)
}
