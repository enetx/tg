package bot

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/types/updates"
)

type Polling struct {
	bot     *Bot
	opts    *ext.PollingOpts
	started bool
}

func (p *Polling) DropPendingUpdates() *Polling {
	p.opts.DropPendingUpdates = true
	return p
}

func (p *Polling) EnableWebhookDeletion() *Polling {
	p.opts.EnableWebhookDeletion = true
	return p
}

func (p *Polling) Timeout(seconds int64) *Polling {
	p.opts.GetUpdatesOpts.Timeout = seconds
	return p
}

func (p *Polling) Limit(n int64) *Polling {
	p.opts.GetUpdatesOpts.Limit = n
	return p
}

func (p *Polling) Offset(offset int64) *Polling {
	p.opts.GetUpdatesOpts.Offset = offset
	return p
}

func (p *Polling) AllowedUpdates(upds ...updates.UpdateType) *Polling {
	p.opts.GetUpdatesOpts.AllowedUpdates = g.TransformSlice(
		g.Slice[updates.UpdateType](upds),
		updates.UpdateType.String,
	)
	return p
}

func (p *Polling) Start() {
	if err := p.bot.updater.StartPolling(p.bot.Raw(), p.opts); err != nil {
		panic("failed to start polling: " + err.Error())
	}

	g.Println("bot started")
	p.bot.updater.Idle()
}

func (p *Polling) Opts() *ext.PollingOpts {
	return p.opts
}
