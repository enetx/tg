package ctx_test

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/constants"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/rich"
	"github.com/enetx/tg/richtext"
)

func TestEphemeralSendParameters(t *testing.T) {
	tests := []struct {
		name string
		send func(*ctx.Context) g.Result[*gotgbot.Message]
	}{
		{"Message", func(c *ctx.Context) g.Result[*gotgbot.Message] {
			return c.SendMessage("private").ReceiverUser(123).CallbackQuery("query").ReplaceCallbackQueryMessage().Send()
		}},
		{"Reply", func(c *ctx.Context) g.Result[*gotgbot.Message] {
			return c.Reply("private").ReplaceCallbackQueryMessage().CallbackQuery("query").ReceiverUser(123).Send()
		}},
		{"Photo", func(c *ctx.Context) g.Result[*gotgbot.Message] {
			return c.SendPhoto(constants.FileIDPrefix + "photo").ReceiverUser(123).CallbackQuery("query").ReplaceCallbackQueryMessage().Send()
		}},
		{"Animation", func(c *ctx.Context) g.Result[*gotgbot.Message] {
			return c.SendAnimation(constants.FileIDPrefix + "animation").ReceiverUser(123).CallbackQuery("query").ReplaceCallbackQueryMessage().Send()
		}},
		{"Audio", func(c *ctx.Context) g.Result[*gotgbot.Message] {
			return c.SendAudio(constants.FileIDPrefix + "audio").ReceiverUser(123).CallbackQuery("query").ReplaceCallbackQueryMessage().Send()
		}},
		{"Document", func(c *ctx.Context) g.Result[*gotgbot.Message] {
			return c.SendDocument(constants.FileIDPrefix + "document").ReceiverUser(123).CallbackQuery("query").ReplaceCallbackQueryMessage().Send()
		}},
		{"LivePhoto", func(c *ctx.Context) g.Result[*gotgbot.Message] {
			return c.SendLivePhoto(constants.FileIDPrefix+"live", constants.FileIDPrefix+"photo").ReceiverUser(123).CallbackQuery("query").ReplaceCallbackQueryMessage().Send()
		}},
		{"Sticker", func(c *ctx.Context) g.Result[*gotgbot.Message] {
			return c.SendSticker(constants.FileIDPrefix + "sticker").ReceiverUser(123).CallbackQuery("query").ReplaceCallbackQueryMessage().Send()
		}},
		{"Video", func(c *ctx.Context) g.Result[*gotgbot.Message] {
			return c.SendVideo(constants.FileIDPrefix + "video").ReceiverUser(123).CallbackQuery("query").ReplaceCallbackQueryMessage().Send()
		}},
		{"VideoNote", func(c *ctx.Context) g.Result[*gotgbot.Message] {
			return c.SendVideoNote(constants.FileIDPrefix + "video_note").ReceiverUser(123).CallbackQuery("query").ReplaceCallbackQueryMessage().Send()
		}},
		{"Voice", func(c *ctx.Context) g.Result[*gotgbot.Message] {
			return c.SendVoice(constants.FileIDPrefix + "voice").ReceiverUser(123).CallbackQuery("query").ReplaceCallbackQueryMessage().Send()
		}},
		{"Contact", func(c *ctx.Context) g.Result[*gotgbot.Message] {
			return c.SendContact("+1234567", "Test").ReceiverUser(123).CallbackQuery("query").ReplaceCallbackQueryMessage().Send()
		}},
		{"Location", func(c *ctx.Context) g.Result[*gotgbot.Message] {
			return c.SendLocation(1, 2).ReceiverUser(123).CallbackQuery("query").ReplaceCallbackQueryMessage().Send()
		}},
		{"Venue", func(c *ctx.Context) g.Result[*gotgbot.Message] {
			return c.SendVenue(1, 2, "Title", "Address").ReceiverUser(123).CallbackQuery("query").ReplaceCallbackQueryMessage().Send()
		}},
		{"RichMessage", func(c *ctx.Context) g.Result[*gotgbot.Message] {
			return c.SendRichMessage(rich.New().Block(rich.Paragraph(richtext.New("private")))).
				CallbackQuery("query").ReplaceCallbackQueryMessage().ReceiverUser(123).Send()
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := apiContext(t, func(r *http.Request) {
				var params gotgbot.EphemeralMessageParameters
				if err := json.Unmarshal([]byte(r.FormValue("ephemeral_message_parameters")), &params); err != nil {
					t.Errorf("Invalid ephemeral_message_parameters: %v", err)
				}
				if params.ReceiverUserId != 123 || params.CallbackQueryId != "query" || !params.ReplaceCallbackQueryMessage {
					t.Errorf("Unexpected ephemeral parameters: %+v", params)
				}
				if r.FormValue("receiver_user_id") != "" || r.FormValue("callback_query_id") != "" {
					t.Error("Ephemeral parameters must be nested in ephemeral_message_parameters")
				}
			})
			if result := tt.send(c); result.IsErr() {
				t.Fatalf("Send failed: %v", result.Err())
			}
		})
	}
}

