package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/input"
)

/*
Example: TG Input Builders

This example demonstrates the complete input builder system for all gotgbot.Input* types:

Input Builder Categories:
1. InputMedia - For media groups and file operations
2. InputMessageContent - For inline query result content
3. InputPaidMedia - For paid media content  
4. InputProfilePhoto - For profile photo operations
5. InputStoryContent - For story posting
6. Other Input types - Stickers, checklists, poll options

All builders follow the same pattern as inline/ builders:
- Embedded gotgbot struct pattern
- Fluent API with method chaining
- HTML()/Markdown() convenience methods where applicable
- Build() method returns appropriate gotgbot interface
*/

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Example: Creating Input Builders - demonstrates all builder types
	b.Command("builders", func(ctx *ctx.Context) error {
		// 1. InputMedia builders
		photoBuilder := input.Photo("https://picsum.photos/800/600").
			Caption("Beautiful photo with HTML formatting").
			HTML().
			HasSpoiler()

		videoBuilder := input.Video("https://example.com/sample.mp4").
			Caption("Sample video with **markdown**").
			Markdown().
			Duration(120).
			Size(1920, 1080).
			SupportsStreaming()

		audioBuilder := input.Audio("https://example.com/song.mp3").
			Caption("🎵 Music track").
			Title("Sample Song").
			Performer("Artist Name").
			Duration(180)

		// 2. InputMessageContent builders
		textContent := input.Text("Hello from <b>input builder</b>!").HTML()
		locationContent := input.Location(40.7128, -74.0060) // NYC
		venueContent := input.Venue(40.7128, -74.0060, "Times Square", "New York, NY")
		contactContent := input.Contact("+1234567890", "John Doe")

		// 3. InputProfilePhoto builders  
		staticPhoto := input.StaticPhoto("photo_file_id_here")
		animatedPhoto := input.AnimatedPhoto("video_file_id_here").MainFrameTimestamp(2.5)

		// 4. InputStoryContent builders
		storyPhoto := input.StoryPhoto("story_photo_file_id") 
		storyVideo := input.StoryVideo("story_video_file_id").Duration(15.0).IsAnimation()

		// 5. InputPaidMedia builders
		paidPhoto := input.PaidPhoto("https://example.com/premium.jpg")
		paidVideo := input.PaidVideo("https://example.com/premium.mp4").Duration(60)

		// Build all types to show they work
		_ = photoBuilder.Build()
		_ = videoBuilder.Build() 
		_ = audioBuilder.Build()
		_ = textContent.Build()
		_ = locationContent.Build()
		_ = venueContent.Build()
		_ = contactContent.Build()
		_ = staticPhoto.Build()
		_ = animatedPhoto.Build()
		_ = storyPhoto.Build()
		_ = storyVideo.Build()
		_ = paidPhoto.Build()
		_ = paidVideo.Build()

		return ctx.SendMessage("✅ All Input Builders created successfully!\n\n" +
			"Created builders for:\n" +
			"📸 InputMedia: Photo, Video, Audio\n" +
			"💬 InputMessageContent: Text, Location, Venue, Contact\n" +
			"👤 InputProfilePhoto: Static, Animated\n" +
			"📱 InputStoryContent: Photo, Video\n" +
			"💰 InputPaidMedia: Photo, Video\n\n" +
			"All builders use fluent API with method chaining!").Send().Err()
	})

	// Example: Usage pattern for developers
	b.Command("pattern", func(ctx *ctx.Context) error {
		message := "🏗️ **Input Builder Pattern**\n\n" +
			"```go\n" +
			"// 1. Create builder\n" +
			"builder := input.Photo(\"url\").\n" +
			"  Caption(\"text\").\n" +
			"  HTML().\n" +
			"  HasSpoiler()\n\n" +
			"// 2. Build to gotgbot type\n" +
			"media := builder.Build()\n\n" +
			"// 3. Use with TG methods\n" +
			"ctx.SendPhoto(media)\n" +
			"```\n\n" +
			"✨ **Benefits:**\n" +
			"• Fluent API with method chaining\n" +
			"• HTML()/Markdown() convenience methods\n" +
			"• Type-safe builder interfaces\n" +
			"• Consistent pattern across all Input* types"

		return ctx.SendMessage(String(message)).Markdown().Send().Err()
	})

	b.Polling().Start()
}