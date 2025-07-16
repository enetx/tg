package main

import (
	"log"

	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Message
	b.On.Message.Text(func(ctx *ctx.Context) error {
		return ctx.Reply("Text received").Send().Err()
	})

	b.On.Message.Photo(func(ctx *ctx.Context) error {
		return ctx.Reply("Photo received").Send().Err()
	})

	b.On.Message.Voice(func(ctx *ctx.Context) error {
		return ctx.Reply("Voice received").Send().Err()
	})

	b.On.Message.Video(func(ctx *ctx.Context) error {
		return ctx.Reply("Video received").Send().Err()
	})

	b.On.Message.Audio(func(ctx *ctx.Context) error {
		return ctx.Reply("Audio received").Send().Err()
	})

	b.On.Message.Sticker(func(ctx *ctx.Context) error {
		return ctx.Reply("Sticker received").Send().Err()
	})

	b.On.Message.Document(func(ctx *ctx.Context) error {
		return ctx.Reply("Document received").Send().Err()
	})

	b.On.Message.Location(func(ctx *ctx.Context) error {
		return ctx.Reply("Location received").Send().Err()
	})

	b.On.Message.Contact(func(ctx *ctx.Context) error {
		return ctx.Reply("Contact received").Send().Err()
	})

	b.On.Message.Poll(func(ctx *ctx.Context) error {
		return ctx.Reply("Poll message received").Send().Err()
	})

	b.On.Message.Reply(func(ctx *ctx.Context) error {
		return ctx.Reply("This is a reply").Send().Err()
	})

	b.On.Message.Any(func(ctx *ctx.Context) error {
		for _, user := range ctx.EffectiveMessage.NewChatMembers {
			_ = ctx.Reply(Format("Welcome, {}!", user.FirstName)).Send()
		}
		return nil
	})

	// Callback
	b.On.Callback.Equal("confirm", func(ctx *ctx.Context) error {
		return ctx.Answer("Confirmed").Send().Err()
	})

	b.On.Callback.Prefix("cb_", func(ctx *ctx.Context) error {
		return ctx.Answer("Callback prefix matched").Send().Err()
	})

	// InlineQuery
	b.On.Inline.Any(func(ctx *ctx.Context) error {
		return ctx.Reply("Inline query received").Send().Err()
	})

	// Poll
	b.On.Poll.Any(func(ctx *ctx.Context) error {
		return ctx.Reply("New poll received").Send().Err()
	})

	// PollAnswer
	b.On.PollAnswer.Any(func(ctx *ctx.Context) error {
		return ctx.Reply("Poll answer received").Send().Err()
	})

	// ChatMember updates
	b.On.ChatMember.
		Joined(func(ctx *ctx.Context) error {
			return ctx.Reply("User joined").Send().Err()
		}).
		Left(func(ctx *ctx.Context) error {
			return ctx.Reply("User left").Send().Err()
		}).
		Banned(func(ctx *ctx.Context) error {
			return ctx.Reply("User was kicked").Send().Err()
		}).
		Unbanned(func(ctx *ctx.Context) error {
			return ctx.Reply("User was unbanned").Send().Err()
		}).
		Promoted(func(ctx *ctx.Context) error {
			return ctx.Reply("User was promoted").Send().Err()
		}).
		Demoted(func(ctx *ctx.Context) error {
			return ctx.Reply("User was demoted").Send().Err()
		})

	// MyChatMember (bot's own status updates)
	b.On.MyChatMember.Any(func(ctx *ctx.Context) error {
		return ctx.Reply("Bot status changed").Send().Err()
	})

	// ChatJoinRequest
	b.On.ChatJoinRequest.Any(func(ctx *ctx.Context) error {
		return ctx.Reply("Join request received").Send().Err()
	})

	// ChosenInlineResult
	b.On.ChosenInlineResult.Any(func(*ctx.Context) error {
		log.Println("Chosen inline result received")
		return nil
	})

	// ShippingQuery
	b.On.Shipping.Any(func(ctx *ctx.Context) error {
		return ctx.Reply("Shipping query received").Send().Err()
	})

	// PreCheckoutQuery
	b.On.PreCheckout.Any(func(ctx *ctx.Context) error {
		return ctx.Reply("Pre-checkout query received").Send().Err()
	})

	// Reactions
	b.On.Reaction.Any(func(ctx *ctx.Context) error {
		return ctx.Reply("Reaction update received").Send().Err()
	})

	// PaidMediaPurchased
	b.On.PaidMedia.PayloadPrefix("buy_", func(ctx *ctx.Context) error {
		return ctx.Reply("Paid media purchased").Send().Err()
	})

	// Start polling
	b.Polling().Start()
}
