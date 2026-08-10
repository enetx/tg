package richtext_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/richtext"
)

func TestNew(t *testing.T) {
	text := richtext.New("Hello")
	if text == nil {
		t.Fatal("Expected New to create a Text builder")
	}

	node, ok := text.Std().(gotgbot.RichTextString)
	if !ok {
		t.Fatal("Expected node to be RichTextString")
	}

	if string(node) != "Hello" {
		t.Error("Expected New to keep the plain text")
	}
}

func TestNew_GString(t *testing.T) {
	if string(richtext.New(g.String("Hello")).Std().(gotgbot.RichTextString)) != "Hello" {
		t.Error("Expected New to accept g.String")
	}
}

func TestJoin(t *testing.T) {
	text := richtext.Join(richtext.New("a"), nil, richtext.New("b"))

	array, ok := text.Std().(gotgbot.RichTextArray)
	if !ok {
		t.Fatal("Expected node to be RichTextArray")
	}

	if len(array) != 2 {
		t.Error("Expected Join to skip nil nodes")
	}
}

func TestJoin_Empty(t *testing.T) {
	array, ok := richtext.Join().Std().(gotgbot.RichTextArray)
	if !ok {
		t.Fatal("Expected node to be RichTextArray")
	}

	if len(array) != 0 {
		t.Error("Expected Join without arguments to produce an empty array")
	}
}

func TestEmoji(t *testing.T) {
	node, ok := richtext.Emoji(g.String("5368324170671202286"), g.String("👍")).Std().(gotgbot.RichTextCustomEmoji)
	if !ok {
		t.Fatal("Expected node to be RichTextCustomEmoji")
	}

	if node.CustomEmojiId != "5368324170671202286" || node.AlternativeText != "👍" {
		t.Error("Expected Emoji fields to be set correctly")
	}
}

func TestMath(t *testing.T) {
	node, ok := richtext.Math(g.String("E = mc^2")).Std().(gotgbot.RichTextMathematicalExpression)
	if !ok {
		t.Fatal("Expected node to be RichTextMathematicalExpression")
	}

	if node.Expression != "E = mc^2" {
		t.Error("Expected Expression to be set correctly")
	}
}

func TestAnchor(t *testing.T) {
	node, ok := richtext.Anchor(g.String("section-1")).Std().(gotgbot.RichTextAnchor)
	if !ok {
		t.Fatal("Expected node to be RichTextAnchor")
	}

	if node.Name != "section-1" {
		t.Error("Expected Name to be set correctly")
	}
}

func TestText_Styles(t *testing.T) {
	tests := []struct {
		name  string
		apply func(*richtext.Text) *richtext.Text
		check func(gotgbot.RichText) bool
	}{
		{"Bold", (*richtext.Text).Bold, func(n gotgbot.RichText) bool { _, ok := n.(gotgbot.RichTextBold); return ok }},
		{"Italic", (*richtext.Text).Italic, func(n gotgbot.RichText) bool { _, ok := n.(gotgbot.RichTextItalic); return ok }},
		{
			"Underline",
			(*richtext.Text).Underline,
			func(n gotgbot.RichText) bool { _, ok := n.(gotgbot.RichTextUnderline); return ok },
		},
		{
			"Strikethrough",
			(*richtext.Text).Strikethrough,
			func(n gotgbot.RichText) bool { _, ok := n.(gotgbot.RichTextStrikethrough); return ok },
		},
		{
			"Spoiler",
			(*richtext.Text).Spoiler,
			func(n gotgbot.RichText) bool { _, ok := n.(gotgbot.RichTextSpoiler); return ok },
		},
		{
			"Subscript",
			(*richtext.Text).Subscript,
			func(n gotgbot.RichText) bool { _, ok := n.(gotgbot.RichTextSubscript); return ok },
		},
		{
			"Superscript",
			(*richtext.Text).Superscript,
			func(n gotgbot.RichText) bool { _, ok := n.(gotgbot.RichTextSuperscript); return ok },
		},
		{"Marked", (*richtext.Text).Marked, func(n gotgbot.RichText) bool { _, ok := n.(gotgbot.RichTextMarked); return ok }},
		{"Code", (*richtext.Text).Code, func(n gotgbot.RichText) bool { _, ok := n.(gotgbot.RichTextCode); return ok }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			text := richtext.New("Hello")
			if tt.apply(text) != text {
				t.Errorf("Expected %s to return the same Text instance", tt.name)
			}

			if !tt.check(text.Std()) {
				t.Errorf("Expected %s to wrap the node", tt.name)
			}
		})
	}
}

func TestText_URL(t *testing.T) {
	node, ok := richtext.New("site").URL(g.String("https://example.com")).Std().(gotgbot.RichTextUrl)
	if !ok {
		t.Fatal("Expected node to be RichTextUrl")
	}

	if node.Url != "https://example.com" {
		t.Error("Expected Url to be set correctly")
	}
}

func TestText_Email(t *testing.T) {
	node, ok := richtext.New("mail").Email(g.String("a@b.com")).Std().(gotgbot.RichTextEmailAddress)
	if !ok {
		t.Fatal("Expected node to be RichTextEmailAddress")
	}

	if node.EmailAddress != "a@b.com" {
		t.Error("Expected EmailAddress to be set correctly")
	}
}

