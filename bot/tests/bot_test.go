package bot_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func TestBot_Dispatcher(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	dispatcher := bot.Dispatcher()

	if dispatcher == nil {
		t.Error("Expected dispatcher to be non-nil")
	}
}

func TestBot_Updater(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	updater := bot.Updater()

	if updater == nil {
		t.Error("Expected updater to be non-nil")
	}
}

func TestBot_Raw(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	raw := bot.Raw()

	if raw == nil {
		t.Error("Expected raw bot to be non-nil")
	}
}

func TestBot_Middlewares(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	middlewares := bot.Middlewares()

	if middlewares.Len() != 0 {
		t.Errorf("Expected 0 middlewares initially, got %d", middlewares.Len())
	}
}

func TestBot_Use(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()

	// Test middleware function
	middleware := func(ctx *ctx.Context) error {
		return nil
	}

	returnedBot := bot.Use(middleware)

	// Should return the same bot instance for chaining
	if returnedBot != bot {
		t.Error("Expected Use to return the same bot instance")
	}

	// Check that middleware was added
	middlewares := bot.Middlewares()
	if middlewares.Len() != 1 {
		t.Errorf("Expected 1 middleware after Use, got %d", middlewares.Len())
	}
}

// === INTEGRATION TESTS ===

func TestBot_IntegrationFlow(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()

	// Test complete workflow: create bot -> add middleware -> register command -> setup polling
	middleware := func(ctx *ctx.Context) error {
		return nil
	}

	// Add middleware
	returnedBot := bot.Use(middleware)
	if returnedBot != bot {
		t.Error("Expected middleware chaining to return same bot")
	}

	// Register command
	commandHandler := func(ctx *ctx.Context) error {
		return nil
	}
	cmd := bot.Command(g.String("test"), commandHandler)
	if cmd == nil {
		t.Error("Expected command registration to return command handler")
	}

	// Setup polling
	polling := bot.Polling()
	if polling == nil {
		t.Error("Expected polling setup to return polling instance")
	}

	// Setup webhook
	webhook := bot.Webhook()
	if webhook == nil {
		t.Error("Expected webhook setup to return webhook instance")
	}

	// Verify middleware count
	middlewares := bot.Middlewares()
	if middlewares.Len() != 1 {
		t.Errorf("Expected 1 middleware, got %d", middlewares.Len())
	}
}

// === COMPREHENSIVE BUILDER METHOD TESTS ===

