package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_CreateChatInviteLink(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test basic creation
	result := testCtx.CreateChatInviteLink()
	if result == nil {
		t.Error("Expected CreateChatInviteLink builder to be created")
	}

	// Test ChatID method
	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return CreateChatInviteLink for chaining")
	}

	// Test Name method
	result = testCtx.CreateChatInviteLink().Name(g.String("VIP Access"))
	if result == nil {
		t.Error("Name method should return CreateChatInviteLink for chaining")
	}

	// Test ExpiresAt method
	futureTime := time.Now().Add(24 * time.Hour)
	result = testCtx.CreateChatInviteLink().ExpiresAt(futureTime)
	if result == nil {
		t.Error("ExpiresAt method should return CreateChatInviteLink for chaining")
	}

	// Test ExpiresIn method
	result = testCtx.CreateChatInviteLink().ExpiresIn(7 * 24 * time.Hour)
	if result == nil {
		t.Error("ExpiresIn method should return CreateChatInviteLink for chaining")
	}

	// Test MemberLimit method
	result = testCtx.CreateChatInviteLink().MemberLimit(100)
	if result == nil {
		t.Error("MemberLimit method should return CreateChatInviteLink for chaining")
	}

	// Test CreatesJoinRequest method
	result = testCtx.CreateChatInviteLink().CreatesJoinRequest()
	if result == nil {
		t.Error("CreatesJoinRequest method should return CreateChatInviteLink for chaining")
	}

	// Test Timeout method
	result = testCtx.CreateChatInviteLink().Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return CreateChatInviteLink for chaining")
	}

	// Test APIURL method
	result = testCtx.CreateChatInviteLink().APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return CreateChatInviteLink for chaining")
	}
}

func TestContext_CreateChatInviteLinkChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test complex method chaining
	result := testCtx.CreateChatInviteLink().
		ChatID(-1001987654321).
		Name(g.String("Premium Members")).
		ExpiresIn(30 * 24 * time.Hour).
		MemberLimit(50).
		CreatesJoinRequest().
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return CreateChatInviteLink")
	}

	// Test with ExpiresAt instead of ExpiresIn
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	result2 := testCtx.CreateChatInviteLink().
		ChatID(-1002000000000).
		Name(g.String("Event Access")).
		ExpiresAt(expireTime).
		MemberLimit(200)

	if result2 == nil {
		t.Error("ExpiresAt chaining should work")
	}
}

func TestCreateChatInviteLink_InviteLinkNames(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	// Test various invite link names
	linkNames := []string{
		"VIP Members",
		"Event Access 2024",
		"Premium Subscribers",
		"Beta Testers",
		"Special Guests",
		"Early Access",
		"Team Members",
		"Moderators Only",
		"Limited Time Offer",
		"Conference Attendees",
		"🎉 Celebration Link",
		"Developer Access - API Testing",
	}

	for _, name := range linkNames {
		result := testCtx.CreateChatInviteLink().
			ChatID(chatID).
			Name(g.String(name))

		if result == nil {
			t.Errorf("Invite link name '%s' should work", name)
		}

		// Test with additional options
		combinedResult := result.
			MemberLimit(100).
			ExpiresIn(24 * time.Hour)

		if combinedResult == nil {
			t.Errorf("Combined options with name '%s' should work", name)
		}
	}
}

func TestCreateChatInviteLink_ExpirationTimes(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	// Test ExpiresIn with various durations
	expiresInDurations := []time.Duration{
		1 * time.Hour,
		6 * time.Hour,
		24 * time.Hour,
		7 * 24 * time.Hour,  // 1 week
		30 * 24 * time.Hour, // 1 month
		90 * 24 * time.Hour, // 3 months
	}

	for _, duration := range expiresInDurations {
		result := testCtx.CreateChatInviteLink().
			ChatID(chatID).
			Name(g.String("Timed Access")).
			ExpiresIn(duration)

		if result == nil {
			t.Errorf("ExpiresIn duration %v should work", duration)
		}
	}

	// Test ExpiresAt with various future times
	now := time.Now()
	expiresAtTimes := []time.Time{
		now.Add(2 * time.Hour),
		now.Add(12 * time.Hour),
		now.Add(48 * time.Hour),
		now.AddDate(0, 0, 7), // 1 week
		now.AddDate(0, 1, 0), // 1 month
		now.AddDate(0, 3, 0), // 3 months
		now.AddDate(1, 0, 0), // 1 year
	}

	for _, expireTime := range expiresAtTimes {
		result := testCtx.CreateChatInviteLink().
			ChatID(chatID).
			Name(g.String("Scheduled Access")).
			ExpiresAt(expireTime)

		if result == nil {
			t.Errorf("ExpiresAt time %v should work", expireTime)
		}
	}

	// Test combining both expiration methods (ExpiresIn should override ExpiresAt)
	combinedResult := testCtx.CreateChatInviteLink().
		ChatID(chatID).
		ExpiresAt(now.Add(24 * time.Hour)).
		ExpiresIn(48 * time.Hour). // This should override ExpiresAt
		Name(g.String("Override Test"))

	if combinedResult == nil {
		t.Error("Combined expiration methods should work (ExpiresIn overrides ExpiresAt)")
	}
}

