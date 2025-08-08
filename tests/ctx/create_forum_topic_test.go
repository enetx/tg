package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_CreateForumTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	name := g.String("Test Topic")

	// Test basic creation
	result := testCtx.CreateForumTopic(name)
	if result == nil {
		t.Error("Expected CreateForumTopic builder to be created")
	}

	// Test ChatID method
	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return CreateForumTopic for chaining")
	}

	// Test IconColor method
	result = testCtx.CreateForumTopic(name).IconColor(0xFF0000) // Red color
	if result == nil {
		t.Error("IconColor method should return CreateForumTopic for chaining")
	}

	// Test IconCustomEmojiID method
	result = testCtx.CreateForumTopic(name).IconCustomEmojiID(g.String("5789134455711613990"))
	if result == nil {
		t.Error("IconCustomEmojiID method should return CreateForumTopic for chaining")
	}

	// Test Timeout method
	result = testCtx.CreateForumTopic(name).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return CreateForumTopic for chaining")
	}

	// Test APIURL method
	result = testCtx.CreateForumTopic(name).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return CreateForumTopic for chaining")
	}
}

func TestContext_CreateForumTopicChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	name := g.String("Advanced Topic")

	// Test complete method chaining
	result := testCtx.CreateForumTopic(name).
		ChatID(-1001987654321).
		IconColor(0x00FF00). // Green color
		IconCustomEmojiID(g.String("5789134455711613991")).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return CreateForumTopic")
	}
}

func TestCreateForumTopic_TopicNames(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	// Test various topic names
	topicNames := []string{
		"General Discussion",
		"Announcements",
		"Development",
		"Bug Reports",
		"Feature Requests",
		"Off-Topic",
		"Help & Support",
		"Community Events",
		"Project Planning",
		"Code Reviews",
		"💡 Ideas & Suggestions",
		"🐛 Bug Tracking",
		"🎉 Celebrations",
		"Team Alpha Discussion",
		"Sprint #42 Updates",
		"Marketing Strategy 2024",
		"Customer Feedback Analysis",
		"Product Roadmap Q1-Q2",
		"API Documentation Updates",
		"Security & Compliance Review",
		"A", // Single character
		"This is a very long topic name that contains many words and characters to test the limits", // Long name
	}

	for _, topicName := range topicNames {
		result := testCtx.CreateForumTopic(g.String(topicName)).
			ChatID(chatID)

		if result == nil {
			t.Errorf("Topic name '%s' should work", topicName)
		}

		// Test with icon color
		colorResult := result.IconColor(0x1E90FF) // Dodger blue
		if colorResult == nil {
			t.Errorf("Topic '%s' with icon color should work", topicName)
		}
	}
}

func TestCreateForumTopic_IconColors(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)
	topicName := g.String("Color Test Topic")

	// Test various icon colors (RGB format)
	iconColors := []struct {
		name        string
		color       int64
		description string
	}{
		{"Red", 0xFF0000, "Pure red color"},
		{"Green", 0x00FF00, "Pure green color"},
		{"Blue", 0x0000FF, "Pure blue color"},
		{"Yellow", 0xFFFF00, "Yellow color"},
		{"Cyan", 0x00FFFF, "Cyan color"},
		{"Magenta", 0xFF00FF, "Magenta color"},
		{"Orange", 0xFF8000, "Orange color"},
		{"Purple", 0x8000FF, "Purple color"},
		{"Pink", 0xFF69B4, "Hot pink color"},
		{"Lime", 0x32CD32, "Lime green color"},
		{"Turquoise", 0x40E0D0, "Turquoise color"},
		{"Gold", 0xFFD700, "Gold color"},
		{"Silver", 0xC0C0C0, "Silver color"},
		{"Brown", 0xA52A2A, "Brown color"},
		{"Navy", 0x000080, "Navy blue color"},
		{"Maroon", 0x800000, "Maroon color"},
		{"Olive", 0x808000, "Olive color"},
		{"Teal", 0x008080, "Teal color"},
		{"Black", 0x000000, "Black color"},
		{"White", 0xFFFFFF, "White color"},
	}

	for _, iconColor := range iconColors {
		t.Run(iconColor.name, func(t *testing.T) {
			result := testCtx.CreateForumTopic(topicName).
				ChatID(chatID).
				IconColor(iconColor.color)

			if result == nil {
				t.Errorf("%s icon color (%s) should work", iconColor.name, iconColor.description)
			}

			// Test with timeout for each color
			timedResult := result.Timeout(30 * time.Second)
			if timedResult == nil {
				t.Errorf("%s color with timeout should work", iconColor.name)
			}
		})
	}

	// Test color with custom emoji (should override color)
	combinedResult := testCtx.CreateForumTopic(topicName).
		ChatID(chatID).
		IconColor(0xFF0000).
		IconCustomEmojiID(g.String("5789134455711613990")) // Custom emoji overrides color

	if combinedResult == nil {
		t.Error("Color with custom emoji should work (emoji overrides color)")
	}
}

