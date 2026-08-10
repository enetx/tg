package main

import (
	"github.com/enetx/g"
	"github.com/enetx/g/fs"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/rich"
	"github.com/enetx/tg/richtext"
)

// Rich messages (Bot API 10.2+): article-like content built from a list of blocks.
// Blocks come from the rich package, inline formatting from the richtext package.
func main() {
	token := g.Slice[g.String](fs.NewFile("../.env").Read().Ok().Trim().Split("=")).Last().Some()
	b := bot.New(token).Build().Unwrap()

	// /article — send a rich message assembled from blocks.
	b.Command("article", func(ctx *ctx.Context) error {
		photo := input.Photo(file.Input(g.String("https://picsum.photos/800/400")).Ok())

		message := rich.New().
			Block(rich.Heading(richtext.New("Rich messages")).Size(1)).
			Block(rich.Paragraph(richtext.Join(
				richtext.New("Built from "),
				richtext.New("blocks").Bold(),
				richtext.New(" and "),
				richtext.New("inline nodes").Italic(),
				richtext.New("."),
			))).
			Block(rich.Divider()).
			Block(rich.List(
				rich.ListItem(rich.Paragraph(richtext.New("Checked item"))).Checked(),
				rich.ListItem(rich.Paragraph(richtext.New("Pending item"))).Checkbox(),
			)).
			Block(rich.Pre(richtext.New("fmt.Println(\"hello\")")).Language(g.String("go"))).
			Block(rich.Math(g.String("E = mc^2"))).
			Block(rich.Table().
				Bordered().
				Striped().
				Row(rich.Cell(richtext.New("Block")).Header(), rich.Cell(richtext.New("Purpose")).Header()).
				Row(rich.Cell(richtext.New("heading")), rich.Cell(richtext.New("section title"))).
				Caption(richtext.New("Supported blocks"))).
			Block(rich.Photo(photo).
				Caption(rich.NewCaption(richtext.New("A random photo")).Credit(richtext.New("picsum.photos")))).
			Block(rich.Quotation(rich.Paragraph(richtext.New("Blocks nest freely."))).
				Credit(richtext.New("the docs"))).
			Block(rich.Details(richtext.New("Show more")).
				Block(rich.Paragraph(richtext.New("Collapsed by default.")))).
			Block(rich.Footer(richtext.New("Sent with enetx/tg")))

		return ctx.SendRichMessage(message).Send().Err()
	})

	// /html — the same message described with HTML instead of blocks.
	b.Command("html", func(ctx *ctx.Context) error {
		message := rich.New().
			HTML(g.String(`<h1>Rich messages</h1><p>Described with <b>HTML</b>.</p>`)).
			SkipEntityDetection()

		return ctx.SendRichMessage(message).Send().Err()
	})

	// /draft — stream a rich draft before persisting the final rich message.
	b.Command("draft", func(ctx *ctx.Context) error {
		draftID := int64(1)

		partial := rich.New().Block(rich.Paragraph(richtext.New("Writing the article")))
		if err := ctx.SendRichMessageDraft(draftID, partial).Send().Err(); err != nil {
			return err
		}

		final := rich.New().Block(rich.Paragraph(richtext.New("Article is ready!").Bold()))

		return ctx.SendRichMessage(final).Send().Err()
	})

	b.Polling().Start()
}
