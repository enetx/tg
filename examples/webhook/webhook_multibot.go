package main

import (
	"io"
	"net/http"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/types/updates"
	// "github.com/valyala/fasthttp"
)

var bots = g.NewMap[g.String, *bot.Bot]()

func main() {
	domain := g.String("https://3b1d-134-19-179-195.ngrok-free.app")

	register("111111111:AAA...A", domain)
	register("222222222:BBB...B", domain)
	register("333333333:CCC...C", domain)

	g.Println("Listening on :8080")

	if err := http.ListenAndServe(":8080", http.HandlerFunc(handler)); err != nil {
		panic(err)
	}

	// if err := fasthttp.ListenAndServe(":8080", handler); err != nil {
	// 	panic(err)
	// }
}

func register(token, domain g.String) {
	path := "/bot/" + token

	b := bot.New(token).Build().Unwrap()

	b.Webhook().
		Domain(domain).
		Path(path).
		SecretToken(token.Hash().MD5()). // use md5 hash of the token as a valid secret
		AllowedUpdates(updates.Message, updates.CallbackQuery).
		DropPending(true).
		MaxConnections(100).
		Register()

	b.On.Message.Text(func(ctx *ctx.Context) error {
		return ctx.SendMessage("Hi from bot: " + token.Hash().MD5()).Send().Err()
	})

	bots.Set(token, b)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	var (
		path  g.String
		token g.String
	)

	// /bot/<token>
	g.String(r.URL.Path).Split("/").Exclude(g.String.Empty).Collect().Unpack(&path, &token)

	if path.Ne("bot") {
		http.NotFound(w, r)
		return
	}

	bot := bots.Get(token)
	if bot.IsNone() {
		http.Error(w, "unknown bot", http.StatusNotFound)
		return
	}

	if g.String(r.Header.Get("X-Telegram-Bot-Api-Secret-Token")).Ne(token.Hash().MD5()) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}

	if err := bot.Some().HandleWebhook(body); err != nil {
		http.Error(w, "invalid update: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// func handler(ctx *fasthttp.RequestCtx) {
// 	if !ctx.IsPost() {
// 		ctx.SetStatusCode(fasthttp.StatusNotFound)
// 		return
// 	}
//
// 	var (
// 		path  g.String
// 		token g.String
// 	)
//
// 	// /bot/<token>
// 	String(ctx.Path()).Split("/").Exclude(String.Empty).Collect().Unpack(&path, &token)
//
// 	if path.Ne("bot") {
// 		ctx.SetStatusCode(fasthttp.StatusNotFound)
// 		return
// 	}
//
// 	b := bots.Get(token)
// 	if b.IsNone() {
// 		ctx.SetStatusCode(fasthttp.StatusNotFound)
// 		ctx.SetBodyString("unknown bot")
// 		return
// 	}
//
// 	if g.String(ctx.Request.Header.Peek("X-Telegram-Bot-Api-Secret-Token")).Ne(token.Hash().MD5()) {
// 		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
// 		return
// 	}
//
// 	if err := b.Some().HandleWebhook(ctx.PostBody()); err != nil {
// 		ctx.SetStatusCode(fasthttp.StatusBadRequest)
// 		ctx.SetBodyString("invalid update: " + err.Error())
// 		return
// 	}
//
// 	ctx.SetStatusCode(fasthttp.StatusOK)
// }