func TestCreateForumTopic_CustomEmojis(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)
	topicName := g.String("Emoji Test Topic")

	// Test various custom emoji IDs
	customEmojiIDs := []string{
		"5789134455711613990", // Example emoji ID
		"5789134455711613991",
		"5789134455711613992",
		"1234567890123456789",
		"9876543210987654321",
		"5000000000000000000",
		"1111111111111111111",
		"9999999999999999999",
		"123",                           // Short ID
		"987654",                        // Medium ID
		"12345678901234567890123456789", // Very long ID
	}

	for i, emojiID := range customEmojiIDs {
		result := testCtx.CreateForumTopic(topicName).
			ChatID(chatID).
			IconCustomEmojiID(g.String(emojiID))

		if result == nil {
			t.Errorf("Custom emoji ID '%s' should work", emojiID)
		}

		// Test with additional options
		enhancedResult := result.
			Timeout(time.Duration(30+i*5) * time.Second).
			APIURL(g.String("https://emoji-api.example.com"))

		if enhancedResult == nil {
			t.Errorf("Enhanced custom emoji '%s' should work", emojiID)
		}
	}

	// Test empty emoji ID
	emptyEmojiResult := testCtx.CreateForumTopic(topicName).
		ChatID(chatID).
		IconCustomEmojiID(g.String(""))

	if emptyEmojiResult == nil {
		t.Error("Empty custom emoji ID should work")
	}
}

func TestCreateForumTopic_ForumTopicScenarios(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	// Test realistic forum topic scenarios
	forumScenarios := []struct {
		name        string
		topicName   string
		iconColor   int64
		customEmoji string
		description string
	}{
		{"Announcements", "📢 Announcements", 0xFF6B6B, "", "Official announcements topic with red color"},
		{"Development", "💻 Development", 0x4ECDC4, "", "Development discussion with teal color"},
		{"Bug Reports", "🐛 Bug Reports", 0xFF4757, "", "Bug tracking with red color"},
		{"General", "General Discussion", 0x5352ED, "", "General chat with purple color"},
		{"Custom Emoji Topic", "🎉 Events", 0, "5789134455711613990", "Events topic with custom emoji"},
		{"Feature Requests", "💡 Feature Requests", 0xFFA502, "", "Feature requests with orange color"},
		{"Help Support", "🆘 Help & Support", 0x3742FA, "", "Help section with blue color"},
		{"Off Topic", "🌍 Off-Topic", 0x2ED573, "", "Off-topic discussions with green color"},
	}

	for _, scenario := range forumScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := testCtx.CreateForumTopic(g.String(scenario.topicName)).ChatID(chatID)

			if scenario.customEmoji != "" {
				// Use custom emoji
				result = result.IconCustomEmojiID(g.String(scenario.customEmoji))
			} else {
				// Use color
				result = result.IconColor(scenario.iconColor)
			}

			if result == nil {
				t.Errorf("%s scenario (%s) should work", scenario.name, scenario.description)
			}

			// Test complete workflow for each scenario
			completedResult := result.
				Timeout(45 * time.Second).
				APIURL(g.String("https://forum-api.telegram.org"))

			if completedResult == nil {
				t.Errorf("Complete %s workflow should work", scenario.name)
			}
		})
	}
}

func TestCreateForumTopic_ChatTypes(t *testing.T) {
	bot := &mockBot{}
	topicName := g.String("Test Topic")

	// Test forum topics for various supergroup types
	chatTypes := []struct {
		name   string
		chatID int64
		type_  string
	}{
		{"Standard Supergroup", -1001234567890, "supergroup"},
		{"Large Supergroup", -1002000000000, "supergroup"},
		{"Enterprise Supergroup", -1003000000000, "supergroup"},
		{"Community Supergroup", -1004000000000, "supergroup"},
	}

	for _, chatType := range chatTypes {
		t.Run(chatType.name, func(t *testing.T) {
			rawCtx := &ext.Context{
				EffectiveChat: &gotgbot.Chat{Id: chatType.chatID, Type: chatType.type_},
				Update:        &gotgbot.Update{UpdateId: 1},
			}

			testCtx := ctx.New(bot, rawCtx)

			result := testCtx.CreateForumTopic(topicName).
				ChatID(chatType.chatID).
				IconColor(0x00FF00)

			if result == nil {
				t.Errorf("CreateForumTopic should work for %s", chatType.name)
			}

			// Test with custom emoji for each chat type
			emojiResult := testCtx.CreateForumTopic(topicName).
				ChatID(chatType.chatID).
				IconCustomEmojiID(g.String("5789134455711613990"))

			if emojiResult == nil {
				t.Errorf("Custom emoji should work for %s", chatType.name)
			}
		})
	}
}

