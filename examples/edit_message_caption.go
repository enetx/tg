package main

import (
	"github.com/enetx/g"
	"github.com/enetx/g/fs"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

func main() {
	token := g.Slice[g.String](fs.NewFile("../.env").Read().Ok().Trim().Split("=")).Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Command to send a photo with caption that can be edited later
	b.Command("photo", func(ctx *ctx.Context) error {
		// Send a photo with initial caption
		result := ctx.SendPhoto("https://picsum.photos/800/600").
			Caption("🖼 <b>Original Caption</b>\n\nThis is the original caption for this beautiful photo.").
			HTML().
			Markup(
				keyboard.Inline().
					Text("✏️ Edit Caption", "edit_caption").
					Text("🔄 Change Style", "change_style").
					Row().
					Text("📝 Add Description", "add_desc").
					Text("🌟 Show Above", "show_above")).
			Send()

		if result.IsErr() {
			return ctx.Reply("Failed to send photo").Send().Err()
		}

		return nil
	})

	// Handle callback queries for caption editing
	b.On.Callback.Equal("edit_caption", func(ctx *ctx.Context) error {
		newCaption := "✨ <b>Updated Caption!</b>\n\n📸 This caption has been <i>edited</i> using EditMessageCaption!\n\n🎯 <u>Features demonstrated:</u>\n• HTML formatting\n• Bold and italic text\n• Underlined text"

		result := ctx.EditMessageCaption(g.String(newCaption)).
			HTML().
			Markup(
				keyboard.Inline().
					Text("🔙 Original", "original_caption").
					Text("🎨 Markdown", "markdown_caption").
					Row().
					Text("📝 Plain Text", "plain_caption")).
			Send()

		if result.IsErr() {
			return ctx.AnswerCallbackQuery("Failed to edit caption").Send().Err()
		}

		return ctx.AnswerCallbackQuery("Caption updated with HTML formatting! ✨").Send().Err()
	})

	b.On.Callback.Equal("change_style", func(ctx *ctx.Context) error {
		newCaption := "*Styled Caption*\n\n_This caption uses_ **Markdown** formatting\\!\n\n• Feature 1\n• Feature 2\n• Feature 3\n\n`Code example: ctx.EditMessageCaption()`"

		result := ctx.EditMessageCaption(g.String(newCaption)).
			Markdown().
			Markup(
				keyboard.Inline().
					Text("🔙 Back to HTML", "edit_caption").
					Text("📱 Plain", "plain_caption")).
			Send()

		if result.IsErr() {
			return ctx.AnswerCallbackQuery("Failed to change style").Send().Err()
		}

		return ctx.AnswerCallbackQuery("Switched to Markdown formatting! 📝").Send().Err()
	})

	b.On.Callback.Equal("add_desc", func(ctx *ctx.Context) error {
		newCaption := `📷 <b>Enhanced Photo Description</b>

🎨 <b>Artistic Details:</b>
• Camera: Professional DSLR
• Resolution: 800x600 pixels
• Style: Modern photography
• Source: Lorem Picsum

💡 <b>Technical Info:</b>
• g.Format: JPEG
• Color: Full spectrum
• Quality: High definition

🔗 <b>Edit Features:</b>
• Caption modification ✅
• HTML formatting ✅
• Inline keyboard ✅
• Multiple styles ✅`

		result := ctx.EditMessageCaption(g.String(newCaption)).
			HTML().
			Markup(
				keyboard.Inline().
					Text("🎯 Show Above Media", "show_above").
					Text("🔄 Reset", "original_caption")).
			Send()

		if result.IsErr() {
			return ctx.AnswerCallbackQuery("Failed to add description").Send().Err()
		}

		return ctx.AnswerCallbackQuery("Added detailed description! 📝").Send().Err()
	})

	b.On.Callback.Equal("show_above", func(ctx *ctx.Context) error {
		newCaption := `⬆️ <b>Caption Above Media!</b>

🎯 This caption is now displayed <b>above</b> the photo instead of below it.

✨ <i>Perfect for:</i>
• Important announcements
• Photo titles
• Context that should be seen first

🔧 <b>Technical note:</b> This feature works for animation, photo and video messages.`

		result := ctx.EditMessageCaption(g.String(newCaption)).
			HTML().
			ShowCaptionAboveMedia().
			Markup(
				keyboard.Inline().
					Text("⬇️ Show Below", "edit_caption").
					Text("🔄 Reset", "original_caption")).
			Send()

		if result.IsErr() {
			return ctx.AnswerCallbackQuery("Failed to show caption above").Send().Err()
		}

		return ctx.AnswerCallbackQuery("Caption moved above media! ⬆️").Send().Err()
	})

	b.On.Callback.Equal("original_caption", func(ctx *ctx.Context) error {
		originalCaption := "🖼 <b>Original Caption</b>\n\nThis is the original caption for this beautiful photo."

		result := ctx.EditMessageCaption(g.String(originalCaption)).
			HTML().
			Markup(
				keyboard.Inline().
					Text("✏️ Edit Caption", "edit_caption").
					Text("🔄 Change Style", "change_style").
					Row().
					Text("📝 Add Description", "add_desc").
					Text("🌟 Show Above", "show_above")).
			Send()

		if result.IsErr() {
			return ctx.AnswerCallbackQuery("Failed to restore original").Send().Err()
		}

		return ctx.AnswerCallbackQuery("Restored original caption! 🔄").Send().Err()
	})

	b.On.Callback.Equal("plain_caption", func(ctx *ctx.Context) error {
		plainCaption := "Simple plain text caption without any formatting. This demonstrates how the caption looks without HTML or Markdown parsing."

		result := ctx.EditMessageCaption(g.String(plainCaption)).
			Markup(
				keyboard.Inline().
					Text("🎨 Add HTML", "edit_caption").
					Text("📝 Add Markdown", "change_style").
					Row().
					Text("🔙 Original", "original_caption")).
			Send()

		if result.IsErr() {
			return ctx.AnswerCallbackQuery("Failed to set plain text").Send().Err()
		}

		return ctx.AnswerCallbackQuery("Switched to plain text! 📄").Send().Err()
	})

	// Command to edit caption of replied message
	b.Command("editcaption", func(ctx *ctx.Context) error {
		if ctx.EffectiveMessage.ReplyToMessage == nil {
			return ctx.Reply("Please reply to a message with media to edit its caption").Send().Err()
		}

		args := ctx.Args()
		if len(args) == 0 {
			return ctx.Reply("Usage: /editcaption <new_caption>").Send().Err()
		}

		newCaption := args.Join(" ")
		replyMsg := ctx.EffectiveMessage.ReplyToMessage

		result := ctx.EditMessageCaption(newCaption).
			MessageID(replyMsg.MessageId).
			HTML().
			Send()

		if result.IsErr() {
			return ctx.Reply(g.Format("Failed to edit caption: {}", result.Err())).Send().Err()
		}

		return ctx.Reply("Caption edited successfully! ✨").Send().Err()
	})

	// Command to demonstrate business connection caption editing
	b.Command("bizphoto", func(ctx *ctx.Context) error {
		return ctx.Reply(`
📋 <b>Business Caption Demo</b>

To test business connection caption editing:

1. Send a photo from your business account
2. Use: <code>/editbizcaption &lt;business_id&gt; &lt;caption&gt;</code>

<i>Note: Requires proper business connection setup</i>
		`).HTML().Send().Err()
	})

	b.Command("editbizcaption", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if len(args) < 2 {
			return ctx.Reply("Usage: /editbizcaption <business_id> <caption>").Send().Err()
		}

		businessID := args[0]
		caption := args.Iter().Skip(1).Collect().Slice().Join(" ")

		if ctx.EffectiveMessage.ReplyToMessage == nil {
			return ctx.Reply("Please reply to a business message").Send().Err()
		}

		result := ctx.EditMessageCaption(caption).
			Business(businessID).
			MessageID(ctx.EffectiveMessage.ReplyToMessage.MessageId).
			HTML().
			Send()

		if result.IsErr() {
			return ctx.Reply(g.Format("Failed to edit business caption: {}", result.Err())).Send().Err()
		}

		return ctx.Reply("Business message caption edited! 💼").Send().Err()
	})

	b.Polling().Start()
}
