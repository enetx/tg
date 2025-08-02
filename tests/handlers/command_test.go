package handlers_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/handlers"
)

func TestNewCommand(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("start")

	command := handlers.NewCommand(bot, cmd, MockHandler)

	if command == nil {
		t.Error("NewCommand should return a Command")
	}
}

func TestNewCommand_LowerCase(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("START") // uppercase

	command := handlers.NewCommand(bot, cmd, MockHandler)

	if command == nil {
		t.Error("NewCommand should return a Command")
	}

	// The command should be converted to lowercase internally
	// We can't directly test this since the fields are private,
	// but we can verify the command was created successfully
}

func TestCommand_AllowEdited(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("help")

	command := handlers.NewCommand(bot, cmd, MockHandler).AllowEdited()

	if command == nil {
		t.Error("AllowEdited should return the same Command")
	}
}

func TestCommand_AllowChannel(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("info")

	command := handlers.NewCommand(bot, cmd, MockHandler).AllowChannel()

	if command == nil {
		t.Error("AllowChannel should return the same Command")
	}
}

func TestCommand_Triggers(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("custom")

	// Test with custom triggers
	command := handlers.NewCommand(bot, cmd, MockHandler).Triggers('!', '?', '.')

	if command == nil {
		t.Error("Triggers should return the same Command")
	}
}

func TestCommand_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("settings")

	command := handlers.NewCommand(bot, cmd, MockHandler).
		AllowEdited().
		AllowChannel().
		Triggers('/', '!').
		AllowEdited() // Test multiple calls

	if command == nil {
		t.Error("Chained methods should return the same Command")
	}
}

func TestCommand_Register(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("register")

	command := handlers.NewCommand(bot, cmd, MockHandler)

	// Register should not return anything (void method)
	command.Register()

	// Verify the command still exists after registration
	if command == nil {
		t.Error("Command should still exist after registration")
	}
}

func TestCommand_RegisterMultipleTimes(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("multi")

	command := handlers.NewCommand(bot, cmd, MockHandler)

	// Register multiple times should work without issues
	command.Register()
	command.Register()
	command.Register()

	// Should not panic or cause issues
}

func TestCommand_EmptyCommand(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("")

	command := handlers.NewCommand(bot, cmd, MockHandler)

	if command == nil {
		t.Error("NewCommand should handle empty command string")
	}
}

func TestCommand_WhitespaceCommand(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("   ")

	command := handlers.NewCommand(bot, cmd, MockHandler)

	if command == nil {
		t.Error("NewCommand should handle whitespace command string")
	}
}

func TestCommand_SpecialCharacterCommand(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("test@command")

	command := handlers.NewCommand(bot, cmd, MockHandler)

	if command == nil {
		t.Error("NewCommand should handle command with special characters")
	}
}

func TestCommand_UnicodeCommand(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("тест") // Cyrillic

	command := handlers.NewCommand(bot, cmd, MockHandler)

	if command == nil {
		t.Error("NewCommand should handle Unicode command")
	}
}

func TestCommand_LongCommand(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("very_long_command_name_that_exceeds_normal_expectations")

	command := handlers.NewCommand(bot, cmd, MockHandler)

	if command == nil {
		t.Error("NewCommand should handle long command names")
	}
}

func TestCommand_NumberCommand(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("123")

	command := handlers.NewCommand(bot, cmd, MockHandler)

	if command == nil {
		t.Error("NewCommand should handle numeric command names")
	}
}

func TestCommand_MixedCaseCommand(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("MiXeD_CaSe_CoMmAnD")

	command := handlers.NewCommand(bot, cmd, MockHandler)

	if command == nil {
		t.Error("NewCommand should handle mixed case command names")
	}
}

func TestCommand_SingleCharacterTriggers(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("test")

	// Test with single character triggers
	triggers := []rune{'/', '!', '?', '.', ':', ';', '@', '#', '$', '%', '^', '&', '*'}

	for _, trigger := range triggers {
		command := handlers.NewCommand(bot, cmd, MockHandler).Triggers(trigger)
		if command == nil {
			t.Errorf("Command with trigger '%c' should be created successfully", trigger)
		}
	}
}

func TestCommand_MultipleTriggers(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("multi_trigger")

	// Test with multiple triggers
	command := handlers.NewCommand(bot, cmd, MockHandler).Triggers('/', '!', '?', '.', ':', ';')

	if command == nil {
		t.Error("Command with multiple triggers should be created successfully")
	}
}

func TestCommand_EmptyTriggers(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("no_triggers")

	// Test with no triggers
	command := handlers.NewCommand(bot, cmd, MockHandler).Triggers()

	if command == nil {
		t.Error("Command with empty triggers should be created successfully")
	}
}

func TestCommand_UnicodeTriggers(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("unicode_triggers")

	// Test with Unicode triggers
	command := handlers.NewCommand(bot, cmd, MockHandler).Triggers('🔥', '💯', '🚀')

	if command == nil {
		t.Error("Command with Unicode triggers should be created successfully")
	}
}

func TestCommand_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("nil_handler")

	command := handlers.NewCommand(bot, cmd, nil)

	if command == nil {
		t.Error("NewCommand should handle nil handler")
	}
}

func TestCommand_WithNilBot(t *testing.T) {
	cmd := g.String("nil_bot")

	command := handlers.NewCommand(nil, cmd, MockHandler)

	if command == nil {
		t.Error("NewCommand should handle nil bot")
	}
}

func TestCommand_AllOptions(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("full_featured")

	// Test command with all options enabled
	command := handlers.NewCommand(bot, cmd, MockHandler).
		AllowEdited().
		AllowChannel().
		Triggers('/', '!', '?')

	command.Register()

	if command == nil {
		t.Error("Fully configured command should work properly")
	}
}

func TestCommand_DuplicateTriggers(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("duplicate_triggers")

	// Test with duplicate triggers
	command := handlers.NewCommand(bot, cmd, MockHandler).Triggers('/', '/', '!', '!', '?', '?')

	if command == nil {
		t.Error("Command with duplicate triggers should be created successfully")
	}
}

func TestCommand_MethodOrderVariations(t *testing.T) {
	bot := NewMockBot()
	cmd := g.String("order_test")

	// Test different orders of method calls
	tests := []struct {
		name  string
		setup func() *handlers.Command
	}{
		{"Triggers first", func() *handlers.Command {
			return handlers.NewCommand(bot, cmd, MockHandler).
				Triggers('/').
				AllowEdited().
				AllowChannel()
		}},
		{"AllowEdited first", func() *handlers.Command {
			return handlers.NewCommand(bot, cmd, MockHandler).
				AllowEdited().
				Triggers('/').
				AllowChannel()
		}},
		{"AllowChannel first", func() *handlers.Command {
			return handlers.NewCommand(bot, cmd, MockHandler).
				AllowChannel().
				AllowEdited().
				Triggers('/')
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			command := test.setup()
			if command == nil {
				t.Errorf("Command setup '%s' should work", test.name)
			}
		})
	}
}

func TestCommand_SpecialCommandNames(t *testing.T) {
	bot := NewMockBot()

	// Test various special command names
	specialCommands := []string{
		"start",
		"help",
		"settings",
		"admin",
		"cancel",
		"ping",
		"status",
		"version",
		"about",
		"contact",
		"support",
		"feedback",
		"bug_report",
		"feature_request",
	}

	for _, cmdName := range specialCommands {
		t.Run(cmdName, func(t *testing.T) {
			command := handlers.NewCommand(bot, g.String(cmdName), MockHandler)
			if command == nil {
				t.Errorf("Command '%s' should be created successfully", cmdName)
			}
		})
	}
}