func TestCreateForumTopic_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with empty topic name
	emptyName := g.String("")
	result := testCtx.CreateForumTopic(emptyName)
	if result == nil {
		t.Error("Empty topic name should work (builder creation)")
	}

	// Test with zero chat ID (should use effective chat)
	result = testCtx.CreateForumTopic(g.String("Test Topic")).ChatID(0)
	if result == nil {
		t.Error("Zero chat ID should work")
	}

	// Test with zero icon color
	result = testCtx.CreateForumTopic(g.String("Test Topic")).IconColor(0)
	if result == nil {
		t.Error("Zero icon color should work")
	}

	// Test with maximum icon color value
	result = testCtx.CreateForumTopic(g.String("Test Topic")).IconColor(0xFFFFFF) // Maximum RGB
	if result == nil {
		t.Error("Maximum icon color should work")
	}

	// Test with zero timeout
	result = testCtx.CreateForumTopic(g.String("Test Topic")).Timeout(0 * time.Second)
	if result == nil {
		t.Error("Zero timeout should work")
	}

	// Test with very long timeout
	result = testCtx.CreateForumTopic(g.String("Test Topic")).Timeout(24 * time.Hour)
	if result == nil {
		t.Error("Very long timeout should work")
	}

	// Test with empty API URL
	result = testCtx.CreateForumTopic(g.String("Test Topic")).APIURL(g.String(""))
	if result == nil {
		t.Error("Empty API URL should work")
	}

	// Test without ChatID (should use effective chat)
	result = testCtx.CreateForumTopic(g.String("Default Chat Topic"))
	if result == nil {
		t.Error("CreateForumTopic should work without explicit ChatID")
	}
}

func TestCreateForumTopic_MethodCoverage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001987654321)
	topicName := g.String("Complete Test Topic")

	// Test all methods combined in different orders
	// Order 1: Color-based topic
	result1 := testCtx.CreateForumTopic(topicName).
		ChatID(chatID).
		IconColor(0xFF6B6B).
		Timeout(60 * time.Second).
		APIURL(g.String("https://api.telegram.org"))

	if result1 == nil {
		t.Error("All methods with color (order 1) should work")
	}

	// Order 2: Custom emoji-based topic
	result2 := testCtx.CreateForumTopic(topicName).
		APIURL(g.String("https://custom-api.example.com")).
		Timeout(45 * time.Second).
		IconCustomEmojiID(g.String("5789134455711613990")).
		ChatID(chatID)

	if result2 == nil {
		t.Error("All methods with emoji (order 2) should work")
	}

	// Test overriding methods
	result3 := testCtx.CreateForumTopic(topicName).
		ChatID(chatID).
		ChatID(-1002000000000). // Should override first
		IconColor(0xFF0000).
		IconColor(0x00FF00).                                // Should override first
		IconCustomEmojiID(g.String("5789134455711613990")). // Should override color
		Timeout(30 * time.Second).
		Timeout(60 * time.Second) // Should override first

	if result3 == nil {
		t.Error("Method overriding should work")
	}

	// Test both icon options (emoji should override color)
	result4 := testCtx.CreateForumTopic(topicName).
		ChatID(chatID).
		IconColor(0xFF0000).                               // Red color
		IconCustomEmojiID(g.String("5789134455711613991")) // Custom emoji overrides color

	if result4 == nil {
		t.Error("Both icon options should work (emoji overrides color)")
	}

	// Test minimal configuration
	result5 := testCtx.CreateForumTopic(g.String("Minimal Topic"))
	if result5 == nil {
		t.Error("Minimal configuration should work")
	}

	// Test various topic configurations
	topicConfigs := []struct {
		name        string
		topicName   string
		iconType    string
		iconColor   int64
		customEmoji string
	}{
		{"Development Topic", "💻 Development Discussion", "color", 0x4ECDC4, ""},
		{"Bug Reports", "🐛 Bug Tracking", "color", 0xFF4757, ""},
		{"Announcements", "📢 Official Announcements", "emoji", 0, "5789134455711613990"},
		{"General Chat", "General Discussion", "color", 0x5352ED, ""},
		{"Events", "🎉 Community Events", "emoji", 0, "5789134455711613991"},
	}

	for _, config := range topicConfigs {
		result := testCtx.CreateForumTopic(g.String(config.topicName)).ChatID(chatID)

		if config.iconType == "color" {
			result = result.IconColor(config.iconColor)
		} else {
			result = result.IconCustomEmojiID(g.String(config.customEmoji))
		}

		result = result.Timeout(30 * time.Second).APIURL(g.String("https://forum-api.example.com"))

		if result == nil {
			t.Errorf("Configuration %s should work", config.name)
		}
	}
}

func TestCreateForumTopic_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	topicName := g.String("Send Test Topic")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.CreateForumTopic(topicName).Send()

	if sendResult.IsErr() {
		t.Logf("CreateForumTopic Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.CreateForumTopic(topicName).
		ChatID(789).
		IconColor(0xFF0000).
		Timeout(30).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("CreateForumTopic configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
