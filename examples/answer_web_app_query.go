package main

import (
	"encoding/json"
	"time"

	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/inline"
	"github.com/enetx/tg/inline/content"
	"github.com/enetx/tg/keyboard"
)

/*
Example: Telegram Mini Apps with AnswerWebAppQuery

This example demonstrates the complete flow of working with Telegram Mini Apps:

1. Bot creates inline button with "web_app" type
2. User clicks button → Telegram opens Mini App in WebView
3. Mini App sends data back using Telegram.WebApp.sendData()
4. Bot receives Message with WebAppData field
5. Bot processes data and calls answerWebAppQuery to send inline result

Mini App Flow:
- User clicks "web_app" button → Mini App opens
- Mini App calls: window.Telegram.WebApp.sendData(JSON.stringify({...}))
- Telegram creates Message.WebAppData with button_text and data
- Bot receives WebAppData and can use answerWebAppQuery

References:
- Mini Apps: https://core.telegram.org/bots/webapps
- answerWebAppQuery: https://core.telegram.org/bots/api#answerwebappquery
- WebAppData: https://core.telegram.org/bots/api#webappdata
- sendData(): https://core.telegram.org/bots/webapps#initializing-mini-apps

Important: answerWebAppQuery requires query_id from Mini App's initDataUnsafe.query_id,
not from WebAppData.Data. The Mini App must include query_id in the data it sends.
*/

// WebAppResponse represents data sent from Mini App via Telegram.WebApp.sendData()
type WebAppResponse struct {
	QueryID String `json:"query_id"` // From initDataUnsafe.query_id in Mini App
	Type    String `json:"type"`     // Custom field to determine response type
	Data    String `json:"data"`     // Actual payload from Mini App
}

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Command to create inline button that opens Mini App
	b.Command("webapp", func(ctx *ctx.Context) error {
		// Create inline button that opens Mini App
		// When clicked, Telegram opens WebView with your Mini App
		button := keyboard.Inline().
			WebApp("🚀 Open Mini App", "https://your-miniapp.com")

		return ctx.SendMessage("Click the button to open the Mini App.\n\n" +
			"📱 The Mini App should call:\n" +
			"```javascript\n" +
			"// Get query_id from init data\n" +
			"const queryId = Telegram.WebApp.initDataUnsafe.query_id;\n\n" +
			"// Send data back to bot\n" +
			"Telegram.WebApp.sendData(JSON.stringify({\n" +
			"  query_id: queryId,\n" +
			"  type: 'article',\n" +
			"  data: 'Hello from Mini App!'\n" +
			"}));\n" +
			"```").
			Markup(button).
			Send().
			Err()
	})

	// Handle data received from Mini App
	// This handler triggers when Mini App calls Telegram.WebApp.sendData()
	// Telegram automatically creates Message.WebAppData field
	b.On.Message.WebAppData(func(ctx *ctx.Context) error {
		webAppData := ctx.EffectiveMessage.WebAppData

		// webAppData.Data contains the string passed to sendData()
		// webAppData.ButtonText contains the text of the web_app button that was pressed

		// Parse JSON data from Mini App
		var response WebAppResponse
		if err := json.Unmarshal([]byte(webAppData.Data), &response); err != nil {
			return ctx.SendMessage("Invalid JSON data from Mini App: " + String(webAppData.Data)).Send().Err()
		}

		// Validate that Mini App included query_id
		// query_id comes from initDataUnsafe.query_id in the Mini App
		if response.QueryID == "" {
			return ctx.SendMessage("No query_id received from Mini App.\n\n" +
				"Mini App should include query_id from initDataUnsafe.query_id").Send().Err()
		}

		// Create appropriate inline result based on type
		var result inline.QueryResult

		switch response.Type {
		case "article":
			result = inline.NewArticle(
				"webapp_article_1",
				"Mini App Result",
				content.Text("✅ Data from Mini App: "+response.Data),
			).
				Description("Generated by Mini App")

		case "photo":
			result = inline.NewPhoto(
				"webapp_photo_1",
				"https://example.com/photo.jpg",
				"https://example.com/thumb.jpg",
			).
				Caption("📸 Photo selected in Mini App: " + response.Data)

		default:
			result = inline.NewArticle(
				"webapp_default_1",
				"Unknown Type",
				content.Text("Received unknown type: "+response.Type),
			)
		}

		// Send inline result using query_id from Mini App
		// This calls Telegram's answerWebAppQuery API method
		// The result will be sent as inline message on behalf of the user
		if r := ctx.AnswerWebAppQuery(response.QueryID, result).Timeout(time.Second * 30).Send(); r.IsErr() {
			return ctx.SendMessage(Format("Failed to answer web app query: {}", r.Err())).Send().Err()
		}

		// Optional: Send confirmation message to user
		// Note: The inline result was already sent via answerWebAppQuery
		return ctx.SendMessage("Mini App data processed successfully!\n\n" +
			"📤 Inline message sent on your behalf.").Send().Err()
	})

	b.Polling().Start()
}
