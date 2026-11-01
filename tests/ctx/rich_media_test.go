package ctx_test

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/rich"
)

func TestSendRichMessageDocumentUpload(t *testing.T) {
	for _, method := range []string{"Send", "EditEphemeral"} {
		for _, html := range []bool{false, true} {
			name := "Block"
			if html {
				name = "HTML"
			}
			t.Run(method+"/"+name, func(t *testing.T) {
				c := apiContext(t, func(r *http.Request) {
					var message struct {
						HTML   string `json:"html"`
						Blocks []struct {
							Document struct {
								Media string `json:"media"`
							} `json:"document"`
						} `json:"blocks"`
						Media []struct {
							ID    string `json:"id"`
							Media struct {
								Media string `json:"media"`
							} `json:"media"`
						} `json:"media"`
					}
					if err := json.Unmarshal([]byte(r.FormValue("rich_message")), &message); err != nil {
						t.Error(err)
						return
					}
					var attachment string
					if html {
						if len(message.Media) != 1 || message.Media[0].ID != "report" ||
							!strings.Contains(message.HTML, "tg://document?id=report") {
							t.Errorf("Expected a document registered for the HTML content: %+v", message)
							return
						}
						attachment = message.Media[0].Media.Media
					} else {
						if len(message.Blocks) != 1 {
							t.Error("Expected one document block")
							return
						}
						attachment = message.Blocks[0].Document.Media
					}
					f, header, err := r.FormFile(strings.TrimPrefix(attachment, "attach://"))
					if err != nil {
						t.Error(err)
						return
					}
					defer f.Close()
					data, err := io.ReadAll(f)
					if err != nil || string(data) != "report content" || header.Filename != "report.txt" {
						t.Errorf("Unexpected document upload: %q, %s, %v", data, header.Filename, err)
					}
				})
				document := input.Document(file.InputFile{
					Doc: gotgbot.InputFileByReader("report.txt", strings.NewReader("report content")),
				})
				message := rich.New()
				if html {
					message.HTML(`<tg-document src="tg://document?id=report"></tg-document>`).Media("report", document)
				} else {
					message.Block(rich.Document(document))
				}
				if method == "EditEphemeral" {
					if result := c.EditEphemeralMessageText("").RichMessage(message).Send(); result.IsErr() {
						t.Fatal(result.Err())
					}
				} else if result := c.SendRichMessage(message).Send(); result.IsErr() {
					t.Fatal(result.Err())
				}
			})
		}
	}
}

func TestSendRichMessageNestedDocumentUploads(t *testing.T) {
	c := apiContext(t, func(r *http.Request) {
		if len(r.MultipartForm.File) != 2 {
			t.Errorf("Expected two separate attachments, got %d", len(r.MultipartForm.File))
			return
		}
		seen := make(map[string]bool)
		for name := range r.MultipartForm.File {
			f, _, err := r.FormFile(name)
			if err != nil {
				t.Error(err)
				continue
			}
			data, err := io.ReadAll(f)
			f.Close()
			if err != nil {
				t.Error(err)
			}
			seen[string(data)] = true
			if !strings.Contains(r.FormValue("rich_message"), "attach://"+name) {
				t.Errorf("Missing JSON reference to attachment %s", name)
			}
		}
		if !seen["first"] || !seen["second"] {
			t.Errorf("Expected both nested documents, got %v", seen)
		}
	})
	document := func(text string) *rich.BlockDocument {
		return rich.Document(input.Document(file.InputFile{
			Doc: gotgbot.InputFileByReader(text+".txt", strings.NewReader(text)),
		}))
	}
	message := rich.New().Block(
		rich.Details(nil).Block(rich.Quotation(document("first"))),
		rich.List(rich.ListItem(rich.Details(nil).Block(document("second")))),
	)
	if result := c.SendRichMessage(message).Send(); result.IsErr() {
		t.Fatal(result.Err())
	}
}
