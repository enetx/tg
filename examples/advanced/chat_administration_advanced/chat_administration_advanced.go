// Package main demonstrates advanced chat administration features in TG Framework.
// This example showcases comprehensive group/channel management, user permissions,
// admin tools, and moderation features.
package main

import (
	"log"
	"os"
	"time"

	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/types/permissions"
	"github.com/enetx/tg/types/roles"
)

// Global bot instance for administration operations
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
	botInstance.Command("admin", handleAdminPanel).Register()

	// User management handlers
	botInstance.On.Callback.Equal("user_management", handleUserManagement)
	botInstance.On.Callback.Equal("ban_user", handleBanUser)
	botInstance.On.Callback.Equal("unban_user", handleUnbanUser)
	botInstance.On.Callback.Equal("restrict_user", handleRestrictUser)
	botInstance.On.Callback.Equal("promote_user", handlePromoteUser)

	// Chat management handlers
	botInstance.On.Callback.Equal("chat_management", handleChatManagement)
	botInstance.On.Callback.Equal("set_title", handleSetTitle)
	botInstance.On.Callback.Equal("set_description", handleSetDescription)
	botInstance.On.Callback.Equal("pin_message", handlePinMessage)
	botInstance.On.Callback.Equal("delete_messages", handleDeleteMessages)

	// Permission management handlers
	botInstance.On.Callback.Equal("permissions", handlePermissions)
	botInstance.On.Callback.Equal("default_permissions", handleDefaultPermissions)
	botInstance.On.Callback.Equal("admin_permissions", handleAdminPermissions)

	// Invite link management
	botInstance.On.Callback.Equal("invite_links", handleInviteLinks)
	botInstance.On.Callback.Equal("create_invite", handleCreateInvite)
	botInstance.On.Callback.Equal("revoke_invite", handleRevokeInvite)

	// Back navigation
	botInstance.On.Callback.Equal("back_admin", handleAdminPanel)

	// Start the bot
	log.Println("🚀 Advanced Chat Administration Example started...")
	botInstance.Polling().AllowedUpdates().Start()
}

// handleAdminPanel provides main administration menu
func handleAdminPanel(ctx *ctx.Context) error {
	// Check if user has admin rights
	admin := ctx.IsAdmin()
	if admin.IsErr() || !admin.Ok() {
		return ctx.Reply("❌ This command is only available to administrators.").Send().Err()
	}

	kb := keyboard.Inline().
		Row().
		Text("👥 User Management", "user_management").
		Text("🏗️ Chat Management", "chat_management").
		Row().
		Text("🔐 Permissions", "permissions").
		Text("🔗 Invite Links", "invite_links")

	return ctx.Reply("🛡️ <b>Advanced Chat Administration</b>\n\n" +
		"Choose an administration category:\n\n" +
		"👥 <b>User Management</b> - Ban, unban, restrict, promote users\n" +
		"🏗️ <b>Chat Management</b> - Chat settings, messages, moderation\n" +
		"🔐 <b>Permissions</b> - Configure chat and user permissions\n" +
		"🔗 <b>Invite Links</b> - Create and manage invitation links").
		HTML().
		Markup(kb).
		Send().Err()
}

// ================ USER MANAGEMENT ================

