package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// VerifyUser represents a request to verify a user.
type VerifyUser struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.VerifyUserOpts
}

// Timeout sets a custom timeout for this request.
func (vu *VerifyUser) Timeout(duration time.Duration) *VerifyUser {
	if vu.opts.RequestOpts == nil {
		vu.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	vu.opts.RequestOpts.Timeout = duration

	return vu
}

// APIURL sets a custom API URL for this request.
func (vu *VerifyUser) APIURL(url String) *VerifyUser {
	if vu.opts.RequestOpts == nil {
		vu.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	vu.opts.RequestOpts.APIURL = url.Std()

	return vu
}

// Send verifies the user.
func (vu *VerifyUser) Send() Result[bool] {
	return ResultOf(vu.ctx.Bot.Raw().VerifyUser(vu.userID, vu.opts))
}

// VerifyChat represents a request to verify a chat.
type VerifyChat struct {
	ctx    *Context
	chatID int64
	opts   *gotgbot.VerifyChatOpts
}

// Timeout sets a custom timeout for this request.
func (vc *VerifyChat) Timeout(duration time.Duration) *VerifyChat {
	if vc.opts.RequestOpts == nil {
		vc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	vc.opts.RequestOpts.Timeout = duration

	return vc
}

// APIURL sets a custom API URL for this request.
func (vc *VerifyChat) APIURL(url String) *VerifyChat {
	if vc.opts.RequestOpts == nil {
		vc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	vc.opts.RequestOpts.APIURL = url.Std()

	return vc
}

// Send verifies the chat.
func (vc *VerifyChat) Send() Result[bool] {
	return ResultOf(vc.ctx.Bot.Raw().VerifyChat(vc.chatID, vc.opts))
}

// RemoveUserVerification represents a request to remove user verification.
type RemoveUserVerification struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.RemoveUserVerificationOpts
}

// Timeout sets a custom timeout for this request.
func (ruv *RemoveUserVerification) Timeout(duration time.Duration) *RemoveUserVerification {
	if ruv.opts.RequestOpts == nil {
		ruv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ruv.opts.RequestOpts.Timeout = duration

	return ruv
}

// APIURL sets a custom API URL for this request.
func (ruv *RemoveUserVerification) APIURL(url String) *RemoveUserVerification {
	if ruv.opts.RequestOpts == nil {
		ruv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ruv.opts.RequestOpts.APIURL = url.Std()

	return ruv
}

// Send removes user verification.
func (ruv *RemoveUserVerification) Send() Result[bool] {
	return ResultOf(ruv.ctx.Bot.Raw().RemoveUserVerification(ruv.userID, ruv.opts))
}

// RemoveChatVerification represents a request to remove chat verification.
type RemoveChatVerification struct {
	ctx    *Context
	chatID int64
	opts   *gotgbot.RemoveChatVerificationOpts
}

// Timeout sets a custom timeout for this request.
func (rcv *RemoveChatVerification) Timeout(duration time.Duration) *RemoveChatVerification {
	if rcv.opts.RequestOpts == nil {
		rcv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rcv.opts.RequestOpts.Timeout = duration

	return rcv
}

// APIURL sets a custom API URL for this request.
func (rcv *RemoveChatVerification) APIURL(url String) *RemoveChatVerification {
	if rcv.opts.RequestOpts == nil {
		rcv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rcv.opts.RequestOpts.APIURL = url.Std()

	return rcv
}

// Send removes chat verification.
func (rcv *RemoveChatVerification) Send() Result[bool] {
	return ResultOf(rcv.ctx.Bot.Raw().RemoveChatVerification(rcv.chatID, rcv.opts))
}
