package main

import (
	"io"
	"net/http"

	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/types/updates"
	// "github.com/valyala/fasthttp"
)

func main() {
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token).Build().Unwrap()

	path := String("/bot/" + token)
	secret := String("my-secret-token")

	bot.Webhook().
		Domain("https://3b1d-134-19-179-195.ngrok-free.app").
		Path(path).
		SecretToken(secret).
		MaxConnections(100).
		DropPending(true).
		AllowedUpdates(updates.Message, updates.CallbackQuery).
		// Certificate("cert.pem").
		Register()

	bot.On.Message.Text(func(ctx *tg.Context) error { return ctx.Message("pong").Send().Err() })

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || String(r.URL.Path).Ne(path) {
			http.NotFound(w, r)
			return
		}

		if String(r.Header.Get("X-Telegram-Bot-Api-Secret-Token")).Ne(secret) {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		if err := bot.HandleWebhook(body); err != nil {
			http.Error(w, "invalid update: "+err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}

	// Println("Listening on https://0.0.0.0:8443 at {}", path)
	// err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", http.HandlerFunc(handler))

	Println("Listening on http://0.0.0.0:8080 at {}", path)
	err := http.ListenAndServe(":8080", http.HandlerFunc(handler))
	if err != nil {
		panic(err)
	}

	// handler := func(ctx *fasthttp.RequestCtx) {
	// 	if !ctx.IsPost() || String(ctx.Path()).Ne(path) {
	// 		ctx.SetStatusCode(fasthttp.StatusNotFound)
	// 		return
	// 	}
	//
	// 	if String(ctx.Request.Header.Peek("X-Telegram-Bot-Api-Secret-Token")).Ne(secret) {
	// 		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
	// 		return
	// 	}
	//
	// 	if err := bot.HandleWebhook(ctx.PostBody()); err != nil {
	// 		ctx.SetStatusCode(fasthttp.StatusBadRequest)
	// 		ctx.SetBodyString("invalid update: " + err.Error())
	// 		return
	// 	}
	//
	// 	ctx.SetStatusCode(fasthttp.StatusOK)
	// }
	//
	// // Println("Listening on https://0.0.0.0:8443 at {}", path)
	// // err := fasthttp.ListenAndServeTLS(":8443", "cert.pem", "key.pem", handler)
	//
	// Println("Listening on http://0.0.0.0:8080 at {}", path)
	// err := fasthttp.ListenAndServe(":8080", handler)
	// if err != nil {
	// 	panic(err)
	// }
}
