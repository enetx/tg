package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
)

func TestChoice(t *testing.T) {
	optionText := g.String("Option A")
	choice := input.Choice(optionText)
	if choice == nil {
		t.Error("Expected PollChoice to be created")
	}
	if !assertPollOption(choice) {
		t.Error("PollChoice should implement PollOption correctly")
	}
}

func TestChoice_HTML(t *testing.T) {
	optionText := g.String("Option A")
	choice := input.Choice(optionText)
	result := choice.HTML()
	if result == nil {
		t.Error("Expected HTML method to return PollChoice")
	}
	if result != choice {
		t.Error("Expected HTML to return same PollChoice instance")
	}

	built := result.Build()
	if v, ok := interface{}(built).(gotgbot.InputPollOption); ok {
		if v.TextParseMode != "HTML" {
			t.Error("Expected TextParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InputPollOption")
	}
}

func TestChoice_Markdown(t *testing.T) {
	optionText := g.String("Option A")
	choice := input.Choice(optionText)
	result := choice.Markdown()
	if result == nil {
		t.Error("Expected Markdown method to return PollChoice")
	}
	if result != choice {
		t.Error("Expected Markdown to return same PollChoice instance")
	}

	built := result.Build()
	if v, ok := interface{}(built).(gotgbot.InputPollOption); ok {
		if v.TextParseMode != "MarkdownV2" {
			t.Error("Expected TextParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InputPollOption")
	}
}

func TestChoice_TextEntities(t *testing.T) {
	optionText := g.String("Option A")
	choice := input.Choice(optionText)
	entities := createTestEntities()
	result := choice.TextEntities(entities)
	if result == nil {
		t.Error("Expected TextEntities method to return PollChoice")
	}
	if result != choice {
		t.Error("Expected TextEntities to return same PollChoice instance")
	}

	built := result.Build()
	if v, ok := interface{}(built).(gotgbot.InputPollOption); ok {
		if len(v.TextEntities) == 0 {
			t.Error("Expected TextEntities to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputPollOption")
	}
}

func TestChoice_Build(t *testing.T) {
	optionText := g.String("Option A")
	choice := input.Choice(optionText)
	built := choice.Build()

	if v, ok := interface{}(built).(gotgbot.InputPollOption); ok {
		if v.Text != optionText.Std() {
			t.Errorf("Expected option text to be %s, got %s", optionText.Std(), v.Text)
		}
	} else {
		t.Error("Expected result to be InputPollOption")
	}
}

func TestChoice_MethodChaining(t *testing.T) {
	optionText := g.String("Option A")
	result := input.Choice(optionText).
		HTML()

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built.Text == "" {
		t.Error("Expected chained Choice to build correctly")
	}

	if _, ok := interface{}(built).(gotgbot.InputPollOption); !ok {
		t.Error("Expected result to be InputPollOption")
	}

	if !assertPollOption(result) {
		t.Error("Expected result to implement PollOption interface")
	}
}

func TestChoice_EmptyText(t *testing.T) {
	emptyText := g.String("")
	choice := input.Choice(emptyText)
	if choice == nil {
		t.Error("Expected PollChoice to be created with empty text")
	}

	built := choice.Build()
	if v, ok := interface{}(built).(gotgbot.InputPollOption); ok {
		if v.Text != "" {
			t.Errorf("Expected empty text, got %s", v.Text)
		}
	} else {
		t.Error("Expected result to be InputPollOption")
	}
}
