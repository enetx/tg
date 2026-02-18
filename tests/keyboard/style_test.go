package keyboard_test

import (
	"testing"

	. "github.com/enetx/tg/keyboard"
)

func TestButtonStyle_String(t *testing.T) {
	tests := []struct {
		name     string
		style    ButtonStyle
		expected string
	}{
		{"Default", ButtonStyleDefault, ""},
		{"Danger", ButtonStyleDanger, "danger"},
		{"Success", ButtonStyleSuccess, "success"},
		{"Primary", ButtonStylePrimary, "primary"},
		{"Unknown", ButtonStyle(999), ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.style.String()
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestButtonStyle_Constants(t *testing.T) {
	if ButtonStyleDefault != 0 {
		t.Errorf("Expected ButtonStyleDefault to be 0, got %d", int(ButtonStyleDefault))
	}

	if ButtonStyleDanger != 1 {
		t.Errorf("Expected ButtonStyleDanger to be 1, got %d", int(ButtonStyleDanger))
	}

	if ButtonStyleSuccess != 2 {
		t.Errorf("Expected ButtonStyleSuccess to be 2, got %d", int(ButtonStyleSuccess))
	}

	if ButtonStylePrimary != 3 {
		t.Errorf("Expected ButtonStylePrimary to be 3, got %d", int(ButtonStylePrimary))
	}
}

func TestButtonStyle_Coverage(t *testing.T) {
	allStyles := []ButtonStyle{
		ButtonStyleDanger,
		ButtonStyleSuccess,
		ButtonStylePrimary,
	}

	for _, style := range allStyles {
		result := style.String()
		if result == "" {
			t.Errorf("ButtonStyle %d returned empty string, expected specific value", int(style))
		}
	}
}

func TestButton_Style(t *testing.T) {
	btn := NewButton()

	result := btn.Style(ButtonStyleDanger)
	if result != btn {
		t.Error("Expected Style to return the same button for chaining")
	}
}

func TestButton_Style_Danger(t *testing.T) {
	btn := NewButton().Text("Delete").Style(ButtonStyleDanger)
	if btn == nil {
		t.Error("Expected button with danger style to be non-nil")
	}

	built := btn.Build()
	if built.Style != "danger" {
		t.Errorf("Expected Style 'danger', got %q", built.Style)
	}
}

func TestButton_Style_Success(t *testing.T) {
	btn := NewButton().Text("Confirm").Style(ButtonStyleSuccess)
	built := btn.Build()

	if built.Style != "success" {
		t.Errorf("Expected Style 'success', got %q", built.Style)
	}
}

func TestButton_Style_Primary(t *testing.T) {
	btn := NewButton().Text("OK").Style(ButtonStylePrimary)
	built := btn.Build()

	if built.Style != "primary" {
		t.Errorf("Expected Style 'primary', got %q", built.Style)
	}
}

func TestButton_Style_Default(t *testing.T) {
	btn := NewButton().Text("Neutral").Style(ButtonStyleDefault)
	built := btn.Build()

	if built.Style != "" {
		t.Errorf("Expected Style '' for default, got %q", built.Style)
	}
}

func TestButton_Style_Chaining(t *testing.T) {
	btn := NewButton().
		Text("Action").
		Callback("action:1").
		Style(ButtonStyleDanger)

	if btn == nil {
		t.Error("Expected chained button to be non-nil")
	}

	built := btn.Build()
	if built.Text != "Action" {
		t.Errorf("Expected text 'Action', got %q", built.Text)
	}
	if built.CallbackData != "action:1" {
		t.Errorf("Expected callback 'action:1', got %q", built.CallbackData)
	}
	if built.Style != "danger" {
		t.Errorf("Expected style 'danger', got %q", built.Style)
	}
}

func TestButton_Style_AllValues(t *testing.T) {
	styles := []struct {
		style    ButtonStyle
		expected string
	}{
		{ButtonStyleDanger, "danger"},
		{ButtonStyleSuccess, "success"},
		{ButtonStylePrimary, "primary"},
		{ButtonStyleDefault, ""},
	}

	for _, tt := range styles {
		btn := NewButton().Text("test").Style(tt.style)
		built := btn.Build()

		if built.Style != tt.expected {
			t.Errorf("For style %d: expected %q, got %q", int(tt.style), tt.expected, built.Style)
		}
	}
}
