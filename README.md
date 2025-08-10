<img width="2562" height="2855" alt="image" src="https://github.com/user-attachments/assets/fe90b972-3cdb-43ee-8584-aac05a414d61" />

# TG - Telegram Bot Framework for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/enetx/tg.svg)](https://pkg.go.dev/github.com/enetx/tg)
[![Go Report Card](https://goreportcard.com/badge/github.com/enetx/tg)](https://goreportcard.com/report/github.com/enetx/tg)
[![Coverage Status](https://coveralls.io/repos/github/enetx/tg/badge.svg?branch=main&service=github)](https://coveralls.io/github/enetx/tg?branch=main)
[![Go](https://github.com/enetx/tg/actions/workflows/go.yml/badge.svg)](https://github.com/enetx/tg/actions/workflows/go.yml)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/enetx/tg)
[![Telegram Bot API Version][TelegramVersionBadge]][TelegramLastVersion]

Modern and elegant wrapper around [gotgbot](https://github.com/PaulSonOfLars/gotgbot) with convenient API and functional programming style support.

## Video Demo

[![TG Framework Demo](https://img.youtube.com/vi/Mb8y4hDj7so/0.jpg)](https://www.youtube.com/watch?v=Mb8y4hDj7so)

## Features

- 🚀 **Simple and intuitive API** with method chaining support
- 🎯 **Type-safe** event handlers for all Telegram update types
- 🔧 **Flexible bot configuration** through Builder pattern
- 📝 **Rich functionality** for messages, media, keyboards and payments
- 🎮 **Built-in FSM support** (finite state machines) for complex dialogues
- 💰 **Telegram Payments and Stars** support with refunds
- 🎲 **Full support** for all Telegram content types and features
- 🔧 **Middleware system** for request filtering and processing
- 📂 **Advanced file handling** with metadata and thumbnails
- 🌐 **Webhook support** with security features

## Quick Start

### Installation

```bash
go mod init your-bot
go get github.com/enetx/tg
```

### Simple Echo Bot

```go
package main

import (
    "github.com/enetx/g"
    "github.com/enetx/tg/bot"
    "github.com/enetx/tg/ctx"
)

func main() {
    token := "token"
    b := bot.New(token).Build().Unwrap()

    // Handle text messages
    b.On.Message.Text(func(ctx *ctx.Context) error {
        return ctx.Reply("Echo: " + g.String(ctx.EffectiveMessage.Text)).Send().Err()
    })

    // Handle /start command
    b.Command("start", func(ctx *ctx.Context) error {
        return ctx.Reply("Welcome to the bot!").Send().Err()
    })

    b.Polling().Start()
}
```

## Message Handlers

Handle different types of Telegram messages:

```go
// Text messages
b.On.Message.Text(func(ctx *ctx.Context) error {
    return ctx.Reply("Text received").Send().Err()
})

// Media messages
b.On.Message.Photo(func(ctx *ctx.Context) error {
    return ctx.Reply("Photo received").Send().Err()
})

b.On.Message.Voice(func(ctx *ctx.Context) error {
    return ctx.Reply("Voice received").Send().Err()
})

b.On.Message.Video(func(ctx *ctx.Context) error {
    return ctx.Reply("Video received").Send().Err()
})

b.On.Message.Document(func(ctx *ctx.Context) error {
    return ctx.Reply("Document received").Send().Err()
})

// Contact and location
b.On.Message.Contact(func(ctx *ctx.Context) error {
    contact := ctx.EffectiveMessage.Contact
    return ctx.Reply("Thanks for sharing your contact!").Send().Err()
})

b.On.Message.Location(func(ctx *ctx.Context) error {
    location := ctx.EffectiveMessage.Location
    return ctx.Reply("Location received").Send().Err()
})
```

## Commands

Register commands with advanced options:

```go
// Basic command
b.Command("start", func(ctx *ctx.Context) error {
    return ctx.SendMessage("Start command triggered!").Send().Err()
})

// Command with custom triggers and options
b.Command("help", func(ctx *ctx.Context) error {
    return ctx.Reply("Help message").Send().Err()
}).
    Triggers('!', '.').  // Allow !help and .help
    AllowEdited().       // Handle edited messages
    AllowChannel().      // Work in channels
    Register()

// Commands work automatically, but you can customize them further
```

## Inline Keyboards

Create interactive inline keyboards:

```go
// Basic inline keyboard
b.Command("menu", func(ctx *ctx.Context) error {
    markup := keyboard.Inline().
        Row().
        Text("Option 1", "opt1").
        Text("Option 2", "opt2").
        Row().
        URL("Visit Site", "https://example.com").
        WebApp("Open App", "https://webapp.com")

    return ctx.Reply("Choose an option:").Markup(markup).Send().Err()
})

// Handle button presses
b.On.Callback.Equal("opt1", func(ctx *ctx.Context) error {
    return ctx.AnswerCallbackQuery("You chose option 1!").Send().Err()
})

b.On.Callback.Prefix("opt", func(ctx *ctx.Context) error {
    data := ctx.Update.CallbackQuery.Data
    return ctx.AnswerCallbackQuery("You clicked: " + g.String(data)).Alert().Send().Err()
})
```

### Dynamic Keyboard Editing

```go
b.On.Callback.Equal("edit", func(ctx *ctx.Context) error {
    // Edit existing keyboard
    markup := keyboard.Inline(ctx.EffectiveMessage.ReplyMarkup).
        Edit(func(btn *keyboard.Button) {
            switch btn.Get.Callback() {
            case "opt1":
                btn.Text("Modified Option 1")
            case "remove":
                btn.Delete()
            }
        })

    return ctx.EditMessageReplyMarkup(markup).Send().Err()
})
```

## Reply Keyboards

Create custom reply keyboards:

```go
b.Command("keyboard", func(ctx *ctx.Context) error {
    markup := keyboard.Reply().
        Row().
        Text("Regular Button").
        Contact("📞 Share Phone").
        Row().
        Location("📍 Send Location").
        WebApp("🌐 Web App", "https://webapp.com").
        Row().
        Poll("📊 Create Poll", "regular")

    return ctx.Reply("Use the keyboard below:").Markup(markup).Send().Err()
})
```

## File Handling

Send various types of media files:

```go
// Photo
b.Command("photo", func(ctx *ctx.Context) error {
    return ctx.SendPhoto("photo.png").
        Caption("Beautiful photo").
        Send().Err()
})

// Document with advanced options
b.Command("doc", func(ctx *ctx.Context) error {
    return ctx.SendDocument("document.pdf").
        Caption("Important document").
        ReplyTo(ctx.EffectiveMessage.MessageId).
        Send().Err()
})

// Video with metadata
b.Command("video", func(ctx *ctx.Context) error {
    return ctx.SendVideo("video.mp4").
        Caption("Cool video").
        Spoiler().
        Timeout(3 * time.Minute). // Custom timeout
        ApplyMetadata().          // Extract video info (ffprobe)
        GenerateThumbnail().      // Auto-generate thumbnail (ffmpeg)
        Send().Err()
})

// Audio with metadata
b.Command("audio", func(ctx *ctx.Context) error {
    return ctx.SendAudio("song.mp3").
        Title("Song Title").
        Performer("Artist Name").
        Duration(180 * time.Second).
        Send().Err()
})
```

## Finite State Machine (FSM)

Create complex multi-step conversations:

```go
import (
	"github.com/enetx/fsm"
	"github.com/enetx/g"
)

// Define states
const (
    StateGetEmail = "get_email"
    StateGetName  = "get_name"
    StateSummary  = "summary"
)

// Store FSM instances per user
var fsmStore = g.NewMapSafe[int64, *fsm.SyncFSM]()

func main() {
    b := bot.New(token).Build().Unwrap()

    // Create FSM template
    template := fsm.New(StateGetEmail).
        Transition(StateGetEmail, "next", StateGetName).
        Transition(StateGetName, "next", StateSummary)

    // Define state handlers
    template.OnEnter(StateGetEmail, func(fctx *fsm.Context) error {
        tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)
        return tgctx.Reply("Enter your email:").Send().Err()
    })

    template.OnEnter(StateGetName, func(fctx *fsm.Context) error {
        email := fctx.Input.(string)
        fctx.Data.Set("email", email)

        tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)
        return tgctx.Reply("Enter your name:").Send().Err()
    })

    template.OnEnter(StateSummary, func(fctx *fsm.Context) error {
        name := fctx.Input.(string)
        email := fctx.Data.Get("email").UnwrapOr("<no email>")

        tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)
        defer fsmStore.Delete(tgctx.EffectiveUser.Id)

        return tgctx.Reply(g.Format("Got name: {} and email: {}", name, email)).Send().Err()
    })

    // Start FSM
    b.Command("register", func(ctx *ctx.Context) error {
        entry := fsmStore.Entry(ctx.EffectiveUser.Id)
        entry.OrSetBy(func() *fsm.SyncFSM { return template.Clone().Sync() })
        fsm := entry.Get().Some()

        fsm.SetState(StateGetEmail)
        fsm.Context().Meta.Set("tgctx", ctx)
        return fsm.CallEnter(StateGetEmail)
    })

    // Handle FSM input
    b.On.Message.Text(func(ctx *ctx.Context) error {
        opt := fsmStore.Get(ctx.EffectiveUser.Id)
        if opt.IsNone() {
            return nil // No active FSM
        }

        fsm := opt.Some()
        fsm.Context().Meta.Set("tgctx", ctx)
        return fsm.Trigger("next", ctx.EffectiveMessage.Text)
    })

    b.Polling().Start()
}
```

## Payments with Telegram Stars

Handle payments using Telegram's Stars system:

```go
// Create invoice
b.Command("buy", func(ctx *ctx.Context) error {
    if ctx.EffectiveChat.Type != "private" {
        return nil
    }

    return ctx.Invoice("Premium Access", "Get premium features", "premium_123", "XTR").
        Price("Premium Plan", 100).  // 100 stars
        Protect().                   // Content protection
        Send().Err()
})

// Handle pre-checkout (validation)
b.On.PreCheckout.Any(func(ctx *ctx.Context) error {
    // Validate payment here if needed
    return ctx.PreCheckout().Ok().Send().Err()
})

// Handle successful payment
b.On.Message.SuccessfulPayment(func(ctx *ctx.Context) error {
    user := ctx.EffectiveUser
    payment := ctx.EffectiveMessage.SuccessfulPayment
    chargeID := payment.TelegramPaymentChargeId

    // Grant premium access here
    g.Println("User {1.FirstName} ({1.Id}) paid {2.TotalAmount} {2.Currency} with payload {2.InvoicePayload}",
        user, payment)

    return ctx.SendMessage(g.Format("Payment complete! Thank you, {}!\nChargeID:\n{}", user.FirstName, chargeID)).
        Send().Err()
})

// Handle refunds
b.Command("refund", func(ctx *ctx.Context) error {
    chargeID := ctx.Args().Get(0).Some()

    if result := ctx.RefundStarPayment(chargeID).Send(); result.IsErr() {
        err := g.String(result.Err().Error())
        if err.Contains("CHARGE_ALREADY_REFUNDED") {
            return ctx.Reply("This payment was already refunded.").Send().Err()
        }
        return ctx.Reply("Refund failed.").Send().Err()
    }

    return ctx.Reply("Refund processed successfully.").Send().Err()
})
```

## Middleware

Add middleware for request processing:

```go
// Global middleware
b.Use(func(ctx *ctx.Context) error {
    // Log all updates
    fmt.Println("Update from user:", ctx.EffectiveUser.Id)
    return nil // Continue processing
})

// Admin-only middleware
adminMiddleware := func(ctx *ctx.Context) error {
	admin := ctx.IsAdmin()
	if admin.IsErr() {
		return admin.Err()
	}

	if !admin.Ok() {
		return ctx.Answer("Access restricted to admins only!").Alert().Send().Err()
	}

    return nil // Continue
}

// Apply middleware to specific handlers
b.On.Callback.Prefix("admin_", adminMiddleware)
```

## Webhook Mode

Set up webhook instead of polling:

```go
import (
    "net/http"
    "io"

    "github.com/enetx/tg/bot"
    "github.com/enetx/tg/types/updates"
)

func main() {
    b := bot.New(token).Build().Unwrap()

    // Register webhook
    err := b.Webhook().
        Domain("https://yourdomain.com").
        Path("/webhook").
        SecretToken("your-secret").
        AllowedUpdates(updates.Message, updates.CallbackQuery).
        Register()
    if err != nil {
        panic(err)
    }

    // Setup HTTP server
    http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
        // Verify secret token
        if r.Header.Get("X-Telegram-Bot-Api-Secret-Token") != "your-secret" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        body, _ := io.ReadAll(r.Body)
        b.HandleWebhook(body)
        w.WriteHeader(http.StatusOK)
    })

    http.ListenAndServe(":8080", nil)
}
```

## Business Account API

Handle business account connections and messages:

```go
// Handle business connection updates
b.On.BusinessConnection.Enabled(func(ctx *ctx.Context) error {
    conn := ctx.Update.BusinessConnection

    // Configure business account
    return ctx.Business(g.String(conn.Id)).SetName("My Business").
        LastName("LLC").
        Send().Err()
})

// Handle business messages
b.On.Message.Business(func(ctx *ctx.Context) error {
    return ctx.Reply("Business message received!").Send().Err()
})

// Handle deleted business messages
b.On.DeletedBusinessMessages.Any(func(ctx *ctx.Context) error {
    deleted := ctx.Update.DeletedBusinessMessages
    // Process message deletions
    return nil
})

// Manage business account settings
b.Command("business_setup", func(ctx *ctx.Context) error {
    connectionId := g.String("your_connection_id")

    // Set profile information
    err := ctx.Business(connectionId).
        SetBio("Professional business account").
        Send().Err()

    if err != nil {
        return err
    }

    // Check star balance
    balance := ctx.Business(connectionId).Balance().GetStarBalance().Send()
    if balance.IsOk() {
        return ctx.Reply("Stars balance: " + g.String(balance.Ok().Amount)).Send().Err()
    }

    return ctx.Reply("Business account configured").Send().Err()
})
```

## Text Entities and Formatting

Format text messages with various entities:

```go
import (
    "github.com/enetx/g"
    "github.com/enetx/tg/entities"
)

// Basic text formatting
b.Command("format", func(ctx *ctx.Context) error {
    text := g.String("Hello bold italic code")

    e := entities.New(text).
        Bold("bold").     // Make "bold" bold
        Italic("italic"). // Make "italic" italic
        Code("code")      // Make "code" monospace

    return ctx.Reply(text).
        Entities(e).
        Send().Err()
})

// Links and spoilers
b.Command("links", func(ctx *ctx.Context) error {
    text := g.String("Click here to visit Google")

    e := entities.New(text).
        URL("here", "https://google.com"). // "here" as hyperlink
        Spoiler("Google")                  // "Google" as spoiler

    return ctx.Reply(text).
        Entities(e).
        Send().Err()
})

// Code blocks with syntax highlighting
b.Command("codeblock", func(ctx *ctx.Context) error {
    code := g.String(`func main() {
    fmt.Println("Hello")
}`)
    codeText := g.Format("Check this Go code:\n{}", code)

    e := entities.New(codeText).
        Pre(code, "go") // Go code with syntax highlighting

    return ctx.Reply(codeText).
        Entities(e).
        Send().Err()
})

// Multiple formatting types
b.Command("mixed", func(ctx *ctx.Context) error {
    text := g.String("Bold italic underline strikethrough spoiler")

    e := entities.New(text).
        Bold("Bold").
        Italic("italic").
        Underline("underline").
        Strikethrough("strikethrough").
        Spoiler("spoiler")

    return ctx.Reply(text).
        Entities(e).
        Send().Err()
})

// Blockquotes
b.Command("quotes", func(ctx *ctx.Context) error {
    text := g.String(`Regular text
This is a blockquote
This is expandable quote`)

    e := entities.New(text).
        Blockquote("This is a blockquote").
        ExpandableBlockquote("This is expandable quote")

    return ctx.Reply(text).
        Entities(e).
        Send().Err()
})
```

## Advanced Features

### Chat Actions

Show typing indicators and other actions:

```go
b.On.Message.Text(func(ctx *ctx.Context) error {
    // Show typing indicator
    ctx.ChatAction().Typing().Send()

    // Process message...
    time.Sleep(2 * time.Second)

    return ctx.Reply("Processed your message").Send().Err()
})
```

### Dice and Games

```go
// Send dice
b.Command("dice", func(ctx *ctx.Context) error {
    return ctx.SendDice().Send().Err()
})

// Send slot machine
b.Command("slot", func(ctx *ctx.Context) error {
    return ctx.SendDice().Slot().Send().Err()
})
```

### Message Editing and Deletion

```go
b.Command("edit", func(ctx *ctx.Context) error {
    // Send initial message
    msg := ctx.Reply("Original message").Send()

    // Edit it
    return ctx.EditMessageText("Edited message").MessageID(msg.Ok().MessageId).Send().Err()
})

b.Command("delete", func(ctx *ctx.Context) error {
    return ctx.DeleteMessage().Send().Err()
})
```

## Bot Configuration

Configure bot with advanced options:

```go
b := bot.New(token).
    APIURL("https://api.telegram.org").  // Custom API URL
    UseTestEnvironment().                // Use test environment
    DisableTokenCheck().                 // Skip token validation
    Build().
    Unwrap()
```

## Error Handling

All methods follow a consistent error handling pattern:

```go
if err := ctx.Reply("Hello").Send().Err(); err != nil {
    log.Printf("Failed to send message: %v", err)
}

// Or chain with result handling
result := ctx.SendPhoto("image.jpg").Send()
if result.IsErr() {
    log.Printf("Failed to send photo: %v", result.Err())
}
```

## API Documentation

Full API documentation is available at [GoDoc](https://pkg.go.dev/github.com/enetx/tg).

## License

MIT License. See `LICENSE` file for details.

## Support

- Create GitHub issues for bug reports
- Use discussions for questions and suggestions
- Explore examples in the `examples/` folder

[TelegramBotAPI]: https://core.telegram.org/bots/api

[TelegramVersionBadge]: https://img.shields.io/static/v1?label=Supported%20Telegram%20Bot%20API&color=29a1d4&logo=telegram&message=v9.1

[TelegramLastVersion]: https://core.telegram.org/bots/api#july-3-2025
