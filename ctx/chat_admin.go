package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/types/permissions"
)

// SetChatTitle represents a request to set the chat title.
type SetChatTitle struct {
	ctx    *Context
	title  String
	opts   *gotgbot.SetChatTitleOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID for this request.
func (sat *SetChatTitle) ChatID(id int64) *SetChatTitle {
	sat.chatID = Some(id)
	return sat
}

// Timeout sets a custom timeout for this request.
func (sat *SetChatTitle) Timeout(duration time.Duration) *SetChatTitle {
	if sat.opts.RequestOpts == nil {
		sat.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sat.opts.RequestOpts.Timeout = duration

	return sat
}

// APIURL sets a custom API URL for this request.
func (sat *SetChatTitle) APIURL(url String) *SetChatTitle {
	if sat.opts.RequestOpts == nil {
		sat.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sat.opts.RequestOpts.APIURL = url.Std()

	return sat
}

// Send executes the SetChatTitle request.
func (sat *SetChatTitle) Send() Result[bool] {
	chatID := sat.chatID.UnwrapOr(sat.ctx.EffectiveChat.Id)
	return ResultOf(sat.ctx.Bot.Raw().SetChatTitle(chatID, sat.title.Std(), sat.opts))
}

// SetChatDescription represents a request to set the chat description.
type SetChatDescription struct {
	ctx         *Context
	description String
	opts        *gotgbot.SetChatDescriptionOpts
	chatID      Option[int64]
}

// ChatID sets the target chat ID for this request.
func (scd *SetChatDescription) ChatID(id int64) *SetChatDescription {
	scd.chatID = Some(id)
	return scd
}

// Timeout sets a custom timeout for this request.
func (scd *SetChatDescription) Timeout(duration time.Duration) *SetChatDescription {
	if scd.opts.RequestOpts == nil {
		scd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scd.opts.RequestOpts.Timeout = duration

	return scd
}

// APIURL sets a custom API URL for this request.
func (scd *SetChatDescription) APIURL(url String) *SetChatDescription {
	if scd.opts.RequestOpts == nil {
		scd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scd.opts.RequestOpts.APIURL = url.Std()

	return scd
}

// Send executes the SetChatDescription request.
func (scd *SetChatDescription) Send() Result[bool] {
	chatID := scd.chatID.UnwrapOr(scd.ctx.EffectiveChat.Id)
	return ResultOf(scd.ctx.Bot.Raw().SetChatDescription(chatID, scd.opts))
}

// SetChatPhoto represents a request to set the chat photo.
type SetChatPhoto struct {
	ctx    *Context
	opts   *gotgbot.SetChatPhotoOpts
	doc    gotgbot.InputFile
	file   *File
	chatID Option[int64]
	err    error
}

// ChatID sets the target chat ID for this request.
func (scp *SetChatPhoto) ChatID(id int64) *SetChatPhoto {
	scp.chatID = Some(id)
	return scp
}

// Timeout sets a custom timeout for this request.
func (scp *SetChatPhoto) Timeout(duration time.Duration) *SetChatPhoto {
	if scp.opts.RequestOpts == nil {
		scp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scp.opts.RequestOpts.Timeout = duration

	return scp
}

// APIURL sets a custom API URL for this request.
func (scp *SetChatPhoto) APIURL(url String) *SetChatPhoto {
	if scp.opts.RequestOpts == nil {
		scp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scp.opts.RequestOpts.APIURL = url.Std()

	return scp
}

// Send executes the SetChatPhoto request.
func (scp *SetChatPhoto) Send() Result[bool] {
	if scp.err != nil {
		return Err[bool](scp.err)
	}

	if scp.file != nil {
		defer scp.file.Close()
	}

	chatID := scp.chatID.UnwrapOr(scp.ctx.EffectiveChat.Id)

	return ResultOf(scp.ctx.Bot.Raw().SetChatPhoto(chatID, scp.doc, scp.opts))
}

// DeleteChatPhoto represents a request to delete the chat photo.
type DeleteChatPhoto struct {
	ctx    *Context
	opts   *gotgbot.DeleteChatPhotoOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID for this request.
func (dcp *DeleteChatPhoto) ChatID(id int64) *DeleteChatPhoto {
	dcp.chatID = Some(id)
	return dcp
}

// Timeout sets a custom timeout for this request.
func (dcp *DeleteChatPhoto) Timeout(duration time.Duration) *DeleteChatPhoto {
	if dcp.opts.RequestOpts == nil {
		dcp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dcp.opts.RequestOpts.Timeout = duration

	return dcp
}

// APIURL sets a custom API URL for this request.
func (dcp *DeleteChatPhoto) APIURL(url String) *DeleteChatPhoto {
	if dcp.opts.RequestOpts == nil {
		dcp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dcp.opts.RequestOpts.APIURL = url.Std()

	return dcp
}

// Send executes the DeleteChatPhoto request.
func (dcp *DeleteChatPhoto) Send() Result[bool] {
	chatID := dcp.chatID.UnwrapOr(dcp.ctx.EffectiveChat.Id)
	return ResultOf(dcp.ctx.Bot.Raw().DeleteChatPhoto(chatID, dcp.opts))
}

// SetChatPermissions represents a request to set default chat permissions.
type SetChatPermissions struct {
	ctx             *Context
	permissions     *gotgbot.ChatPermissions
	autoPermissions bool
	opts            *gotgbot.SetChatPermissionsOpts
	chatID          Option[int64]
}

// ChatID sets the target chat ID for this request.
func (scp *SetChatPermissions) ChatID(id int64) *SetChatPermissions {
	scp.chatID = Some(id)
	return scp
}

// AutoPermissions uses chat default permissions instead of independent permissions.
func (scp *SetChatPermissions) AutoPermissions() *SetChatPermissions {
	scp.autoPermissions = true
	return scp
}

// Timeout sets a custom timeout for this request.
func (scp *SetChatPermissions) Timeout(duration time.Duration) *SetChatPermissions {
	if scp.opts.RequestOpts == nil {
		scp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scp.opts.RequestOpts.Timeout = duration

	return scp
}

// APIURL sets a custom API URL for this request.
func (scp *SetChatPermissions) APIURL(url String) *SetChatPermissions {
	if scp.opts.RequestOpts == nil {
		scp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scp.opts.RequestOpts.APIURL = url.Std()

	return scp
}

// Permissions sets the allowed permissions for the chat.
func (scp *SetChatPermissions) Permissions(perms ...permissions.Permission) *SetChatPermissions {
	scp.permissions = permissions.Permissions(perms...)
	return scp
}

// Send executes the SetChatPermissions request.
func (scp *SetChatPermissions) Send() Result[bool] {
	if scp.permissions == nil {
		return Err[bool](Errorf("permissions are required"))
	}

	chatID := scp.chatID.UnwrapOr(scp.ctx.EffectiveChat.Id)
	scp.opts.UseIndependentChatPermissions = !scp.autoPermissions

	return ResultOf(scp.ctx.Bot.Raw().SetChatPermissions(chatID, *scp.permissions, scp.opts))
}

// SetChatAdministratorCustomTitle represents a request to set a custom title for an administrator.
type SetChatAdministratorCustomTitle struct {
	ctx         *Context
	userID      int64
	customTitle String
	opts        *gotgbot.SetChatAdministratorCustomTitleOpts
	chatID      Option[int64]
}

// ChatID sets the target chat ID for this request.
func (scact *SetChatAdministratorCustomTitle) ChatID(id int64) *SetChatAdministratorCustomTitle {
	scact.chatID = Some(id)
	return scact
}

// Timeout sets a custom timeout for this request.
func (scact *SetChatAdministratorCustomTitle) Timeout(duration time.Duration) *SetChatAdministratorCustomTitle {
	if scact.opts.RequestOpts == nil {
		scact.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scact.opts.RequestOpts.Timeout = duration

	return scact
}

// APIURL sets a custom API URL for this request.
func (scact *SetChatAdministratorCustomTitle) APIURL(url String) *SetChatAdministratorCustomTitle {
	if scact.opts.RequestOpts == nil {
		scact.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scact.opts.RequestOpts.APIURL = url.Std()

	return scact
}

// Send executes the SetChatAdministratorCustomTitle request.
func (scact *SetChatAdministratorCustomTitle) Send() Result[bool] {
	chatID := scact.chatID.UnwrapOr(scact.ctx.EffectiveChat.Id)
	return ResultOf(
		scact.ctx.Bot.Raw().SetChatAdministratorCustomTitle(chatID, scact.userID, scact.customTitle.Std(), scact.opts),
	)
}

// PinChatMessage represents a request to pin a message.
type PinChatMessage struct {
	ctx       *Context
	messageID int64
	opts      *gotgbot.PinChatMessageOpts
	chatID    Option[int64]
}

// ChatID sets the target chat ID for this request.
func (pcm *PinChatMessage) ChatID(id int64) *PinChatMessage {
	pcm.chatID = Some(id)
	return pcm
}

// Business sets the business connection ID for the pin action.
func (pcm *PinChatMessage) Business(id String) *PinChatMessage {
	pcm.opts.BusinessConnectionId = id.Std()
	return pcm
}

// Silent sets whether to disable notification.
func (pcm *PinChatMessage) Silent() *PinChatMessage {
	pcm.opts.DisableNotification = true
	return pcm
}

// Timeout sets a custom timeout for this request.
func (pcm *PinChatMessage) Timeout(duration time.Duration) *PinChatMessage {
	if pcm.opts.RequestOpts == nil {
		pcm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	pcm.opts.RequestOpts.Timeout = duration

	return pcm
}

// APIURL sets a custom API URL for this request.
func (pcm *PinChatMessage) APIURL(url String) *PinChatMessage {
	if pcm.opts.RequestOpts == nil {
		pcm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	pcm.opts.RequestOpts.APIURL = url.Std()

	return pcm
}

// Send executes the PinChatMessage request.
func (pcm *PinChatMessage) Send() Result[bool] {
	chatID := pcm.chatID.UnwrapOr(pcm.ctx.EffectiveChat.Id)
	return ResultOf(pcm.ctx.Bot.Raw().PinChatMessage(chatID, pcm.messageID, pcm.opts))
}

// UnpinChatMessage represents a request to unpin a message.
type UnpinChatMessage struct {
	ctx    *Context
	opts   *gotgbot.UnpinChatMessageOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID for this request.
func (ucm *UnpinChatMessage) ChatID(id int64) *UnpinChatMessage {
	ucm.chatID = Some(id)
	return ucm
}

// MessageID sets the specific message ID to unpin.
func (ucm *UnpinChatMessage) MessageID(messageID int64) *UnpinChatMessage {
	ucm.opts.MessageId = &messageID
	return ucm
}

// Business sets the business connection ID for the unpin action.
func (ucm *UnpinChatMessage) Business(id String) *UnpinChatMessage {
	ucm.opts.BusinessConnectionId = id.Std()
	return ucm
}

// Timeout sets a custom timeout for this request.
func (ucm *UnpinChatMessage) Timeout(duration time.Duration) *UnpinChatMessage {
	if ucm.opts.RequestOpts == nil {
		ucm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ucm.opts.RequestOpts.Timeout = duration

	return ucm
}

// APIURL sets a custom API URL for this request.
func (ucm *UnpinChatMessage) APIURL(url String) *UnpinChatMessage {
	if ucm.opts.RequestOpts == nil {
		ucm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ucm.opts.RequestOpts.APIURL = url.Std()

	return ucm
}

// Send executes the UnpinChatMessage request.
func (ucm *UnpinChatMessage) Send() Result[bool] {
	chatID := ucm.chatID.UnwrapOr(ucm.ctx.EffectiveChat.Id)
	return ResultOf(ucm.ctx.Bot.Raw().UnpinChatMessage(chatID, ucm.opts))
}

// UnpinAllChatMessages represents a request to unpin all messages.
type UnpinAllChatMessages struct {
	ctx    *Context
	opts   *gotgbot.UnpinAllChatMessagesOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID for this request.
func (uacm *UnpinAllChatMessages) ChatID(id int64) *UnpinAllChatMessages {
	uacm.chatID = Some(id)
	return uacm
}

// Timeout sets a custom timeout for this request.
func (uacm *UnpinAllChatMessages) Timeout(duration time.Duration) *UnpinAllChatMessages {
	if uacm.opts.RequestOpts == nil {
		uacm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	uacm.opts.RequestOpts.Timeout = duration

	return uacm
}

// APIURL sets a custom API URL for this request.
func (uacm *UnpinAllChatMessages) APIURL(url String) *UnpinAllChatMessages {
	if uacm.opts.RequestOpts == nil {
		uacm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	uacm.opts.RequestOpts.APIURL = url.Std()

	return uacm
}

// Send executes the UnpinAllChatMessages request.
func (uacm *UnpinAllChatMessages) Send() Result[bool] {
	chatID := uacm.chatID.UnwrapOr(uacm.ctx.EffectiveChat.Id)
	return ResultOf(uacm.ctx.Bot.Raw().UnpinAllChatMessages(chatID, uacm.opts))
}

// GetChatAdministrators represents a request to get chat administrators.
type GetChatAdministrators struct {
	ctx    *Context
	opts   *gotgbot.GetChatAdministratorsOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID for this request.
func (gca *GetChatAdministrators) ChatID(id int64) *GetChatAdministrators {
	gca.chatID = Some(id)
	return gca
}

// Timeout sets a custom timeout for this request.
func (gca *GetChatAdministrators) Timeout(duration time.Duration) *GetChatAdministrators {
	if gca.opts.RequestOpts == nil {
		gca.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gca.opts.RequestOpts.Timeout = duration

	return gca
}

// APIURL sets a custom API URL for this request.
func (gca *GetChatAdministrators) APIURL(url String) *GetChatAdministrators {
	if gca.opts.RequestOpts == nil {
		gca.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gca.opts.RequestOpts.APIURL = url.Std()

	return gca
}

// Send executes the GetChatAdministrators request.
func (gca *GetChatAdministrators) Send() Result[Slice[gotgbot.ChatMember]] {
	chatID := gca.chatID.UnwrapOr(gca.ctx.EffectiveChat.Id)

	members, err := gca.ctx.Bot.Raw().GetChatAdministrators(chatID, gca.opts)
	return ResultOf(Slice[gotgbot.ChatMember](members), err)
}

// GetChatMemberCount represents a request to get the chat member count.
type GetChatMemberCount struct {
	ctx    *Context
	opts   *gotgbot.GetChatMemberCountOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID for this request.
func (gcm *GetChatMemberCount) ChatID(id int64) *GetChatMemberCount {
	gcm.chatID = Some(id)
	return gcm
}

// Timeout sets a custom timeout for this request.
func (gcm *GetChatMemberCount) Timeout(duration time.Duration) *GetChatMemberCount {
	if gcm.opts.RequestOpts == nil {
		gcm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gcm.opts.RequestOpts.Timeout = duration

	return gcm
}

// APIURL sets a custom API URL for this request.
func (gcm *GetChatMemberCount) APIURL(url String) *GetChatMemberCount {
	if gcm.opts.RequestOpts == nil {
		gcm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gcm.opts.RequestOpts.APIURL = url.Std()

	return gcm
}

// Send executes the GetChatMemberCount request.
func (gcm *GetChatMemberCount) Send() Result[Int] {
	chatID := gcm.chatID.UnwrapOr(gcm.ctx.EffectiveChat.Id)
	count, err := gcm.ctx.Bot.Raw().GetChatMemberCount(chatID, gcm.opts)

	return ResultOf(Int(count), err)
}
