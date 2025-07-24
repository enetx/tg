package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// Bot defines the minimal interface required to perform business account operations.
type Bot interface {
	Raw() *gotgbot.Bot
}

// Account provides high-level methods for managing a Telegram Business account.
type Account struct {
	bot    Bot
	connID String
}

// NewAccount creates a new Account instance bound to the given bot and connection ID.
func NewAccount(bot Bot, connectionID String) *Account {
	return &Account{
		bot:    bot,
		connID: connectionID,
	}
}

// Name returns a builder for setting the account's first and last name.
func (a *Account) SetName(firstName String) *SetName {
	return &SetName{
		account:   a,
		firstName: firstName,
		opts:      new(gotgbot.SetBusinessAccountNameOpts),
	}
}

// SetName is a request builder for setting the account's name.
type SetName struct {
	account   *Account
	firstName String
	opts      *gotgbot.SetBusinessAccountNameOpts
}

// LastName sets the optional last name.
func (sn *SetName) LastName(lastName String) *SetName {
	sn.opts.LastName = lastName.Std()
	return sn
}

// Timeout sets a custom timeout for this request.
func (sn *SetName) Timeout(duration time.Duration) *SetName {
	if sn.opts.RequestOpts == nil {
		sn.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sn.opts.RequestOpts.Timeout = duration

	return sn
}

// APIURL sets a custom API URL for this request.
func (sn *SetName) APIURL(url String) *SetName {
	if sn.opts.RequestOpts == nil {
		sn.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sn.opts.RequestOpts.APIURL = url.Std()

	return sn
}

// Send executes the SetName request.
func (sn *SetName) Send() Result[bool] {
	return ResultOf(sn.account.bot.Raw().SetBusinessAccountName(
		sn.account.connID.Std(),
		sn.firstName.Std(),
		sn.opts,
	))
}

// SetUsername returns a builder for setting the account's username.
func (a *Account) SetUsername(username String) *SetUsername {
	return &SetUsername{
		account: a,
		opts:    &gotgbot.SetBusinessAccountUsernameOpts{Username: username.Std()},
	}
}

// SetUsername is a request builder for setting the account's username.
type SetUsername struct {
	account *Account
	opts    *gotgbot.SetBusinessAccountUsernameOpts
}

// Timeout sets a custom timeout for this request.
func (su *SetUsername) Timeout(duration time.Duration) *SetUsername {
	if su.opts.RequestOpts == nil {
		su.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	su.opts.RequestOpts.Timeout = duration

	return su
}

// APIURL sets a custom API URL for this request.
func (su *SetUsername) APIURL(url String) *SetUsername {
	if su.opts.RequestOpts == nil {
		su.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	su.opts.RequestOpts.APIURL = url.Std()

	return su
}

// Send executes the SetUsername request.
func (su *SetUsername) Send() Result[bool] {
	return ResultOf(su.account.bot.Raw().SetBusinessAccountUsername(
		su.account.connID.Std(),
		su.opts,
	))
}

// SetBio returns a builder for setting the account's biography text.
func (a *Account) SetBio(bio String) *SetBio {
	return &SetBio{
		account: a,
		opts:    &gotgbot.SetBusinessAccountBioOpts{Bio: bio.Std()},
	}
}

// SetBio is a request builder for setting the business account bio.
type SetBio struct {
	account *Account
	opts    *gotgbot.SetBusinessAccountBioOpts
}

// Timeout sets a custom timeout for this request.
func (sb *SetBio) Timeout(duration time.Duration) *SetBio {
	if sb.opts.RequestOpts == nil {
		sb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sb.opts.RequestOpts.Timeout = duration

	return sb
}

// APIURL sets a custom API URL for this request.
func (sb *SetBio) APIURL(url String) *SetBio {
	if sb.opts.RequestOpts == nil {
		sb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sb.opts.RequestOpts.APIURL = url.Std()

	return sb
}

// Send executes the SetBio request.
func (sb *SetBio) Send() Result[bool] {
	return ResultOf(sb.account.bot.Raw().SetBusinessAccountBio(
		sb.account.connID.Std(),
		sb.opts,
	))
}

// SetPhoto returns a builder for setting the account's profile photo.
func (a *Account) SetPhoto(photo String) *SetPhoto {
	return &SetPhoto{
		account: a,
		photo:   photo,
		opts:    new(gotgbot.SetBusinessAccountProfilePhotoOpts),
	}
}

// SetAnimatedPhoto returns a builder for setting the account's animated profile photo.
func (a *Account) SetAnimatedPhoto(animation String) *SetAnimatedPhoto {
	return &SetAnimatedPhoto{
		account:   a,
		animation: animation,
		opts:      new(gotgbot.SetBusinessAccountProfilePhotoOpts),
	}
}

// SetAnimatedPhoto is a request builder for setting the business account animated profile photo.
type SetAnimatedPhoto struct {
	account            *Account
	animation          String
	mainFrameTimestamp Option[float64]
	opts               *gotgbot.SetBusinessAccountProfilePhotoOpts
}

// MainFrame sets the timestamp in seconds of the frame that will be used as the static profile photo.
func (sap *SetAnimatedPhoto) MainFrame(timestamp float64) *SetAnimatedPhoto {
	sap.mainFrameTimestamp = Some(timestamp)
	return sap
}

// Public marks the profile photo as publicly visible.
func (sap *SetAnimatedPhoto) Public() *SetAnimatedPhoto {
	sap.opts.IsPublic = true
	return sap
}

// Timeout sets a custom timeout for this request.
func (sap *SetAnimatedPhoto) Timeout(duration time.Duration) *SetAnimatedPhoto {
	if sap.opts.RequestOpts == nil {
		sap.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sap.opts.RequestOpts.Timeout = duration

	return sap
}

// APIURL sets a custom API URL for this request.
func (sap *SetAnimatedPhoto) APIURL(url String) *SetAnimatedPhoto {
	if sap.opts.RequestOpts == nil {
		sap.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sap.opts.RequestOpts.APIURL = url.Std()

	return sap
}

// Send executes the SetAnimatedPhoto request.
func (sap *SetAnimatedPhoto) Send() Result[bool] {
	animated := gotgbot.InputProfilePhotoAnimated{
		Animation: sap.animation.Std(),
	}

	if sap.mainFrameTimestamp.IsSome() {
		animated.MainFrameTimestamp = sap.mainFrameTimestamp.Some()
	}

	return ResultOf(sap.account.bot.Raw().SetBusinessAccountProfilePhoto(
		sap.account.connID.Std(),
		animated,
		sap.opts,
	))
}

// SetPhoto is a request builder for setting the business account profile photo.
type SetPhoto struct {
	account *Account
	photo   String
	opts    *gotgbot.SetBusinessAccountProfilePhotoOpts
}

// Public marks the profile photo as publicly visible.
func (sp *SetPhoto) Public() *SetPhoto {
	sp.opts.IsPublic = true
	return sp
}

// Timeout sets a custom timeout for this request.
func (sp *SetPhoto) Timeout(duration time.Duration) *SetPhoto {
	if sp.opts.RequestOpts == nil {
		sp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sp.opts.RequestOpts.Timeout = duration

	return sp
}

// APIURL sets a custom API URL for this request.
func (sp *SetPhoto) APIURL(url String) *SetPhoto {
	if sp.opts.RequestOpts == nil {
		sp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sp.opts.RequestOpts.APIURL = url.Std()

	return sp
}

// Send executes the SetPhoto request.
func (sp *SetPhoto) Send() Result[bool] {
	return ResultOf(sp.account.bot.Raw().SetBusinessAccountProfilePhoto(
		sp.account.connID.Std(),
		gotgbot.InputProfilePhotoStatic{Photo: sp.photo.Std()},
		sp.opts,
	))
}

// RemovePhoto returns a builder for removing the account's profile photo.
func (a *Account) RemovePhoto() *RemovePhoto {
	return &RemovePhoto{
		account: a,
		opts:    new(gotgbot.RemoveBusinessAccountProfilePhotoOpts),
	}
}

// RemovePhoto is a request builder for removing the business account profile photo.
type RemovePhoto struct {
	account *Account
	opts    *gotgbot.RemoveBusinessAccountProfilePhotoOpts
}

// Public removes the public profile photo if present.
func (rp *RemovePhoto) Public() *RemovePhoto {
	rp.opts.IsPublic = true
	return rp
}

// Timeout sets a custom timeout for this request.
func (rp *RemovePhoto) Timeout(duration time.Duration) *RemovePhoto {
	if rp.opts.RequestOpts == nil {
		rp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rp.opts.RequestOpts.Timeout = duration

	return rp
}

// APIURL sets a custom API URL for this request.
func (rp *RemovePhoto) APIURL(url String) *RemovePhoto {
	if rp.opts.RequestOpts == nil {
		rp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rp.opts.RequestOpts.APIURL = url.Std()

	return rp
}

// Send executes the RemovePhoto request.
func (rp *RemovePhoto) Send() Result[bool] {
	return ResultOf(rp.account.bot.Raw().RemoveBusinessAccountProfilePhoto(
		rp.account.connID.Std(),
		rp.opts,
	))
}

// Get returns a builder for fetching the current business connection state.
func (a *Account) Get() *Get {
	return &Get{
		account: a,
		opts:    new(gotgbot.GetBusinessConnectionOpts),
	}
}

// Get is a request builder for retrieving business connection info.
type Get struct {
	account *Account
	opts    *gotgbot.GetBusinessConnectionOpts
}

// Timeout sets a custom timeout for this request.
func (g *Get) Timeout(duration time.Duration) *Get {
	if g.opts.RequestOpts == nil {
		g.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	g.opts.RequestOpts.Timeout = duration

	return g
}

// APIURL sets a custom API URL for this request.
func (g *Get) APIURL(url String) *Get {
	if g.opts.RequestOpts == nil {
		g.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	g.opts.RequestOpts.APIURL = url.Std()

	return g
}

// Send executes the Get request.
func (g *Get) Send() Result[*gotgbot.BusinessConnection] {
	return ResultOf(g.account.bot.Raw().GetBusinessConnection(
		g.account.connID.Std(),
		g.opts,
	))
}

// Balance returns a Balance handler for managing stars and gifts.
func (a *Account) Balance() *Balance {
	return &Balance{
		bot:    a.bot,
		connID: a.connID,
	}
}

// Message returns a Message handler for interacting with business messages.
func (a *Account) Message() *Message {
	return &Message{
		bot:    a.bot,
		connID: a.connID,
	}
}
