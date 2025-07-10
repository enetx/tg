package tg

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	. "github.com/enetx/g"
)

type Command struct {
	bot          *Bot
	command      String
	handler      Handler
	name         String
	triggers     []rune
	allowEdited  bool
	allowChannel bool
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
	c.bot.dispatcher.RemoveHandlerFromGroup(c.name.Std(), 0)

	cmd := handlers.Command{
		Triggers:     c.triggers,
		AllowEdited:  c.allowEdited,
		AllowChannel: c.allowChannel,
		Command:      c.command.Std(),
		Response:     wrap(c.bot, c.handler),
	}

	c.bot.dispatcher.AddHandlerToGroup(cmd, 0)
}
