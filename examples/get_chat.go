package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Command to get current chat info
	b.Command("chatinfo", func(ctx *ctx.Context) error {
		result := ctx.GetChat().Send()

		if result.IsErr() {
			return ctx.Reply(g.Format("Error getting chat info: {}", result.Err())).Send().Err()
		}

		chat := result.Ok()
		info := g.Format(`
📊 <b>Chat Information</b>

🆔 <b>ID:</b> <code>{}</code>
📝 <b>Type:</b> {}
👥 <b>Title:</b> {}
📖 <b>Description:</b> {}
🔗 <b>Username:</b> @{}
📍 <b>Location:</b> {}
🎭 <b>Emoji status:</b> {}
		`,
			chat.Id,
			chat.Type,
			chat.Title,
			chat.Description,
			chat.Username,
			func() g.String {
				if chat.Location != nil {
					return g.Format("{}, {}", chat.Location.Location.Latitude, chat.Location.Location.Longitude)
				}
				return "N/A"
			}(),
			func() string {
				if chat.EmojiStatusCustomEmojiId != "" {
					return chat.EmojiStatusCustomEmojiId
				}
				return "None"
			}(),
		)

		return ctx.Reply(info).HTML().Send().Err()
	})

	// Command to get info about another chat
	b.Command("getchat", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if len(args) == 0 {
			return ctx.Reply("Usage: /getchat <chat_id_or_username>").Send().Err()
		}

		chatArg := args[0]

		// Try to parse as int64 first, if fails use as username
		if chatID := chatArg.ToInt().UnwrapOrDefault().Int64(); chatID != 0 {
			result := ctx.GetChat().ChatID(chatID).Send()

			if result.IsErr() {
				return ctx.Reply(g.Format("Error getting chat info: {}", result.Err())).Send().Err()
			}

			chat := result.Ok()
			return ctx.Reply(g.Format("Chat: {} ({})", chat.Title, chat.Type)).Send().Err()
		}

		return ctx.Reply("Please provide a valid chat ID (number)").Send().Err()
	})

	// Command to check if user is admin in current chat
	b.Command("checkadmin", func(ctx *ctx.Context) error {
		chatInfo := ctx.GetChat().Send()

		if chatInfo.IsErr() {
			return ctx.Reply("Failed to get chat info").Send().Err()
		}

		chat := chatInfo.Ok()

		// Check different chat types and permissions
		switch chat.Type {
		case "private":
			return ctx.Reply("This is a private chat - no admin concept").Send().Err()
		case "group", "supergroup":
			permissions := chat.Permissions
			if permissions != nil {
				return ctx.Reply(g.Format(`
🔒 <b>Chat Permissions:</b>

💬 Can send messages: {}
🖼 Can send media: {}
📊 Can send polls: {}
🔗 Can add web page previews: {}
👥 Can invite users: {}
📌 Can pin messages: {}
📝 Can change info: {}
				`,
					permissions.CanSendMessages,
					permissions.CanSendPhotos && permissions.CanSendVideos,
					permissions.CanSendPolls,
					permissions.CanAddWebPagePreviews,
					permissions.CanInviteUsers,
					permissions.CanPinMessages,
					permissions.CanChangeInfo,
				)).HTML().Send().Err()
			}
		case "channel":
			return ctx.Reply("This is a channel").Send().Err()
		}

		return ctx.Reply("Unknown chat type").Send().Err()
	})

	b.Polling().Start()
}
