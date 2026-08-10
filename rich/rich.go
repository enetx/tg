// Package rich provides builders for Telegram rich formatted messages.
//
// A rich message is an ordered list of blocks. Blocks carry rich text nodes
// built with the richtext package, media built with the input package, or other
// nested blocks.
//
// See: https://core.telegram.org/bots/api#inputrichmessage
package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/input"
)

// Block represents an interface for all rich message block builders.
type Block interface {
	Build() gotgbot.InputRichBlock
}

// Compile-time interface checks
var (
	_ Block = (*BlockParagraph)(nil)
	_ Block = (*BlockHeading)(nil)
	_ Block = (*BlockPre)(nil)
	_ Block = (*BlockFooter)(nil)
	_ Block = (*BlockDivider)(nil)
	_ Block = (*BlockMath)(nil)
	_ Block = (*BlockAnchor)(nil)
	_ Block = (*BlockList)(nil)
	_ Block = (*BlockQuotation)(nil)
	_ Block = (*BlockPullQuotation)(nil)
	_ Block = (*BlockCollage)(nil)
	_ Block = (*BlockSlideshow)(nil)
	_ Block = (*BlockTable)(nil)
	_ Block = (*BlockDetails)(nil)
	_ Block = (*BlockMap)(nil)
	_ Block = (*BlockAnimation)(nil)
	_ Block = (*BlockAudio)(nil)
	_ Block = (*BlockPhoto)(nil)
	_ Block = (*BlockVideo)(nil)
	_ Block = (*BlockVoiceNote)(nil)
	_ Block = (*BlockThinking)(nil)

	// MessageContent checks
	_ input.MessageContent = (*Content)(nil)
)
