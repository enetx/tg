package suggested

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

// PostParameters provides a builder for creating SuggestedPostParameters.
type PostParameters struct {
	params *gotgbot.SuggestedPostParameters
}

// New creates a new PostParameters builder.
func New() *PostParameters {
	return &PostParameters{
		params: new(gotgbot.SuggestedPostParameters),
	}
}

// PriceStars sets the price in Telegram Stars (XTR).
// Amount must be between 5 and 100000.
func (p *PostParameters) PriceStars(amount int64) *PostParameters {
	p.params.Price = &gotgbot.SuggestedPostPrice{
		Currency: "XTR",
		Amount:   amount,
	}

	return p
}

// PriceTon sets the price in Toncoins (TON).
// Amount is in nanotoncoins, must be between 10000000 and 10000000000000.
func (p *PostParameters) PriceTon(nanotons int64) *PostParameters {
	p.params.Price = &gotgbot.SuggestedPostPrice{
		Currency: "TON",
		Amount:   nanotons,
	}

	return p
}

// SendDate sets the proposed send date of the post using a specific time.
// The date must be between 5 minutes and 30 days in the future.
func (p *PostParameters) SendDate(sendTime time.Time) *PostParameters {
	p.params.SendDate = sendTime.Unix()
	return p
}

// SendAfter sets the proposed send date relative to now.
// Duration must be between 5 minutes and 30 days.
func (p *PostParameters) SendAfter(duration time.Duration) *PostParameters {
	p.params.SendDate = time.Now().Add(duration).Unix()
	return p
}

// Build returns the SuggestedPostParameters for use with the Telegram API.
func (p *PostParameters) Build() *gotgbot.SuggestedPostParameters {
	if p.params.Price == nil && p.params.SendDate == 0 {
		return nil
	}

	return p.params
}

// Std is an alias for Build for consistency with other builders.
func (p *PostParameters) Std() *gotgbot.SuggestedPostParameters {
	return p.Build()
}
