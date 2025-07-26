package business

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// Balance provides balance management
type Balance struct {
	bot    Bot
	connID String
}

// GetStarBalance creates a request to retrieve the current star balance.
func (b *Balance) GetStarBalance() *GetStarBalance {
	return &GetStarBalance{
		bot:    b.bot,
		connID: b.connID,
		opts:   new(gotgbot.GetBusinessAccountStarBalanceOpts),
	}
}

// TransferStars creates a request to transfer a specific amount of stars.
func (b *Balance) TransferStars(amount int64) *TransferStars {
	return &TransferStars{
		bot:    b.bot,
		connID: b.connID,
		amount: amount,
		opts:   new(gotgbot.TransferBusinessAccountStarsOpts),
	}
}

// GetGifts creates a request to retrieve owned gifts.
func (b *Balance) GetGifts() *GetGifts {
	return &GetGifts{
		bot:    b.bot,
		connID: b.connID,
		opts:   new(gotgbot.GetBusinessAccountGiftsOpts),
	}
}
