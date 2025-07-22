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
func (n *SetName) LastName(lastName String) *SetName {
	n.opts.LastName = lastName.Std()
	return n
}

// Timeout sets a custom timeout for this request.
func (n *SetName) Timeout(duration time.Duration) *SetName {
	if n.opts.RequestOpts == nil {
		n.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	n.opts.RequestOpts.Timeout = duration

	return n
}

// APIURL sets a custom API URL for this request.
func (n *SetName) APIURL(url String) *SetName {
	if n.opts.RequestOpts == nil {
		n.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	n.opts.RequestOpts.APIURL = url.Std()

	return n
}

// Send executes the SetName request.
func (n *SetName) Send() Result[bool] {
	return ResultOf(n.account.bot.Raw().SetBusinessAccountName(
		n.account.connID.Std(),
		n.firstName.Std(),
		n.opts,
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
func (u *SetUsername) Timeout(duration time.Duration) *SetUsername {
	if u.opts.RequestOpts == nil {
		u.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	u.opts.RequestOpts.Timeout = duration

	return u
}

// APIURL sets a custom API URL for this request.
func (u *SetUsername) APIURL(url String) *SetUsername {
	if u.opts.RequestOpts == nil {
		u.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	u.opts.RequestOpts.APIURL = url.Std()

	return u
}

// Send executes the SetUsername request.
func (u *SetUsername) Send() Result[bool] {
	return ResultOf(u.account.bot.Raw().SetBusinessAccountUsername(
		u.account.connID.Std(),
		u.opts,
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
func (b *SetBio) Timeout(duration time.Duration) *SetBio {
	if b.opts.RequestOpts == nil {
		b.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	b.opts.RequestOpts.Timeout = duration

	return b
}

// APIURL sets a custom API URL for this request.
func (b *SetBio) APIURL(url String) *SetBio {
	if b.opts.RequestOpts == nil {
		b.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	b.opts.RequestOpts.APIURL = url.Std()

	return b
}

// Send executes the SetBio request.
func (b *SetBio) Send() Result[bool] {
	return ResultOf(b.account.bot.Raw().SetBusinessAccountBio(
		b.account.connID.Std(),
		b.opts,
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
func (p *SetAnimatedPhoto) MainFrame(timestamp float64) *SetAnimatedPhoto {
	p.mainFrameTimestamp = Some(timestamp)
	return p
}

// Public marks the profile photo as publicly visible.
func (p *SetAnimatedPhoto) Public() *SetAnimatedPhoto {
	p.opts.IsPublic = true
	return p
}

// Timeout sets a custom timeout for this request.
func (p *SetAnimatedPhoto) Timeout(duration time.Duration) *SetAnimatedPhoto {
	if p.opts.RequestOpts == nil {
		p.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	p.opts.RequestOpts.Timeout = duration

	return p
}

// APIURL sets a custom API URL for this request.
func (p *SetAnimatedPhoto) APIURL(url String) *SetAnimatedPhoto {
	if p.opts.RequestOpts == nil {
		p.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	p.opts.RequestOpts.APIURL = url.Std()

	return p
}

// Send executes the SetAnimatedPhoto request.
func (p *SetAnimatedPhoto) Send() Result[bool] {
	animated := gotgbot.InputProfilePhotoAnimated{
		Animation: p.animation.Std(),
	}

	if p.mainFrameTimestamp.IsSome() {
		animated.MainFrameTimestamp = p.mainFrameTimestamp.Some()
	}

	return ResultOf(p.account.bot.Raw().SetBusinessAccountProfilePhoto(
		p.account.connID.Std(),
		animated,
		p.opts,
	))
}

// SetPhoto is a request builder for setting the business account profile photo.
type SetPhoto struct {
	account *Account
	photo   String
	opts    *gotgbot.SetBusinessAccountProfilePhotoOpts
}

// Public marks the profile photo as publicly visible.
func (p *SetPhoto) Public() *SetPhoto {
	p.opts.IsPublic = true
	return p
}

// Timeout sets a custom timeout for this request.
func (p *SetPhoto) Timeout(duration time.Duration) *SetPhoto {
	if p.opts.RequestOpts == nil {
		p.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	p.opts.RequestOpts.Timeout = duration

	return p
}

// APIURL sets a custom API URL for this request.
func (p *SetPhoto) APIURL(url String) *SetPhoto {
	if p.opts.RequestOpts == nil {
		p.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	p.opts.RequestOpts.APIURL = url.Std()

	return p
}

// Send executes the SetPhoto request.
func (p *SetPhoto) Send() Result[bool] {
	return ResultOf(p.account.bot.Raw().SetBusinessAccountProfilePhoto(
		p.account.connID.Std(),
		gotgbot.InputProfilePhotoStatic{Photo: p.photo.Std()},
		p.opts,
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
func (r *RemovePhoto) Public() *RemovePhoto {
	r.opts.IsPublic = true
	return r
}

// Timeout sets a custom timeout for this request.
func (r *RemovePhoto) Timeout(duration time.Duration) *RemovePhoto {
	if r.opts.RequestOpts == nil {
		r.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	r.opts.RequestOpts.Timeout = duration

	return r
}

// APIURL sets a custom API URL for this request.
func (r *RemovePhoto) APIURL(url String) *RemovePhoto {
	if r.opts.RequestOpts == nil {
		r.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	r.opts.RequestOpts.APIURL = url.Std()

	return r
}

// Send executes the RemovePhoto request.
func (r *RemovePhoto) Send() Result[bool] {
	return ResultOf(r.account.bot.Raw().RemoveBusinessAccountProfilePhoto(
		r.account.connID.Std(),
		r.opts,
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
