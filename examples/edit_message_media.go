package main

import (
	"github.com/enetx/g"
	"github.com/enetx/g/fs"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

func main() {
	token := g.Slice[g.String](fs.NewFile("../.env").Read().Ok().Trim().Split("=")).Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Command to send initial media that can be edited later
	b.Command("media", func(ctx *ctx.Context) error {
		// Send a photo with interactive buttons to edit media
		result := ctx.SendPhoto("https://picsum.photos/800/600").
			Caption("🖼 <b>Original Photo</b>\n\nClick buttons below to change this media to different types:").
			HTML().
			Markup(
				keyboard.Inline().
					Text("🎥 Change to Video", "change_video").
					Text("🎵 Change to Audio", "change_audio").
					Row().
					Text("📄 Change to Document", "change_document").
					Text("🎬 Change to Animation", "change_animation")).
			Send()

		if result.IsErr() {
			return ctx.Reply("Failed to send initial media").Send().Err()
		}

		return nil
	})

	// Handle callback to change to video
	b.On.Callback.Equal("change_video", func(ctx *ctx.Context) error {
		result := ctx.EditMessageMedia(
			input.Video(file.Input("https://filesamples.com/samples/video/mp4/sample_640x360.mp4").Ok()).
				Caption("🎥 <b>Changed to Video!</b>\n\nThis message media has been updated using EditMessageMedia.\n\n📋 <b>Features:</b>\n• Media type changed from photo to video\n• Caption updated with HTML formatting\n• Inline keyboard preserved").
				HTML(),
		).
			Markup(
				keyboard.Inline().
					Text("🔙 Back to Photo", "back_photo").
					Text("🎵 To Audio", "change_audio").
					Row().
					Text("📄 To Document", "change_document")).
			Send()

		if result.IsErr() {
			return ctx.AnswerCallbackQuery(g.Format("Failed to change to video: {}", result.Err())).Send().Err()
		}

		return ctx.AnswerCallbackQuery("Changed to video! 🎥").Send().Err()
	})

	// Handle callback to change to audio
	b.On.Callback.Equal("change_audio", func(ctx *ctx.Context) error {
		result := ctx.EditMessageMedia(
			input.Audio(file.Input("https://www.soundjay.com/misc/sounds/bell-ringing-05.wav").Ok()).
				Caption("🎵 <b>Changed to Audio!</b>\n\nNow this is an audio file with a caption.").
				HTML().
				Title("Sample Bell Sound").
				Performer("SoundJay")).
			Markup(
				keyboard.Inline().
					Text("🔙 Back to Photo", "back_photo").
					Text("🎥 To Video", "change_video").
					Row().
					Text("📄 To Document", "change_document")).
			Send()

		if result.IsErr() {
			return ctx.AnswerCallbackQuery("Failed to change to audio").Send().Err()
		}

		return ctx.AnswerCallbackQuery("Changed to audio! 🎵").Send().Err()
	})

	// Handle callback to change to document
	b.On.Callback.Equal("change_document", func(ctx *ctx.Context) error {
		result := ctx.EditMessageMedia(
			input.Document(file.Input("https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf").Ok()).
				Caption("📄 <b>Changed to Document!</b>\n\nThis is now a PDF document instead of the original photo.").
				HTML()).
			Markup(
				keyboard.Inline().
					Text("🔙 Back to Photo", "back_photo").
					Text("🎥 To Video", "change_video").
					Row().
					Text("🎬 To Animation", "change_animation")).
			Send()

		if result.IsErr() {
			return ctx.AnswerCallbackQuery("Failed to change to document").Send().Err()
		}

		return ctx.AnswerCallbackQuery("Changed to document! 📄").Send().Err()
	})

	// Handle callback to change to animation
	b.On.Callback.Equal("change_animation", func(ctx *ctx.Context) error {
		result := ctx.EditMessageMedia(
			input.Animation(file.Input("https://media.giphy.com/media/3oEjI6SIIHBdRxXI40/giphy.gif").Ok()).
				Caption("🎬 <b>Changed to Animation!</b>\n\nNow showing an animated GIF instead!").
				HTML()).
			Markup(
				keyboard.Inline().
					Text("🔙 Back to Photo", "back_photo").
					Text("🎥 To Video", "change_video").
					Row().
					Text("📄 To Document", "change_document")).
			Send()

		if result.IsErr() {
			return ctx.AnswerCallbackQuery("Failed to change to animation").Send().Err()
		}

		return ctx.AnswerCallbackQuery("Changed to animation! 🎬").Send().Err()
	})

	// Handle callback to go back to original photo
	b.On.Callback.Equal("back_photo", func(ctx *ctx.Context) error {
		result := ctx.EditMessageMedia(
			input.Photo(file.Input("https://picsum.photos/800/600").Ok()).
				Caption("🖼 <b>Back to Original Photo</b>\n\nReturned to the original photo media.").
				HTML()).
			Markup(
				keyboard.Inline().
					Text("🎥 Change to Video", "change_video").
					Text("🎵 Change to Audio", "change_audio").
					Row().
					Text("📄 Change to Document", "change_document").
					Text("🎬 Change to Animation", "change_animation")).
			Send()

		if result.IsErr() {
			return ctx.AnswerCallbackQuery("Failed to restore photo").Send().Err()
		}

		return ctx.AnswerCallbackQuery("Restored original photo! 🖼").Send().Err()
	})

	// Command to edit media of replied message using file_id
	b.Command("editmedia", func(ctx *ctx.Context) error {
		if ctx.EffectiveMessage.ReplyToMessage == nil {
			return ctx.Reply("Please reply to a message with media to edit it").Send().Err()
		}

		args := ctx.Args()
		if len(args) < 2 {
			return ctx.Reply("Usage: /editmedia <type> <file_id_or_url>").Send().Err()
		}

		mediaType := args[0]
		mediaSource := args[1]
		replyMsg := ctx.EffectiveMessage.ReplyToMessage

		var media input.Media

		switch mediaType.Std() {
		case "photo":
			media = input.Photo(file.Input(mediaSource).Ok()).
				Caption("🖼 Photo updated via command").
				HTML()
		case "video":
			media = input.Video(file.Input(mediaSource).Ok()).
				Caption("🎥 Video updated via command").
				HTML()
		case "audio":
			media = input.Audio(file.Input(mediaSource).Ok()).
				Caption("🎵 Audio updated via command").
				HTML()
		case "document":
			media = input.Document(file.Input(mediaSource).Ok()).
				Caption("📄 Document updated via command").
				HTML()
		case "animation":
			media = input.Animation(file.Input(mediaSource).Ok()).
				Caption("🎬 Animation updated via command").
				HTML()
		default:
			return ctx.Reply("Invalid media type. Use: photo, video, audio, document, or animation").Send().Err()
		}

		if r := ctx.EditMessageMedia(media).MessageID(replyMsg.MessageId).Send(); r.IsErr() {
			return ctx.Reply(g.Format("Failed to edit media: {}", r.Err())).Send().Err()
		}

		return ctx.Reply("Media edited successfully! ✨").Send().Err()
	})

	// Command to demonstrate business connection media editing
	b.Command("bizmediademo", func(ctx *ctx.Context) error {
		return ctx.Reply(`
📋 <b>Business Media Demo</b>

To test business connection media editing:

1. Send media from your business account
2. Use: <code>/editbizmedia &lt;business_id&gt; &lt;type&gt; &lt;media&gt;</code>

<b>Example:</b>
<code>/editbizmedia abc123 photo https://example.com/new.jpg</code>

<i>Note: Requires proper business connection setup</i>
		`).HTML().Send().Err()
	})

	b.Command("editbizmedia", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if len(args) < 3 {
			return ctx.Reply("Usage: /editbizmedia <business_id> <type> <media>").Send().Err()
		}

		businessID := args[0]
		mediaType := args[1]
		mediaSource := args[2]

		if ctx.EffectiveMessage.ReplyToMessage == nil {
			return ctx.Reply("Please reply to a business message").Send().Err()
		}

		var media input.Media

		switch mediaType.Std() {
		case "photo":
			media = input.Photo(file.Input(mediaSource).Ok()).
				Caption("💼 Business photo updated!").
				HTML()
		case "video":
			media = input.Video(file.Input(mediaSource).Ok()).
				Caption("💼 Business video updated!").
				HTML()
		default:
			return ctx.Reply("Supported types: photo, video").Send().Err()
		}

		result := ctx.EditMessageMedia(media).
			Business(businessID).
			MessageID(ctx.EffectiveMessage.ReplyToMessage.MessageId).
			Send()

		if result.IsErr() {
			return ctx.Reply(g.Format("Failed to edit business media: {}", result.Err())).Send().Err()
		}

		return ctx.Reply("Business message media edited! 💼").Send().Err()
	})

	b.Polling().Start()
}
