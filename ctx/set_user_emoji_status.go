package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetUserEmojiStatus represents a request to set user emoji status.
type SetUserEmojiStatus struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.SetUserEmojiStatusOpts
}

// EmojiStatusCustomEmojiID sets the custom emoji identifier for the status.
func (sues *SetUserEmojiStatus) EmojiStatusCustomEmojiID(emojiID g.String) *SetUserEmojiStatus {
	sues.opts.EmojiStatusCustomEmojiId = emojiID.Std()
	return sues
}

// RemoveStatus removes the emoji status by setting empty emoji ID.
func (sues *SetUserEmojiStatus) RemoveStatus() *SetUserEmojiStatus {
	sues.opts.EmojiStatusCustomEmojiId = ""
	return sues
}

// EmojiStatusExpirationDate sets the expiration date for the emoji status.
func (sues *SetUserEmojiStatus) EmojiStatusExpirationDate(expirationDate int64) *SetUserEmojiStatus {
	sues.opts.EmojiStatusExpirationDate = expirationDate
	return sues
}

// ExpiresAt sets the expiration time for the emoji status.
func (sues *SetUserEmojiStatus) ExpiresAt(t time.Time) *SetUserEmojiStatus {
	sues.opts.EmojiStatusExpirationDate = t.Unix()
	return sues
}

// ExpiresIn sets the emoji status to expire after the given duration.
func (sues *SetUserEmojiStatus) ExpiresIn(duration time.Duration) *SetUserEmojiStatus {
	sues.opts.EmojiStatusExpirationDate = time.Now().Add(duration).Unix()
	return sues
}

// Timeout sets a custom timeout for this request.
func (sues *SetUserEmojiStatus) Timeout(duration time.Duration) *SetUserEmojiStatus {
	if sues.opts.RequestOpts == nil {
		sues.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sues.opts.RequestOpts.Timeout = duration

	return sues
}

// APIURL sets a custom API URL for this request.
func (sues *SetUserEmojiStatus) APIURL(url g.String) *SetUserEmojiStatus {
	if sues.opts.RequestOpts == nil {
		sues.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sues.opts.RequestOpts.APIURL = url.Std()

	return sues
}

// Send sets the user emoji status.
func (sues *SetUserEmojiStatus) Send() g.Result[bool] {
	return g.ResultOf(sues.ctx.Bot.Raw().SetUserEmojiStatus(sues.userID, sues.opts))
}
