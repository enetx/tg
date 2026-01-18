package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetChatGifts represents a request to get gifts received by a chat.
type GetChatGifts struct {
	ctx    *Context
	opts   *gotgbot.GetChatGiftsOpts
	chatID g.Option[int64]
}

// ChatID sets the target chat ID for this request.
func (gcg *GetChatGifts) ChatID(id int64) *GetChatGifts {
	gcg.chatID = g.Some(id)
	return gcg
}

// ExcludeUnsaved excludes gifts that aren't saved to the chat's profile page.
func (gcg *GetChatGifts) ExcludeUnsaved() *GetChatGifts {
	gcg.opts.ExcludeUnsaved = true
	return gcg
}

// ExcludeSaved excludes gifts that are saved to the chat's profile page.
func (gcg *GetChatGifts) ExcludeSaved() *GetChatGifts {
	gcg.opts.ExcludeSaved = true
	return gcg
}

// ExcludeUnlimited excludes gifts that can be purchased an unlimited number of times.
func (gcg *GetChatGifts) ExcludeUnlimited() *GetChatGifts {
	gcg.opts.ExcludeUnlimited = true
	return gcg
}

// ExcludeLimitedUpgradable excludes limited gifts that can be upgraded to unique.
func (gcg *GetChatGifts) ExcludeLimitedUpgradable() *GetChatGifts {
	gcg.opts.ExcludeLimitedUpgradable = true
	return gcg
}

// ExcludeLimitedNonUpgradable excludes limited gifts that can't be upgraded to unique.
func (gcg *GetChatGifts) ExcludeLimitedNonUpgradable() *GetChatGifts {
	gcg.opts.ExcludeLimitedNonUpgradable = true
	return gcg
}

// ExcludeFromBlockchain excludes gifts assigned from the TON blockchain.
func (gcg *GetChatGifts) ExcludeFromBlockchain() *GetChatGifts {
	gcg.opts.ExcludeFromBlockchain = true
	return gcg
}

// ExcludeUnique excludes unique gifts.
func (gcg *GetChatGifts) ExcludeUnique() *GetChatGifts {
	gcg.opts.ExcludeUnique = true
	return gcg
}

// SortByPrice sorts results by gift price instead of send date.
func (gcg *GetChatGifts) SortByPrice() *GetChatGifts {
	gcg.opts.SortByPrice = true
	return gcg
}

// Offset sets the pagination offset.
func (gcg *GetChatGifts) Offset(offset g.String) *GetChatGifts {
	gcg.opts.Offset = offset.Std()
	return gcg
}

// Limit sets the maximum number of gifts to return (1–100).
func (gcg *GetChatGifts) Limit(limit int64) *GetChatGifts {
	gcg.opts.Limit = limit
	return gcg
}

// Timeout sets a custom timeout for this request.
func (gcg *GetChatGifts) Timeout(duration time.Duration) *GetChatGifts {
	if gcg.opts.RequestOpts == nil {
		gcg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gcg.opts.RequestOpts.Timeout = duration

	return gcg
}

// APIURL sets a custom API URL for this request.
func (gcg *GetChatGifts) APIURL(url g.String) *GetChatGifts {
	if gcg.opts.RequestOpts == nil {
		gcg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gcg.opts.RequestOpts.APIURL = url.Std()

	return gcg
}

// Send executes the GetChatGifts request and returns the chat's gifts.
func (gcg *GetChatGifts) Send() g.Result[*gotgbot.OwnedGifts] {
	chatID := gcg.chatID.UnwrapOr(gcg.ctx.EffectiveChat.Id)
	return g.ResultOf(gcg.ctx.Bot.Raw().GetChatGifts(chatID, gcg.opts))
}
