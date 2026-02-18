// Package main demonstrates comprehensive bot configuration and management.
// This example shows how to set up bot commands, profile, webhooks, and advanced settings.
package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

// Global bot instance for configuration
var botInstance *bot.Bot

func main() {
	// Get bot token from environment
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	// Create bot instance
	botInstance = bot.New(token).Build().Unwrap()

	// Start command handler
	botInstance.Command("start", handleStart).Register()

	// Configuration menu handlers
	botInstance.On.Callback.Equal("config_commands", handleCommandsConfig)
	botInstance.On.Callback.Equal("config_profile", handleProfileConfig)
	botInstance.On.Callback.Equal("config_webhook", handleWebhookConfig)

	// Command management handlers
	botInstance.On.Callback.Equal("cmd_set", handleSetCommands)
	botInstance.On.Callback.Equal("cmd_get", handleGetCommands)
	botInstance.On.Callback.Equal("cmd_delete", handleDeleteCommands)

	// Profile management handlers
	botInstance.On.Callback.Equal("profile_info", handleProfileInfo)

	// Webhook configuration handlers
	botInstance.On.Callback.Equal("webhook_info", handleWebhookInfo)

	// Back navigation
	botInstance.On.Callback.Equal("back_main", handleStart)

	// Start the bot
	log.Println("🚀 Bot Configuration Example started...")
	botInstance.Polling().AllowedUpdates().Start()
}

// handleStart provides main configuration menu
func handleStart(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("⚙️ Commands Config", "config_commands").
		Text("👤 Profile Config", "config_profile").
		Row().
		Text("🌐 Webhook Config", "config_webhook")

	return ctx.Reply("🤖 <b>Bot Configuration Center</b>\n\n" +
		"Choose a configuration category:\n\n" +
		"⚙️ <b>Commands Config</b> - Manage bot commands and scopes\n" +
		"👤 <b>Profile Config</b> - View bot profile information\n" +
		"🌐 <b>Webhook Config</b> - Webhook vs polling setup").
		HTML().
		Markup(kb).
		Send().Err()
}

// ================ COMMANDS CONFIGURATION ================

