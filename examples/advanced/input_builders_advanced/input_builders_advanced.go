// Package main demonstrates advanced usage of all input builders in TG Framework.
// This example showcases every input builder type with real-world use cases,
// metadata extraction, formatting options, and integration patterns.
package main

import (
	"log"
	"os"

	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

func main() {
	// Get bot token from environment
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	// Create bot instance
	b := bot.New(token).Build().Unwrap()

	// Start command handler
	b.Command("start", handleStart).Register()

	// Media builders demonstration
	b.On.Callback.Equal("demo_media", handleMediaDemo)
	b.On.Callback.Equal("media_photo", handlePhotoBuilder)
	b.On.Callback.Equal("media_group", handleMediaGroupBuilder)

	// Poll builders demonstration
	b.On.Callback.Equal("demo_poll", handlePollDemo)
	b.On.Callback.Equal("poll_simple", handleSimplePoll)

	// Back navigation
	b.On.Callback.Equal("start", handleStart)

	// Start the bot
	log.Println("🚀 Advanced Input Builders Example Bot started...")
	b.Polling().AllowedUpdates().Start()
}

// handleStart provides main menu for input builder demonstrations
func handleStart(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("📸 Media Builders", "demo_media").
		Text("📊 Poll Builders", "demo_poll")

	return ctx.Reply("🎯 <b>Advanced Input Builders Showcase</b>\n\n" +
		"Choose a category to see advanced input builder examples:\n\n" +
		"📸 <b>Media Builders</b> - Photos, videos, documents with metadata\n" +
		"📊 <b>Poll Builders</b> - Interactive polls and quizzes").
		HTML().
		Markup(kb).
		Send().Err()
}

// ================ MEDIA BUILDERS ================

func handleMediaDemo(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("📷 Photo Builder", "media_photo").
		Text("📚 Media Group", "media_group").
		Row().
		Text("🔙 Back", "start")

	return ctx.EditMessageText("📸 <b>Media Builders Examples</b>\n\n" +
		"These examples show how to use input.Media builders for different content types:\n\n" +
		"📷 <b>Photo Builder</b> - Images with captions and effects\n" +
		"📚 <b>Media Group</b> - Albums with mixed media types").
		HTML().
		Markup(kb).
		Send().Err()
}

func handlePhotoBuilder(ctx *ctx.Context) error {
	// Create file input
	photoFile := file.Input("https://picsum.photos/800/600").Unwrap()

	// Demonstrate photo builder with various options
	photo := input.Photo(photoFile).
		Caption("🌅 <b>Beautiful Landscape</b>\n\n" +
			"This photo demonstrates:\n" +
			"• HTML formatting in captions\n" +
			"• Spoiler effects").
		HTML().
		Spoiler()

	// Send using MediaGroup for demonstration
	ctx.MediaGroup().Photo(photo).Send()

	return ctx.AnswerCallbackQuery("📷 Photo builder example sent!").Send().Err()
}

func handleMediaGroupBuilder(ctx *ctx.Context) error {
	// Demonstrate complex media group with different types
	file1 := file.Input("https://picsum.photos/400/300?random=1").Unwrap()
	file2 := file.Input("https://picsum.photos/400/300?random=2").Unwrap()

	photo1 := input.Photo(file1).
		Caption("Photo 1 with caption")

	photo2 := input.Photo(file2).
		Caption("<b>Photo 2</b> with HTML").
		HTML()

	ctx.MediaGroup().
		Photo(photo1).
		Photo(photo2).
		Silent().
		Send()

	return ctx.AnswerCallbackQuery("📚 Media group example sent!").Send().Err()
}

// ================ POLL BUILDERS ================

func handlePollDemo(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("📊 Simple Poll", "poll_simple").
		Row().
		Text("🔙 Back", "start")

	return ctx.EditMessageText("📊 <b>Poll Builders Examples</b>\n\n" +
		"Create interactive polls with rich formatting:\n\n" +
		"📊 <b>Simple Poll</b> - Basic voting").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleSimplePoll(ctx *ctx.Context) error {
	// Create simple poll with basic options
	ctx.SendPoll("🌟 What's your favorite programming language?").
		Option(input.Choice("Go")).
		Option(input.Choice("Python")).
		Option(input.Choice("JavaScript")).
		Option(input.Choice("Rust")).
		MultipleAnswers().
		Anonymous().
		Send()

	return ctx.AnswerCallbackQuery("📊 Simple poll created!").Send().Err()
}