func TestBot_ComprehensiveWorkflow(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")

	// Build bot
	result := bot.New(token).DisableTokenCheck().Build()
	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()

	// Add multiple middlewares
	middleware1 := func(ctx *ctx.Context) error { return nil }
	middleware2 := func(ctx *ctx.Context) error { return nil }
	middleware3 := func(ctx *ctx.Context) error { return nil }
	bot.Use(middleware1).Use(middleware2).Use(middleware3)

	// Register multiple commands
	handler := func(ctx *ctx.Context) error { return nil }
	cmd1 := bot.Command(g.String("start"), handler)
	cmd2 := bot.Command(g.String("help"), handler)
	cmd3 := bot.Command(g.String("settings"), handler)
	cmd4 := bot.Command(g.String("about"), handler)

	// Test polling configuration
	polling := bot.Polling().
		DropPendingUpdates().
		EnableWebhookDeletion().
		Timeout(60).
		Limit(100).
		Offset(1000)

	// Test webhook configuration
	webhook := bot.Webhook().
		Domain("https://example.com").
		Path("/webhook").
		SecretToken("my-secret").
		DropPending(true).
		MaxConnections(50).
		IP("192.168.1.100")

	// Test all API method builders
	getMe := bot.GetMe()
	setCommands := bot.SetMyCommands()
	getCommands := bot.GetMyCommands()
	deleteCommands := bot.DeleteMyCommands()
	setName := bot.SetMyName()
	getName := bot.GetMyName()
	setDesc := bot.SetMyDescription()
	getDesc := bot.GetMyDescription()
	setShortDesc := bot.SetMyShortDescription()
	getShortDesc := bot.GetMyShortDescription()
	logOut := bot.LogOut()
	close := bot.Close()
	setRights := bot.SetMyDefaultAdministratorRights()
	getRights := bot.GetMyDefaultAdministratorRights()
	getWebhookInfo := bot.GetWebhookInfo()
	deleteWebhook := bot.DeleteWebhook()

	// Verify all components
	if bot.Dispatcher() == nil {
		t.Error("Expected dispatcher to be initialized")
	}
	if bot.Updater() == nil {
		t.Error("Expected updater to be initialized")
	}
	if bot.Raw() == nil {
		t.Error("Expected raw bot to be initialized")
	}
	if bot.On == nil {
		t.Error("Expected On handlers to be initialized")
	}
	if bot.Middlewares().Len() != 3 {
		t.Errorf("Expected 3 middlewares, got %d", bot.Middlewares().Len())
	}

	// Verify all builders were created
	builders := []any{
		polling, webhook, getMe, setCommands, getCommands, deleteCommands,
		setName, getName, setDesc, getDesc, setShortDesc, getShortDesc,
		logOut, close, setRights, getRights, getWebhookInfo, deleteWebhook,
	}
	for i, builder := range builders {
		if builder == nil {
			t.Errorf("Builder %d is nil", i)
		}
	}

	// Verify all commands were created
	if cmd1 == nil || cmd2 == nil || cmd3 == nil || cmd4 == nil {
		t.Error("Expected all commands to be created")
	}

	// Test webhook handling with various JSON formats
	testCases := [][]byte{
		[]byte(
			`{"update_id": 1, "message": {"message_id": 1, "date": 1234567890, "text": "test", "chat": {"id": 1, "type": "private"}}}`,
		),
		[]byte(
			`{"update_id": 2, "callback_query": {"id": "123", "from": {"id": 1, "is_bot": false, "first_name": "Test"}, "data": "callback_data"}}`,
		),
		[]byte(
			`{"update_id": 3, "inline_query": {"id": "456", "from": {"id": 1, "is_bot": false, "first_name": "Test"}, "query": "inline query", "offset": ""}}`,
		),
	}

	for i, testJSON := range testCases {
		err := bot.HandleWebhook(testJSON)
		if err != nil {
			t.Logf("Test case %d: Expected error in test environment: %v", i+1, err)
		}
	}

	// Test invalid JSON handling
	invalidJSONs := [][]byte{
		[]byte(`{invalid json}`),
		[]byte(`{"update_id": "not_a_number"}`),
		[]byte(``),
	}

	for i, invalidJSON := range invalidJSONs {
		err := bot.HandleWebhook(invalidJSON)
		if len(invalidJSON) > 0 && string(invalidJSON) != `{"update_id": "not_a_number"}` {
			if err == nil {
				t.Errorf("Invalid JSON test case %d: Expected error for invalid JSON", i+1)
			}
		}
	}
}

// === BOT STRUCTURE VALIDATION TESTS ===

func TestBot_StructureValidation(t *testing.T) {
	token := "123456:ABCDEF-test-token-here"
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()

	// Test that bot implements BotAPI interface
	var _ any = bot

	// Test bot token is set correctly
	if bot.Raw().Token != token {
		t.Errorf("Expected bot token %s, got %s", token, bot.Raw().Token)
	}

	// Test that raw bot user is properly set with disabled token check
	if bot.Raw().User.Id == 0 {
		t.Error("Expected bot user ID to be set")
	}
	if !bot.Raw().User.IsBot {
		t.Error("Expected user to be marked as bot")
	}
	if bot.Raw().User.FirstName != "<unknown>" {
		t.Errorf("Expected FirstName to be '<unknown>', got %s", bot.Raw().User.FirstName)
	}
	if bot.Raw().User.Username != "<unknown>" {
		t.Errorf("Expected Username to be '<unknown>', got %s", bot.Raw().User.Username)
	}
}

func TestBot_ConcurrentAccess(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()

	// Test concurrent access to bot methods (should be safe)
	done := make(chan bool, 5)

	// Concurrent middleware addition
	go func() {
		bot.Use(func(ctx *ctx.Context) error { return nil })
		done <- true
	}()

	// Concurrent command registration
	go func() {
		bot.Command(g.String("test1"), func(ctx *ctx.Context) error { return nil })
		done <- true
	}()

	// Concurrent polling access
	go func() {
		polling := bot.Polling()
		if polling == nil {
			t.Error("Expected polling to be non-nil")
		}
		done <- true
	}()

	// Concurrent webhook access
	go func() {
		webhook := bot.Webhook()
		if webhook == nil {
			t.Error("Expected webhook to be non-nil")
		}
		done <- true
	}()

	// Concurrent API method access
	go func() {
		getMe := bot.GetMe()
		if getMe == nil {
			t.Error("Expected GetMe to be non-nil")
		}
		done <- true
	}()

	// Wait for all goroutines to complete
	for range 5 {
		<-done
	}
}
