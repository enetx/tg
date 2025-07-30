package business

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// Account provides high-level methods for managing a Telegram Business account.
type Account struct {
	bot    Bot
	connID g.String
}

// NewAccount creates a new Account instance bound to the given bot and connection ID.
func NewAccount(bot Bot, connectionID g.String) *Account {
	return &Account{
		bot:    bot,
		connID: connectionID,
	}
}

// Name returns a builder for setting the account's first and last name.
func (a *Account) SetName(firstName g.String) *SetName {
	return &SetName{
		account:   a,
		firstName: firstName,
		opts:      new(gotgbot.SetBusinessAccountNameOpts),
	}
}

// SetUsername returns a builder for setting the account's username.
func (a *Account) SetUsername(username g.String) *SetUsername {
	return &SetUsername{
		account: a,
		opts:    &gotgbot.SetBusinessAccountUsernameOpts{Username: username.Std()},
	}
}

// SetBio returns a builder for setting the account's biography text.
func (a *Account) SetBio(bio g.String) *SetBio {
	return &SetBio{
		account: a,
		opts:    &gotgbot.SetBusinessAccountBioOpts{Bio: bio.Std()},
	}
}

// SetPhoto returns a builder for setting the account's profile photo.
func (a *Account) SetPhoto(photo g.String) *SetPhoto {
	return &SetPhoto{
		account: a,
		photo:   photo,
		opts:    new(gotgbot.SetBusinessAccountProfilePhotoOpts),
	}
}

// SetAnimatedPhoto returns a builder for setting the account's animated profile photo.
func (a *Account) SetAnimatedPhoto(animation g.String) *SetAnimatedPhoto {
	return &SetAnimatedPhoto{
		account:   a,
		animation: animation,
		opts:      new(gotgbot.SetBusinessAccountProfilePhotoOpts),
	}
}

// RemovePhoto returns a builder for removing the account's profile photo.
func (a *Account) RemovePhoto() *RemovePhoto {
	return &RemovePhoto{
		account: a,
		opts:    new(gotgbot.RemoveBusinessAccountProfilePhotoOpts),
	}
}

// GetConnection returns a builder for fetching the current business connection state.
func (a *Account) GetConnection() *GetConnection {
	return &GetConnection{
		account: a,
		opts:    new(gotgbot.GetBusinessConnectionOpts),
	}
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
