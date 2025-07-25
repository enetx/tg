package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// ShippingOptionBuilder helps build shipping options with prices.
type ShippingOptionBuilder struct {
	asq    *AnswerShippingQuery
	id     String
	title  String
	prices Slice[gotgbot.LabeledPrice]
}

// Price adds a labeled price to the shipping option.
func (sob *ShippingOptionBuilder) Price(label String, amount int64) *ShippingOptionBuilder {
	sob.prices.Push(gotgbot.LabeledPrice{Label: label.Std(), Amount: amount})
	return sob
}

// Done finishes building the shipping option and adds it to the query.
func (sob *ShippingOptionBuilder) Done() *AnswerShippingQuery {
	option := gotgbot.ShippingOption{
		Id:     sob.id.Std(),
		Title:  sob.title.Std(),
		Prices: sob.prices,
	}

	sob.asq.options.Push(option)

	return sob.asq
}

// AnswerShippingQuery represents a request to answer a shipping query.
type AnswerShippingQuery struct {
	ctx     *Context
	ok      bool
	options Slice[gotgbot.ShippingOption]
	opts    *gotgbot.AnswerShippingQueryOpts
}

// Ok marks the shipping query as successful and sets shipping options.
func (asq *AnswerShippingQuery) Ok() *AnswerShippingQuery {
	asq.ok = true
	return asq
}

// Error marks the shipping query as failed with the specified error message.
func (asq *AnswerShippingQuery) Error(text String) *AnswerShippingQuery {
	asq.ok = false
	asq.opts.ErrorMessage = text.Std()

	return asq
}

// Option adds a shipping option to the query response.
func (asq *AnswerShippingQuery) Option(id, title String) *ShippingOptionBuilder {
	return &ShippingOptionBuilder{
		asq:   asq,
		id:    id,
		title: title,
	}
}

// AddOption adds a pre-built shipping option to the query response.
func (asq *AnswerShippingQuery) AddOption(option gotgbot.ShippingOption) *AnswerShippingQuery {
	asq.options.Push(option)
	return asq
}

// Options sets multiple shipping options at once.
func (asq *AnswerShippingQuery) Options(options Slice[gotgbot.ShippingOption]) *AnswerShippingQuery {
	asq.options = options
	return asq
}

// Timeout sets a custom timeout for this request.
func (asq *AnswerShippingQuery) Timeout(duration time.Duration) *AnswerShippingQuery {
	if asq.opts.RequestOpts == nil {
		asq.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	asq.opts.RequestOpts.Timeout = duration

	return asq
}

// APIURL sets a custom API URL for this request.
func (asq *AnswerShippingQuery) APIURL(url String) *AnswerShippingQuery {
	if asq.opts.RequestOpts == nil {
		asq.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	asq.opts.RequestOpts.APIURL = url.Std()

	return asq
}

// Send answers the shipping query and returns the result.
func (asq *AnswerShippingQuery) Send() Result[bool] {
	query := asq.ctx.Update.ShippingQuery
	if query == nil {
		return Err[bool](Errorf("no shipping query"))
	}

	if asq.ok {
		asq.opts.ShippingOptions = asq.options
	}

	return ResultOf(query.Answer(asq.ctx.Bot.Raw(), asq.ok, asq.opts))
}
