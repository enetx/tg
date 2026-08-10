package main

import (
	"time"

	"github.com/enetx/g"
	"github.com/enetx/g/fs"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

func main() {
	token := g.Slice[g.String](fs.NewFile("../.env").Read().Ok().Trim().Split("=")).Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Create poll with stop button
	createPoll := func(ctx *ctx.Context) {
		ctx.SendPoll("🚀 What's your favorite programming language?").
			Option(input.Choice("Go")).
			Option(input.Choice("Python")).
			Option(input.Choice("JavaScript")).
			Option(input.Choice("Rust")).
			Anonymous().
			MultipleAnswers().
			Markup(
				keyboard.Inline().
					Text("⏹️ Stop Poll", "stop_poll")).
			Send()
	}

	// Start command creates a new poll
	b.Command("start", func(ctx *ctx.Context) error {
		createPoll(ctx)
		return nil
	})

	// Poll command creates poll with auto-stop
	b.Command("poll", func(ctx *ctx.Context) error {
		result := ctx.SendPoll("⏱️ Quick poll (auto-stops in 5 seconds)").
			Option(input.Choice("Option A")).
			Option(input.Choice("Option B")).
			Option(input.Choice("Option C")).
			Send()

		if result.IsOk() {
			// Auto-stop after 5 seconds
			time.AfterFunc(5*time.Second, func() {
				ctx.StopPoll().
					MessageID(result.Ok().MessageId).
					Markup(
						keyboard.Inline().
							Text("⏰ Poll Expired", "expired")).
					Send()
			})
		}

		return nil
	})

	// Handle stop poll callback
	b.On.Callback.Equal("stop_poll", func(ctx *ctx.Context) error {
		result := ctx.StopPoll().
			Markup(
				keyboard.Inline().
					Text("✅ Poll Stopped", "stopped")).
			Send()

		if result.IsOk() {
			return ctx.AnswerCallbackQuery("Poll stopped successfully").Send().Err()
		}

		return ctx.AnswerCallbackQuery("Failed to stop poll").Alert().Send().Err()
	})

	// Stop command for replying to polls
	b.Command("stop", func(ctx *ctx.Context) error {
		if ctx.EffectiveMessage.ReplyToMessage == nil {
			return ctx.Reply("Reply to a poll message with /stop").Send().Err()
		}

		if ctx.EffectiveMessage.ReplyToMessage.Poll == nil {
			return ctx.Reply("The replied message doesn't contain a poll").Send().Err()
		}

		result := ctx.StopPoll().
			MessageID(ctx.EffectiveMessage.ReplyToMessage.MessageId).
			Send()

		if result.IsOk() {
			poll := result.Ok()
			return ctx.Reply(g.Format("✅ Poll stopped! Total votes: {}", poll.TotalVoterCount)).Send().Err()
		}

		return ctx.Reply("❌ Failed to stop poll").Send().Err()
	})

	b.Polling().DropPendingUpdates().Start()
}
