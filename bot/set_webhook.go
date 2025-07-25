package bot

import (
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/types/updates"
)

type SetWebhook struct {
	bot    *Bot
	domain String
	path   String
	opt    *gotgbot.SetWebhookOpts
	cert   *os.File
}

func (w *SetWebhook) Certificate(path String) *SetWebhook {
	file, err := os.Open(path.Std())
	if err != nil {
		panic("failed to open certificate file: " + err.Error())
	}

	w.opt.Certificate = gotgbot.InputFileByReader(path.Std(), file)
	w.cert = file

	return w
}

func (w *SetWebhook) Domain(s String) *SetWebhook {
	w.domain = s
	return w
}

func (w *SetWebhook) Path(s String) *SetWebhook {
	w.path = s
	return w
}

func (w *SetWebhook) SecretToken(s String) *SetWebhook {
	w.opt.SecretToken = s.Std()
	return w
}

func (w *SetWebhook) DropPending(b bool) *SetWebhook {
	w.opt.DropPendingUpdates = b
	return w
}

func (w *SetWebhook) MaxConnections(n int) *SetWebhook {
	w.opt.MaxConnections = int64(n)
	return w
}

func (w *SetWebhook) IP(ip String) *SetWebhook {
	w.opt.IpAddress = ip.Std()
	return w
}

func (w *SetWebhook) AllowedUpdates(upds ...updates.UpdateType) *SetWebhook {
	w.opt.AllowedUpdates = TransformSlice(Slice[updates.UpdateType](upds), updates.UpdateType.String)
	return w
}

func (w *SetWebhook) Register() error {
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
