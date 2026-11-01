package ctx_test

import (
	"net/http"
	"testing"

	"github.com/enetx/tg/rich"
	"github.com/enetx/tg/richtext"
	"github.com/enetx/tg/types/roles"
)

func TestMessageDraftStopParameters(t *testing.T) {
	for _, richDraft := range []bool{false, true} {
		name := "Text"
		if richDraft {
			name = "Rich"
		}
		t.Run(name, func(t *testing.T) {
			c := apiContext(t, func(r *http.Request) {
				if r.FormValue("draft_id") != "42" || r.FormValue("can_stop") != "true" ||
					r.FormValue("keep_on_stop") != "true" {
					t.Error("Expected draft ID and stop parameters")
				}
			})
			if richDraft {
				message := rich.New().Block(rich.Paragraph(richtext.New("partial")))
				if result := c.SendRichMessageDraft(42, message).CanStop().KeepOnStop().Send(); result.IsErr() {
					t.Fatal(result.Err())
				}
			} else if result := c.SendMessageDraft(42).Text("partial").CanStop().KeepOnStop().Send(); result.IsErr() {
				t.Fatal(result.Err())
			}
		})
	}
}

func TestPromoteChatMemberWelcomeMessages(t *testing.T) {
	c := apiContext(t, func(r *http.Request) {
		if r.FormValue("can_send_welcome_messages") != "true" || r.FormValue("can_manage_tags") != "true" {
			t.Error("Expected welcome message and tag management rights")
		}
	})
	if result := c.PromoteChatMember(123).Roles(roles.SendWelcomeMessages, roles.ManageTags).Send(); result.IsErr() {
		t.Fatal(result.Err())
	}
}
