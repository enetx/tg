// Package effects defines Telegram message effect identifiers.
// These can be used in MessageEffectId for special effects in private chats.
package effects

type EffectType int

const (
	Fire        EffectType = iota // 🔥 Fire / Sparkles animation
	ThumbsUp                      // 👍 Thumbs up animation
	ThumbsDown                    // 👎 Thumbs down animation
	Heart                         // ❤️ Heart/Confetti animation
	Celebration                   // 🎉 Balloons/Confetti
	Poop                          // 💩 Funny scream face animation
)

// String returns the Telegram MessageEffectId string associated with the effect.
func (e EffectType) String() string {
	switch e {
	case Fire:
		return "5104841245755180586"
	case ThumbsUp:
		return "5107584321108051014"
	case ThumbsDown:
		return "5104858069142078462"
	case Heart:
		return "️5044134455711629726"
	case Celebration:
		return "5046509860389126442"
	case Poop:
		return "5046589136895476101"
	default:
		return "unknown"
	}
}
