package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
)

// Command represents a command handler for bot commands.
type Command struct {
	bot          core.BotAPI
	command      String
	handler      Handler
	name         String
	triggers     []rune
	allowEdited  bool
	allowChannel bool
}

// NewCommand creates a new command handler for the specified command.
func NewCommand(bot core.BotAPI, cmd String, handler Handler) *Command {
	return &Command{
		bot:          bot,
		command:      cmd.Lower(),
		handler:      handler,
		name:         "command_" + cmd.Lower(),
		triggers:     []rune{'/'},
		allowEdited:  false,
		allowChannel: false,
	}
}

// AllowEdited configures the command to handle edited messages containing the command.
func (c *Command) AllowEdited() *Command {
	c.allowEdited = true
	return c
}

// AllowChannel configures the command to handle commands from channels.
func (c *Command) AllowChannel() *Command {
	c.allowChannel = true
	return c
}

// Triggers sets custom trigger characters for the command instead of the default '/'.
func (c *Command) Triggers(r ...rune) *Command {
	c.triggers = r
	return c
}

// Register registers the command handler with the bot dispatcher.
func (c *Command) Register() {
	c.bot.Dispatcher().RemoveHandlerFromGroup(c.name.Std(), 0)

	cmd := handlers.Command{
		Triggers:     c.triggers,
		AllowEdited:  c.allowEdited,
		AllowChannel: c.allowChannel,
		Command:      c.command.Std(),
		Response:     wrap(c.bot, middlewares(c.bot), c.handler),
	}

	c.bot.Dispatcher().AddHandlerToGroup(namedHandler{
		name:    c.name.Std(),
		Handler: cmd,
	}, 0)
}
