package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// DeleteMyCommands represents a request to delete bot commands.
type DeleteMyCommands struct {
	bot  *Bot
	opts *gotgbot.DeleteMyCommandsOpts
}

// Scope sets the scope for which to delete the commands.
func (dmc *DeleteMyCommands) Scope(scope gotgbot.BotCommandScope) *DeleteMyCommands {
	dmc.opts.Scope = scope
	return dmc
}

// ScopeDefault deletes commands for all users.
func (dmc *DeleteMyCommands) ScopeDefault() *DeleteMyCommands {
	dmc.opts.Scope = gotgbot.BotCommandScopeDefault{}
	return dmc
}

// ScopeAllPrivateChats deletes commands for all private chats.
func (dmc *DeleteMyCommands) ScopeAllPrivateChats() *DeleteMyCommands {
	dmc.opts.Scope = gotgbot.BotCommandScopeAllPrivateChats{}
	return dmc
}

// ScopeAllGroupChats deletes commands for all group chats.
func (dmc *DeleteMyCommands) ScopeAllGroupChats() *DeleteMyCommands {
	dmc.opts.Scope = gotgbot.BotCommandScopeAllGroupChats{}
	return dmc
}

// ScopeAllChatAdministrators deletes commands for all chat administrators.
func (dmc *DeleteMyCommands) ScopeAllChatAdministrators() *DeleteMyCommands {
	dmc.opts.Scope = gotgbot.BotCommandScopeAllChatAdministrators{}
	return dmc
}

// ScopeChat deletes commands for a specific chat.
func (dmc *DeleteMyCommands) ScopeChat(chatID int64) *DeleteMyCommands {
	dmc.opts.Scope = gotgbot.BotCommandScopeChat{ChatId: chatID}
	return dmc
}

// ScopeChatAdministrators deletes commands for administrators of a specific chat.
func (dmc *DeleteMyCommands) ScopeChatAdministrators(chatID int64) *DeleteMyCommands {
	dmc.opts.Scope = gotgbot.BotCommandScopeChatAdministrators{ChatId: chatID}
	return dmc
}

// ScopeChatMember deletes commands for a specific member of a specific chat.
func (dmc *DeleteMyCommands) ScopeChatMember(chatID, userID int64) *DeleteMyCommands {
	dmc.opts.Scope = gotgbot.BotCommandScopeChatMember{
		ChatId: chatID,
		UserId: userID,
	}

	return dmc
}

// LanguageCode sets the language code for the commands.
func (dmc *DeleteMyCommands) LanguageCode(code g.String) *DeleteMyCommands {
	dmc.opts.LanguageCode = code.Std()
	return dmc
}

// Timeout sets a custom timeout for this request.
func (dmc *DeleteMyCommands) Timeout(duration time.Duration) *DeleteMyCommands {
	if dmc.opts.RequestOpts == nil {
		dmc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dmc.opts.RequestOpts.Timeout = duration

	return dmc
}

// APIURL sets a custom API URL for this request.
func (dmc *DeleteMyCommands) APIURL(url g.String) *DeleteMyCommands {
	if dmc.opts.RequestOpts == nil {
		dmc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dmc.opts.RequestOpts.APIURL = url.Std()

	return dmc
}

// Send deletes the bot commands.
func (dmc *DeleteMyCommands) Send() g.Result[bool] {
	return g.ResultOf(dmc.bot.raw.DeleteMyCommands(dmc.opts))
}
