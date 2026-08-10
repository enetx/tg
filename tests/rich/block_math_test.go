package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/rich"
)

func TestMath(t *testing.T) {
	block := rich.Math(g.String("E = mc^2"))
	if block == nil {
		t.Fatal("Expected Math to create a BlockMath builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockMathematicalExpression)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockMathematicalExpression")
	}

	if built.Expression != "E = mc^2" {
		t.Error("Expected Expression to be set correctly")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeMathematicalExpression {
		t.Error("Expected block type to be mathematical_expression")
	}
}

func TestBlockMath_Expression(t *testing.T) {
	block := rich.Math(g.String("a"))

	result := block.Expression(g.String("b"))
	if result != block {
		t.Error("Expected Expression to return same BlockMath instance")
	}

	built := result.Build().(gotgbot.InputRichBlockMathematicalExpression)
	if built.Expression != "b" {
		t.Error("Expected Expression to be set correctly")
	}
}
