package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// verification system examples

	// Verify a user
	b.Command("verifyuser", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /verifyuser <user_id>").Send().Err()
		}

		userID := args[0].ToInt().Unwrap().Int64()

		result := ctx.VerifyUser(userID).Send()
		if result.IsErr() {
			return result.Err()
		}

		if result.Ok() {
			return ctx.Reply("User " + args[0] + " verified successfully!").Send().Err()
		}
		return ctx.Reply("User " + args[0] + " not verified.").Send().Err()
	})

	// Verify a chat
	b.Command("verifychat", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /verifychat <chat_id>").Send().Err()
		}

		chatID := args[0].ToInt().Unwrap().Int64()

		result := ctx.VerifyChat(chatID).Send()
		if result.IsErr() {
			return result.Err()
		}

		if result.Ok() {
			return ctx.Reply("Chat " + args[0] + " verified successfully!").Send().Err()
		}
		return ctx.Reply("Chat " + args[0] + " not verified.").Send().Err()
	})

	// Remove user verification
	b.Command("removeuserverify", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /removeuserverify <user_id>").Send().Err()
		}

		userID := args[0].ToInt().Unwrap().Int64()

		result := ctx.RemoveUserVerification(userID).Send()
		if result.IsErr() {
			return result.Err()
		}

		if result.Ok() {
			return ctx.Reply("User " + args[0] + " verification removed!").Send().Err()
		}
		return ctx.Reply("Failed to remove user " + args[0] + " verification.").Send().Err()
	})

	// Remove chat verification
	b.Command("removechatverify", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /removechatverify <chat_id>").Send().Err()
		}

		chatID := args[0].ToInt().Unwrap().Int64()

		result := ctx.RemoveChatVerification(chatID).Send()
		if result.IsErr() {
			return result.Err()
		}

		if result.Ok() {
			return ctx.Reply("Chat " + args[0] + " verification removed!").Send().Err()
		}
		return ctx.Reply("Failed to remove chat " + args[0] + " verification.").Send().Err()
	})
}
