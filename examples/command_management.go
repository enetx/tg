package main

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Set basic commands for all users
	b.Command("setcommands", func(ctx *ctx.Context) error {
		result := b.SetMyCommands().
			AddCommand("start", "Start the bot").
			AddCommand("help", "Get help information").
			AddCommand("settings", "Bot settings").
			AddCommand("about", "About this bot").
			ScopeDefault().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Default commands set successfully!").Send().Err()
	})

	// Set commands for private chats only
	b.Command("setprivatecommands", func(ctx *ctx.Context) error {
		result := b.SetMyCommands().
			AddCommand("profile", "View your profile").
			AddCommand("notifications", "Manage notifications").
			AddCommand("privacy", "Privacy settings").
			ScopeAllPrivateChats().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Private chat commands set successfully!").Send().Err()
	})

	// Set commands for group chats only
	b.Command("setgroupcommands", func(ctx *ctx.Context) error {
		result := b.SetMyCommands().
			AddCommand("rules", "Show group rules").
			AddCommand("stats", "Group statistics").
			AddCommand("poll", "Create a poll").
			AddCommand("kick", "Kick a user (admin only)").
			ScopeAllGroupChats().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Group chat commands set successfully!").Send().Err()
	})

	// Set admin-only commands
	b.Command("setadmincommands", func(ctx *ctx.Context) error {
		result := b.SetMyCommands().
			AddCommand("ban", "Ban a user").
			AddCommand("unban", "Unban a user").
			AddCommand("promote", "Promote to admin").
			AddCommand("demote", "Remove admin rights").
			AddCommand("settings", "Group settings").
			ScopeAllChatAdministrators().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Administrator commands set successfully!").Send().Err()
	})

	// Set commands for specific chat
	b.Command("setchatcommands", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /setchatcommands <chat_id>").Send().Err()
		}

		chatID := args[0].ToInt().Unwrap().Int64()

		result := b.SetMyCommands().
			AddCommand("welcome", "Set welcome message").
			AddCommand("goodbye", "Set goodbye message").
			AddCommand("customrule", "Add custom rule").
			ScopeChat(chatID).
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Chat-specific commands set successfully!").Send().Err()
	})

	// Set commands for specific user in specific chat
	b.Command("setusercommands", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /setusercommands <chat_id> <user_id>").Send().Err()
		}

		chatID := args[0].ToInt().Unwrap().Int64()
		userID := args[1].ToInt().Unwrap().Int64()

		result := b.SetMyCommands().
			AddCommand("vip", "VIP user features").
			AddCommand("premium", "Premium commands").
			AddCommand("special", "Special user actions").
			ScopeChatMember(chatID, userID).
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("User-specific commands set successfully!").Send().Err()
	})

	// Set multilingual commands
	b.Command("setlangcommands", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /setlangcommands <language_code> (e.g. ru, es, fr)").Send().Err()
		}

		languageCode := args[0]

		var commands *bot.SetMyCommands

		switch languageCode.Std() {
		case "ru":
			commands = b.SetMyCommands().
				AddCommand("start", "Запустить бота").
				AddCommand("pomosch", "Получить помощь").
				AddCommand("nastroyki", "Настройки бота")
		case "es":
			commands = b.SetMyCommands().
				AddCommand("start", "Iniciar el bot").
				AddCommand("ayuda", "Obtener ayuda").
				AddCommand("configuracion", "Configuración del bot")
		case "fr":
			commands = b.SetMyCommands().
				AddCommand("start", "Démarrer le bot").
				AddCommand("aide", "Obtenir de l'aide").
				AddCommand("parametres", "Paramètres du bot")
		default:
			return ctx.Reply("Unsupported language code. Use: ru, es, fr").Send().Err()
		}

		result := commands.
			LanguageCode(languageCode).
			ScopeDefault().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Commands set for language: " + languageCode).Send().Err()
	})

	// Get current commands
	b.Command("getcommands", func(ctx *ctx.Context) error {
		result := b.GetMyCommands().
			ScopeDefault().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		commands := result.Ok()
		if len(commands) == 0 {
			return ctx.Reply("No commands are currently set.").Send().Err()
		}

		response := "Current commands:\n"
		for _, cmd := range commands {
			response += "/" + cmd.Command + " - " + cmd.Description + "\n"
		}

		return ctx.Reply(g.String(response)).Send().Err()
	})

	// Get commands for specific scope
	b.Command("getgroupcommands", func(ctx *ctx.Context) error {
		result := b.GetMyCommands().
			ScopeAllGroupChats().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		commands := result.Ok()
		if len(commands) == 0 {
			return ctx.Reply("No group commands are currently set.").Send().Err()
		}

		response := "Group commands:\n"
		for _, cmd := range commands {
			response += "/" + cmd.Command + " - " + cmd.Description + "\n"
		}

		return ctx.Reply(g.String(response)).Send().Err()
	})

	// Get commands for specific language
	b.Command("getlangcommands", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /getlangcommands <language_code>").Send().Err()
		}

		languageCode := args[0]

		result := b.GetMyCommands().
			LanguageCode(languageCode).
			ScopeDefault().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		commands := result.Ok()
		if len(commands) == 0 {
			return ctx.Reply("No commands set for language: " + languageCode).Send().Err()
		}

		response := "Commands for " + languageCode.Std() + ":\n"
		for _, cmd := range commands {
			response += "/" + cmd.Command + " - " + cmd.Description + "\n"
		}

		return ctx.Reply(g.String(response)).Send().Err()
	})

	// Delete all commands
	b.Command("deletecommands", func(ctx *ctx.Context) error {
		result := b.DeleteMyCommands().
			ScopeDefault().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("All default commands deleted successfully!").Send().Err()
	})

	// Delete group commands
	b.Command("deletegroupcommands", func(ctx *ctx.Context) error {
		result := b.DeleteMyCommands().
			ScopeAllGroupChats().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("All group commands deleted successfully!").Send().Err()
	})

	// Delete admin commands
	b.Command("deleteadmincommands", func(ctx *ctx.Context) error {
		result := b.DeleteMyCommands().
			ScopeAllChatAdministrators().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("All administrator commands deleted successfully!").Send().Err()
	})

	// Delete commands for specific language
	b.Command("deletelangcommands", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /deletelangcommands <language_code>").Send().Err()
		}

		languageCode := args[0]

		result := b.DeleteMyCommands().
			LanguageCode(languageCode).
			ScopeDefault().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Commands deleted for language: " + languageCode).Send().Err()
	})

	// Set menu button to default
	b.Command("setdefaultmenu", func(ctx *ctx.Context) error {
		result := ctx.SetChatMenuButton().
			DefaultMenu().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Chat menu button set to default!").Send().Err()
	})

	// Set menu button to commands
	b.Command("setcommandsmenu", func(ctx *ctx.Context) error {
		result := ctx.SetChatMenuButton().
			CommandsMenu().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Chat menu button set to commands!").Send().Err()
	})

	// Set web app menu button
	b.Command("setwebappmenu", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /setwebappmenu <button_text> <web_app_url>").Send().Err()
		}

		buttonText := args[0]
		webAppURL := args[1]

		webApp := gotgbot.WebAppInfo{
			Url: webAppURL.Std(),
		}

		result := ctx.SetChatMenuButton().
			WebAppMenu(buttonText, webApp).
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Web app menu button set successfully!").Send().Err()
	})

	// Set menu button for specific chat
	b.Command("setchatmenu", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /setchatmenu <chat_id>").Send().Err()
		}

		chatID := args[0].ToInt().Unwrap().Int64()

		result := ctx.SetChatMenuButton().
			ChatID(chatID).
			CommandsMenu().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Menu button set for chat: " + args[0]).Send().Err()
	})

	// Get current menu button
	b.Command("getmenubutton", func(ctx *ctx.Context) error {
		result := ctx.GetChatMenuButton().Send()
		if result.IsErr() {
			return result.Err()
		}

		menuButton := result.Ok()
		var buttonType string

		switch menuButton.(type) {
		case gotgbot.MenuButtonDefault:
			buttonType = "Default"
		case gotgbot.MenuButtonCommands:
			buttonType = "Commands"
		case gotgbot.MenuButtonWebApp:
			webApp := menuButton.(gotgbot.MenuButtonWebApp)
			buttonType = "Web App: " + webApp.Text + " (" + webApp.WebApp.Url + ")"
		default:
			buttonType = "Unknown"
		}

		return ctx.Reply("Current menu button type: " + g.String(buttonType)).Send().Err()
	})

	// Advanced: Set comprehensive command structure
	b.Command("setupbot", func(ctx *ctx.Context) error {
		// Set default commands
		defaultResult := b.SetMyCommands().
			AddCommand("start", "Start the bot").
			AddCommand("help", "Get help").
			AddCommand("about", "About this bot").
			ScopeDefault().
			Send()

		if defaultResult.IsErr() {
			return defaultResult.Err()
		}

		// Set private commands
		privateResult := b.SetMyCommands().
			AddCommand("profile", "Your profile").
			AddCommand("settings", "Personal settings").
			AddCommand("notifications", "Notification settings").
			ScopeAllPrivateChats().
			Send()

		if privateResult.IsErr() {
			return privateResult.Err()
		}

		// Set group commands
		groupResult := b.SetMyCommands().
			AddCommand("rules", "Group rules").
			AddCommand("stats", "Group statistics").
			AddCommand("poll", "Create poll").
			ScopeAllGroupChats().
			Send()

		if groupResult.IsErr() {
			return groupResult.Err()
		}

		// Set admin commands
		adminResult := b.SetMyCommands().
			AddCommand("ban", "Ban user").
			AddCommand("unban", "Unban user").
			AddCommand("promote", "Promote user").
			AddCommand("settings", "Group settings").
			ScopeAllChatAdministrators().
			Send()

		if adminResult.IsErr() {
			return adminResult.Err()
		}

		// Set commands menu button
		menuResult := ctx.SetChatMenuButton().
			CommandsMenu().
			Send()

		if menuResult.IsErr() {
			return menuResult.Err()
		}

		return ctx.Reply("Complete bot command structure set up successfully!\n\n" +
			"• Default commands: start, help, about\n" +
			"• Private commands: profile, settings, notifications\n" +
			"• Group commands: rules, stats, poll\n" +
			"• Admin commands: ban, unban, promote, settings\n" +
			"• Menu button: Commands").Send().Err()
	})

	// Command to reset everything
	b.Command("resetcommands", func(c *ctx.Context) error {
		// Delete all command scopes
		scopes := []func(b *bot.Bot) *bot.DeleteMyCommands{
			func(b *bot.Bot) *bot.DeleteMyCommands { return b.DeleteMyCommands().ScopeDefault() },
			func(b *bot.Bot) *bot.DeleteMyCommands { return b.DeleteMyCommands().ScopeAllPrivateChats() },
			func(b *bot.Bot) *bot.DeleteMyCommands { return b.DeleteMyCommands().ScopeAllGroupChats() },
			func(b *bot.Bot) *bot.DeleteMyCommands { return b.DeleteMyCommands().ScopeAllChatAdministrators() },
		}

		for _, scopeFunc := range scopes {
			result := scopeFunc(b).Send()
			if result.IsErr() {
				return result.Err()
			}
		}

		// Reset menu button to default
		menuResult := c.SetChatMenuButton().
			DefaultMenu().
			Send()

		if menuResult.IsErr() {
			return menuResult.Err()
		}

		return c.Reply("All commands and menu button reset to default!").Send().Err()
	})

	b.Polling().AllowedUpdates().Start()
}
