package rich_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/rich"
)

func TestNew(t *testing.T) {
	message := rich.New()
	if message == nil {
		t.Fatal("Expected New to create a Message builder")
	}

	built := message.Std()
	if built.Blocks != nil || built.Media != nil {
		t.Error("Expected a new Message to carry no blocks and no media")
	}
}

func TestMessage_Block(t *testing.T) {
	message := rich.New()

	result := message.Block(rich.Paragraph(testText("Hello")), nil, rich.Divider())
	if result != message {
		t.Error("Expected Block to return same Message instance")
	}

	if len(result.Std().Blocks) != 2 {
		t.Error("Expected Block to skip nil blocks")
	}
}

func TestMessage_HTML(t *testing.T) {
	message := rich.New()

	result := message.HTML(g.String("<b>bold</b>"))
	if result != message {
		t.Error("Expected HTML to return same Message instance")
	}

	if result.Std().Html != "<b>bold</b>" {
		t.Error("Expected Html to be set correctly")
	}
}

func TestMessage_Markdown(t *testing.T) {
	message := rich.New()

	result := message.Markdown(g.String("*bold*"))
	if result != message {
		t.Error("Expected Markdown to return same Message instance")
	}

	if result.Std().Markdown != "*bold*" {
		t.Error("Expected Markdown to be set correctly")
	}
}

func TestMessage_Media(t *testing.T) {
	message := rich.New()

	result := message.Media(g.String("photo-1"), testPhoto()).Media(g.String("skipped"), nil)
	if result != message {
		t.Error("Expected Media to return same Message instance")
	}

	media := result.Std().Media
	if len(media) != 1 {
		t.Fatal("Expected Media to skip nil media")
	}

	if media[0].Id != "photo-1" || media[0].Media == nil {
		t.Error("Expected Media fields to be set correctly")
	}
}

func TestMessage_RTL(t *testing.T) {
	message := rich.New()

	result := message.RTL()
	if result != message {
		t.Error("Expected RTL to return same Message instance")
	}

	if !result.Std().IsRtl {
		t.Error("Expected IsRtl to be set correctly")
	}
}

func TestMessage_SkipEntityDetection(t *testing.T) {
	message := rich.New()

	result := message.SkipEntityDetection()
	if result != message {
		t.Error("Expected SkipEntityDetection to return same Message instance")
	}

	if !result.Std().SkipEntityDetection {
		t.Error("Expected SkipEntityDetection to be set correctly")
	}
}

func TestMessage_StdNil(t *testing.T) {
	var message *rich.Message

	built := message.Std()
	if built.Blocks != nil || built.Html != "" {
		t.Error("Expected Std on a nil Message to return the zero value")
	}
}
