package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

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