func handleUserManagement(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("🚫 Ban User", "ban_user").
		Text("✅ Unban User", "unban_user").
		Row().
		Text("🔇 Restrict User", "restrict_user").
		Text("⭐ Promote User", "promote_user").
		Row().
		Text("🔙 Back", "back_admin")

	return ctx.EditMessageText("👥 <b>User Management</b>\n\n" +
		"Manage users in your chat:\n\n" +
		"🚫 <b>Ban User</b> - Remove user from chat\n" +
		"✅ <b>Unban User</b> - Allow banned user to return\n" +
		"🔇 <b>Restrict User</b> - Limit user permissions\n" +
		"⭐ <b>Promote User</b> - Grant admin privileges\n\n" +
		"<i>Usage: Reply to a user's message and use these commands</i>").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleBanUser(ctx *ctx.Context) error {
	// Check if this is a reply to a message
	if ctx.EffectiveMessage.ReplyToMessage == nil {
		return ctx.Reply("Please reply to a user's message to ban them.").Send().Err()
	}

	targetUser := ctx.EffectiveMessage.ReplyToMessage.From
	if targetUser == nil {
		return ctx.Reply("❌ Cannot identify user to ban.").Send().Err()
	}

	// Ban the user
	result := ctx.BanChatMember(targetUser.Id).Send()
	if result.IsErr() {
		return ctx.Reply(String("❌ Failed to ban user: " + result.Err().Error())).Send().Err()
	}

	return ctx.Reply(String("🚫 <b>User Banned Successfully</b>\n\n" +
		"<b>User:</b> " + targetUser.FirstName + "\n" +
		"<b>ID:</b> " + Int(targetUser.Id).String().Std() + "\n\n" +
		"The user has been permanently banned from this chat.")).
		HTML().Send().Err()
}

func handleUnbanUser(ctx *ctx.Context) error {
	// Check if this is a reply to a message
	if ctx.EffectiveMessage.ReplyToMessage == nil {
		return ctx.Reply("Please reply to a user's message to unban them.").Send().Err()
	}

	targetUser := ctx.EffectiveMessage.ReplyToMessage.From
	if targetUser == nil {
		return ctx.Reply("❌ Cannot identify user to unban.").Send().Err()
	}

	// Unban the user
	result := ctx.UnbanChatMember(targetUser.Id).Send()
	if result.IsErr() {
		return ctx.Reply(String("❌ Failed to unban user: " + result.Err().Error())).Send().Err()
	}

	return ctx.Reply(String("✅ <b>User Unbanned Successfully</b>\n\n" +
		"<b>User:</b> " + targetUser.FirstName + "\n" +
		"<b>ID:</b> " + Int(targetUser.Id).String().Std() + "\n\n" +
		"The user can now join the chat again.")).
		HTML().Send().Err()
}

func handleRestrictUser(ctx *ctx.Context) error {
	// Check if this is a reply to a message
	if ctx.EffectiveMessage.ReplyToMessage == nil {
		return ctx.Reply("Please reply to a user's message to restrict them.").Send().Err()
	}

	targetUser := ctx.EffectiveMessage.ReplyToMessage.From
	if targetUser == nil {
		return ctx.Reply("❌ Cannot identify user to restrict.").Send().Err()
	}

	// Restrict user for 1 hour with limited permissions - only basic messaging allowed
	result := ctx.RestrictChatMember(targetUser.Id).
		For(1 * time.Hour).
		Permissions(permissions.SendMessages).
		Send()

	if result.IsErr() {
		return ctx.Reply(String("❌ Failed to restrict user: " + result.Err().Error())).Send().Err()
	}

	return ctx.Reply(String("🔇 <b>User Restricted Successfully</b>\n\n" +
		"<b>User:</b> " + targetUser.FirstName + "\n" +
		"<b>ID:</b> " + Int(targetUser.Id).String().Std() + "\n" +
		"<b>Duration:</b> 1 hour\n\n" +
		"The user's permissions have been limited.")).
		HTML().Send().Err()
}

func handlePromoteUser(ctx *ctx.Context) error {
	// Check if this is a reply to a message
	if ctx.EffectiveMessage.ReplyToMessage == nil {
		return ctx.Reply("Please reply to a user's message to promote them.").Send().Err()
	}

	targetUser := ctx.EffectiveMessage.ReplyToMessage.From
	if targetUser == nil {
		return ctx.Reply("❌ Cannot identify user to promote.").Send().Err()
	}

	// Promote user with basic admin rights
	result := ctx.PromoteChatMember(targetUser.Id).
		Roles(
			roles.ManageChat,
			roles.DeleteMessages,
			roles.ManageVideoChats,
			roles.RestrictMembers,
			roles.ChangeInfo,
			roles.InviteUsers,
			roles.PinMessages,
			roles.ManageTopics,
		).
		Send()

	if result.IsErr() {
		return ctx.Reply(String("❌ Failed to promote user: " + result.Err().Error())).Send().Err()
	}

	return ctx.Reply(String("⭐ <b>User Promoted Successfully</b>\n\n" +
		"<b>User:</b> " + targetUser.FirstName + "\n" +
		"<b>ID:</b> " + Int(targetUser.Id).String().Std() + "\n\n" +
		"The user has been granted administrator privileges.")).
		HTML().Send().Err()
}

// ================ CHAT MANAGEMENT ================

func handleChatManagement(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("📝 Set Title", "set_title").
		Text("📄 Set Description", "set_description").
		Row().
		Text("📌 Pin Message", "pin_message").
		Text("🗑️ Delete Messages", "delete_messages").
		Row().
		Text("🔙 Back", "back_admin")

	return ctx.EditMessageText("🏗️ <b>Chat Management</b>\n\n" +
		"Manage chat settings and content:\n\n" +
		"📝 <b>Set Title</b> - Change chat title\n" +
		"📄 <b>Set Description</b> - Update chat description\n" +
		"📌 <b>Pin Message</b> - Pin important messages\n" +
		"🗑️ <b>Delete Messages</b> - Remove messages\n\n" +
		"<b>Example Usage:</b>\n" +
		"<code>/admin</code> - Access admin panel\n" +
		"Reply to messages for actions").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleSetTitle(ctx *ctx.Context) error {
	// This would typically require user input, showing example
	result := ctx.SetChatTitle("🎯 Advanced Bot Demo Chat").Send()
	if result.IsErr() {
		return ctx.Reply(String("❌ Failed to set title: " + result.Err().Error())).Send().Err()
	}

	return ctx.Reply("📝 <b>Chat Title Updated</b>\n\n" +
		"The chat title has been successfully changed.\n\n" +
		"<i>In a real implementation, you would prompt the user for input.</i>").
		HTML().Send().Err()
}

func handleSetDescription(ctx *ctx.Context) error {
	// Example description update
	description := "🤖 This is an advanced bot demonstration chat showcasing comprehensive Telegram Bot API features including:\n\n" +
		"• Advanced chat administration\n" +
		"• User management and permissions\n" +
		"• Media handling and file processing\n" +
		"• Interactive keyboards and callbacks\n" +
		"• Business account integration\n\n" +
		"Powered by TG Framework for Go"

	result := ctx.SetChatDescription(String(description)).Send()
	if result.IsErr() {
		return ctx.Reply(String("❌ Failed to set description: " + result.Err().Error())).Send().Err()
	}

	return ctx.Reply("📄 <b>Chat Description Updated</b>\n\n" +
		"The chat description has been successfully updated with comprehensive information.").
		HTML().Send().Err()
}

func handlePinMessage(ctx *ctx.Context) error {
	// Check if this is a reply to a message
	if ctx.EffectiveMessage.ReplyToMessage == nil {
		return ctx.Reply("Please reply to a message to pin it.").Send().Err()
	}

	messageID := ctx.EffectiveMessage.ReplyToMessage.MessageId
	result := ctx.PinChatMessage(messageID).Send()

	if result.IsErr() {
		return ctx.Reply(String("❌ Failed to pin message: " + result.Err().Error())).Send().Err()
	}

	return ctx.Reply("📌 <b>Message Pinned Successfully</b>\n\n" +
		"The selected message has been pinned to the chat.").
		HTML().Send().Err()
}

func handleDeleteMessages(ctx *ctx.Context) error {
	// Check if this is a reply to a message
	if ctx.EffectiveMessage.ReplyToMessage == nil {
		return ctx.Reply("Please reply to a message to delete it.").Send().Err()
	}

	messageID := ctx.EffectiveMessage.ReplyToMessage.MessageId
	result := ctx.DeleteMessage().MessageID(messageID).Send()

	if result.IsErr() {
		return ctx.Reply(String("❌ Failed to delete message: " + result.Err().Error())).Send().Err()
	}

	return ctx.Reply("🗑️ <b>Message Deleted Successfully</b>\n\n" +
		"The selected message has been removed from the chat.").
		HTML().Send().Err()
}

// ================ PERMISSIONS MANAGEMENT ================

func handlePermissions(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("🔧 Default Permissions", "default_permissions").
		Text("👑 Admin Permissions", "admin_permissions").
		Row().
		Text("🔙 Back", "back_admin")

	return ctx.EditMessageText("🔐 <b>Permissions Management</b>\n\n" +
		"Configure chat and user permissions:\n\n" +
		"🔧 <b>Default Permissions</b> - Set default user rights\n" +
		"👑 <b>Admin Permissions</b> - Configure administrator rights\n\n" +
		"<b>Permission Types:</b>\n" +
		"• Message permissions (send text, media, polls)\n" +
		"• Chat permissions (invite users, change info)\n" +
		"• Admin permissions (ban, promote, manage)").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleDefaultPermissions(ctx *ctx.Context) error {
	// Set restrictive default permissions - allow basic media but restrict advanced features
	result := ctx.SetChatPermissions().
		Permissions(
			permissions.SendMessages,
			permissions.SendPhotos,
			permissions.SendVideos,
			permissions.SendAudios,
			permissions.SendDocuments,
			permissions.SendVoiceNotes,
			permissions.SendVideoNotes,
		).
		Send()

	if result.IsErr() {
		return ctx.Reply(String("❌ Failed to set permissions: " + result.Err().Error())).Send().Err()
	}

	return ctx.Reply("🔧 <b>Default Permissions Updated</b>\n\n" +
		"<b>Allowed:</b>\n" +
		"• Send messages, photos, videos, audio, documents\n" +
		"• Send voice and video notes\n\n" +
		"<b>Restricted:</b>\n" +
		"• Send polls and other messages\n" +
		"• Add web page previews\n" +
		"• Change chat info\n" +
		"• Invite users\n" +
		"• Pin messages\n" +
		"• Manage topics").
		HTML().Send().Err()
}

func handleAdminPermissions(ctx *ctx.Context) error {
	return ctx.Reply("👑 <b>Admin Permissions Overview</b>\n\n" +
		"<b>Available Admin Rights:</b>\n\n" +
		"🔹 <b>IsAnonymous</b> - Hide admin identity\n" +
		"🔹 <b>CanManageChat</b> - Manage chat settings\n" +
		"🔹 <b>CanDeleteMessages</b> - Delete any messages\n" +
		"🔹 <b>CanManageVideoChats</b> - Manage voice/video chats\n" +
		"🔹 <b>CanRestrictMembers</b> - Ban and restrict users\n" +
		"🔹 <b>CanPromoteMembers</b> - Add new admins\n" +
		"🔹 <b>CanChangeInfo</b> - Edit chat info\n" +
		"🔹 <b>CanInviteUsers</b> - Add new members\n" +
		"🔹 <b>CanPostMessages</b> - Post in channels\n" +
		"🔹 <b>CanEditMessages</b> - Edit any messages\n" +
		"🔹 <b>CanPinMessages</b> - Pin/unpin messages\n" +
		"🔹 <b>CanManageTopics</b> - Manage forum topics\n\n" +
		"Use the promote user function to assign these rights.").
		HTML().Send().Err()
}

// ================ INVITE LINKS MANAGEMENT ================

func handleInviteLinks(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("➕ Create Invite", "create_invite").
		Text("❌ Revoke Invite", "revoke_invite").
		Row().
		Text("🔙 Back", "back_admin")

	return ctx.EditMessageText("🔗 <b>Invite Links Management</b>\n\n" +
		"Create and manage invitation links:\n\n" +
		"➕ <b>Create Invite</b> - Generate new invitation link\n" +
		"❌ <b>Revoke Invite</b> - Disable existing link\n\n" +
		"<b>Link Features:</b>\n" +
		"• Set expiration date\n" +
		"• Limit member count\n" +
		"• Require approval\n" +
		"• Custom name for organization").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleCreateInvite(ctx *ctx.Context) error {
	// Create invite link with 1-week expiry and 50 member limit
	result := ctx.CreateChatInviteLink().
		Name("Weekly Demo Invite").
		ExpiresIn(7 * 24 * time.Hour).
		MemberLimit(50).
		Send()

	if result.IsErr() {
		return ctx.Reply(String("❌ Failed to create invite link: " + result.Err().Error())).Send().Err()
	}

	inviteLink := result.Ok()

	return ctx.Reply(String("➕ <b>Invite Link Created Successfully</b>\n\n" +
		"<b>Link:</b> <code>" + inviteLink.InviteLink + "</code>\n" +
		"<b>Name:</b> " + inviteLink.Name + "\n" +
		"<b>Creator:</b> " + inviteLink.Creator.FirstName + "\n" +
		"<b>Member Limit:</b> " + Int(inviteLink.MemberLimit).String().Std() + "\n" +
		"<b>Expires:</b> " + time.Unix(inviteLink.ExpireDate, 0).Format("2006-01-02 15:04:05") + "\n\n" +
		"Share this link to invite new members!")).
		HTML().Send().Err()
}

func handleRevokeInvite(ctx *ctx.Context) error {
	return ctx.Reply("❌ <b>Revoke Invite Link</b>\n\n" +
		"To revoke an invite link:\n\n" +
		"<b>Method 1:</b> Use <code>ctx.RevokeChatInviteLink(inviteLink)</code>\n" +
		"<b>Method 2:</b> Get all links with <code>ctx.ExportChatInviteLink()</code> and revoke specific ones\n\n" +
		"<i>In a real implementation, you would list existing links and allow selection for revocation.</i>").
		HTML().Send().Err()
}
