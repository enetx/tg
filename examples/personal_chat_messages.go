package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

// Personal chat messages (Bot API 10.0): fetch the last messages (1-20) from the chat
// a user currently has pinned to their profile.
func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// /personal — show how many recent profile-chat messages the caller has.
	b.Command("personal", func(ctx *ctx.Context) error {
		result := ctx.GetUserPersonalChatMessages(ctx.EffectiveUser.Id, 20).Send()
		if result.IsErr() {
			return ctx.Reply(g.Format("Failed to fetch messages: {}", result.Err())).Send().Err()
		}

		messages := result.Ok()
		return ctx.Reply(g.Format("Fetched {} message(s) from your personal chat.", messages.Len())).Send().Err()
	})

	b.Polling().Start()
}
