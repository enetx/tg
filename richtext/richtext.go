// Package richtext provides a builder for Telegram rich text nodes.
//
// A rich text node is a tree: leaf nodes carry the actual characters, and every
// formatting method wraps the current node into a new one. Nodes are combined
// with Join and consumed by the rich package via Std.
//
// See: https://core.telegram.org/bots/api#richtext
package richtext

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// Text provides a builder for creating rich text nodes.
type Text struct {
	node gotgbot.RichText
}

// New creates a new Text builder holding a plain string node.
func New[T ~string](text T) *Text {
	return &Text{node: gotgbot.RichTextString(text)}
}

// Join creates a new Text builder holding the given nodes in order.
// Nil nodes are skipped.
func Join(texts ...*Text) *Text {
	nodes := g.Slice[*Text](texts).
		Iter().
		Exclude(func(text *Text) bool { return text == nil }).
		Collect().
		Slice()

	return &Text{node: gotgbot.RichTextArray(g.TransformSlice(nodes, (*Text).Std))}
}

// Emoji creates a new Text builder holding a custom emoji node.
func Emoji(customEmojiID, alternative g.String) *Text {
	return &Text{node: gotgbot.RichTextCustomEmoji{
		CustomEmojiId:   customEmojiID.Std(),
		AlternativeText: alternative.Std(),
	}}
}

// Math creates a new Text builder holding a mathematical expression node.
func Math(expression g.String) *Text {
	return &Text{node: gotgbot.RichTextMathematicalExpression{Expression: expression.Std()}}
}

// Anchor creates a new Text builder holding an anchor node that can be
// referenced by AnchorLink.
func Anchor(name g.String) *Text {
	return &Text{node: gotgbot.RichTextAnchor{Name: name.Std()}}
}

// Bold wraps the current node in bold formatting.
func (t *Text) Bold() *Text {
	t.node = gotgbot.RichTextBold{Text: t.node}
	return t
}

// Italic wraps the current node in italic formatting.
func (t *Text) Italic() *Text {
	t.node = gotgbot.RichTextItalic{Text: t.node}
	return t
}

// Underline wraps the current node in underline formatting.
func (t *Text) Underline() *Text {
	t.node = gotgbot.RichTextUnderline{Text: t.node}
	return t
}

// Strikethrough wraps the current node in strikethrough formatting.
func (t *Text) Strikethrough() *Text {
	t.node = gotgbot.RichTextStrikethrough{Text: t.node}
	return t
}

// Spoiler wraps the current node in spoiler formatting.
func (t *Text) Spoiler() *Text {
	t.node = gotgbot.RichTextSpoiler{Text: t.node}
	return t
}

// Subscript wraps the current node in subscript formatting.
func (t *Text) Subscript() *Text {
	t.node = gotgbot.RichTextSubscript{Text: t.node}
	return t
}

// Superscript wraps the current node in superscript formatting.
func (t *Text) Superscript() *Text {
	t.node = gotgbot.RichTextSuperscript{Text: t.node}
	return t
}

// Marked wraps the current node in marked (highlighted) formatting.
func (t *Text) Marked() *Text {
	t.node = gotgbot.RichTextMarked{Text: t.node}
	return t
}

// Code wraps the current node in monospace formatting.
func (t *Text) Code() *Text {
	t.node = gotgbot.RichTextCode{Text: t.node}
	return t
}

// URL turns the current node into a hyperlink pointing to url.
func (t *Text) URL(url g.String) *Text {
	t.node = gotgbot.RichTextUrl{Text: t.node, Url: url.Std()}
	return t
}

// Email turns the current node into an email address link.
func (t *Text) Email(address g.String) *Text {
	t.node = gotgbot.RichTextEmailAddress{Text: t.node, EmailAddress: address.Std()}
	return t
}

// Phone turns the current node into a phone number link.
func (t *Text) Phone(number g.String) *Text {
	t.node = gotgbot.RichTextPhoneNumber{Text: t.node, PhoneNumber: number.Std()}
	return t
}

// BankCard turns the current node into a bank card number.
func (t *Text) BankCard(number g.String) *Text {
	t.node = gotgbot.RichTextBankCardNumber{Text: t.node, BankCardNumber: number.Std()}
	return t
}

// Mention turns the current node into a mention of the given username.
func (t *Text) Mention(username g.String) *Text {
	t.node = gotgbot.RichTextMention{Text: t.node, Username: username.Std()}
	return t
}

// TextMention turns the current node into a mention of a user without a username.
func (t *Text) TextMention(user gotgbot.User) *Text {
	t.node = gotgbot.RichTextTextMention{Text: t.node, User: user}
	return t
}

// Hashtag turns the current node into a hashtag.
func (t *Text) Hashtag(hashtag g.String) *Text {
	t.node = gotgbot.RichTextHashtag{Text: t.node, Hashtag: hashtag.Std()}
	return t
}

// Cashtag turns the current node into a cashtag.
func (t *Text) Cashtag(cashtag g.String) *Text {
	t.node = gotgbot.RichTextCashtag{Text: t.node, Cashtag: cashtag.Std()}
	return t
}

// BotCommand turns the current node into a bot command.
func (t *Text) BotCommand(command g.String) *Text {
	t.node = gotgbot.RichTextBotCommand{Text: t.node, BotCommand: command.Std()}
	return t
}

// DateTime turns the current node into a date and time value.
// An optional layout describes how clients must render the value.
func (t *Text) DateTime(when time.Time, format ...g.String) *Text {
	datetime := gotgbot.RichTextDateTime{Text: t.node, UnixTime: when.Unix()}
	if len(format) > 0 {
		datetime.DateTimeFormat = format[0].Std()
	}

	t.node = datetime

	return t
}

// AnchorLink turns the current node into a link to the anchor with the given name.
func (t *Text) AnchorLink(anchorName g.String) *Text {
	t.node = gotgbot.RichTextAnchorLink{Text: t.node, AnchorName: anchorName.Std()}
	return t
}

// Reference turns the current node into a reference definition with the given name.
func (t *Text) Reference(name g.String) *Text {
	t.node = gotgbot.RichTextReference{Text: t.node, Name: name.Std()}
	return t
}

// ReferenceLink turns the current node into a link to the reference with the given name.
func (t *Text) ReferenceLink(referenceName g.String) *Text {
	t.node = gotgbot.RichTextReferenceLink{Text: t.node, ReferenceName: referenceName.Std()}
	return t
}

// Std returns the underlying gotgbot.RichText node.
func (t *Text) Std() gotgbot.RichText {
	if t == nil {
		return nil
	}

	return t.node
}
