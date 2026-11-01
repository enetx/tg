package rich_test

import (
	"encoding/json"
	"testing"

	"github.com/enetx/tg/constants"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/rich"
)

func TestBlockDocument(t *testing.T) {
	media := input.Document(file.Input(constants.FileIDPrefix + "document").Ok())
	block := rich.Document(nil).Document(media).Document(nil).Caption(rich.NewCaption(testText("caption")))
	data, err := json.Marshal(block.Build())
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `{"type":"document","document":{"type":"document","media":"document"},"caption":{"text":"caption"}}` {
		t.Errorf("Unexpected document block JSON: %s", data)
	}
}