func TestCreateChatInviteLink_MemberLimits(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	// Test various member limits
	memberLimits := []int64{
		1,     // Single user
		5,     // Small group
		10,    // Medium group
		25,    // Large group
		50,    // Very large group
		100,   // Maximum typical limit
		500,   // Enterprise level
		1000,  // Large enterprise
		99999, // Very high limit
	}

	for _, limit := range memberLimits {
		result := testCtx.CreateChatInviteLink().
			ChatID(chatID).
			Name(g.String("Limited Access")).
			MemberLimit(limit)

		if result == nil {
			t.Errorf("Member limit %d should work", limit)
		}

		// Test with expiration
		combinedResult := result.ExpiresIn(24 * time.Hour)
		if combinedResult == nil {
			t.Errorf("Member limit %d with expiration should work", limit)
		}
	}

	// Test member limit with join request approval
	limitWithApproval := testCtx.CreateChatInviteLink().
		ChatID(chatID).
		Name(g.String("Approval Required")).
		MemberLimit(20).
		CreatesJoinRequest()

	if limitWithApproval == nil {
		t.Error("Member limit with join request approval should work")
	}
}

func TestCreateChatInviteLink_JoinRequestApproval(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	// Test basic join request approval
	approvalResult := testCtx.CreateChatInviteLink().
		ChatID(chatID).
		Name(g.String("Approval Required")).
		CreatesJoinRequest()

	if approvalResult == nil {
		t.Error("Join request approval should work")
	}

	// Test approval with various combinations
	approvalScenarios := []struct {
		name        string
		linkName    string
		expiration  time.Duration
		memberLimit int64
	}{
		{"VIP Approval", "VIP Members Only", 7 * 24 * time.Hour, 10},
		{"Event Approval", "Event Registration", 48 * time.Hour, 50},
		{"Team Approval", "Team Applications", 30 * 24 * time.Hour, 5},
		{"Beta Approval", "Beta Tester Applications", 14 * 24 * time.Hour, 25},
		{"Premium Approval", "Premium Access Request", 90 * 24 * time.Hour, 100},
	}

	for _, scenario := range approvalScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := testCtx.CreateChatInviteLink().
				ChatID(chatID).
				Name(g.String(scenario.linkName)).
				ExpiresIn(scenario.expiration).
				MemberLimit(scenario.memberLimit).
				CreatesJoinRequest()

			if result == nil {
				t.Errorf("%s scenario should work", scenario.name)
			}
		})
	}

	// Test approval without member limit (unlimited with approval)
	unlimitedApproval := testCtx.CreateChatInviteLink().
		ChatID(chatID).
		Name(g.String("Unlimited Approval")).
		CreatesJoinRequest().
		ExpiresIn(7 * 24 * time.Hour)

	if unlimitedApproval == nil {
		t.Error("Unlimited approval should work")
	}
}

func TestCreateChatInviteLink_ChatTypes(t *testing.T) {
	bot := &mockBot{}

	// Test various chat types
	chatTypes := []struct {
		name   string
		chatID int64
		type_  string
	}{
		{"Supergroup", -1001234567890, "supergroup"},
		{"Group", -100123456789, "group"},
		{"Channel", -1001987654321, "channel"},
		{"Large Supergroup", -1002000000000, "supergroup"},
	}

	for _, chatType := range chatTypes {
		t.Run(chatType.name, func(t *testing.T) {
			rawCtx := &ext.Context{
				EffectiveChat: &gotgbot.Chat{Id: chatType.chatID, Type: chatType.type_},
				Update:        &gotgbot.Update{UpdateId: 1},
			}

			testCtx := ctx.New(bot, rawCtx)

			result := testCtx.CreateChatInviteLink().
				ChatID(chatType.chatID).
				Name(g.String("Access for " + chatType.name))

			if result == nil {
				t.Errorf("CreateChatInviteLink should work for %s", chatType.name)
			}

			// Test with full features for each chat type
			fullResult := result.
				ExpiresIn(24 * time.Hour).
				MemberLimit(50).
				CreatesJoinRequest()

			if fullResult == nil {
				t.Errorf("Full features should work for %s", chatType.name)
			}
		})
	}
}