func handleCommandsConfig(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("➕ Set Commands", "cmd_set").
		Text("📋 Get Commands", "cmd_get").
		Row().
		Text("🗑️ Delete Commands", "cmd_delete").
		Row().
		Text("🔙 Back", "back_main")

	return ctx.EditMessageText("⚙️ <b>Commands Configuration</b>\n\n" +
		"Manage your bot's command menu:\n\n" +
		"➕ <b>Set Commands</b> - Define available commands\n" +
		"📋 <b>Get Commands</b> - View current commands\n" +
		"🗑️ <b>Delete Commands</b> - Remove all commands").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleSetCommands(ctx *ctx.Context) error {
	// Set basic bot commands using bot instance with explicit default scope
	result := botInstance.SetMyCommands().
		AddCommand("start", "Start the bot").
		AddCommand("help", "Show help information").
		AddCommand("settings", "Bot settings").
		AddCommand("about", "About this bot").
		ScopeDefault().
		Send()

	if result.IsErr() {
		return ctx.Reply(g.String("❌ Failed to set commands: " + result.Err().Error())).Send().Err()
	}

	return ctx.Reply("✅ <b>Commands Set Successfully!</b>\n\n" +
		"<pre>" +
		"/start - Start the bot\n" +
		"/help - Show help information\n" +
		"/settings - Bot settings\n" +
		"/about - About this bot" +
		"</pre>\n\n" +
		"Users will now see these commands in their menu.").
		HTML().
		Send().Err()
}

func handleGetCommands(ctx *ctx.Context) error {
	result := botInstance.GetMyCommands().ScopeDefault().Send()

	if result.IsErr() {
		return ctx.Reply(g.Format("❌ Failed to get commands: {}", result.Err())).Send().Err()
	}

	commands := result.Ok()
	if len(commands) == 0 {
		return ctx.Reply("📋 <b>No Commands Set</b>\n\nThe bot has no commands configured.").
			HTML().Send().Err()
	}

	var text strings.Builder
	text.WriteString("📋 <b>Current Bot Commands:</b>\n\n")
	for _, cmd := range commands {
		text.WriteString("• /" + cmd.Command + " - " + cmd.Description + "\n")
	}

	return ctx.Reply(g.String(text.String())).HTML().Send().Err()
}

func handleDeleteCommands(ctx *ctx.Context) error {
	result := botInstance.DeleteMyCommands().ScopeDefault().Send()

	if result.IsErr() {
		return ctx.Reply(g.String("❌ Failed to delete commands: " + result.Err().Error())).Send().Err()
	}

	return ctx.Reply("🗑️ <b>Commands Deleted</b>\n\n" +
		"All bot commands have been removed from the menu.").
		HTML().Send().Err()
}

// ================ PROFILE CONFIGURATION ================

func handleProfileConfig(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("ℹ️ Profile Info", "profile_info").
		Row().
		Text("🔙 Back", "back_main")

	return ctx.EditMessageText("👤 <b>Profile Configuration</b>\n\n" +
		"View your bot's profile information:\n\n" +
		"ℹ️ <b>Profile Info</b> - Current bot profile details\n\n" +
		"<b>Configuration Methods:</b>\n" +
		"<pre><code class=\"language-go\">" +
		"// Set bot name\n" +
		"bot.SetMyName(\"My Bot\").Send()\n\n" +
		"// Set description\n" +
		"bot.SetMyDescription(\"Bot description\").Send()\n\n" +
		"// Set short description\n" +
		"bot.SetMyShortDescription(\"Short desc\").Send()" +
		"</code></pre>").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleProfileInfo(ctx *ctx.Context) error {
	// Get bot information
	result := botInstance.GetMe().Send()

	if result.IsErr() {
		return ctx.Reply(g.String("❌ Failed to get bot info: " + result.Err().Error())).Send().Err()
	}

	bot := result.Ok()

	text := "ℹ️ <b>Bot Profile Information</b>\n\n" +
		"<b>Username:</b> @" + bot.Username + "\n" +
		"<b>First Name:</b> " + bot.FirstName + "\n" +
		"<b>ID:</b> " + g.Int(bot.Id).String().Std() + "\n" +
		"<b>Can Join Groups:</b> " + strconv.FormatBool(bot.CanJoinGroups) + "\n" +
		"<b>Can Read Messages:</b> " + strconv.FormatBool(bot.CanReadAllGroupMessages) + "\n" +
		"<b>Supports Inline:</b> " + strconv.FormatBool(bot.SupportsInlineQueries) + "\n\n" +
		"<i>Use bot configuration methods to update profile settings.</i>"

	return ctx.Reply(g.String(text)).HTML().Send().Err()
}

// ================ WEBHOOK CONFIGURATION ================

func handleWebhookConfig(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("ℹ️ Webhook Info", "webhook_info").
		Row().
		Text("🔙 Back", "back_main")

	return ctx.EditMessageText("🌐 <b>Webhook Configuration</b>\n\n" +
		"Manage webhook vs polling setup:\n\n" +
		"ℹ️ <b>Webhook Info</b> - View current webhook status\n\n" +
		"<b>Configuration Examples:</b>\n" +
		"<pre><code class=\"language-go\">" +
		"// Set webhook\n" +
		"bot.SetWebhook(\"https://yourbot.com/webhook\").\n" +
		"    MaxConnections(100).\n" +
		"    DropPendingUpdates().\n" +
		"    Send()\n\n" +
		"// Delete webhook (switch to polling)\n" +
		"bot.DeleteWebhook().\n" +
		"    DropPendingUpdates().\n" +
		"    Send()\n\n" +
		"// Start webhook server\n" +
		"bot.Webhook(\":8443\").Start()" +
		"</code></pre>").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleWebhookInfo(ctx *ctx.Context) error {
	result := botInstance.GetWebhookInfo().Send()

	if result.IsErr() {
		return ctx.Reply(g.String("❌ Failed to get webhook info: " + result.Err().Error())).Send().Err()
	}

	info := result.Ok()

	text := g.String("ℹ️ <b>Webhook Information</b>\n\n")

	if info.Url == "" {
		text += "<b>Status:</b> Polling Mode\n" +
			"<b>Pending Updates:</b> " + g.Int(info.PendingUpdateCount).String() + "\n\n" +
			"The bot is currently using polling mode to receive updates."
	} else {
		text += "<b>Status:</b> Webhook Mode\n" +
			"<b>URL:</b> <code>" + g.String(info.Url) + "</code>\n" +
			"<b>Pending Updates:</b> " + g.Int(info.PendingUpdateCount).String() + "\n" +
			"<b>Max Connections:</b> " + g.Int(info.MaxConnections).String() + "\n"

		if info.LastErrorDate != 0 {
			text += "<b>Last Error:</b> <i>" + g.String(info.LastErrorMessage) + "</i>\n"
		}
	}

	return ctx.Reply(text).HTML().Send().Err()
}
