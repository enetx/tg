package rich_test

import (
	"encoding/json"
	"testing"

	"github.com/enetx/tg/rich"
)

func TestExpandableQuotation(t *testing.T) {
	block := rich.ExpandableQuotation(testText("old")).Text(testText("quote")).Credit(testText("source"))
	data, err := json.Marshal(block.Build())
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `{"type":"expandable_blockquote","text":"quote","credit":"source"}` {
		t.Errorf("Unexpected expandable quotation JSON: %s", data)
	}
}
