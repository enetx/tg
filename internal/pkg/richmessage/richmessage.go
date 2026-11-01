package richmessage

import (
	"fmt"
	"mime/multipart"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

// richMessage adds multipart attachment support to gotgbot.InputRichMessage.
type richMessage struct{ gotgbot.InputRichMessage }

func (m richMessage) Attach(name string, w *multipart.Writer) error {
	for i, media := range m.Media {
		if media.Media != nil {
			if err := media.Media.Attach(fmt.Sprintf("%s_media_%d", name, i), w); err != nil {
				return err
			}
		}
	}

	return attachRichBlocks(m.Blocks, name, w)
}

func attachRichBlocks(blocks []gotgbot.InputRichBlock, name string, w *multipart.Writer) error {
	for i, block := range blocks {
		if err := attachRichBlock(block, fmt.Sprintf("%s_block_%d", name, i), w); err != nil {
			return err
		}
	}

	return nil
}

func attachRichBlock(block gotgbot.InputRichBlock, name string, w *multipart.Writer) error {
	if media, ok := block.(gotgbot.Attach); ok {
		return media.Attach(name, w)
	}

	switch block := block.(type) {
	case gotgbot.InputRichBlockBlockQuotation:
		return attachRichBlocks(block.Blocks, name, w)
	case gotgbot.InputRichBlockCollage:
		return attachRichBlocks(block.Blocks, name, w)
	case gotgbot.InputRichBlockDetails:
		return attachRichBlocks(block.Blocks, name, w)
	case gotgbot.InputRichBlockSlideshow:
		return attachRichBlocks(block.Blocks, name, w)
	case gotgbot.InputRichBlockList:
		for i, item := range block.Items {
			if err := attachRichBlocks(item.Blocks, fmt.Sprintf("%s_item_%d", name, i), w); err != nil {
				return err
			}
		}
	case *gotgbot.InputRichBlockBlockQuotation:
		return attachRichBlock(*block, name, w)
	case *gotgbot.InputRichBlockCollage:
		return attachRichBlock(*block, name, w)
	case *gotgbot.InputRichBlockDetails:
		return attachRichBlock(*block, name, w)
	case *gotgbot.InputRichBlockSlideshow:
		return attachRichBlock(*block, name, w)
	case *gotgbot.InputRichBlockList:
		return attachRichBlock(*block, name, w)
	}

	return nil
}

// RequestOpts adds multipart attachment support for a rich message to the request options.
func RequestOpts(opts *gotgbot.RequestOpts, message gotgbot.InputRichMessage) *gotgbot.RequestOpts {
	if opts == nil {
		opts = new(gotgbot.RequestOpts)
	}

	if opts.OverrideParams == nil {
		opts.OverrideParams = make(map[string]any)
	}

	opts.OverrideParams["rich_message"] = richMessage{message}
	return opts
}
