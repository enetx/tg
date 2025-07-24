package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/types/permissions"
)

func main() {
	b := bot.New("YOUR_BOT_TOKEN").Build().Unwrap()

	// Set chat title
	b.Command("settitle", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /settitle <new_title>").Send().Err()
		}

		title := args.Join(" ")
		return ctx.SetChatTitle(title).Send().Err()
	})

	// Set chat description
	b.Command("setdescription", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /setdescription <description>").Send().Err()
		}

		description := args.Join(" ")
		return ctx.SetChatDescription(description).Send().Err()
	})

	// Set chat photo
	b.Command("setphoto", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /setphoto <photo_path>").Send().Err()
		}

		photo := args.Get(0).Unwrap()
		return ctx.SetChatPhoto(photo).Send().Err()
	})

	// Delete chat photo
	b.Command("deletephoto", func(ctx *ctx.Context) error {
		return ctx.DeleteChatPhoto().Send().Err()
	})

	// Restrict all permissions
	b.Command("restrictall", func(ctx *ctx.Context) error {
		// No permissions - restricts everything
		return ctx.SetChatPermissions().
			Permissions().
			Send().Err()
	})

	// Allow basic messaging permissions
	b.Command("allowbasic", func(ctx *ctx.Context) error {
		return ctx.SetChatPermissions().
			Permissions(
				permissions.SendMessages,
				permissions.SendPhotos,
				permissions.SendVideos,
			).
			Send().Err()
	})

	// Allow all permissions
	b.Command("allowall", func(ctx *ctx.Context) error {
		return ctx.SetChatPermissions().
			Permissions(
				permissions.SendMessages,
				permissions.SendAudios,
				permissions.SendDocuments,
				permissions.SendPhotos,
				permissions.SendVideos,
				permissions.SendVideoNotes,
				permissions.SendVoiceNotes,
				permissions.SendPolls,
				permissions.SendOtherMessages,
				permissions.AddWebPagePreviews,
				permissions.ChangeInfo,
				permissions.InviteUsers,
				permissions.PinMessages,
				permissions.ManageTopics,
			).
			Send().Err()
	})

	// Set permissions with auto permissions mode
	b.Command("autopermissions", func(ctx *ctx.Context) error {
		return ctx.SetChatPermissions().
			Permissions(permissions.SendMessages, permissions.SendPhotos).
			AutoPermissions().
			Send().Err()
	})

	// Set administrator custom title
	b.Command("setadmintitle", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /setadmintitle <user_id> <custom_title>").Send().Err()
		}

		userID := args[0].ToInt().Unwrap().Int64()
		title := args.Iter().Skip(1).Collect().Join(" ")

		return ctx.SetChatAdministratorCustomTitle(userID, title).Send().Err()
	})

	// Pin message
	b.Command("pin", func(ctx *ctx.Context) error {
		if ctx.EffectiveMessage.ReplyToMessage == nil {
			return ctx.Reply("Reply to a message to pin it").Send().Err()
		}

		messageID := ctx.EffectiveMessage.ReplyToMessage.MessageId
		return ctx.PinChatMessage(messageID).Silent().Send().Err()
	})

	// Pin message with business connection
	b.Command("pinbiz", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 || ctx.EffectiveMessage.ReplyToMessage == nil {
			return ctx.Reply("Usage: /pinbiz <business_connection_id> (reply to message)").Send().Err()
		}

		businessConnID := args[0]
		messageID := ctx.EffectiveMessage.ReplyToMessage.MessageId

		return ctx.PinChatMessage(messageID).Business(businessConnID).Silent().Send().Err()
	})

	// Unpin message
	b.Command("unpin", func(ctx *ctx.Context) error {
		if ctx.EffectiveMessage.ReplyToMessage == nil {
			return ctx.Reply("Reply to a message to unpin it").Send().Err()
		}

		messageID := ctx.EffectiveMessage.ReplyToMessage.MessageId
		return ctx.UnpinChatMessage().MessageID(messageID).Send().Err()
	})

	// Unpin latest pinned message
	b.Command("unpinlatest", func(ctx *ctx.Context) error {
		return ctx.UnpinChatMessage().Send().Err()
	})

	// Unpin with business connection
	b.Command("unpinbiz", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /unpinbiz <business_connection_id> [message_id]").Send().Err()
		}

		businessConnID := args[0]
		unpin := ctx.UnpinChatMessage().Business(businessConnID)

		if args.Len() > 1 {
			messageID := args[1].ToInt().Unwrap().Int64()
			unpin = unpin.MessageID(messageID)
		}

		return unpin.Send().Err()
	})

	// Unpin all messages
	b.Command("unpinall", func(ctx *ctx.Context) error {
		return ctx.UnpinAllChatMessages().Send().Err()
	})

	// Get chat administrators
	b.Command("admins", func(ctx *ctx.Context) error {
		admins := ctx.GetChatAdministrators().Send()
		if admins.IsErr() {
			return admins.Err()
		}

		builder := NewBuilder()
		builder.WriteString("Chat Administrators:\n")

		for admin := range admins.Ok().Iter() {
			builder.WriteString(Format("• {.GetUser.FirstName}", admin))
			if username := admin.GetUser().Username; username != "" {
				builder.WriteString(Format("(@{})", username))
			}
			builder.WriteRune('\n')
		}

		return ctx.Reply(builder.String()).Send().Err()
	})

	// Get chat member count
	b.Command("membercount", func(ctx *ctx.Context) error {
		count := ctx.GetChatMemberCount().Send()
		if count.IsErr() {
			return count.Err()
		}

		return ctx.Reply("Chat has " + count.Ok().String() + " members").Send().Err()
	})

	// Examples with different chat IDs
	b.Command("adminsinother", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /adminsinother <chat_id>").Send().Err()
		}

		chatID := args[0].ToInt().Unwrap().Int64()

		admins := ctx.GetChatAdministrators().ChatID(chatID).Send()
		if admins.IsErr() {
			return admins.Err()
		}

		return ctx.Reply("Found " + admins.Ok().Len().String() + " administrators in that chat").
			Send().Err()
	})

	b.Polling().AllowedUpdates().Start()
}
