package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestQuotation(t *testing.T) {
	block := rich.Quotation(rich.Paragraph(testText("Quoted")), nil)
	if block == nil {
		t.Fatal("Expected Quotation to create a BlockQuotation builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockBlockQuotation)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockBlockQuotation")
	}

	if len(built.Blocks) != 1 {
		t.Error("Expected Quotation to skip nil blocks")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeBlockquote {
		t.Error("Expected block type to be blockquote")
	}
}

func TestBlockQuotation_Block(t *testing.T) {
	block := rich.Quotation()

	result := block.Block(rich.Paragraph(testText("a")), rich.Divider())
	if result != block {
		t.Error("Expected Block to return same BlockQuotation instance")
	}

	built := result.Build().(gotgbot.InputRichBlockBlockQuotation)
	if len(built.Blocks) != 2 {
		t.Error("Expected Block to append blocks")
	}
}

func TestBlockQuotation_Credit(t *testing.T) {
	block := rich.Quotation()

	result := block.Credit(testText(testCredit))
	if result != block {
		t.Error("Expected Credit to return same BlockQuotation instance")
	}

	built := result.Build().(gotgbot.InputRichBlockBlockQuotation)
	if string(built.Credit.(gotgbot.RichTextString)) != testCredit.Std() {
		t.Error("Expected Credit to be set correctly")
	}
}

func TestBlockQuotation_Empty(t *testing.T) {
	built := rich.Quotation().Build().(gotgbot.InputRichBlockBlockQuotation)
	if built.Blocks != nil {
		t.Error("Expected an empty quotation to carry no blocks")
	}
}
