package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// GetMyCommands represents a request to get bot commands.
type GetMyCommands struct {
	bot  *Bot
	opts *gotgbot.GetMyCommandsOpts
}

// Scope sets the scope for which to get the commands.
func (gmc *GetMyCommands) Scope(scope gotgbot.BotCommandScope) *GetMyCommands {
	gmc.opts.Scope = scope
	return gmc
}

// ScopeDefault gets commands for all users.
func (gmc *GetMyCommands) ScopeDefault() *GetMyCommands {
	gmc.opts.Scope = gotgbot.BotCommandScopeDefault{}
	return gmc
}

// ScopeAllPrivateChats gets commands for all private chats.
func (gmc *GetMyCommands) ScopeAllPrivateChats() *GetMyCommands {
	gmc.opts.Scope = gotgbot.BotCommandScopeAllPrivateChats{}
	return gmc
}

// ScopeAllGroupChats gets commands for all group chats.
func (gmc *GetMyCommands) ScopeAllGroupChats() *GetMyCommands {
	gmc.opts.Scope = gotgbot.BotCommandScopeAllGroupChats{}
	return gmc
}

// ScopeAllChatAdministrators gets commands for all chat administrators.
func (gmc *GetMyCommands) ScopeAllChatAdministrators() *GetMyCommands {
	gmc.opts.Scope = gotgbot.BotCommandScopeAllChatAdministrators{}
	return gmc
}

// ScopeChat gets commands for a specific chat.
func (gmc *GetMyCommands) ScopeChat(chatID int64) *GetMyCommands {
	gmc.opts.Scope = gotgbot.BotCommandScopeChat{ChatId: chatID}
	return gmc
}

// ScopeChatAdministrators gets commands for administrators of a specific chat.
func (gmc *GetMyCommands) ScopeChatAdministrators(chatID int64) *GetMyCommands {
	gmc.opts.Scope = gotgbot.BotCommandScopeChatAdministrators{ChatId: chatID}
	return gmc
}

// ScopeChatMember gets commands for a specific member of a specific chat.
func (gmc *GetMyCommands) ScopeChatMember(chatID, userID int64) *GetMyCommands {
	gmc.opts.Scope = gotgbot.BotCommandScopeChatMember{
		ChatId: chatID,
		UserId: userID,
	}

	return gmc
}

// LanguageCode sets the language code for the commands.
func (gmc *GetMyCommands) LanguageCode(code String) *GetMyCommands {
	gmc.opts.LanguageCode = code.Std()
	return gmc
}

// Timeout sets a custom timeout for this request.
func (gmc *GetMyCommands) Timeout(duration time.Duration) *GetMyCommands {
	if gmc.opts.RequestOpts == nil {
		gmc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmc.opts.RequestOpts.Timeout = duration

	return gmc
}

// APIURL sets a custom API URL for this request.
func (gmc *GetMyCommands) APIURL(url String) *GetMyCommands {
	if gmc.opts.RequestOpts == nil {
		gmc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmc.opts.RequestOpts.APIURL = url.Std()

	return gmc
}

// Send gets the bot commands.
func (gmc *GetMyCommands) Send() Result[[]gotgbot.BotCommand] {
	return ResultOf(gmc.bot.raw.GetMyCommands(gmc.opts))
}
