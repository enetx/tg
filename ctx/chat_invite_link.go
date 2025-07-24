package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// CreateChatInviteLink represents a request to create a new chat invite link.
type CreateChatInviteLink struct {
	ctx    *Context
	opts   *gotgbot.CreateChatInviteLinkOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID.
func (ccil *CreateChatInviteLink) ChatID(chatID int64) *CreateChatInviteLink {
	ccil.chatID = Some(chatID)
	return ccil
}

// Name sets the invite link name.
func (ccil *CreateChatInviteLink) Name(name String) *CreateChatInviteLink {
	ccil.opts.Name = name.Std()
	return ccil
}

// ExpireDate sets the date when the link will expire.
func (ccil *CreateChatInviteLink) ExpireDate(expireDate int64) *CreateChatInviteLink {
	ccil.opts.ExpireDate = expireDate
	return ccil
}

// MemberLimit sets the maximum number of users that can be members simultaneously.
func (ccil *CreateChatInviteLink) MemberLimit(limit int64) *CreateChatInviteLink {
	ccil.opts.MemberLimit = limit
	return ccil
}

// CreatesJoinRequest sets whether users joining via this link need to be approved.
func (ccil *CreateChatInviteLink) CreatesJoinRequest() *CreateChatInviteLink {
	ccil.opts.CreatesJoinRequest = true
	return ccil
}

// Timeout sets a custom timeout for this request.
func (ccil *CreateChatInviteLink) Timeout(duration time.Duration) *CreateChatInviteLink {
	if ccil.opts.RequestOpts == nil {
		ccil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ccil.opts.RequestOpts.Timeout = duration

	return ccil
}

// APIURL sets a custom API URL for this request.
func (ccil *CreateChatInviteLink) APIURL(url String) *CreateChatInviteLink {
	if ccil.opts.RequestOpts == nil {
		ccil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ccil.opts.RequestOpts.APIURL = url.Std()

	return ccil
}

// Send creates the chat invite link and returns the result.
func (ccil *CreateChatInviteLink) Send() Result[*gotgbot.ChatInviteLink] {
	chatID := ccil.chatID.UnwrapOr(ccil.ctx.EffectiveChat.Id)
	return ResultOf(ccil.ctx.Bot.Raw().CreateChatInviteLink(chatID, ccil.opts))
}

// EditChatInviteLink represents a request to edit an existing chat invite link.
type EditChatInviteLink struct {
	ctx        *Context
	inviteLink String
	opts       *gotgbot.EditChatInviteLinkOpts
	chatID     Option[int64]
}

// ChatID sets the target chat ID.
func (ecil *EditChatInviteLink) ChatID(chatID int64) *EditChatInviteLink {
	ecil.chatID = Some(chatID)
	return ecil
}

// Name sets the invite link name.
func (ecil *EditChatInviteLink) Name(name String) *EditChatInviteLink {
	ecil.opts.Name = name.Std()
	return ecil
}

// ExpireDate sets the date when the link will expire.
func (ecil *EditChatInviteLink) ExpireDate(expireDate int64) *EditChatInviteLink {
	ecil.opts.ExpireDate = expireDate
	return ecil
}

// MemberLimit sets the maximum number of users that can be members simultaneously.
func (ecil *EditChatInviteLink) MemberLimit(limit int64) *EditChatInviteLink {
	ecil.opts.MemberLimit = limit
	return ecil
}

// CreatesJoinRequest sets whether users joining via this link need to be approved.
func (ecil *EditChatInviteLink) CreatesJoinRequest() *EditChatInviteLink {
	ecil.opts.CreatesJoinRequest = true
	return ecil
}

// Timeout sets a custom timeout for this request.
func (ecil *EditChatInviteLink) Timeout(duration time.Duration) *EditChatInviteLink {
	if ecil.opts.RequestOpts == nil {
		ecil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ecil.opts.RequestOpts.Timeout = duration

	return ecil
}

// APIURL sets a custom API URL for this request.
func (ecil *EditChatInviteLink) APIURL(url String) *EditChatInviteLink {
	if ecil.opts.RequestOpts == nil {
		ecil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ecil.opts.RequestOpts.APIURL = url.Std()

	return ecil
}

// Send edits the chat invite link and returns the result.
func (ecil *EditChatInviteLink) Send() Result[*gotgbot.ChatInviteLink] {
	chatID := ecil.chatID.UnwrapOr(ecil.ctx.EffectiveChat.Id)
	return ResultOf(ecil.ctx.Bot.Raw().EditChatInviteLink(chatID, ecil.inviteLink.Std(), ecil.opts))
}

// RevokeChatInviteLink represents a request to revoke a chat invite link.
type RevokeChatInviteLink struct {
	ctx        *Context
	inviteLink String
	opts       *gotgbot.RevokeChatInviteLinkOpts
	chatID     Option[int64]
}

// ChatID sets the target chat ID.
func (rcil *RevokeChatInviteLink) ChatID(chatID int64) *RevokeChatInviteLink {
	rcil.chatID = Some(chatID)
	return rcil
}

// Timeout sets a custom timeout for this request.
func (rcil *RevokeChatInviteLink) Timeout(duration time.Duration) *RevokeChatInviteLink {
	if rcil.opts.RequestOpts == nil {
		rcil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rcil.opts.RequestOpts.Timeout = duration

	return rcil
}

// APIURL sets a custom API URL for this request.
func (rcil *RevokeChatInviteLink) APIURL(url String) *RevokeChatInviteLink {
	if rcil.opts.RequestOpts == nil {
		rcil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rcil.opts.RequestOpts.APIURL = url.Std()

	return rcil
}

// Send revokes the chat invite link and returns the result.
func (rcil *RevokeChatInviteLink) Send() Result[*gotgbot.ChatInviteLink] {
	chatID := rcil.chatID.UnwrapOr(rcil.ctx.EffectiveChat.Id)
	return ResultOf(rcil.ctx.Bot.Raw().RevokeChatInviteLink(chatID, rcil.inviteLink.Std(), rcil.opts))
}

// ApproveChatJoinRequest represents a request to approve a chat join request.
type ApproveChatJoinRequest struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.ApproveChatJoinRequestOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID.
func (acjr *ApproveChatJoinRequest) ChatID(chatID int64) *ApproveChatJoinRequest {
	acjr.chatID = Some(chatID)
	return acjr
}

// Timeout sets a custom timeout for this request.
func (acjr *ApproveChatJoinRequest) Timeout(duration time.Duration) *ApproveChatJoinRequest {
	if acjr.opts.RequestOpts == nil {
		acjr.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	acjr.opts.RequestOpts.Timeout = duration

	return acjr
}

// APIURL sets a custom API URL for this request.
func (acjr *ApproveChatJoinRequest) APIURL(url String) *ApproveChatJoinRequest {
	if acjr.opts.RequestOpts == nil {
		acjr.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	acjr.opts.RequestOpts.APIURL = url.Std()

	return acjr
}

// Send approves the chat join request and returns the result.
func (acjr *ApproveChatJoinRequest) Send() Result[bool] {
	chatID := acjr.chatID.UnwrapOr(acjr.ctx.EffectiveChat.Id)
	return ResultOf(acjr.ctx.Bot.Raw().ApproveChatJoinRequest(chatID, acjr.userID, acjr.opts))
}

// DeclineChatJoinRequest represents a request to decline a chat join request.
type DeclineChatJoinRequest struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.DeclineChatJoinRequestOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID.
func (dcjr *DeclineChatJoinRequest) ChatID(chatID int64) *DeclineChatJoinRequest {
	dcjr.chatID = Some(chatID)
	return dcjr
}

// Timeout sets a custom timeout for this request.
func (dcjr *DeclineChatJoinRequest) Timeout(duration time.Duration) *DeclineChatJoinRequest {
	if dcjr.opts.RequestOpts == nil {
		dcjr.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dcjr.opts.RequestOpts.Timeout = duration

	return dcjr
}

// APIURL sets a custom API URL for this request.
func (dcjr *DeclineChatJoinRequest) APIURL(url String) *DeclineChatJoinRequest {
	if dcjr.opts.RequestOpts == nil {
		dcjr.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dcjr.opts.RequestOpts.APIURL = url.Std()

	return dcjr
}

// Send declines the chat join request and returns the result.
func (dcjr *DeclineChatJoinRequest) Send() Result[bool] {
	chatID := dcjr.chatID.UnwrapOr(dcjr.ctx.EffectiveChat.Id)
	return ResultOf(dcjr.ctx.Bot.Raw().DeclineChatJoinRequest(chatID, dcjr.userID, dcjr.opts))
}
