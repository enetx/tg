package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/types/permissions"
	"github.com/enetx/tg/types/roles"
)

// Member tags (Bot API 9.5): admins can attach a short tag to regular members,
// grant the "manage tags" admin right, and allow members to edit their own tag.
func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// /tag <text> — reply to a user to set their member tag (0-16 chars, no emoji).
	b.Command("tag", func(ctx *ctx.Context) error {
		if ctx.EffectiveMessage.ReplyToMessage == nil || ctx.EffectiveMessage.ReplyToMessage.From == nil {
			return ctx.Reply("Reply to the member you want to tag.").Send().Err()
		}

		userID := ctx.EffectiveMessage.ReplyToMessage.From.Id
		tag := ctx.Args().Join(" ")

		result := ctx.SetChatMemberTag(userID).Tag(tag).Send()
		if result.IsErr() {
			return ctx.Reply(g.Format("Failed to set tag: {}", result.Err())).Send().Err()
		}

		return ctx.Reply(g.Format("Tag set to: {}", tag)).Send().Err()
	})

	// /untag — reply to a user to remove their tag (empty tag clears it).
	b.Command("untag", func(ctx *ctx.Context) error {
		if ctx.EffectiveMessage.ReplyToMessage == nil || ctx.EffectiveMessage.ReplyToMessage.From == nil {
			return ctx.Reply("Reply to the member whose tag you want to clear.").Send().Err()
		}

		userID := ctx.EffectiveMessage.ReplyToMessage.From.Id

		return ctx.SetChatMemberTag(userID).Tag(g.String("")).Send().Err()
	})

	// /grant — promote the replied user to an admin who can manage members' tags.
	b.Command("grant", func(ctx *ctx.Context) error {
		if ctx.EffectiveMessage.ReplyToMessage == nil || ctx.EffectiveMessage.ReplyToMessage.From == nil {
			return ctx.Reply("Reply to the user you want to promote.").Send().Err()
		}

		userID := ctx.EffectiveMessage.ReplyToMessage.From.Id

		return ctx.PromoteChatMember(userID).
			Roles(roles.ManageTags, roles.DeleteMessages).
			Send().Err()
	})

	// /selftag — let regular members react to messages and edit their own tag.
	b.Command("selftag", func(ctx *ctx.Context) error {
		return ctx.SetChatPermissions().
			Permissions(
				permissions.SendMessages,
				permissions.ReactToMessages,
				permissions.EditTag,
			).
			Send().Err()
	})

	b.Polling().Start()
}
