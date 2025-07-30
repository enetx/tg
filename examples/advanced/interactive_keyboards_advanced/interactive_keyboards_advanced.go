// Package main demonstrates advanced interactive keyboard features in TG Framework.
// This example showcases inline keyboards, callback query handling, dynamic menus,
// pagination, multi-level navigation, and keyboard state management.
package main

import (
	"github.com/enetx/g"
	"github.com/enetx/g/cmp"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

// Global bot instance for keyboard operations
var botInstance *bot.Bot

// Global toggle buttons for settings
var (
	nameToggle = keyboard.NewButton().Callback("toggle_name").
			On("✅ Name: John Doe").
			Off("❌ Name: Hidden")

	emailToggle = keyboard.NewButton().Callback("toggle_email").
			On("✅ Email Notifications").
			Off("❌ Email Notifications")

	publicToggle = keyboard.NewButton().Callback("toggle_public").
			On("✅ Public Profile").
			Off("❌ Public Profile")

	autosaveToggle = keyboard.NewButton().Callback("toggle_autosave").
			On("✅ Auto-save").
			Off("❌ Auto-save")

	pushToggle = keyboard.NewButton().Callback("toggle_push").
			On("✅ Push Notifications").
			Off("❌ Push Notifications")

	emailAlertsToggle = keyboard.NewButton().Callback("toggle_email_alerts").
				On("✅ Email Alerts").
				Off("❌ Email Alerts")

	smsToggle = keyboard.NewButton().Callback("toggle_sms").
			On("✅ SMS Notifications").
			Off("❌ SMS Notifications")

	soundsToggle = keyboard.NewButton().Callback("toggle_sounds").
			On("✅ In-App Sounds").
			Off("❌ In-App Sounds")

	privateToggle = keyboard.NewButton().Callback("toggle_private").
			On("✅ Private Messages").
			Off("❌ Private Messages")

	locationToggle = keyboard.NewButton().Callback("toggle_location").
			On("✅ Location Sharing").
			Off("❌ Location Sharing")

	contactsToggle = keyboard.NewButton().Callback("toggle_contacts").
			On("✅ Contact Sync").
			Off("❌ Contact Sync")

	analyticsToggle = keyboard.NewButton().Callback("toggle_analytics").
			On("✅ Data Analytics").
			Off("❌ Data Analytics")
)

// Global keyboard instances for settings
var (
	accountKeyboard      *keyboard.InlineKeyboard
	notificationKeyboard *keyboard.InlineKeyboard
	privacyKeyboard      *keyboard.InlineKeyboard
)

// Demo data structures for complex keyboards
type MenuItem struct {
	ID          g.String
	Title       g.String
	Description g.String
	Price       g.Int
	Category    g.String
	Available   bool
}

type UserSession struct {
	CurrentPage    g.Int
	SelectedItems  g.Slice[g.String]
	FilterCategory g.String
	SortOrder      g.String
	ViewMode       g.String
}

// Demo menu items
var menuItems = g.SliceOf(
	MenuItem{"pizza_1", "Margherita Pizza", "Classic tomato, mozzarella, basil", 1200, "pizza", true},
	MenuItem{"pizza_2", "Pepperoni Pizza", "Tomato, mozzarella, pepperoni", 1400, "pizza", true},
	MenuItem{"pizza_3", "Vegetarian Pizza", "Tomato, mozzarella, vegetables", 1300, "pizza", true},
	MenuItem{"burger_1", "Classic Burger", "Beef patty, lettuce, tomato, onion", 800, "burger", true},
	MenuItem{"burger_2", "Cheese Burger", "Beef patty, cheese, lettuce, tomato", 900, "burger", false},
	MenuItem{"pasta_1", "Spaghetti Carbonara", "Pasta with eggs, cheese, pancetta", 1100, "pasta", true},
	MenuItem{"pasta_2", "Penne Arrabbiata", "Pasta with spicy tomato sauce", 1000, "pasta", true},
	MenuItem{"drink_1", "Coca Cola", "Classic soft drink", 300, "drink", true},
	MenuItem{"drink_2", "Orange Juice", "Fresh orange juice", 400, "drink", true},
	MenuItem{"dessert_1", "Tiramisu", "Italian coffee dessert", 600, "dessert", true},
)

// User sessions storage using Map (in production, use a database)
var userSessions = g.NewMap[int64, *UserSession]()

func main() {
	// Get bot token from environment
	// token := os.Getenv("BOT_TOKEN")
	// if token == "" {
	// 	log.Fatal("BOT_TOKEN environment variable is required")
	// }

	token := g.NewFile("../../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()

	// Create bot instance
	botInstance = bot.New(token).Build().Unwrap()

	// Initialize toggle keyboards
	initializeToggleKeyboards()

	// Start command handler
	botInstance.Command("keyboards", handleKeyboardDemo).Register()

	// Main menu handlers
	botInstance.On.Callback.Equal("basic_keyboards", handleBasicKeyboards)
	botInstance.On.Callback.Equal("dynamic_menus", handleDynamicMenus)
	botInstance.On.Callback.Equal("pagination_demo", handlePaginationDemo)
	botInstance.On.Callback.Equal("multi_level_nav", handleMultiLevelNavigation)
	botInstance.On.Callback.Equal("keyboard_states", handleKeyboardStates)

	// Basic keyboard handlers
	botInstance.On.Callback.Equal("simple_buttons", handleSimpleButtons)
	botInstance.On.Callback.Equal("url_buttons", handleURLButtons)
	botInstance.On.Callback.Equal("switch_inline", handleSwitchInline)
	botInstance.On.Callback.Equal("login_button", handleLoginButton)

	// Simple button action handlers
	botInstance.On.Callback.Equal("action_success", func(ctx *ctx.Context) error {
		return ctx.AnswerCallbackQuery("✅ Success action triggered!").Alert().Send().Err()
	})
	botInstance.On.Callback.Equal("action_warning", func(ctx *ctx.Context) error {
		return ctx.AnswerCallbackQuery("⚠️ Warning: This is a warning message").Alert().Send().Err()
	})
	botInstance.On.Callback.Equal("action_error", func(ctx *ctx.Context) error {
		return ctx.AnswerCallbackQuery("❌ Error: g.Something went wrong").Alert().Send().Err()
	})
	botInstance.On.Callback.Equal("action_info", func(ctx *ctx.Context) error {
		return ctx.AnswerCallbackQuery("ℹ️ Info: Additional information displayed").Alert().Send().Err()
	})
	botInstance.On.Callback.Equal("action_help", func(ctx *ctx.Context) error {
		return ctx.AnswerCallbackQuery("❓ Help: This is a help message").Alert().Send().Err()
	})

	// Dynamic menu handlers
	botInstance.On.Callback.Equal("restaurant_menu", handleRestaurantMenu)
	botInstance.On.Callback.Prefix("category_", handleCategoryFilter)
	botInstance.On.Callback.Prefix("sort_", handleSortOrder)

	// IMPORTANT: Exact matches must come BEFORE prefix matches
	botInstance.On.Callback.Equal("view_cart", handleViewCart)
	botInstance.On.Callback.Prefix("view_", handleViewMode)
	botInstance.On.Callback.Prefix("item_", handleItemSelection)

	// Pagination handlers
	botInstance.On.Callback.Prefix("page_", handlePageNavigation)

	// Multi-level navigation handlers
	botInstance.On.Callback.Equal("settings_menu", handleSettingsMenu)
	botInstance.On.Callback.Equal("account_settings", handleAccountSettings)
	botInstance.On.Callback.Equal("notification_settings", handleNotificationSettings)
	botInstance.On.Callback.Equal("privacy_settings", handlePrivacySettings)
	botInstance.On.Callback.Prefix("toggle_", handleSettingToggle)

	// Keyboard state handlers
	botInstance.On.Callback.Equal("quiz_demo", handleQuizDemo)
	botInstance.On.Callback.Prefix("answer_", handleQuizAnswer)
	botInstance.On.Callback.Equal("quiz_results", handleQuizResults)
	botInstance.On.Callback.Equal("reset_quiz", handleResetQuiz)

	// Shopping cart handlers
	botInstance.On.Callback.Prefix("add_", handleAddToCart)
	botInstance.On.Callback.Prefix("remove_", handleRemoveFromCart)

	// view_cart handler moved above to avoid prefix conflict
	botInstance.On.Callback.Equal("clear_cart", handleClearCart)
	botInstance.On.Callback.Equal("checkout", handleCheckout)

	// Document handlers
	botInstance.On.Callback.Prefix("doc_", handleDocumentView)
	botInstance.On.Callback.Equal("current_page", handleCurrentPage)

	// Demo login handler
	botInstance.On.Callback.Equal("demo_login", handleDemoLogin)
	botInstance.On.Callback.Equal("real_login_info", handleRealLoginInfo)

	// Navigation handlers
	botInstance.On.Callback.Equal("back_main", handleKeyboardDemo)
	botInstance.On.Callback.Prefix("back_", handleBackNavigation)

	// Start the bot
	g.Println("🚀 Interactive Keyboards Example started...")
	botInstance.Polling().AllowedUpdates().Start()
}

// handleKeyboardDemo provides main interactive keyboards menu
func handleKeyboardDemo(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("🎮 Basic Keyboards", "basic_keyboards").
		Text("📱 Dynamic Menus", "dynamic_menus").
		Row().
		Text("📄 Pagination Demo", "pagination_demo").
		Text("🏗️ Multi-Level Navigation", "multi_level_nav").
		Row().
		Text("🎯 Keyboard States", "keyboard_states")

	return ctx.Reply("⌨️ <b>Interactive Keyboards Showcase</b>\n\n" +
		"Comprehensive keyboard interaction examples:\n\n" +
		"🎮 <b>Basic Keyboards</b> - Simple button types and layouts\n" +
		"📱 <b>Dynamic Menus</b> - Context-aware menu generation\n" +
		"📄 <b>Pagination Demo</b> - Large dataset navigation\n" +
		"🏗️ <b>Multi-Level Navigation</b> - Nested menu systems\n" +
		"🎯 <b>Keyboard States</b> - Stateful interaction patterns\n\n" +
		"<b>Key Features:</b>\n" +
		"• Callback query handling\n" +
		"• Dynamic keyboard generation\n" +
		"• State management\n" +
		"• User session tracking\n" +
		"• Complex navigation flows\n\n" +
		"<i>Master advanced keyboard interactions with real-world examples.</i>").
		HTML().
		Markup(kb).
		Send().Err()
}

// ================ BASIC KEYBOARDS ================

func handleBasicKeyboards(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("🔘 Simple Buttons", "simple_buttons").
		Text("🔗 URL Buttons", "url_buttons").
		Row().
		Text("↗️ Switch Inline", "switch_inline").
		Text("🔐 Login Button", "login_button").
		Row().
		Text("🔙 Back", "back_main")

	return ctx.EditMessageText("🎮 <b>Basic Keyboard Types</b>\n\n" +
		"Fundamental keyboard button types:\n\n" +
		"🔘 <b>Simple Buttons</b> - Basic callback buttons\n" +
		"🔗 <b>URL Buttons</b> - External link buttons\n" +
		"↗️ <b>Switch Inline</b> - Inline query buttons\n" +
		"🔐 <b>Login Button</b> - Telegram Login Widget\n\n" +
		"<b>Basic Button Features:</b>\n" +
		"• Immediate callback responses\n" +
		"• External URL navigation\n" +
		"• Inline query switching\n" +
		"• Authentication integration\n" +
		"• Custom callback data\n\n" +
		"<i>Basic buttons form the foundation of keyboard interactions.</i>").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleSimpleButtons(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("✅ Success", "action_success").
		Text("⚠️ Warning", "action_warning").
		Text("❌ Error", "action_error").
		Row().
		Text("ℹ️ Info", "action_info").
		Text("❓ Help", "action_help").
		Row().
		Text("🔙 Back", "basic_keyboards")

	return ctx.EditMessageText("🔘 <b>Simple Callback Buttons</b>\n\n" +
		"Click buttons to see different callback responses:\n\n" +
		"<b>Button Types:</b>\n" +
		"• <b>Success</b> - Positive action confirmation\n" +
		"• <b>Warning</b> - Caution alert message\n" +
		"• <b>Error</b> - Error state notification\n" +
		"• <b>Info</b> - Informational popup\n" +
		"• <b>Help</b> - Assistance message\n\n" +
		"<b>Callback Features:</b>\n" +
		"• <code>ctx.AnswerCallbackQuery()</code> - Basic response\n" +
		"• <code>.Alert()</code> - Display as alert popup\n" +
		"• <code>.Cache(duration)</code> - Client-side caching\n" +
		"• Custom callback data handling\n\n" +
		"<i>Try clicking the buttons to see different callback responses!</i>").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleURLButtons(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		URL("🌐 Telegram", "https://telegram.org").
		URL("📚 TG Framework", "https://github.com/enetx/tg").
		Row().
		URL("🐹 Go Language", "https://golang.org").
		URL("📖 Bot API", "https://core.telegram.org/bots/api").
		Row().
		Text("🔙 Back", "basic_keyboards")

	return ctx.EditMessageText("🔗 <b>URL Buttons</b>\n\n" +
		"External link buttons that open URLs:\n\n" +
		"<b>Available Links:</b>\n" +
		"• <b>Telegram</b> - Official Telegram website\n" +
		"• <b>TG Framework</b> - GitHub repository\n" +
		"• <b>Go Language</b> - Official Go website\n" +
		"• <b>Bot API</b> - Telegram Bot API documentation\n\n" +
		"<b>URL Button Features:</b>\n" +
		"• <code>keyboard.URL(text, url)</code> - Create URL button\n" +
		"• Opens in default browser\n" +
		"• No callback data required\n" +
		"• Supports any valid URL\n" +
		"• Can be mixed with callback buttons\n\n" +
		"<i>Click buttons to open external websites in your browser.</i>").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleSwitchInline(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		SwitchInlineQuery("🔍 Search Public", "search ").
		SwitchInlineQueryCurrentChat("💬 Search Here", "local ").
		Row().
		SwitchInlineQuery("📝 Share Text", "share ").
		SwitchInlineQueryCurrentChat("🎯 Quick Action", "action ").
		Row().
		Text("🔙 Back", "basic_keyboards")

	return ctx.EditMessageText("↗️ <b>Switch Inline Query Buttons</b>\n\n" +
		"Buttons that trigger inline queries:\n\n" +
		"<b>Inline Button Types:</b>\n" +
		"• <b>Search Public</b> - Switch to any chat with query\n" +
		"• <b>Search Here</b> - Switch inline in current chat\n" +
		"• <b>Share Text</b> - Share content to other chats\n" +
		"• <b>Quick Action</b> - Perform action in current chat\n\n" +
		"<b>Switch Inline Features:</b>\n" +
		"• <code>SwitchInlineQuery(text, query)</code> - Any chat\n" +
		"• <code>SwitchInlineQueryCurrentChat(text, query)</code> - Current chat\n" +
		"• Pre-filled query text\n" +
		"• Seamless user experience\n" +
		"• Works with inline query handlers\n\n" +
		"<i>Click buttons to see inline query switching in action!</i>").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleLoginButton(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("🔐 Demo Login", "demo_login").
		Text("🌐 Real Login (Info)", "real_login_info").
		Row().
		Text("🔙 Back", "basic_keyboards")

	err := ctx.EditMessageText("🔐 <b>Telegram Login Button</b>\n\n" +
		"Secure authentication with Telegram Login:\n\n" +
		"<b>Button Types:</b>\n" +
		"• <b>🔐 Demo Login:</b> Works immediately - callback button\n" +
		"• <b>🌐 Real Login:</b> Requires bot domain setup\n\n" +
		"<b>LoginURL Button Features:</b>\n" +
		"• Secure OAuth-like authentication\n" +
		"• User data verification\n" +
		"• No password required\n" +
		"• Server-side validation\n\n" +
		"<b>Setup Requirements:</b>\n" +
		"1. Configure bot domain in @BotFather\n" +
		"2. Set up webhook endpoint\n" +
		"3. Implement login verification\n\n" +
		"<b>Demo Note:</b>\n" +
		"The 'Demo Login' button works as a regular callback.\n" +
		"The 'Real Login' button requires proper domain setup.\n\n" +
		"<i>Try the Demo Login button - it works immediately!</i>").
		HTML().
		Markup(kb).
		Send().Err()

	return err
}

// ================ DYNAMIC MENUS ================

func handleDynamicMenus(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("🍽️ Restaurant Menu", "restaurant_menu").
		Row().
		Text("🔙 Back", "back_main")

	return ctx.EditMessageText("📱 <b>Dynamic Menu Generation</b>\n\n" +
		"Context-aware menus that adapt to user state:\n\n" +
		"🍽️ <b>Restaurant Menu</b> - Complete ordering system\n\n" +
		"<b>Dynamic Features:</b>\n" +
		"• Real-time menu updates\n" +
		"• Category-based filtering\n" +
		"• Availability status\n" +
		"• Price calculations\n" +
		"• Shopping cart integration\n" +
		"• Multi-step ordering\n\n" +
		"<b>State Management:</b>\n" +
		"• User session tracking\n" +
		"• Filter preferences\n" +
		"• Selection memory\n" +
		"• Cart persistence\n\n" +
		"<i>Experience dynamic keyboard generation with real-time updates.</i>").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleRestaurantMenu(ctx *ctx.Context) error {
	userID := ctx.EffectiveUser.Id
	session := getUserSession(userID)

	// Build filter and sort options
	kb := keyboard.Inline().
		Row().
		Text("🍕 Pizza", "category_pizza").
		Text("🍔 Burgers", "category_burger").
		Text("🍝 Pasta", "category_pasta").
		Row().
		Text("🥤 Drinks", "category_drink").
		Text("🍰 Desserts", "category_dessert").
		Text("📋 All", "category_all").
		Row().
		Text("💰 Price ↑", "sort_price_asc").
		Text("💰 Price ↓", "sort_price_desc").
		Text("📝 Name", "sort_name").
		Row()

	// Add current filter info
	filterText := g.String("All Categories")
	if !session.FilterCategory.Empty() && session.FilterCategory.Ne("all") {
		filterText = session.FilterCategory.Title()
	}

	// Get filtered items
	filteredItems := getFilteredItems(session)

	// Add items to keyboard using functional approach (first 6 items)
	itemCount := g.Int(0)
	filteredItems.Iter().
		Take(6).
		ForEach(func(item MenuItem) {
			available := g.String("✅")
			if !item.Available {
				available = "❌"
			}

			kb.Text(g.Format("{} {} - {}",
				available,
				item.Title,
				(g.Float(item.Price)/100).RoundDecimal(2)),
				"item_"+item.ID,
			)

			if itemCount%2 == 1 {
				kb.Row()
			}

			itemCount++
		})

	if itemCount%2 == 1 {
		kb.Row()
	}

	kb.Row().
		Text("🛒 View Cart", "view_cart").
		Text("🔙 Back", "dynamic_menus")

	menuText := "🍽️ <b>Restaurant Menu</b>\n\n" +
		"<b>Filter:</b> " + filterText + "\n" +
		"<b>Sort:</b> " + getSortDisplayName(session.SortOrder) + "\n" +
		"<b>Available Items:</b> " + filteredItems.Len().String() + "\n\n"

	if filteredItems.Empty() {
		menuText += "<i>No items match current filter.</i>\n\n"
	} else {
		menuText += "<b>Menu Items:</b>\n"
		filteredItems.Iter().
			Take(6).
			ForEach(func(item MenuItem) {
				status := g.String("✅ Available")
				if !item.Available {
					status = "❌ Unavailable"
				}

				menuText += g.Format("• <b>{}</b> - {} ({})\n", item.Title, (item.Price.Float() / 100).RoundDecimal(2), status)
			})
		menuText += "\n"
	}

	menuText += "<b>Actions:</b>\n" +
		"• Select category to filter items\n" +
		"• Choose sort order for display\n" +
		"• Click items to add to cart\n" +
		"• View cart to see selections\n\n" +
		"<i>Dynamic menu updates based on your preferences!</i>"

	return ctx.EditMessageText(menuText).
		HTML().
		Markup(kb).
		Send().Err()
}

func handleCategoryFilter(ctx *ctx.Context) error {
	userID := ctx.EffectiveUser.Id
	session := getUserSession(userID)

	// Extract category from callback data
	session.FilterCategory = g.String(ctx.Update.CallbackQuery.Data).StripPrefix("category_")

	// Update menu with new filter
	return handleRestaurantMenu(ctx)
}

func handleSortOrder(ctx *ctx.Context) error {
	userID := ctx.EffectiveUser.Id
	session := getUserSession(userID)

	// Extract sort order from callback data
	session.SortOrder = g.String(ctx.Update.CallbackQuery.Data).StripPrefix("sort_")

	// Update menu with new sort order
	return handleRestaurantMenu(ctx)
}

func handleItemSelection(ctx *ctx.Context) error {
	userID := ctx.EffectiveUser.Id
	session := getUserSession(userID)

	// Extract item ID from callback data
	itemID := g.String(ctx.Update.CallbackQuery.Data).StripPrefix("item_")

	// Find the item
	var selectedItem *MenuItem

	if idx := menuItems.IndexBy(func(item MenuItem) bool { return item.ID.Eq(itemID) }); idx.Gte(0) {
		selectedItem = &menuItems[idx]
	}

	if selectedItem == nil {
		return ctx.AnswerCallbackQuery("❌ Item not found").Alert().Send().Err()
	}

	if !selectedItem.Available {
		return ctx.AnswerCallbackQuery("❌ Item is currently unavailable").Alert().Send().Err()
	}

	session.SelectedItems.Push(itemID)

	// Show confirmation with current cart count
	cartCount := session.SelectedItems.Len()
	return ctx.AnswerCallbackQuery(g.Format("✅ {} added to cart! ({} items)", selectedItem.Title, cartCount)).
		Send().
		Err()
}

func handleViewCart(ctx *ctx.Context) error {
	userID := ctx.EffectiveUser.Id
	session := getUserSession(userID)

	if session.SelectedItems.Empty() {
		// Show empty cart view
		kb := keyboard.Inline().
			Row().
			Text("🔙 Back to Menu", "restaurant_menu")

		ctx.AnswerCallbackQuery("🛒 Your cart is empty").Send()

		return ctx.EditMessageText(g.String("🛒 <b>Shopping Cart</b>\n\n" +
			"<i>Your cart is empty. Add some items from the menu!</i>\n\n" +
			"<b>Available Actions:</b>\n" +
			"• Browse menu items\n" +
			"• Add items to cart\n" +
			"• View cart contents\n\n" +
			"<i>Start shopping to see items here.</i>")).
			HTML().
			Markup(kb).
			Send().Err()
	}

	// Count items
	itemCounts := session.SelectedItems.Iter().Counter()

	cartText := g.String("🛒 <b>Shopping Cart</b>\n\n")
	total := g.Int(0)

	// Build cart text using functional approach
	itemCounts.ForEach(func(itemID g.String, count g.Int) {
		menuItems.Iter().
			Filter(func(item MenuItem) bool { return item.ID == itemID }).
			Take(1).
			ForEach(func(item MenuItem) {
				itemTotal := item.Price * count
				cartText += g.Format("• <b>{}</b> x{} - {}\n",
					item.Title,
					count,
					(g.Float(itemTotal) / 100).RoundDecimal(2))
				total += itemTotal
			})
	})

	cartText += g.Format("\n<b>Total: {}</b>\n\n", (g.Float(total) / 100).RoundDecimal(2))

	kb := keyboard.Inline().
		Row().
		Text("💳 Checkout", "checkout").
		Text("🗑️ Clear Cart", "clear_cart").
		Row().
		Text("🔙 Back to Menu", "restaurant_menu")

	// Add remove buttons for each item type
	if itemCounts.Collect().NotEmpty() {
		cartText += "<b>Remove Items:</b>\n"
		for itemID := range itemCounts {
			for _, item := range menuItems {
				if item.ID == itemID {
					kb.Row().Text("➖ "+item.Title, "remove_"+itemID)
					break
				}
			}
		}
	}

	cartText += "\n<i>Manage your cart or proceed to checkout.</i>"

	return ctx.EditMessageText(cartText).
		HTML().
		Markup(kb).
		Send().Err()
}

// ================ PAGINATION DEMO ================

func handlePaginationDemo(ctx *ctx.Context) error {
	return handlePageNavigation(ctx)
}

func handlePageNavigation(ctx *ctx.Context) error {
	userID := ctx.EffectiveUser.Id
	session := getUserSession(userID)

	data := g.String(ctx.Update.CallbackQuery.Data)

	// Parse page number from callback data
	if data.StartsWith("page_") {
		if page := data.StripPrefix("page_").ToInt(); page.IsOk() {
			session.CurrentPage = page.Ok()
		}
	}

	// Pagination settings
	itemsPerPage := 5
	totalItems := 50 // Demo: 50 items total
	totalPages := (totalItems + itemsPerPage - 1) / itemsPerPage

	// Ensure page is within bounds
	if session.CurrentPage < 1 {
		session.CurrentPage = 1
	}
	if session.CurrentPage.Std() > totalPages {
		session.CurrentPage = g.Int(totalPages)
	}

	// Build pagination keyboard
	kb := keyboard.Inline()

	// Add items for current page
	startItem := (session.CurrentPage.Std() - 1) * itemsPerPage
	for i := 0; i < itemsPerPage && startItem+i < totalItems; i++ {
		itemNum := startItem + i + 1
		kb.Row().Text(g.Format("📄 Document {}", itemNum), g.Format("doc_{}", itemNum))
	}

	// Build pagination controls
	paginationRow := kb.Row()

	// Previous page button
	if session.CurrentPage > 1 {
		paginationRow.Text("⬅️ Prev", g.Format("page_{}", session.CurrentPage-1))
	}

	// Page indicator
	paginationRow.Text(g.Format("📄 {}/{}", session.CurrentPage, totalPages), "current_page")

	// Next page button
	if session.CurrentPage.Std() < totalPages {
		paginationRow.Text("Next ➡️", g.Format("page_{}", session.CurrentPage+1))
	}

	// Navigation controls
	kb.Row().
		Text("⏮️ First", "page_1").
		Text("🔄 Refresh", g.Format("page_{}", session.CurrentPage)).
		Text("Last ⏭️", g.Format("page_{}", totalPages)).
		Row().
		Text("🔙 Back", "back_main")

	paginationText := "📄 <b>Pagination Demo</b>\n\n" +
		g.Format("<b>Page {} of {}</b>\n", session.CurrentPage, totalPages) +
		g.Format(
			"<b>Showing items {}-{} of {}</b>\n\n",
			startItem+1,
			min(startItem+itemsPerPage, totalItems),
			totalItems,
		)

	paginationText += "<b>Current Page Items:</b>\n"
	for i := 0; i < itemsPerPage && startItem+i < totalItems; i++ {
		itemNum := startItem + i + 1
		paginationText += g.Format("• Document {} - Sample content item\n", itemNum)
	}

	paginationText += "\n<b>Pagination Features:</b>\n" +
		"• Dynamic page calculation\n" +
		"• Boundary validation\n" +
		"• Navigation controls\n" +
		"• State persistence\n" +
		"• Responsive design\n\n" +
		"<i>Navigate through large datasets efficiently!</i>"

	return ctx.EditMessageText(paginationText).
		HTML().
		Markup(kb).
		Send().Err()
}

// ================ MULTI-LEVEL NAVIGATION ================

func handleMultiLevelNavigation(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("⚙️ Settings Menu", "settings_menu").
		Row().
		Text("🔙 Back", "back_main")

	return ctx.EditMessageText("🏗️ <b>Multi-Level Navigation</b>\n\n" +
		"Complex navigation with nested menus:\n\n" +
		"⚙️ <b>Settings Menu</b> - Comprehensive settings system\n\n" +
		"<b>Navigation Features:</b>\n" +
		"• Hierarchical menu structure\n" +
		"• Breadcrumb navigation\n" +
		"• State preservation\n" +
		"• Deep linking support\n" +
		"• Context-aware back buttons\n\n" +
		"<b>Menu Levels:</b>\n" +
		"• Level 1: Main categories\n" +
		"• Level 2: Subcategories\n" +
		"• Level 3: Individual settings\n" +
		"• Level 4: Setting values\n\n" +
		"<i>Navigate through complex menu hierarchies seamlessly.</i>").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleSettingsMenu(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("👤 Account Settings", "account_settings").
		Row().
		Text("🔔 Notifications", "notification_settings").
		Row().
		Text("🔒 Privacy Settings", "privacy_settings").
		Row().
		Text("🔙 Back", "multi_level_nav")

	return ctx.EditMessageText("⚙️ <b>Settings Menu</b>\n\n" +
		"Configure your account preferences:\n\n" +
		"👤 <b>Account Settings</b> - Profile and account options\n" +
		"🔔 <b>Notifications</b> - Notification preferences\n" +
		"🔒 <b>Privacy Settings</b> - Privacy and security options\n\n" +
		"<b>Settings Categories:</b>\n" +
		"• Personal information management\n" +
		"• Communication preferences\n" +
		"• Security configurations\n" +
		"• Privacy controls\n\n" +
		"<i>Access and modify all your settings from one place.</i>").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleAccountSettings(ctx *ctx.Context) error {
	return ctx.EditMessageText("👤 <b>Account Settings</b>\n\n" +
		"Manage your account configuration:\n\n" +
		"<b>Setting Controls:</b>\n" +
		"• ✅ = Enabled setting\n" +
		"• ❌ = Disabled setting\n" +
		"• Click to toggle states\n" +
		"• Changes saved automatically\n\n" +
		"<i>Click settings to toggle their values.</i>").
		HTML().
		Markup(accountKeyboard).
		Send().Err()
}

func handleNotificationSettings(ctx *ctx.Context) error {
	return ctx.EditMessageText("🔔 <b>Notification Settings</b>\n\n" +
		"Control how you receive notifications:\n\n" +
		"<b>Notification Types:</b>\n" +
		"• System messages and updates\n" +
		"• Security alerts\n" +
		"• Feature announcements\n" +
		"• Promotional content\n\n" +
		"<i>Customize your notification experience.</i>").
		HTML().
		Markup(notificationKeyboard).
		Send().Err()
}

func handlePrivacySettings(ctx *ctx.Context) error {
	return ctx.EditMessageText("🔒 <b>Privacy Settings</b>\n\n" +
		"Control your privacy and data sharing:\n\n" +
		"<b>Privacy Features:</b>\n" +
		"• End-to-end encryption\n" +
		"• Minimal data collection\n" +
		"• Granular permissions\n" +
		"• Transparency controls\n\n" +
		"<i>Maintain full control over your privacy.</i>").
		HTML().
		Markup(privacyKeyboard).
		Send().Err()
}

func handleSettingToggle(ctx *ctx.Context) error {
	callbackData := ctx.Update.CallbackQuery.Data

	// Toggle the appropriate button and update keyboard
	switch callbackData {
	case "toggle_name":
		nameToggle.Flip()
		return ctx.EditMessageReplyMarkup(accountKeyboard).Send().Err()
	case "toggle_email":
		emailToggle.Flip()
		return ctx.EditMessageReplyMarkup(accountKeyboard).Send().Err()
	case "toggle_public":
		publicToggle.Flip()
		return ctx.EditMessageReplyMarkup(accountKeyboard).Send().Err()
	case "toggle_autosave":
		autosaveToggle.Flip()
		return ctx.EditMessageReplyMarkup(accountKeyboard).Send().Err()
	case "toggle_push":
		pushToggle.Flip()
		return ctx.EditMessageReplyMarkup(notificationKeyboard).Send().Err()
	case "toggle_email_alerts":
		emailAlertsToggle.Flip()
		return ctx.EditMessageReplyMarkup(notificationKeyboard).Send().Err()
	case "toggle_sms":
		smsToggle.Flip()
		return ctx.EditMessageReplyMarkup(notificationKeyboard).Send().Err()
	case "toggle_sounds":
		soundsToggle.Flip()
		return ctx.EditMessageReplyMarkup(notificationKeyboard).Send().Err()
	case "toggle_private":
		privateToggle.Flip()
		return ctx.EditMessageReplyMarkup(privacyKeyboard).Send().Err()
	case "toggle_location":
		locationToggle.Flip()
		return ctx.EditMessageReplyMarkup(privacyKeyboard).Send().Err()
	case "toggle_contacts":
		contactsToggle.Flip()
		return ctx.EditMessageReplyMarkup(privacyKeyboard).Send().Err()
	case "toggle_analytics":
		analyticsToggle.Flip()
		return ctx.EditMessageReplyMarkup(privacyKeyboard).Send().Err()
	}

	return ctx.AnswerCallbackQuery("Setting toggled!").Send().Err()
}

// ================ KEYBOARD STATES ================

func handleKeyboardStates(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("🧠 Quiz Demo", "quiz_demo").
		Row().
		Text("🔙 Back", "back_main")

	return ctx.EditMessageText("🎯 <b>Keyboard State Management</b>\n\n" +
		"Stateful interactions with persistent data:\n\n" +
		"🧠 <b>Quiz Demo</b> - Interactive quiz with progress tracking\n\n" +
		"<b>State Features:</b>\n" +
		"• Question progression\n" +
		"• Answer tracking\n" +
		"• Score calculation\n" +
		"• Progress indicators\n" +
		"• Session persistence\n\n" +
		"<b>State Management:</b>\n" +
		"• User session storage\n" +
		"• Progress tracking\n" +
		"• Answer validation\n" +
		"• g.Results calculation\n\n" +
		"<i>Experience complex stateful interactions with keyboards.</i>").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleQuizDemo(ctx *ctx.Context) error {
	userID := ctx.EffectiveUser.Id
	session := getUserSession(userID)

	// Initialize quiz state if needed
	if session.ViewMode == "" {
		session.ViewMode = "quiz_active"
		session.CurrentPage = 0                        // Current question
		session.SelectedItems = g.NewSlice[g.String]() // User answers
	}

	questions := g.Slice[struct {
		Question g.String
		Options  g.Slice[g.String]
		Correct  g.Int
	}]{
		{
			"What is the capital of France?",
			g.Slice[g.String]{"London", "Berlin", "Paris", "Madrid"},
			2,
		},
		{
			"Which programming language is TG Framework built with?",
			g.Slice[g.String]{"Python", "Go", "JavaScript", "Java"},
			1,
		},
		{
			"What does API stand for?",
			g.Slice[g.String]{
				"Application Programming Interface",
				"Automated Program Integration",
				"Advanced Programming Instructions",
				"Application Process Integration",
			},
			0,
		},
	}

	currentQ := session.CurrentPage

	// Check if quiz is completed
	if currentQ.Gte(questions.Len()) {
		return handleQuizResults(ctx)
	}

	question := questions[currentQ]

	kb := keyboard.Inline()

	// Add answer options
	for i, option := range question.Options {
		kb.Row().Text(g.Format("{}) {}", g.String(rune('A'+i)), option), g.Format("answer_{}_{}", currentQ, i))
	}

	kb.Row().
		Text("🔄 Reset Quiz", "reset_quiz").
		Text("🔙 Back", "keyboard_states")

	quizText := "🧠 <b>Interactive Quiz</b>\n\n" +
		g.Format("<b>Question {} of {}</b>\n\n", currentQ+1, len(questions)) +
		"<b>" + question.Question + "</b>\n\n" +
		"<b>Options:</b>\n"

	for i, option := range question.Options {
		quizText += g.Format("{}) {}\n", g.String(rune('A'+i)), option)
	}

	quizText += "\n<b>Progress:</b>\n"
	for i := range questions {
		if i < currentQ.Std() {
			quizText += "✅ "
		} else if i == currentQ.Std() {
			quizText += "⏳ "
		} else {
			quizText += "⏸️ "
		}
	}

	quizText += "\n\n<i>Select your answer to continue to the next question.</i>"

	return ctx.EditMessageText(quizText).
		HTML().
		Markup(kb).
		Send().Err()
}

func handleQuizAnswer(ctx *ctx.Context) error {
	userID := ctx.EffectiveUser.Id
	session := getUserSession(userID)

	// Parse answer from callback data: answer_questionIndex_answerIndex
	parts := g.String(ctx.Update.CallbackQuery.Data).StripPrefix("answer_").Split("_").Collect()

	if parts.Len().Ne(2) {
		return ctx.AnswerCallbackQuery("❌ Invalid answer format").Alert().Send().Err()
	}

	questionIndex := parts[0].ToInt().Unwrap()
	answerIndex := parts[1].ToInt().Unwrap()

	// Store the answer
	for session.SelectedItems.Len().Lte(questionIndex) {
		session.SelectedItems.Push("")
	}

	if questionIndex.Lt(session.SelectedItems.Len()) {
		session.SelectedItems.Set(questionIndex, answerIndex.String())

		// items := session.SelectedItems.Clone()
		// items[questionIndex] = answerIndex.String()
		// session.SelectedItems = items
	}

	// Move to next question
	session.CurrentPage = questionIndex + 1

	// Continue with quiz
	return handleQuizDemo(ctx)
}

func handleQuizResults(ctx *ctx.Context) error {
	userID := ctx.EffectiveUser.Id
	session := getUserSession(userID)

	questions := g.Slice[struct {
		Question g.String
		Options  g.Slice[g.String]
		Correct  g.Int
	}]{
		{
			"What is the capital of France?",
			g.Slice[g.String]{"London", "Berlin", "Paris", "Madrid"},
			2,
		},
		{
			"Which programming language is TG Framework built with?",
			g.Slice[g.String]{"Python", "Go", "JavaScript", "Java"},
			1,
		},
		{
			"What does API stand for?",
			g.Slice[g.String]{
				"Application Programming Interface",
				"Automated Program Integration",
				"Advanced Programming Instructions",
				"Application Process Integration",
			},
			0,
		},
	}

	// Calculate score
	correctAnswers := 0
	for i := g.Int(0); i < session.SelectedItems.Len() && i < questions.Len(); i++ {
		if userAnswer := session.SelectedItems.Get(i); userAnswer.IsSome() {
			if answerIndex := userAnswer.Some().ToInt(); answerIndex.IsOk() {
				if answerIndex.Ok() == questions[i].Correct {
					correctAnswers++
				}
			}
		}
	}

	percentage := (correctAnswers * 100) / len(questions)

	kb := keyboard.Inline().
		Row().
		Text("🔄 Retake Quiz", "reset_quiz").
		Text("🔙 Back", "keyboard_states")

	resultsText := "🎉 <b>Quiz g.Results</b>\n\n" +
		g.Format("<b>Score: {}/{} ({}%)</b>\n\n", correctAnswers, len(questions), percentage)

	// Performance message
	if percentage >= 80 {
		resultsText += "🏆 <b>Excellent!</b> You're well-informed!\n\n"
	} else if percentage >= 60 {
		resultsText += "👍 <b>Good job!</b> You know your stuff!\n\n"
	} else {
		resultsText += "📚 <b>Keep learning!</b> Practice makes perfect!\n\n"
	}

	resultsText += "<b>Answer Review:</b>\n"
	for i, question := range questions {
		userAnswerIndex := g.Int(-1)
		if i < session.SelectedItems.Len().Std() {
			userAnswerIndex = session.SelectedItems[i].ToInt().Ok()
		}

		correctIcon := "❌"
		if userAnswerIndex.Eq(question.Correct) {
			correctIcon = "✅"
		}

		questionText := question.Question
		if questionText.Len().Gt(50) {
			questionText.Truncate(50)
		}

		resultsText += g.Format("{} <b>Q{}:</b> {}\n", correctIcon, i+1, questionText)
		resultsText += g.Format("   <b>Correct:</b> {}) {}\n",
			g.String(rune('A'+question.Correct)),
			question.Options[question.Correct])

		if userAnswerIndex >= 0 && userAnswerIndex < question.Options.Len() && userAnswerIndex.Ne(question.Correct) {
			resultsText += g.Format("   <b>Your answer:</b> {}) {}\n",
				g.String(rune('A'+userAnswerIndex)),
				question.Options[userAnswerIndex])
		}

		resultsText += "\n"
	}

	resultsText += "<i>Great job completing the quiz! Try again to improve your score.</i>"

	return ctx.EditMessageText(resultsText).
		HTML().
		Markup(kb).
		Send().Err()
}

func handleResetQuiz(ctx *ctx.Context) error {
	userID := ctx.EffectiveUser.Id
	session := getUserSession(userID)

	// Reset quiz state
	session.ViewMode = ""
	session.CurrentPage = 0
	session.SelectedItems = g.NewSlice[g.String]()

	return handleQuizDemo(ctx)
}

// ================ HELPER FUNCTIONS ================

func getUserSession(userID int64) *UserSession {
	if session := userSessions.Get(userID); session.IsSome() {
		return session.Unwrap()
	}

	// Create new session
	session := &UserSession{
		CurrentPage:    1,
		SelectedItems:  g.NewSlice[g.String](),
		FilterCategory: "all",
		SortOrder:      "name",
		ViewMode:       "list",
	}

	userSessions.Set(userID, session)
	return session
}

func getFilteredItems(session *UserSession) g.Slice[MenuItem] {
	filtered := menuItems.Iter().
		Filter(func(item MenuItem) bool {
			return session.FilterCategory.Eq("all") ||
				session.FilterCategory.Empty() ||
				item.Category.Eq(session.FilterCategory)
		}).
		Collect()

	result := filtered.Clone()
	switch session.SortOrder.Std() {
	case "price_asc":
		result.SortBy(func(a, b MenuItem) cmp.Ordering { return a.Price.Cmp(b.Price) })
	case "price_desc":
		result.SortBy(func(a, b MenuItem) cmp.Ordering { return b.Price.Cmp(a.Price) })
	case "name":
		result.SortBy(func(a, b MenuItem) cmp.Ordering { return a.Title.Cmp(b.Title) })
	}

	return result
}

func getSortDisplayName(sortOrder g.String) g.String {
	switch sortOrder.Std() {
	case "price_asc":
		return "Price (Low to High)"
	case "price_desc":
		return "Price (High to Low)"
	case "name":
		return "Name (A-Z)"
	default:
		return "Default"
	}
}

func handleBackNavigation(ctx *ctx.Context) error {
	// Extract destination from callback data
	destination := g.String(ctx.Update.CallbackQuery.Data).StripPrefix("back_")

	switch destination {
	case "main":
		return handleKeyboardDemo(ctx)
	case "basic":
		return handleBasicKeyboards(ctx)
	case "dynamic":
		return handleDynamicMenus(ctx)
	case "pagination":
		return handlePaginationDemo(ctx)
	case "multilevel":
		return handleMultiLevelNavigation(ctx)
	case "states":
		return handleKeyboardStates(ctx)
	default:
		return handleKeyboardDemo(ctx)
	}
}

func handleAddToCart(ctx *ctx.Context) error {
	userID := ctx.EffectiveUser.Id
	session := getUserSession(userID)

	// Extract item ID from callback data
	itemID := g.String(ctx.Update.CallbackQuery.Data).StripPrefix("add_")

	// Add to cart
	session.SelectedItems.Push(itemID)

	return ctx.AnswerCallbackQuery("✅ Item added to cart!").Send().Err()
}

func handleRemoveFromCart(ctx *ctx.Context) error {
	userID := ctx.EffectiveUser.Id
	session := getUserSession(userID)

	// Extract item ID from callback data
	itemID := g.String(ctx.Update.CallbackQuery.Data).StripPrefix("remove_")

	// Remove one instance of the item from cart
	found := false
	session.SelectedItems = session.SelectedItems.
		Iter().
		Exclude(func(s g.String) bool {
			if !found && s.Eq(itemID) {
				found = true
				return true
			}
			return false
		}).
		Collect()

	// Update the cart view
	return handleViewCart(ctx)
}

func handleCheckout(ctx *ctx.Context) error {
	return ctx.AnswerCallbackQuery("💳 Proceeding to checkout...").Alert().Send().Err()
}

func handleViewMode(ctx *ctx.Context) error {
	return ctx.AnswerCallbackQuery("👁️ View mode changed!").Send().Err()
}

func handleClearCart(ctx *ctx.Context) error {
	userID := ctx.EffectiveUser.Id
	session := getUserSession(userID)

	// Clear all items from cart
	session.SelectedItems = g.NewSlice[g.String]()

	// Update the cart view
	return handleViewCart(ctx)
}

func handleDocumentView(ctx *ctx.Context) error {
	// Extract document ID from callback data
	docID := g.String(ctx.Update.CallbackQuery.Data).StripPrefix("doc_")
	return ctx.AnswerCallbackQuery("📄 Opening document " + docID).Send().Err()
}

func handleCurrentPage(ctx *ctx.Context) error {
	userID := ctx.EffectiveUser.Id
	session := getUserSession(userID)

	return ctx.AnswerCallbackQuery(g.Format("📄 Current page: {}", session.CurrentPage)).Send().Err()
}

func handleDemoLogin(ctx *ctx.Context) error {
	user := ctx.EffectiveUser

	return ctx.AnswerCallbackQuery(g.Format("🔐 Demo login successful!\nWelcome, {}!", user.FirstName)).
		Alert().
		Send().
		Err()
}

func handleRealLoginInfo(ctx *ctx.Context) error {
	return ctx.AnswerCallbackQuery("🌐 Real LoginURL requires bot domain setup in @BotFather.\n\n" +
		"Steps:\n" +
		"1. Contact @BotFather\n" +
		"2. Use /setdomain command\n" +
		"3. Set your domain\n" +
		"4. Configure webhook endpoint").
		Alert().
		Send().
		Err()
}

// Initialize toggle keyboards with buttons
func initializeToggleKeyboards() {
	// Set initial states (some enabled, some disabled)
	nameToggle.SetActive(true)     // Name visible by default
	emailToggle.SetActive(true)    // Email notifications on
	publicToggle.SetActive(false)  // Public profile off
	autosaveToggle.SetActive(true) // Auto-save on

	pushToggle.SetActive(true)         // Push notifications on
	emailAlertsToggle.SetActive(false) // Email alerts off
	smsToggle.SetActive(true)          // SMS on
	soundsToggle.SetActive(true)       // Sounds on

	privateToggle.SetActive(true)    // Private messages on
	locationToggle.SetActive(false)  // Location sharing off
	contactsToggle.SetActive(true)   // Contact sync on
	analyticsToggle.SetActive(false) // Analytics off

	// Account settings keyboard
	accountKeyboard = keyboard.Inline().
		Row().
		Button(nameToggle).
		Row().
		Button(emailToggle).
		Row().
		Button(publicToggle).
		Row().
		Button(autosaveToggle).
		Row().
		Text("🔙 Back to Settings", "settings_menu")

	// Notification settings keyboard
	notificationKeyboard = keyboard.Inline().
		Row().
		Button(pushToggle).
		Row().
		Button(emailAlertsToggle).
		Row().
		Button(smsToggle).
		Row().
		Button(soundsToggle).
		Row().
		Text("🔙 Back to Settings", "settings_menu")

	// Privacy settings keyboard
	privacyKeyboard = keyboard.Inline().
		Row().
		Button(privateToggle).
		Row().
		Button(locationToggle).
		Row().
		Button(contactsToggle).
		Row().
		Button(analyticsToggle).
		Row().
		Text("🔙 Back to Settings", "settings_menu")
}
