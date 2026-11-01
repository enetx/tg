package ctx_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/tg/ctx"
)

type apiBot struct {
	mockBot
	raw *gotgbot.Bot
}

func (b *apiBot) Raw() *gotgbot.Bot { return b.raw }

func apiContext(t *testing.T, check func(*http.Request)) *ctx.Context {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(1 << 20); err != nil {
			t.Errorf("ParseMultipartForm failed: %v", err)
		}
		if r.MultipartForm != nil {
			defer r.MultipartForm.RemoveAll()
		}

		check(r)
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(r.URL.Path, "Draft") || strings.Contains(r.URL.Path, "editEphemeral") ||
			strings.Contains(r.URL.Path, "promoteChatMember") {
			fmt.Fprint(w, `{"ok":true,"result":true}`)
		} else {
			fmt.Fprint(w, `{"ok":true,"result":{"message_id":789,"date":1,"chat":{"id":456,"type":"supergroup"}}}`)
		}
	}))
	t.Cleanup(server.Close)

	b := &apiBot{raw: &gotgbot.Bot{
		Token: "123:test",
		BotClient: &gotgbot.BaseBotClient{
			Client:             *server.Client(),
			DefaultRequestOpts: &gotgbot.RequestOpts{APIURL: server.URL, Timeout: 5 * time.Second},
		},
	}}
	message := &gotgbot.Message{
		MessageId:          789,
		EphemeralMessageId: 555,
		ReceiverUser:       &gotgbot.User{Id: 123},
		Chat:               gotgbot.Chat{Id: 456, Type: "supergroup"},
	}

	return ctx.New(b, &ext.Context{
		Update:           &gotgbot.Update{Message: message},
		EffectiveChat:    &message.Chat,
		EffectiveMessage: message,
	})
}
