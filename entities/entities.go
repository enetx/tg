package entities

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// Entities provides a builder for creating MessageEntity slices from text.
type Entities struct {
	entities g.Slice[gotgbot.MessageEntity]
	text     g.String
}

// New creates a new Entities builder bound to the given source text.
func New[T ~string](text T) *Entities {
	return &Entities{
		text:     g.String(text),
		entities: g.NewSlice[gotgbot.MessageEntity](),
	}
}

// Bold marks the first occurrence of sub as bold.
func (e *Entities) Bold(sub g.String) *Entities {
	return e.match(sub, func(offset, length int64) {
		e.entities.Push(gotgbot.MessageEntity{Type: "bold", Offset: offset, Length: length})
	})
}

// Italic marks the first occurrence of sub as italic.
func (e *Entities) Italic(sub g.String) *Entities {
	return e.match(sub, func(offset, length int64) {
		e.entities.Push(gotgbot.MessageEntity{Type: "italic", Offset: offset, Length: length})
	})
}

// Underline marks the first occurrence of sub as underlined.
func (e *Entities) Underline(sub g.String) *Entities {
	return e.match(sub, func(offset, length int64) {
		e.entities.Push(gotgbot.MessageEntity{Type: "underline", Offset: offset, Length: length})
	})
}

// Strikethrough marks the first occurrence of sub as strikethrough.
func (e *Entities) Strikethrough(sub g.String) *Entities {
	return e.match(sub, func(offset, length int64) {
		e.entities.Push(gotgbot.MessageEntity{Type: "strikethrough", Offset: offset, Length: length})
	})
}

// Spoiler marks the first occurrence of sub as spoiler.
func (e *Entities) Spoiler(sub g.String) *Entities {
	return e.match(sub, func(offset, length int64) {
		e.entities.Push(gotgbot.MessageEntity{Type: "spoiler", Offset: offset, Length: length})
	})
}

// Code marks the first occurrence of sub as inline code.
func (e *Entities) Code(sub g.String) *Entities {
	return e.match(sub, func(offset, length int64) {
		e.entities.Push(gotgbot.MessageEntity{Type: "code", Offset: offset, Length: length})
	})
}

// Pre marks the first occurrence of sub as preformatted code.
func (e *Entities) Pre(sub g.String, language ...g.String) *Entities {
	return e.match(sub, func(offset, length int64) {
		entity := gotgbot.MessageEntity{Type: "pre", Offset: offset, Length: length}
		if len(language) > 0 {
			entity.Language = language[0].Std()
		}

		e.entities.Push(entity)
	})
}

// URL marks the first occurrence of sub as a hyperlink.
func (e *Entities) URL(sub, url g.String) *Entities {
	return e.match(sub, func(offset, length int64) {
		e.entities.Push(gotgbot.MessageEntity{
			Type:   "text_link",
			Offset: offset,
			Length: length,
			Url:    url.Std(),
		})
	})
}

// Mention marks the first occurrence of sub as a user mention.
func (e *Entities) Mention(sub g.String, userID int64) *Entities {
	return e.match(sub, func(offset, length int64) {
		e.entities.Push(gotgbot.MessageEntity{
			Type:   "text_mention",
			Offset: offset,
			Length: length,
			User:   &gotgbot.User{Id: userID},
		})
	})
}

// CustomEmoji marks the first occurrence of sub as a custom emoji.
func (e *Entities) CustomEmoji(sub, emojiID g.String) *Entities {
	return e.match(sub, func(offset, length int64) {
		e.entities.Push(gotgbot.MessageEntity{
			Type:          "custom_emoji",
			Offset:        offset,
			Length:        length,
			CustomEmojiId: emojiID.Std(),
		})
	})
}

// Blockquote marks the first occurrence of sub as a blockquote.
func (e *Entities) Blockquote(sub g.String) *Entities {
	return e.match(sub, func(offset, length int64) {
		e.entities.Push(gotgbot.MessageEntity{Type: "blockquote", Offset: offset, Length: length})
	})
}

// ExpandableBlockquote marks the first occurrence of sub as an expandable blockquote.
func (e *Entities) ExpandableBlockquote(sub g.String) *Entities {
	return e.match(sub, func(offset, length int64) {
		e.entities.Push(gotgbot.MessageEntity{Type: "expandable_blockquote", Offset: offset, Length: length})
	})
}

// match finds the first occurrence of sub in text and applies fn with offset and length.
func (e *Entities) match(sub g.String, fn func(offset, length int64)) *Entities {
	offset := e.text.Index(sub)
	if offset >= 0 {
		fn(int64(offset), int64(len(sub)))
	}

	return e
}

// Add inserts a raw MessageEntity.
func (e *Entities) Add(entity gotgbot.MessageEntity) *Entities {
	e.entities.Push(entity)
	return e
}

// Import appends existing MessageEntity slice.
func (e *Entities) Import(src []gotgbot.MessageEntity) *Entities {
	if src != nil {
		e.entities.Push(src...)
	}

	return e
}

// Clear removes all entities.
func (e *Entities) Clear() *Entities {
	e.entities = e.entities[:0]
	return e
}

// Count returns the number of entities.
func (e *Entities) Count() g.Int {
	return e.entities.Len()
}

// Std returns the underlying MessageEntity slice.
func (e *Entities) Std() []gotgbot.MessageEntity {
	return e.entities
}
