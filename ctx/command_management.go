package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// SetMyCommands represents a request to set bot commands.
type SetMyCommands struct {
	ctx      *Context
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

	return ResultOf(smc.ctx.Bot.Raw().SetMyCommands(smc.commands, smc.opts))
}

// GetMyCommands represents a request to get bot commands.
type GetMyCommands struct {
	ctx  *Context
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
	return ResultOf(gmc.ctx.Bot.Raw().GetMyCommands(gmc.opts))
}

// DeleteMyCommands represents a request to delete bot commands.
type DeleteMyCommands struct {
	ctx  *Context
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
func (dmc *DeleteMyCommands) LanguageCode(code String) *DeleteMyCommands {
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
func (dmc *DeleteMyCommands) APIURL(url String) *DeleteMyCommands {
	if dmc.opts.RequestOpts == nil {
		dmc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dmc.opts.RequestOpts.APIURL = url.Std()

	return dmc
}

// Send deletes the bot commands.
func (dmc *DeleteMyCommands) Send() Result[bool] {
	return ResultOf(dmc.ctx.Bot.Raw().DeleteMyCommands(dmc.opts))
}

// SetChatMenuButton represents a request to set the menu button of a chat.
type SetChatMenuButton struct {
	ctx        *Context
	chatID     Option[*int64]
	menuButton gotgbot.MenuButton
	opts       *gotgbot.SetChatMenuButtonOpts
}

// ChatID sets the target chat ID.
func (scmb *SetChatMenuButton) ChatID(chatID int64) *SetChatMenuButton {
	scmb.chatID = Some(&chatID)
	return scmb
}

// MenuButton sets the menu button.
func (scmb *SetChatMenuButton) MenuButton(button gotgbot.MenuButton) *SetChatMenuButton {
	scmb.menuButton = button
	return scmb
}

// DefaultMenu sets the default menu button.
func (scmb *SetChatMenuButton) DefaultMenu() *SetChatMenuButton {
	scmb.menuButton = gotgbot.MenuButtonDefault{}
	return scmb
}

// WebAppMenu sets a web app menu button.
func (scmb *SetChatMenuButton) WebAppMenu(text String, webApp gotgbot.WebAppInfo) *SetChatMenuButton {
	scmb.menuButton = gotgbot.MenuButtonWebApp{
		Text:   text.Std(),
		WebApp: webApp,
	}

	return scmb
}

// CommandsMenu sets the commands menu button.
func (scmb *SetChatMenuButton) CommandsMenu() *SetChatMenuButton {
	scmb.menuButton = gotgbot.MenuButtonCommands{}
	return scmb
}

// Timeout sets a custom timeout for this request.
func (scmb *SetChatMenuButton) Timeout(duration time.Duration) *SetChatMenuButton {
	if scmb.opts.RequestOpts == nil {
		scmb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scmb.opts.RequestOpts.Timeout = duration

	return scmb
}

// APIURL sets a custom API URL for this request.
func (scmb *SetChatMenuButton) APIURL(url String) *SetChatMenuButton {
	if scmb.opts.RequestOpts == nil {
		scmb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scmb.opts.RequestOpts.APIURL = url.Std()

	return scmb
}

// Send sets the chat menu button.
func (scmb *SetChatMenuButton) Send() Result[bool] {
	scmb.opts.ChatId = scmb.chatID.UnwrapOrDefault()
	scmb.opts.MenuButton = scmb.menuButton

	return ResultOf(scmb.ctx.Bot.Raw().SetChatMenuButton(scmb.opts))
}

// GetChatMenuButton represents a request to get the menu button of a chat.
type GetChatMenuButton struct {
	ctx    *Context
	chatID Option[*int64]
	opts   *gotgbot.GetChatMenuButtonOpts
}

// ChatID sets the target chat ID.
func (gcmb *GetChatMenuButton) ChatID(chatID int64) *GetChatMenuButton {
	gcmb.chatID = Some(&chatID)
	return gcmb
}

// Timeout sets a custom timeout for this request.
func (gcmb *GetChatMenuButton) Timeout(duration time.Duration) *GetChatMenuButton {
	if gcmb.opts.RequestOpts == nil {
		gcmb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gcmb.opts.RequestOpts.Timeout = duration

	return gcmb
}

// APIURL sets a custom API URL for this request.
func (gcmb *GetChatMenuButton) APIURL(url String) *GetChatMenuButton {
	if gcmb.opts.RequestOpts == nil {
		gcmb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gcmb.opts.RequestOpts.APIURL = url.Std()

	return gcmb
}

// Send gets the chat menu button.
func (gcmb *GetChatMenuButton) Send() Result[gotgbot.MenuButton] {
	gcmb.opts.ChatId = gcmb.chatID.UnwrapOrDefault()
	return ResultOf(gcmb.ctx.Bot.Raw().GetChatMenuButton(gcmb.opts))
}
