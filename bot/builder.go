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

// BotBuilder provides a fluent interface for configuring and building Telegram bots.
type BotBuilder struct {
	token String
	opts  *gotgbot.BotOpts
}

// UseTestEnvironment configures the bot to use Telegram's test environment.
func (b *BotBuilder) UseTestEnvironment() *BotBuilder {
	if base, ok := b.opts.BotClient.(*gotgbot.BaseBotClient); ok {
		base.UseTestEnvironment = true
	}

	return b
}

// DisableTokenCheck disables the automatic bot token validation during build.
func (b *BotBuilder) DisableTokenCheck() *BotBuilder {
	b.opts.DisableTokenCheck = true
	return b
}

// UseClient sets a custom HTTP client for the bot.
func (b *BotBuilder) UseClient(c *http.Client) *BotBuilder {
	b.opts.BotClient = &gotgbot.BaseBotClient{Client: *c}
	return b
}

// Timeout sets the timeout for bot builder operations.
func (b *BotBuilder) Timeout(duration time.Duration) *BotBuilder {
	b.opts.RequestOpts.Timeout = duration
	return b
}

// APIURL sets the API URL for bot builder operations.
func (b *BotBuilder) APIURL(url String) *BotBuilder {
	b.opts.RequestOpts.APIURL = url.Std()
	return b
}

// DefaultTimeout sets the default timeout for all bot requests.
func (b *BotBuilder) DefaultTimeout(duration time.Duration) *BotBuilder {
	if base, ok := b.opts.BotClient.(*gotgbot.BaseBotClient); ok {
		base.DefaultRequestOpts.Timeout = duration
	}

	return b
}

// DefaultAPIURL sets the default API URL for all bot requests.
func (b *BotBuilder) DefaultAPIURL(url String) *BotBuilder {
	if base, ok := b.opts.BotClient.(*gotgbot.BaseBotClient); ok {
		base.DefaultRequestOpts.APIURL = url.Std()
	}

	return b
}

// Build creates and initializes a new Bot instance with the configured settings.
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
