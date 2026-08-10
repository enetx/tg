package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
)

// Message provides a builder for creating rich formatted messages.
type Message struct {
	message *gotgbot.InputRichMessage
	blocks  g.Slice[Block]
	media   g.Slice[gotgbot.InputRichMessageMedia]
}

// New creates a new empty rich Message builder.
func New() *Message {
	return &Message{
		message: new(gotgbot.InputRichMessage),
		blocks:  g.NewSlice[Block](),
		media:   g.NewSlice[gotgbot.InputRichMessageMedia](),
	}
}

// Block appends blocks to the message content.
func (m *Message) Block(blocks ...Block) *Message {
	g.Slice[Block](blocks).
		Iter().
		Exclude(func(block Block) bool { return block == nil }).
		ForEach(func(block Block) { m.blocks.Push(block) })

	return m
}

// HTML sets the message content described using HTML formatting.
// Media referenced from the markup is added with Media.
func (m *Message) HTML(html g.String) *Message {
	m.message.Html = html.Std()
	return m
}

// Markdown sets the message content described using Markdown formatting.
// Media referenced from the markup is added with Media.
func (m *Message) Markdown(markdown g.String) *Message {
	m.message.Markdown = markdown.Std()
	return m
}

// Media registers media referenced from the HTML or Markdown content by the
// given identifier, as used in tg://photo?id=, tg://video?id= and tg://audio?id= links.
func (m *Message) Media(id g.String, media input.Media) *Message {
	if media != nil {
		m.media.Push(gotgbot.InputRichMessageMedia{Id: id.Std(), Media: media.Build()})
	}

	return m
}

// RTL marks the message as right-to-left.
func (m *Message) RTL() *Message {
	m.message.IsRtl = true
	return m
}

// SkipEntityDetection disables automatic detection of entities in the text.
func (m *Message) SkipEntityDetection() *Message {
	m.message.SkipEntityDetection = true
	return m
}

// Std returns the underlying gotgbot.InputRichMessage.
func (m *Message) Std() gotgbot.InputRichMessage {
	if m == nil {
		return gotgbot.InputRichMessage{}
	}

	message := *m.message

	if !m.blocks.IsEmpty() {
		message.Blocks = g.TransformSlice(m.blocks, Block.Build)
	}

	if !m.media.IsEmpty() {
		message.Media = m.media
	}

	return message
}
