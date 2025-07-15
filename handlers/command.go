package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"

	"github.com/enetx/tg/core"

	. "github.com/enetx/g"
)

type Command struct {
	bot          core.BotAPI
	command      String
	handler      Handler
	name         String
	triggers     []rune
	allowEdited  bool
	allowChannel bool
}

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

func (c *Command) AllowEdited() *Command {
	c.allowEdited = true
	return c
}

func (c *Command) AllowChannel() *Command {
	c.allowChannel = true
	return c
}

func (c *Command) Triggers(r ...rune) *Command {
	c.triggers = r
	return c
}

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
