package bot

import (
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/types/updates"
)

type SetWebhook struct {
	bot    *Bot
	domain g.String
	path   g.String
	opts   *gotgbot.SetWebhookOpts
	cert   *os.File
}

func (w *SetWebhook) Certificate(path g.String) *SetWebhook {
	file, err := os.Open(path.Std())
	if err != nil {
		panic("failed to open certificate file: " + err.Error())
	}

	w.opts.Certificate = gotgbot.InputFileByReader(path.Std(), file)
	w.cert = file

	return w
}

func (w *SetWebhook) Domain(s g.String) *SetWebhook {
	w.domain = s
	return w
}

func (w *SetWebhook) Path(s g.String) *SetWebhook {
	w.path = s
	return w
}

func (w *SetWebhook) SecretToken(s g.String) *SetWebhook {
	w.opts.SecretToken = s.Std()
	return w
}

func (w *SetWebhook) DropPending(b bool) *SetWebhook {
	w.opts.DropPendingUpdates = b
	return w
}

func (w *SetWebhook) MaxConnections(n int) *SetWebhook {
	w.opts.MaxConnections = int64(n)
	return w
}

func (w *SetWebhook) IP(ip g.String) *SetWebhook {
	w.opts.IpAddress = ip.Std()
	return w
}

func (w *SetWebhook) AllowedUpdates(upds ...updates.UpdateType) *SetWebhook {
	w.opts.AllowedUpdates = g.TransformSlice(g.Slice[updates.UpdateType](upds), updates.UpdateType.String)
	return w
}

func (w *SetWebhook) Register() g.Result[bool] {
	if w.cert != nil {
		defer w.cert.Close()
	}

	if w.domain.IsEmpty() || w.path.IsEmpty() {
		return g.Err[bool](g.Errorf("Webhook domain and path must be set"))
	}

	url := w.domain.StripSuffix("/") + "/" + w.path.StripPrefix("/")

	return g.ResultOf(w.bot.Raw().SetWebhook(url.Std(), w.opts))
}

func (w *SetWebhook) Opts() *gotgbot.SetWebhookOpts {
	return w.opts
}
