package preview

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type Preview struct {
	opts *gotgbot.LinkPreviewOptions
}

// New creates a new Preview builder.
func New() *Preview {
	return &Preview{opts: new(gotgbot.LinkPreviewOptions)}
}

// Import copies settings from another LinkPreviewOptions object.
func (p *Preview) Import(src *gotgbot.LinkPreviewOptions) *Preview {
	if src != nil {
		*p.opts = *src
	}

	return p
}

// URL sets a custom preview URL.
func (p *Preview) URL(url String) *Preview {
	p.opts.Url = url.Std()
	return p
}

// Disable disables the link preview.
func (p *Preview) Disable() *Preview {
	p.opts.IsDisabled = true
	return p
}

// Above shows the preview above the message text.
func (p *Preview) Above() *Preview {
	p.opts.ShowAboveText = true
	return p
}

// Large prefers large media in the preview.
func (p *Preview) Large() *Preview {
	p.opts.PreferLargeMedia = true
	return p
}

// Small prefers small media in the preview.
func (p *Preview) Small() *Preview {
	p.opts.PreferSmallMedia = true
	return p
}

// Std returns the built LinkPreviewOptions object.
func (p *Preview) Std() *gotgbot.LinkPreviewOptions {
	return p.opts
}