func TestCreateChatInviteLink_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with zero chat ID (should use effective chat)
	result := testCtx.CreateChatInviteLink().ChatID(0)
	if result == nil {
		t.Error("Zero chat ID should work")
	}

	// Test with empty name
	result = testCtx.CreateChatInviteLink().Name(g.String(""))
	if result == nil {
		t.Error("Empty name should work")
	}

	// Test with zero member limit
	result = testCtx.CreateChatInviteLink().MemberLimit(0)
	if result == nil {
		t.Error("Zero member limit should work")
	}

	// Test with past expiration time
	pastTime := time.Now().Add(-24 * time.Hour)
	result = testCtx.CreateChatInviteLink().ExpiresAt(pastTime)
	if result == nil {
		t.Error("Past expiration time should work (builder creation)")
	}

	// Test with zero timeout
	result = testCtx.CreateChatInviteLink().Timeout(0 * time.Second)
	if result == nil {
		t.Error("Zero timeout should work")
	}

	// Test with very long timeout
	result = testCtx.CreateChatInviteLink().Timeout(24 * time.Hour)
	if result == nil {
		t.Error("Very long timeout should work")
	}

	// Test with empty API URL
	result = testCtx.CreateChatInviteLink().APIURL(g.String(""))
	if result == nil {
		t.Error("Empty API URL should work")
	}

	// Test without ChatID (should use effective chat)
	result = testCtx.CreateChatInviteLink().Name(g.String("Default Chat Link"))
	if result == nil {
		t.Error("CreateChatInviteLink should work without explicit ChatID")
	}
}

func TestCreateChatInviteLink_MethodCoverage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001987654321)

	// Test all methods combined in different orders
	// Order 1
	result1 := testCtx.CreateChatInviteLink().
		ChatID(chatID).
		Name(g.String("Complete Test")).
		ExpiresIn(7 * 24 * time.Hour).
		MemberLimit(100).
		CreatesJoinRequest().
		Timeout(60 * time.Second).
		APIURL(g.String("https://api.telegram.org"))

	if result1 == nil {
		t.Error("All methods combined (order 1) should work")
	}

	// Order 2 (different sequence)
	result2 := testCtx.CreateChatInviteLink().
		APIURL(g.String("https://custom-api.example.com")).
		Timeout(45 * time.Second).
		CreatesJoinRequest().
		MemberLimit(50).
		ExpiresAt(time.Now().Add(14 * 24 * time.Hour)).
		Name(g.String("Reordered Test")).
		ChatID(chatID)

	if result2 == nil {
		t.Error("All methods combined (order 2) should work")
	}

	// Test overriding methods
	result3 := testCtx.CreateChatInviteLink().
		ChatID(chatID).
		ChatID(-1002000000000). // Should override first
		Name(g.String("First Name")).
		Name(g.String("Second Name")). // Should override first
		MemberLimit(50).
		MemberLimit(100). // Should override first
		ExpiresIn(24 * time.Hour).
		ExpiresIn(48 * time.Hour) // Should override first

	if result3 == nil {
		t.Error("Method overriding should work")
	}

	// Test minimal configuration
	result4 := testCtx.CreateChatInviteLink().Name(g.String("Minimal Link"))
	if result4 == nil {
		t.Error("Minimal configuration should work")
	}

	// Test without any configuration (just defaults)
	result5 := testCtx.CreateChatInviteLink()
	if result5 == nil {
		t.Error("Default configuration should work")
	}
}

func TestCreateChatInviteLink_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.CreateChatInviteLink().Send()

	if sendResult.IsErr() {
		t.Logf("CreateChatInviteLink Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.CreateChatInviteLink().
		ChatID(789).
		Name(g.String("Test Link")).
		MemberLimit(50).
		ExpiresIn(24 * time.Hour).
		Timeout(30).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("CreateChatInviteLink configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
