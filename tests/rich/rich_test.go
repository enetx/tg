package rich_test

import (
	"testing"

	"github.com/enetx/tg/input"
	"github.com/enetx/tg/rich"
)

func TestBlock_Interface(t *testing.T) {
	blocks := []rich.Block{
		rich.Paragraph(testText("p")),
		rich.Heading(testText("h")),
		rich.Pre(testText("code")),
		rich.Footer(testText("f")),
		rich.Divider(),
		rich.Math("E = mc^2"),
		rich.Anchor("anchor"),
		rich.List(),
		rich.Quotation(),
		rich.PullQuotation(testText("q")),
		rich.Collage(),
		rich.Slideshow(),
		rich.Table(),
		rich.Details(testText("summary")),
		rich.Map(1, 2),
		rich.Animation(testAnimation()),
		rich.Audio(testAudio()),
		rich.Photo(testPhoto()),
		rich.Video(testVideo()),
		rich.VoiceNote(testVoiceNote()),
		rich.Thinking(testText("t")),
	}

	if len(blocks) != 21 {
		t.Errorf("Expected 21 block builders, got %d", len(blocks))
	}

	for i, block := range blocks {
		if block.Build() == nil {
			t.Errorf("Expected block %d to build a non-nil InputRichBlock", i)
		}
	}
}

func TestContent_MessageContentInterface(t *testing.T) {
	var content input.MessageContent = rich.NewContent(rich.New())
	if content.Build() == nil {
		t.Error("Expected Content to implement input.MessageContent")
	}
}
