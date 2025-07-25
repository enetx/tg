package main

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Command to get member info about yourself
	b.Command("me", func(ctx *ctx.Context) error {
		result := ctx.GetChatMember(ctx.EffectiveUser.Id).Send()

		if result.IsErr() {
			return ctx.Reply(Format("Error getting member info: {}", result.Err())).Send().Err()
		}

		member := result.Ok()
		status := getMemberStatus(member)
		user := member.GetUser()

		info := Format(`
👤 <b>Your Status in This Chat</b>

🆔 <b>User ID:</b> <code>{}</code>
👤 <b>Name:</b> {} {}
🏷 <b>Username:</b> @{}
🎖 <b>Status:</b> {}
		`,
			user.Id,
			user.FirstName,
			user.LastName,
			user.Username,
			status,
		)

		return ctx.Reply(info).HTML().Send().Err()
	})

	b.Polling().Start()
}

// Helper function to get member status string
func getMemberStatus(member gotgbot.ChatMember) String {
	switch member.(type) {
	case *gotgbot.ChatMemberOwner:
		return "👑 Owner"
	case *gotgbot.ChatMemberAdministrator:
		return "🛡 Administrator"
	case *gotgbot.ChatMemberMember:
		return "👤 Member"
	case *gotgbot.ChatMemberRestricted:
		return "🚫 Restricted"
	case *gotgbot.ChatMemberLeft:
		return "👋 Left"
	case *gotgbot.ChatMemberBanned:
		return "⛔ Banned"
	default:
		return "❓ Unknown"
	}
}
