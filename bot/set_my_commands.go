package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// SetMyCommands represents a request to set bot commands.
type SetMyCommands struct {
	bot      *Bot
	commands Slice[gotgbot.BotCommand]
	opts     *gotgbot.SetMyCommandsOpts
}

// AddCommand adds a command to the command list.
func (smc *SetMyCommands) AddCommand(command, description String) *SetMyCommands {
	smc.commands.Push(gotgbot.BotCommand{
		Command:     command.Std(),
		Description: description.Std(),
	})

	return smc
}

// Commands sets the entire command list at once.
func (smc *SetMyCommands) Commands(commands []gotgbot.BotCommand) *SetMyCommands {
	smc.commands = commands
	return smc
}

// Scope sets the scope for which the commands are relevant.
func (smc *SetMyCommands) Scope(scope gotgbot.BotCommandScope) *SetMyCommands {
	smc.opts.Scope = scope
	return smc
}

// ScopeDefault sets commands for all users.
func (smc *SetMyCommands) ScopeDefault() *SetMyCommands {
	smc.opts.Scope = gotgbot.BotCommandScopeDefault{}
	return smc
}

// ScopeAllPrivateChats sets commands for all private chats.
func (smc *SetMyCommands) ScopeAllPrivateChats() *SetMyCommands {
	smc.opts.Scope = gotgbot.BotCommandScopeAllPrivateChats{}
	return smc
}

// ScopeAllGroupChats sets commands for all group chats.
func (smc *SetMyCommands) ScopeAllGroupChats() *SetMyCommands {
	smc.opts.Scope = gotgbot.BotCommandScopeAllGroupChats{}
	return smc
}

// ScopeAllChatAdministrators sets commands for all chat administrators.
func (smc *SetMyCommands) ScopeAllChatAdministrators() *SetMyCommands {
	smc.opts.Scope = gotgbot.BotCommandScopeAllChatAdministrators{}
	return smc
}

// ScopeChat sets commands for a specific chat.
func (smc *SetMyCommands) ScopeChat(chatID int64) *SetMyCommands {
	smc.opts.Scope = gotgbot.BotCommandScopeChat{ChatId: chatID}
	return smc
}

// ScopeChatAdministrators sets commands for administrators of a specific chat.
func (smc *SetMyCommands) ScopeChatAdministrators(chatID int64) *SetMyCommands {
	smc.opts.Scope = gotgbot.BotCommandScopeChatAdministrators{ChatId: chatID}
	return smc
}

// ScopeChatMember sets commands for a specific member of a specific chat.
func (smc *SetMyCommands) ScopeChatMember(chatID, userID int64) *SetMyCommands {
	smc.opts.Scope = gotgbot.BotCommandScopeChatMember{
		ChatId: chatID,
		UserId: userID,
	}

	return smc
}

// LanguageCode sets the language code for the commands.
func (smc *SetMyCommands) LanguageCode(code String) *SetMyCommands {
	smc.opts.LanguageCode = code.Std()
	return smc
}

// Timeout sets a custom timeout for this request.
func (smc *SetMyCommands) Timeout(duration time.Duration) *SetMyCommands {
	if smc.opts.RequestOpts == nil {
		smc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smc.opts.RequestOpts.Timeout = duration

	return smc
}

// APIURL sets a custom API URL for this request.
func (smc *SetMyCommands) APIURL(url String) *SetMyCommands {
	if smc.opts.RequestOpts == nil {
		smc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smc.opts.RequestOpts.APIURL = url.Std()

	return smc
}

// Send sets the bot commands.
func (smc *SetMyCommands) Send() Result[bool] {
	if len(smc.commands) == 0 {
		return Err[bool](Errorf("no commands specified"))
	}

	if len(smc.commands) > 100 {
		return Err[bool](Errorf("too many commands: {} (maximum 100)", len(smc.commands)))
	}

	return ResultOf(smc.bot.raw.SetMyCommands(smc.commands, smc.opts))
}