func TestText_Phone(t *testing.T) {
	node, ok := richtext.New("call").Phone(g.String("+1234567890")).Std().(gotgbot.RichTextPhoneNumber)
	if !ok {
		t.Fatal("Expected node to be RichTextPhoneNumber")
	}

	if node.PhoneNumber != "+1234567890" {
		t.Error("Expected PhoneNumber to be set correctly")
	}
}

func TestText_BankCard(t *testing.T) {
	node, ok := richtext.New("card").BankCard(g.String("4111111111111111")).Std().(gotgbot.RichTextBankCardNumber)
	if !ok {
		t.Fatal("Expected node to be RichTextBankCardNumber")
	}

	if node.BankCardNumber != "4111111111111111" {
		t.Error("Expected BankCardNumber to be set correctly")
	}
}

func TestText_Mention(t *testing.T) {
	node, ok := richtext.New("me").Mention(g.String("enetx")).Std().(gotgbot.RichTextMention)
	if !ok {
		t.Fatal("Expected node to be RichTextMention")
	}

	if node.Username != "enetx" {
		t.Error("Expected Username to be set correctly")
	}
}

func TestText_TextMention(t *testing.T) {
	node, ok := richtext.New("me").TextMention(gotgbot.User{Id: 42}).Std().(gotgbot.RichTextTextMention)
	if !ok {
		t.Fatal("Expected node to be RichTextTextMention")
	}

	if node.User.Id != 42 {
		t.Error("Expected User to be set correctly")
	}
}

func TestText_Hashtag(t *testing.T) {
	node, ok := richtext.New("tag").Hashtag(g.String("#go")).Std().(gotgbot.RichTextHashtag)
	if !ok {
		t.Fatal("Expected node to be RichTextHashtag")
	}

	if node.Hashtag != "#go" {
		t.Error("Expected Hashtag to be set correctly")
	}
}

func TestText_Cashtag(t *testing.T) {
	node, ok := richtext.New("cash").Cashtag(g.String("$USD")).Std().(gotgbot.RichTextCashtag)
	if !ok {
		t.Fatal("Expected node to be RichTextCashtag")
	}

	if node.Cashtag != "$USD" {
		t.Error("Expected Cashtag to be set correctly")
	}
}

func TestText_BotCommand(t *testing.T) {
	node, ok := richtext.New("cmd").BotCommand(g.String("/start")).Std().(gotgbot.RichTextBotCommand)
	if !ok {
		t.Fatal("Expected node to be RichTextBotCommand")
	}

	if node.BotCommand != "/start" {
		t.Error("Expected BotCommand to be set correctly")
	}
}

func TestText_DateTime(t *testing.T) {
	when := time.Unix(1700000000, 0)

	node, ok := richtext.New("when").DateTime(when).Std().(gotgbot.RichTextDateTime)
	if !ok {
		t.Fatal("Expected node to be RichTextDateTime")
	}

	if node.UnixTime != when.Unix() {
		t.Error("Expected UnixTime to be set correctly")
	}

	if node.DateTimeFormat != "" {
		t.Error("Expected DateTimeFormat to stay empty without an explicit format")
	}
}

func TestText_DateTimeWithFormat(t *testing.T) {
	node := richtext.New("when").
		DateTime(time.Unix(1700000000, 0), g.String("dd.MM.yyyy")).
		Std().(gotgbot.RichTextDateTime)

	if node.DateTimeFormat != "dd.MM.yyyy" {
		t.Error("Expected DateTimeFormat to be set correctly")
	}
}

func TestText_AnchorLink(t *testing.T) {
	node, ok := richtext.New("jump").AnchorLink(g.String("section-1")).Std().(gotgbot.RichTextAnchorLink)
	if !ok {
		t.Fatal("Expected node to be RichTextAnchorLink")
	}

	if node.AnchorName != "section-1" {
		t.Error("Expected AnchorName to be set correctly")
	}
}

func TestText_Reference(t *testing.T) {
	node, ok := richtext.New("ref").Reference(g.String("rfc")).Std().(gotgbot.RichTextReference)
	if !ok {
		t.Fatal("Expected node to be RichTextReference")
	}

	if node.Name != "rfc" {
		t.Error("Expected Name to be set correctly")
	}
}

func TestText_ReferenceLink(t *testing.T) {
	node, ok := richtext.New("ref").ReferenceLink(g.String("rfc")).Std().(gotgbot.RichTextReferenceLink)
	if !ok {
		t.Fatal("Expected node to be RichTextReferenceLink")
	}

	if node.ReferenceName != "rfc" {
		t.Error("Expected ReferenceName to be set correctly")
	}
}

func TestText_Nesting(t *testing.T) {
	node, ok := richtext.New("Hello").Bold().Italic().Std().(gotgbot.RichTextItalic)
	if !ok {
		t.Fatal("Expected the outermost node to be RichTextItalic")
	}

	if _, ok := node.Text.(gotgbot.RichTextBold); !ok {
		t.Error("Expected the inner node to be RichTextBold")
	}
}

func TestText_StdNil(t *testing.T) {
	var text *richtext.Text
	if text.Std() != nil {
		t.Error("Expected Std on a nil Text to return nil")
	}
}
