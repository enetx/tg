package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Handle /create_invite command
	b.Command("create_invite", func(ctx *ctx.Context) error {
		// Create a new invite link
		result := ctx.CreateChatInviteLink().
			Name("My Special Invite").
			MemberLimit(50).
			// CreatesJoinRequest().
			Send()

		if result.IsErr() {
			return ctx.Reply(Format("Failed to create invite link: {}", result.Err())).Send().Err()
		}

		inviteLink := result.Ok()
		return ctx.Reply("Created invite link: " + String(inviteLink.InviteLink)).Send().Err()
	})

	// Handle /edit_invite command
	b.Command("edit_invite", func(ctx *ctx.Context) error {
		// Assuming we have a stored invite link (in real app, you'd store this)
		inviteLink := String("https://t.me/+example")

		result := ctx.EditChatInviteLink(inviteLink).
			Name("Updated Invite Link").
			MemberLimit(100).
			Send()

		if result.IsErr() {
			return ctx.Reply(Format("Failed to edit invite link: {}", result.Err())).Send().Err()
		}

		editedLink := result.Ok()
		return ctx.Reply("Edited invite link: " + String(editedLink.InviteLink)).Send().Err()
	})

	// Handle /revoke_invite command
	b.Command("revoke_invite", func(ctx *ctx.Context) error {
		// Assuming we have a stored invite link
		inviteLink := String("https://t.me/+example")

		result := ctx.RevokeChatInviteLink(inviteLink).Send()

		if result.IsErr() {
			return ctx.Reply(Format("Failed to revoke invite link: {}", result.Err())).Send().Err()
		}

		revokedLink := result.Ok()
		return ctx.Reply("Revoked invite link: " + String(revokedLink.InviteLink)).Send().Err()
	})

	// Handle chat join requests
	b.On.ChatJoinRequest.Any(func(ctx *ctx.Context) error {
		user := ctx.Update.ChatJoinRequest.From

		// Auto-approve join requests for demonstration
		result := ctx.ApproveChatJoinRequest(user.Id).Send()

		if result.IsErr() {
			Println("Failed to approve join request from {}:{}", user.FirstName, result.Err())
			return result.Err()
		}

		Println("Approved join request from {1.FirstName} {1.LastName}", user)
		return nil
	})

	// Handle /decline_request command (admin only)
	b.Command("decline_request", func(ctx *ctx.Context) error {
		// In a real bot, you'd get the user ID from command arguments
		// For demo purposes, we'll use a placeholder
		if len(ctx.Args()) < 1 {
			return ctx.Reply("Usage: /decline_request <user_id>").Send().Err()
		}

		userID := ctx.Args()[0].ToInt().UnwrapOrDefault().Int64()
		if userID == 0 {
			return ctx.Reply("Invalid user ID").Send().Err()
		}

		result := ctx.DeclineChatJoinRequest(userID).Send()

		if result.IsErr() {
			return ctx.Reply(Format("Failed to decline join request: {}", result.Err())).Send().Err()
		}

		return ctx.Reply("Successfully declined join request").Send().Err()
	})

	b.Polling().Start()
}