func TestSendMessageWithoutEphemeralParameters(t *testing.T) {
	c := apiContext(t, func(r *http.Request) {
		if _, ok := r.Form["ephemeral_message_parameters"]; ok {
			t.Error("Regular messages must omit ephemeral parameters")
		}
	})
	if result := c.SendMessage("regular").Send(); result.IsErr() {
		t.Fatal(result.Err())
	}
}

func TestEditEphemeralMessageRichContent(t *testing.T) {
	c := apiContext(t, func(r *http.Request) {
		if r.FormValue("chat_id") != "456" || r.FormValue("receiver_user_id") != "123" ||
			r.FormValue("ephemeral_message_id") != "555" {
			t.Error("Expected effective ephemeral message identifiers")
		}
		var message struct {
			Blocks []struct {
				Type string `json:"type"`
			} `json:"blocks"`
		}
		if err := json.Unmarshal([]byte(r.FormValue("rich_message")), &message); err != nil {
			t.Errorf("Invalid rich_message: %v", err)
		}
		if len(message.Blocks) != 1 || message.Blocks[0].Type != "paragraph" {
			t.Error("Expected rich paragraph content")
		}
		if _, ok := r.Form["text"]; ok {
			t.Error("Empty text must be omitted when editing rich content")
		}
	})
	message := rich.New().Block(rich.Paragraph(richtext.New("updated")))
	if result := c.EditEphemeralMessageText("").RichMessage(message).Send(); result.IsErr() {
		t.Fatal(result.Err())
	}
}

func TestEditEphemeralMessageTextPreserved(t *testing.T) {
	c := apiContext(t, func(r *http.Request) {
		if r.FormValue("text") != "updated" {
			t.Error("Expected the edited text to be sent")
		}
	})
	if result := c.EditEphemeralMessageText("updated").Send(); result.IsErr() {
		t.Fatal(result.Err())
	}
}

func TestEditEphemeralMessageCaptionAboveMedia(t *testing.T) {
	c := apiContext(t, func(r *http.Request) {
		if r.FormValue("caption") != "updated" || r.FormValue("show_caption_above_media") != "true" {
			t.Error("Expected caption and show_caption_above_media")
		}
	})
	if result := c.EditEphemeralMessageCaption("updated").ShowCaptionAboveMedia().Send(); result.IsErr() {
		t.Fatal(result.Err())
	}
}

func TestEditEphemeralMessageMediaUpload(t *testing.T) {
	c := apiContext(t, func(r *http.Request) {
		var media struct {
			Type  string `json:"type"`
			Media string `json:"media"`
		}
		if err := json.Unmarshal([]byte(r.FormValue("media")), &media); err != nil {
			t.Errorf("Invalid media: %v", err)
		}
		if media.Type != "document" || !strings.HasPrefix(media.Media, "attach://") {
			t.Errorf("Expected an attached document, got %+v", media)
			return
		}
		f, header, err := r.FormFile(strings.TrimPrefix(media.Media, "attach://"))
		if err != nil {
			t.Error(err)
			return
		}
		defer f.Close()
		data, err := io.ReadAll(f)
		if err != nil || string(data) != "new document" || header.Filename != "test.txt" {
			t.Errorf("Unexpected upload: %q, %s, %v", data, header.Filename, err)
		}
	})
	document := input.Document(file.InputFile{Doc: gotgbot.InputFileByReader("test.txt", strings.NewReader("new document"))})
	if result := c.EditEphemeralMessageMedia(document).Send(); result.IsErr() {
		t.Fatal(result.Err())
	}
}
