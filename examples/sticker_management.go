package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Create new sticker set
	b.Command("createstickerSet", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 5 {
			return ctx.Reply("Usage: /createstickerSet <user_id> <name> <title> <sticker_file> <emoji>").Send().Err()
		}

		userID := args[0].ToInt().Unwrap().Int64()
		name := args[1]
		title := args[2]
		stickerFile := args[3]
		emoji := args[4]

		result := ctx.CreateNewStickerSet(userID, name, title).
			StickerType("regular").
			Sticker(stickerFile, "static", []g.String{emoji}).Add().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Sticker set created successfully!").Send().Err()
	})

	// Add sticker to existing set
	b.Command("addsticker", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 4 {
			return ctx.Reply("Usage: /addsticker <user_id> <set_name> <sticker_file> <emoji>").Send().Err()
		}

		userID := args[0].ToInt().Unwrap().Int64()
		setName := args[1]
		stickerFile := args[2]
		emoji := args[3]

		result := ctx.AddStickerToSet(userID, setName).
			File(stickerFile).
			Format("static").
			EmojiList([]g.String{emoji}).
			Keywords([]g.String{"custom", "sticker"}).
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Sticker added to set successfully!").Send().Err()
	})

	// Get sticker set info
	b.Command("getstickerSet", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /getstickerSet <set_name>").Send().Err()
		}

		setName := args[0]

		result := ctx.GetStickerSet(setName).Send()
		if result.IsErr() {
			return result.Err()
		}

		stickerSet := result.Ok()
		info := "Sticker Set Info:\n" +
			"Name: " + stickerSet.Name + "\n" +
			"Title: " + stickerSet.Title + "\n" +
			"Count: " + g.String(len(stickerSet.Stickers)).Std()

		return ctx.Reply(g.String(info)).Send().Err()
	})

	// Delete sticker from set
	b.Command("deletesticker", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /deletesticker <sticker_file_id>").Send().Err()
		}

		stickerID := args[0]

		result := ctx.DeleteStickerFromSet(stickerID).Send()
		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Sticker deleted from set successfully!").Send().Err()
	})

	// Set sticker position
	b.Command("setstickerposition", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /setstickerposition <sticker_file_id> <position>").Send().Err()
		}

		stickerID := args[0]
		position := args[1].ToInt().Unwrap().Int64()

		result := ctx.SetStickerPositionInSet(stickerID, position).Send()
		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Sticker position updated successfully!").Send().Err()
	})

	// Set sticker emoji list
	b.Command("setstickeremoji", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /setstickeremoji <sticker_file_id> <emoji1> [emoji2] ...").Send().Err()
		}

		stickerID := args[0]
		var emojis g.Slice[g.String]

		args.Iter().Skip(1).ForEach(func(emoji g.String) {
			emojis.Push(emoji)
		})

		result := ctx.SetStickerEmojiList(stickerID).
			EmojiList(emojis).
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Sticker emoji list updated successfully!").Send().Err()
	})

	// Set sticker keywords
	b.Command("setstickerkeywords", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /setstickerkeywords <sticker_file_id> <keyword1> [keyword2] ...").Send().Err()
		}

		stickerID := args[0]
		var keywords g.Slice[g.String]

		args.Iter().Skip(1).ForEach(func(keyword g.String) {
			keywords.Push(keyword)
		})

		result := ctx.SetStickerKeywords(stickerID).
			Keywords(keywords).
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Sticker keywords updated successfully!").Send().Err()
	})

	// Set sticker mask position
	b.Command("setstickermask", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 5 {
			return ctx.Reply("Usage: /setstickermask <sticker_file_id> <point> <x_shift> <y_shift> <scale>").
				Send().
				Err()
		}

		stickerID := args[0]
		point := args[1]
		xShift := args[2].ToFloat().Unwrap()
		yShift := args[3].ToFloat().Unwrap()
		scale := args[4].ToFloat().Unwrap()

		result := ctx.SetStickerMaskPosition(stickerID).
			MaskPosition(point, xShift.Std(), yShift.Std(), scale.Std()).
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Sticker mask position updated successfully!").Send().Err()
	})

	// Set sticker set thumbnail
	b.Command("setsetthumb", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 3 {
			return ctx.Reply("Usage: /setsetthumb <set_name> <user_id> <thumbnail_file>").Send().Err()
		}

		setName := args[0]
		userID := args[1].ToInt().Unwrap().Int64()
		thumbFile := args[2]

		result := ctx.SetStickerSetThumbnail(setName, userID).
			Thumbnail(thumbFile).
			Format("static").
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Sticker set thumbnail updated successfully!").Send().Err()
	})

	// Upload sticker file
	b.Command("uploadstickerfile", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 3 {
			return ctx.Reply("Usage: /uploadstickerfile <user_id> <sticker_file> <format>").Send().Err()
		}

		userID := args[0].ToInt().Unwrap().Int64()
		stickerFile := args[1]
		format := args[2]

		result := ctx.UploadStickerFile(userID, format).
			File(stickerFile).
			Send()

		if result.IsErr() {
			return result.Err()
		}

		uploadedFile := result.Ok()

		return ctx.Reply("File uploaded successfully! File ID: " + g.String(uploadedFile.FileId)).Send().Err()
	})

	// Delete entire sticker set
	b.Command("deletestickerSet", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /deletestickerSet <set_name>").Send().Err()
		}

		setName := args[0]

		result := ctx.DeleteStickerSet(setName).Send()
		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Sticker set deleted successfully!").Send().Err()
	})

	// Get custom emoji stickers
	b.Command("getcustomemoji", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /getcustomemoji <emoji_id1> [emoji_id2] ...").Send().Err()
		}

		result := ctx.GetCustomEmojiStickers(args).Send()
		if result.IsErr() {
			return result.Err()
		}

		info := "Found " + result.Ok().Len().String() + " custom emoji stickers"

		return ctx.Reply(info).Send().Err()
	})

	// Advanced: Create animated sticker set
	b.Command("createaniset", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 5 {
			return ctx.Reply("Usage: /createaniset <user_id> <name> <title> <tgs_file> <emoji>").Send().Err()
		}

		userID := args[0].ToInt().Unwrap().Int64()
		name := args[1]
		title := args[2]
		tgsFile := args[3]
		emoji := args[4]

		result := ctx.CreateNewStickerSet(userID, name, title).
			StickerType("regular").
			Sticker(tgsFile, "animated", []g.String{emoji}).
			Keywords([]g.String{"animated", "custom"}).
			Add().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Animated sticker set created successfully!").Send().Err()
	})

	// Advanced: Create video sticker set
	b.Command("createvideoset", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 5 {
			return ctx.Reply("Usage: /createvideoset <user_id> <name> <title> <webm_file> <emoji>").Send().Err()
		}

		userID := args[0].ToInt().Unwrap().Int64()
		name := args[1]
		title := args[2]
		webmFile := args[3]
		emoji := args[4]

		result := ctx.CreateNewStickerSet(userID, name, title).
			StickerType("regular").
			Sticker(webmFile, "video", []g.String{emoji}).
			Keywords([]g.String{"video", "custom"}).
			Add().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Video sticker set created successfully!").Send().Err()
	})

	// Advanced: Create mask sticker set
	b.Command("createmaskset", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 9 {
			return ctx.Reply("Usage: /createmaskset <user_id> <name> <title> <sticker_file> <emoji> <point> <x_shift> <y_shift> <scale>").
				Send().
				Err()
		}

		userID := args[0].ToInt().Unwrap().Int64()
		name := args[1]
		title := args[2]
		stickerFile := args[3]
		emoji := args[4]
		point := args[5]
		xShift := args[6].ToFloat().Unwrap()
		yShift := args[7].ToFloat().Unwrap()
		scale := args[8].ToFloat().Unwrap()

		result := ctx.CreateNewStickerSet(userID, name, title).
			StickerType("mask").
			Sticker(stickerFile, "static", []g.String{emoji}).
			MaskPosition(point, xShift.Std(), yShift.Std(), scale.Std()).
			Add().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Mask sticker set created successfully!").Send().Err()
	})

	// Multi-sticker set creation with different properties
	b.Command("multistickerset", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 7 {
			return ctx.Reply("Usage: /multistickerset <user_id> <name> <title> <file1> <emoji1> <file2> <emoji2>").
				Send().
				Err()
		}

		userID := args[0].ToInt().Unwrap().Int64()
		name := args[1]
		title := args[2]

		result := ctx.CreateNewStickerSet(userID, name, title).
			StickerType("regular").
			Sticker(args[3], "static", []g.String{args[4]}).
			Keywords([]g.String{"first", "custom", "happy"}).
			Add().
			Sticker(args[5], "static", []g.String{args[6]}).
			Keywords([]g.String{"second", "custom", "sad"}).
			Add().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Multi-sticker set created successfully!").Send().Err()
	})

	b.Polling().AllowedUpdates().Start()
}
