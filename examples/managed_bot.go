package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/types/updates"
)

// Managed bots (Bot API 9.6 / 10.0): a bot that is allowed to manage other bots
// (enabled in the @BotFather Mini App) can fetch/replace their tokens, configure
// their access settings, and react to managed-bot lifecycle updates.
func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// React to any managed-bot lifecycle update (creation / token update / owner update).
	b.On.ManagedBot.Any(func(ctx *ctx.Context) error {
		mb := ctx.Update.ManagedBot
		return ctx.SendMessage(
			g.Format("Managed bot {} is owned by user {}", mb.Bot.Id, mb.User.Id),
		).To(mb.User.Id).Send().Err()
	})

	// Narrow the handler to a single owner or a single managed bot.
	b.On.ManagedBot.OwnedByUserID(123456789, func(ctx *ctx.Context) error {
		return ctx.SendMessage("A managed bot was updated for the VIP owner").
			To(ctx.Update.ManagedBot.User.Id).Send().Err()
	})

	// /mbtoken — reply to the managed-bot owner to fetch the bot token.
	b.Command("mbtoken", func(ctx *ctx.Context) error {
		if ctx.EffectiveMessage.ReplyToMessage == nil || ctx.EffectiveMessage.ReplyToMessage.From == nil {
			return ctx.Reply("Reply to the managed-bot owner.").Send().Err()
		}

		userID := ctx.EffectiveMessage.ReplyToMessage.From.Id

		result := ctx.GetManagedBotToken(userID).Send()
		if result.IsErr() {
			return ctx.Reply(g.Format("Failed to get token: {}", result.Err())).Send().Err()
		}

		return ctx.Reply(g.Format("Token: <code>{}</code>", result.Ok())).HTML().Send().Err()
	})

	// /mbrevoke — revoke the current token and generate a new one.
	b.Command("mbrevoke", func(ctx *ctx.Context) error {
		if ctx.EffectiveMessage.ReplyToMessage == nil || ctx.EffectiveMessage.ReplyToMessage.From == nil {
			return ctx.Reply("Reply to the managed-bot owner.").Send().Err()
		}

		userID := ctx.EffectiveMessage.ReplyToMessage.From.Id

		result := ctx.ReplaceManagedBotToken(userID).Send()
		if result.IsErr() {
			return ctx.Reply(g.Format("Failed to replace token: {}", result.Err())).Send().Err()
		}

		return ctx.Reply("Token replaced. The previous token is now revoked.").Send().Err()
	})

	// /mbaccess — restrict the managed bot to its owner plus two extra users, then read it back.
	b.Command("mbaccess", func(ctx *ctx.Context) error {
		userID := ctx.EffectiveUser.Id

		set := ctx.SetManagedBotAccessSettings(userID, true).
			AddedUserIDs(111111, 222222).
			Send()
		if set.IsErr() {
			return ctx.Reply(g.Format("Failed to set access: {}", set.Err())).Send().Err()
		}

		get := ctx.GetManagedBotAccessSettings(userID).Send()
		if get.IsErr() {
			return ctx.Reply(g.Format("Failed to read access: {}", get.Err())).Send().Err()
		}

		settings := get.Ok()
		return ctx.Reply(g.Format(
			"Access restricted: {}, additional users: {}",
			settings.IsAccessRestricted, len(settings.AddedUsers),
		)).Send().Err()
	})

	// /preparebtn — store a prepared keyboard button that asks the user to create a managed bot.
	b.Command("preparebtn", func(ctx *ctx.Context) error {
		result := ctx.SavePreparedKeyboardButton(ctx.EffectiveUser.Id).
			Text("Create my bot").
			RequestManagedBot(1).
			SuggestedName("My Assistant").
			SuggestedUsername("my_assistant_bot").
			Send()
		if result.IsErr() {
			return ctx.Reply(g.Format("Failed to save button: {}", result.Err())).Send().Err()
		}

		return ctx.Reply(g.Format("Prepared button id: <code>{}</code>", result.Ok().Id)).HTML().Send().Err()
	})

	b.Polling().AllowedUpdates(updates.All...).Start()
}
