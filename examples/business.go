package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/types/updates"
)

func main() {
	token := "YOUR_BOT_TOKEN"
	b := bot.New(token).Build().Unwrap()

	// Handle business connection events
	b.On.BusinessConnection.Any(func(ctx *ctx.Context) error {
		conn := ctx.Update.BusinessConnection
		user := conn.User.FirstName
		id := conn.Id

		if !conn.IsEnabled {
			g.Println("Business account disconnected: {}", user)
			return nil
		}

		g.Println("Business account connected: {}", user)
		g.Println("Connection ID: {}", id)

		if rights := conn.Rights; rights != nil {
			switch {
			case rights.CanEditName:
				g.Println("Bot can edit business account name")
				fallthrough
			case rights.CanManageStories:
				g.Println("Bot can manage stories")
				fallthrough
			case rights.CanTransferStars:
				g.Println("Bot can transfer stars")
			}
		}

		return nil
	})

	// Handle business connection specifically for enabled accounts
	b.On.BusinessConnection.Enabled(func(ctx *ctx.Context) error {
		conn := ctx.Update.BusinessConnection
		g.Println("New business connection enabled from {}", conn.User.FirstName)
		return nil
	})

	// Handle business messages (messages from connected business account)
	b.On.Message.Business(func(ctx *ctx.Context) error {
		g.Println("Received business message: {}", ctx.EffectiveMessage.Text)
		return ctx.Reply("Business message received!").Send().Err()
	})

	// Handle deleted business messages
	b.On.BusinessMessagesDeleted.Any(func(ctx *ctx.Context) error {
		deleted := ctx.Update.DeletedBusinessMessages

		g.Println("Messages deleted from business chat {} (connection: {})",
			deleted.Chat.Title,
			deleted.BusinessConnectionId)

		g.Println("Deleted {} messages", len(deleted.MessageIds))

		return nil
	})

	// Command to manage business account settings
	b.Command("business", func(ctx *ctx.Context) error {
		account := ctx.Business("your_business_connection_id")

		// Example: Update business account name with optional last name
		if err := account.SetName("My Business").LastName("Inc").Send().Err(); err != nil {
			return ctx.Reply("Failed to update business name").Send().Err()
		}

		// Example: Update business bio
		if err := account.SetBio("We provide excellent services!").Send().Err(); err != nil {
			return ctx.Reply("Failed to update bio").Send().Err()
		}

		// Example: Get star balance
		if result := account.Balance().GetStarBalance().Send(); result.IsOk() {
			stars := result.Ok()
			return ctx.Reply(g.Format("Business star balance: {} stars", stars.Amount)).Send().Err()
		}

		return ctx.Reply("Business account updated successfully!").Send().Err()
	})

	// Command to handle star operations
	b.Command("stars", func(ctx *ctx.Context) error {
		account := ctx.Business("your_business_connection_id")

		// Get current balance
		result := account.Balance().GetStarBalance().Send()
		if result.IsErr() {
			return ctx.Reply("Failed to get star balance").Send().Err()
		}

		balance := result.Ok()
		return ctx.Reply(g.Format("Current balance: {} stars", balance.Amount)).Send().Err()
	})

	// Command to read business messages (mark as read)
	b.Command("read", func(ctx *ctx.Context) error {
		account := ctx.Business("your_business_connection_id")

		// You would get these from the actual message context
		chatID := ctx.EffectiveChat.Id
		messageID := ctx.EffectiveMessage.MessageId

		if err := account.Message().Read(chatID, messageID).Send().Err(); err != nil {
			return ctx.Reply("Failed to mark message as read").Send().Err()
		}

		return ctx.Reply("Message marked as read").Send().Err()
	})

	// Start the bot with business updates enabled
	b.Polling().
		AllowedUpdates(
			updates.Message,
			updates.BusinessConnection,
			updates.BusinessMessage,
			updates.EditedBusinessMessage,
			updates.DeletedBusinessMessages,
		).
		Start()
}
