package main

import (
	"fmt"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/inline"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/types/updates"
)

func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Start command with inline query example
	b.Command("start", func(ctx *ctx.Context) error {
		return ctx.Reply("Send an inline query (@your_bot_username query) to test different result types!\n\n" +
			"Try these queries:\n" +
			"• article - Text article with thumbnail\n" +
			"• photo - Photo result\n" +
			"• video - Video result\n" +
			"• audio - Audio result\n" +
			"• location - Location result\n" +
			"• venue - Venue result\n" +
			"• contact - Contact result\n" +
			"• game - Game result\n" +
			"• gif - GIF animation\n" +
			"• cached - Cached content\n" +
			"• keyboard - g.Result with inline keyboard\n" +
			"• content - Different message content types").Send().Err()
	})

	// Handle callback queries from inline keyboards
	b.On.Callback.Prefix("opt", func(ctx *ctx.Context) error {
		data := ctx.Update.CallbackQuery.Data
		return ctx.AnswerCallbackQuery(g.Format("You selected: {}", data)).Alert().Send().Err()
	})

	// Handle inline queries with different result types
	b.On.Inline.Any(func(ctx *ctx.Context) error {
		query := ctx.Update.InlineQuery.Query
		queryID := ctx.Update.InlineQuery.Id

		fmt.Printf("Debug: Received inline query: '%s' (ID: %s)\n", query, queryID)

		var results g.Slice[inline.QueryResult]

		switch {
		case query == "article":
			// Article result example
			article := inline.Article(
				"article_1",
				"Sample Article",
				input.Text("This is the content of the article"),
			).
				Description("A sample article with thumbnail").
				ThumbnailURL("https://via.placeholder.com/150").
				URL("https://example.com")

			results.Push(article)

		case query == "photo":
			// Photo result example
			photo := inline.NewPhoto(
				"photo_1",
				"https://via.placeholder.com/800x600/FF0000/FFFFFF?text=Sample+Photo",
				"https://via.placeholder.com/150x150/FF0000/FFFFFF?text=Thumb",
			).
				Title("Sample Photo").
				Description("A red sample photo").
				Caption("This is a sample photo caption").
				Size(800, 600)

			results.Push(photo)

		case query == "video":
			// Video result example
			video := inline.NewVideo(
				"video_1",
				"https://sample-videos.com/zip/10/mp4/SampleVideo_1280x720_1mb.mp4",
				"video/mp4",
				"https://via.placeholder.com/150x150/0000FF/FFFFFF?text=Video",
				"Sample Video",
			).
				Description("A sample video").
				Caption("Sample video with caption").
				Duration(30).
				Size(1280, 720)

			results.Push(video)

		case query == "audio":
			// Audio result example
			audio := inline.NewAudio(
				"audio_1",
				"https://www.soundjay.com/misc/sounds/bell-ringing-05.mp3",
				"Bell Sound",
			).
				Performer("Sound Effects").
				Duration(5).
				Caption("Bell ringing sound effect")

			results.Push(audio)

		case query == "location":
			// Location result example
			location := inline.NewLocation(
				"location_1",
				40.7128, -74.0060, // NYC coordinates
				"New York City",
			).
				ThumbnailURL("https://via.placeholder.com/150x150/00FF00/FFFFFF?text=NYC")

			results.Push(location)

		case query == "venue":
			// Venue result example
			venue := inline.NewVenue(
				"venue_1",
				40.7589, -73.9851, // Times Square coordinates
				"Times Square",
				"Manhattan, NY 10036, USA",
			).
				GooglePlaceID("ChIJmQJIxlVYwokRLgeuocVOGVU").
				ThumbnailURL("https://via.placeholder.com/150x150/FFFF00/000000?text=TS")

			results.Push(venue)

		case query == "contact":
			// Contact result example
			contact := inline.NewContact(
				"contact_1",
				"+1234567890",
				"John",
			).
				LastName("Doe").
				VCard("BEGIN:VCARD\nVERSION:3.0\nFN:John Doe\nTEL:+1234567890\nEND:VCARD").
				ThumbnailURL("https://via.placeholder.com/150x150/800080/FFFFFF?text=JD")

			results.Push(contact)

		case query == "game":
			// Game result example
			game := inline.NewGame(
				"game_1",
				"my_game_short_name",
			).
				Markup(keyboard.Inline().
					Row().
					Game("🎮 Play Game").
					Row().
					URL("🔗 Learn More", "https://example.com/game-info"))

			results.Push(game)

		case query == "gif":
			// GIF result example
			gif := inline.NewGif(
				"gif_1",
				"https://media.giphy.com/media/l0MYt5jPR6QX5pnqM/giphy.gif",
				"https://via.placeholder.com/150x150/FF69B4/FFFFFF?text=GIF",
			).
				Title("Funny GIF").
				Caption("This is a funny animated GIF").
				Size(480, 270).
				Duration(3)

			results.Push(gif)

		case query == "cached":
			// Cached content examples (requires existing file_id)
			cachedPhoto := inline.NewCachedPhoto(
				"cached_photo_1",
				"AgACAgIAAxkBAAIBYmF...", // Replace with actual file_id
			).
				Title("Cached Photo").
				Description("A cached photo from previous upload").
				Caption("This photo was cached")

			cachedSticker := inline.NewCachedSticker(
				"cached_sticker_1",
				"CAACAgIAAxkBAAIBY2F...", // Replace with actual sticker file_id
			)

			results.Push(cachedPhoto, cachedSticker)

		case query == "keyboard":
			// Simple inline keyboard for testing
			kbrd := keyboard.Inline().
				Text("Option 1", "opt1").
				Text("Option 2", "opt2")

			// Create article with inline keyboard
			article := inline.Article(
				"article_keyboard",
				"🎹 Article with Keyboard",
				input.Text("Click the buttons below:"),
			).
				Description("Interactive article with buttons").
				ThumbnailURL("https://via.placeholder.com/150x150/4169E1/FFFFFF?text=KB").
				Markup(kbrd)

			results.Push(article)

		case query == "content":
			// Different input message content types

			// Text content
			textArticle := inline.Article(
				"text_content",
				"Text Content",
				input.Text("*Bold text* and _italic text_").Markdown(),
			)

			// Location content
			locationArticle := inline.Article(
				"location_content",
				"Send Location",
				input.Location(51.5074, -0.1278), // London coordinates
			).
				Description("This will send a location instead of text")

			// Venue content
			venueArticle := inline.Article(
				"venue_content",
				"Send Venue",
				input.Venue(
					48.8566, 2.3522, // Paris coordinates
					"Eiffel Tower",
					"Champ de Mars, 5 Avenue Anatole France, 75007 Paris, France",
				),
			).
				Description("This will send a venue instead of text")

			// Contact content
			contactArticle := inline.Article(
				"contact_content",
				"Send Contact",
				input.Contact("+33123456789", "Pierre"),
			).
				Description("This will send a contact instead of text")

			results.Push(textArticle, locationArticle, venueArticle, contactArticle)

		default:
			// Default results for empty or unknown query
			defaultArticle := inline.Article(
				"help",
				"Inline Query Help",
				input.Text(
					"Try these queries: article, photo, video, audio, location, venue, contact, game, gif, cached, keyboard, content",
				),
			).
				Description("Available inline query examples").
				ThumbnailURL("https://via.placeholder.com/150x150/4169E1/FFFFFF?text=Help")

			results.Push(defaultArticle)
		}

		// Answer the inline query
		fmt.Printf("Debug: Sending %d results for query '%s'\n", results.Len(), query)

		return ctx.AnswerInlineQuery(g.String(queryID)).
			Results(results...).
			CacheFor(0 * time.Second).
			Send().Err()
	})

	b.Polling().AllowedUpdates(updates.All...).Start()
}
