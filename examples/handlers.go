package main

import (
	"log"

	. "github.com/enetx/g"
	"github.com/enetx/tg"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token)

	// Message
	bot.On.Message.Text(func(ctx *tg.Context) error {
		return ctx.Reply("Text received").Send().Err()
	})

	bot.On.Message.Photo(func(ctx *tg.Context) error {
		return ctx.Reply("Photo received").Send().Err()
	})

	bot.On.Message.Voice(func(ctx *tg.Context) error {
		return ctx.Reply("Voice received").Send().Err()
	})

	bot.On.Message.Video(func(ctx *tg.Context) error {
		return ctx.Reply("Video received").Send().Err()
	})

	bot.On.Message.Audio(func(ctx *tg.Context) error {
		return ctx.Reply("Audio received").Send().Err()
	})

	bot.On.Message.Sticker(func(ctx *tg.Context) error {
		return ctx.Reply("Sticker received").Send().Err()
	})

	bot.On.Message.Document(func(ctx *tg.Context) error {
		return ctx.Reply("Document received").Send().Err()
	})

	bot.On.Message.Location(func(ctx *tg.Context) error {
		return ctx.Reply("Location received").Send().Err()
	})

	bot.On.Message.Contact(func(ctx *tg.Context) error {
		return ctx.Reply("Contact received").Send().Err()
	})

	bot.On.Message.Poll(func(ctx *tg.Context) error {
		return ctx.Reply("Poll message received").Send().Err()
	})

	bot.On.Message.Reply(func(ctx *tg.Context) error {
		return ctx.Reply("This is a reply").Send().Err()
	})

	bot.On.Message.Any(func(ctx *tg.Context) error {
		for _, user := range ctx.EffectiveMessage.NewChatMembers {
			_ = ctx.Reply(Format("Welcome, {}!", user.FirstName)).Send()
		}
		return nil
	})

	// Callback
	bot.On.Callback.Equal("confirm", func(ctx *tg.Context) error {
		return ctx.Answer("Confirmed").Send().Err()
	})

	bot.On.Callback.Prefix("cb_", func(ctx *tg.Context) error {
		return ctx.Answer("Callback prefix matched").Send().Err()
	})

	// InlineQuery
	bot.On.Inline.Any(func(ctx *tg.Context) error {
		return ctx.Reply("Inline query received").Send().Err()
	})

	// Poll
	bot.On.Poll.Any(func(ctx *tg.Context) error {
		return ctx.Reply("New poll received").Send().Err()
	})

	// PollAnswer
	bot.On.PollAnswer.Any(func(ctx *tg.Context) error {
		return ctx.Reply("Poll answer received").Send().Err()
	})

	// ChatMember updates
	bot.On.ChatMember.
		Joined(func(ctx *tg.Context) error {
			return ctx.Reply("User joined").Send().Err()
		}).
		Left(func(ctx *tg.Context) error {
			return ctx.Reply("User left").Send().Err()
		}).
		Kicked(func(ctx *tg.Context) error {
			return ctx.Reply("User was kicked").Send().Err()
		}).
		Unbanned(func(ctx *tg.Context) error {
			return ctx.Reply("User was unbanned").Send().Err()
		}).
		Promoted(func(ctx *tg.Context) error {
			return ctx.Reply("User was promoted").Send().Err()
		}).
		Demoted(func(ctx *tg.Context) error {
			return ctx.Reply("User was demoted").Send().Err()
		})

	// MyChatMember (bot's own status updates)
	bot.On.MyChatMember.Any(func(ctx *tg.Context) error {
		return ctx.Reply("Bot status changed").Send().Err()
	})

	// ChatJoinRequest
	bot.On.ChatJoinRequest.Any(func(ctx *tg.Context) error {
		return ctx.Reply("Join request received").Send().Err()
	})

	// ChosenInlineResult
	bot.On.ChosenInlineResult.Any(func(*tg.Context) error {
		log.Println("Chosen inline result received")
		return nil
	})

	// ShippingQuery
	bot.On.Shipping.Any(func(ctx *tg.Context) error {
		return ctx.Reply("Shipping query received").Send().Err()
	})

	// PreCheckoutQuery
	bot.On.PreCheckout.Any(func(ctx *tg.Context) error {
		return ctx.Reply("Pre-checkout query received").Send().Err()
	})

	// Reactions
	bot.On.Reaction.Any(func(ctx *tg.Context) error {
		return ctx.Reply("Reaction update received").Send().Err()
	})

	// PaidMediaPurchased
	bot.On.PaidMedia.PayloadPrefix("buy_", func(ctx *tg.Context) error {
		return ctx.Reply("Paid media purchased").Send().Err()
	})

	// Start polling
	bot.Polling().Start()
}
