package bot

import (
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/types/updates"
)

type Webhook struct {
	bot    *Bot
	domain String
	path   String
	opt    *gotgbot.SetWebhookOpts
	cert   *os.File
}

func (w *Webhook) Certificate(path String) *Webhook {
	file, err := os.Open(path.Std())
	if err != nil {
		panic("failed to open certificate file: " + err.Error())
	}

	w.opt.Certificate = gotgbot.InputFileByReader(path.Std(), file)
	w.cert = file

	return w
}

func (w *Webhook) Domain(s String) *Webhook {
	w.domain = s
	return w
}

func (w *Webhook) Path(s String) *Webhook {
	w.path = s
	return w
}

func (w *Webhook) SecretToken(s String) *Webhook {
	w.opt.SecretToken = s.Std()
	return w
}

func (w *Webhook) DropPending(b bool) *Webhook {
	w.opt.DropPendingUpdates = b
	return w
}

func (w *Webhook) MaxConnections(n int) *Webhook {
	w.opt.MaxConnections = int64(n)
	return w
}

func (w *Webhook) IP(ip String) *Webhook {
	w.opt.IpAddress = ip.Std()
	return w
}

func (w *Webhook) AllowedUpdates(upds ...updates.UpdateType) *Webhook {
	w.opt.AllowedUpdates = TransformSlice(Slice[updates.UpdateType](upds), updates.UpdateType.String)
	return w
}

func (w *Webhook) Register() error {
	if w.cert != nil {
		defer w.cert.Close()
	}

	if w.domain.Empty() || w.path.Empty() {
		return Errorf("Webhook domain and path must be set")
	}

	url := w.domain.StripSuffix("/") + "/" + w.path.StripPrefix("/")

	_, err := w.bot.Raw().SetWebhook(url.Std(), w.opt)
	return err
}
