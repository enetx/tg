package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/g/ref"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/types/effects"
)

// SendPoll represents a request to send a poll.
type SendPoll struct {
	ctx         *Context
	question    g.String
	chatID      g.Option[int64]
	options     g.Slice[input.PollOption]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
	opts        *gotgbot.SendPollOpts
}

// QuestionHTML sets the question parse mode to HTML.
func (sp *SendPoll) QuestionHTML() *SendPoll {
	sp.opts.QuestionParseMode = "HTML"
	return sp
}

// QuestionMarkdown sets the question parse mode to MarkdownV2.
func (sp *SendPoll) QuestionMarkdown() *SendPoll {
	sp.opts.QuestionParseMode = "MarkdownV2"
	return sp
}

// QuestionEntities sets custom entities for the poll question.
func (sp *SendPoll) QuestionEntities(e *entities.Entities) *SendPoll {
	sp.opts.QuestionEntities = e.Std()
	return sp
}

// Media adds media to the poll question (photo, video, animation, audio, document,
// live photo, location, or venue).
func (sp *SendPoll) Media(media input.PollMedia) *SendPoll {
	sp.opts.Media = media.BuildPollMedia()
	return sp
}

// Description sets the poll description, 0-200 characters.
func (sp *SendPoll) Description(text g.String) *SendPoll {
	sp.opts.Description = text.Std()
	return sp
}

// DescriptionHTML sets the description parse mode to HTML.
func (sp *SendPoll) DescriptionHTML() *SendPoll {
	sp.opts.DescriptionParseMode = "HTML"
	return sp
}

// DescriptionMarkdown sets the description parse mode to MarkdownV2.
func (sp *SendPoll) DescriptionMarkdown() *SendPoll {
	sp.opts.DescriptionParseMode = "MarkdownV2"
	return sp
}

// DescriptionEntities sets custom entities for the poll description.
func (sp *SendPoll) DescriptionEntities(e *entities.Entities) *SendPoll {
	sp.opts.DescriptionEntities = e.Std()
	return sp
}

// ExplanationEntities sets custom entities for the poll explanation.
func (sp *SendPoll) ExplanationEntities(e *entities.Entities) *SendPoll {
	sp.opts.ExplanationEntities = e.Std()
	return sp
}

// After schedules the poll to be sent after the specified duration.
func (sp *SendPoll) After(duration time.Duration) *SendPoll {
	sp.after = g.Some(duration)
	return sp
}

// DeleteAfter schedules the poll message to be deleted after the specified duration.
func (sp *SendPoll) DeleteAfter(duration time.Duration) *SendPoll {
	sp.deleteAfter = g.Some(duration)
	return sp
}

// To sets the target chat ID for the poll.
func (sp *SendPoll) To(id int64) *SendPoll {
	sp.chatID = g.Some(id)
	return sp
}

// Option adds a poll option with the specified text.
func (sp *SendPoll) Option(option input.PollOption) *SendPoll {
	sp.options.Push(option)
	return sp
}

// Anonymous makes the poll anonymous.
func (sp *SendPoll) Anonymous() *SendPoll {
	sp.opts.IsAnonymous = ref.Of(true)
	return sp
}

// Business sets the business connection ID for the poll.
func (sp *SendPoll) Business(id g.String) *SendPoll {
	sp.opts.BusinessConnectionId = id.Std()
	return sp
}

// Thread sets the message thread ID for the poll.
func (sp *SendPoll) Thread(id int64) *SendPoll {
	sp.opts.MessageThreadId = id
	return sp
}

// AllowPaidBroadcast allows the poll to be sent in paid broadcast channels.
func (sp *SendPoll) AllowPaidBroadcast() *SendPoll {
	sp.opts.AllowPaidBroadcast = true
	return sp
}

// Effect sets a message effect for the poll.
func (sp *SendPoll) Effect(effect effects.EffectType) *SendPoll {
	sp.opts.MessageEffectId = effect.String()
	return sp
}

// MultipleAnswers allows users to select multiple answers.
func (sp *SendPoll) MultipleAnswers() *SendPoll {
	sp.opts.AllowsMultipleAnswers = true
	return sp
}

// AllowRevoting allows users to change their vote after it has been cast.
func (sp *SendPoll) AllowRevoting() *SendPoll {
	sp.opts.AllowsRevoting = ref.Of(true)
	return sp
}

// ShuffleOptions shuffles the order of the options each time the poll is displayed.
func (sp *SendPoll) ShuffleOptions() *SendPoll {
	sp.opts.ShuffleOptions = true
	return sp
}

// AllowAddingOptions allows users to add their own options to the poll.
func (sp *SendPoll) AllowAddingOptions() *SendPoll {
	sp.opts.AllowAddingOptions = true
	return sp
}

// HideResultsUntilClosed hides the poll results until the poll is closed.
func (sp *SendPoll) HideResultsUntilClosed() *SendPoll {
	sp.opts.HideResultsUntilCloses = true
	return sp
}

// MembersOnly restricts voting to the members of the chat where the poll is sent.
func (sp *SendPoll) MembersOnly() *SendPoll {
	sp.opts.MembersOnly = true
	return sp
}

// CountryCodes sets the list of two-letter ISO 3166-1 alpha-2 country codes whose
// residents are allowed to vote in the poll.
func (sp *SendPoll) CountryCodes(codes ...g.String) *SendPoll {
	sp.opts.CountryCodes = g.TransformSlice(codes, g.String.Std)
	return sp
}

// Protect enables content protection for the poll.
func (sp *SendPoll) Protect() *SendPoll {
	sp.opts.ProtectContent = true
	return sp
}

// Quiz converts the poll to a quiz with the specified correct option indexes.
// Pass one or more 0-based identifiers of the correct answer options (monotonically increasing).
func (sp *SendPoll) Quiz(correct ...int) *SendPoll {
	sp.opts.Type = "quiz"
	sp.opts.CorrectOptionIds = g.TransformSlice(correct, func(c int) int64 { return int64(c) })
	return sp
}

// Explanation sets an explanation text for quiz answers.
func (sp *SendPoll) Explanation(text g.String) *SendPoll {
	sp.opts.Explanation = text.Std()
	return sp
}

// ExplanationHTML sets the explanation parse mode to HTML.
func (sp *SendPoll) ExplanationHTML() *SendPoll {
	sp.opts.ExplanationParseMode = "HTML"
	return sp
}

// ExplanationMarkdown sets the explanation parse mode to MarkdownV2.
func (sp *SendPoll) ExplanationMarkdown() *SendPoll {
	sp.opts.ExplanationParseMode = "MarkdownV2"
	return sp
}

// ExplanationMedia adds media to the quiz explanation (photo, video, animation, audio,
// document, live photo, location, or venue).
func (sp *SendPoll) ExplanationMedia(media input.PollMedia) *SendPoll {
	sp.opts.ExplanationMedia = media.BuildPollMedia()
	return sp
}

// Silent disables notification for the poll.
func (sp *SendPoll) Silent() *SendPoll {
	sp.opts.DisableNotification = true
	return sp
}

// ClosesIn sets the poll to close after the specified duration.
func (sp *SendPoll) ClosesIn(duration time.Duration) *SendPoll {
	sp.opts.OpenPeriod = int64(duration.Seconds())
	return sp
}

// ClosesAt sets the poll to close at the specified time.
func (sp *SendPoll) ClosesAt(t time.Time) *SendPoll {
	sp.opts.CloseDate = t.Unix()
	return sp
}

// Closed marks the poll as already closed.
func (sp *SendPoll) Closed() *SendPoll {
	sp.opts.IsClosed = true
	return sp
}

// Markup sets the reply markup keyboard for the poll.
func (sp *SendPoll) Markup(kb keyboard.Keyboard) *SendPoll {
	sp.opts.ReplyMarkup = kb.Markup()
	return sp
}

// Reply sets reply parameters using the reply builder.
func (sp *SendPoll) Reply(params *reply.Parameters) *SendPoll {
	if params != nil {
		sp.opts.ReplyParameters = params.Std()
	}
	return sp
}

// Timeout sets a custom timeout for this request.
func (sp *SendPoll) Timeout(duration time.Duration) *SendPoll {
	if sp.opts.RequestOpts == nil {
		sp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sp.opts.RequestOpts.Timeout = duration

	return sp
}

// APIURL sets a custom API URL for this request.
func (sp *SendPoll) APIURL(url g.String) *SendPoll {
	if sp.opts.RequestOpts == nil {
		sp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sp.opts.RequestOpts.APIURL = url.Std()

	return sp
}

// Send sends the poll to Telegram and returns the result.
func (sp *SendPoll) Send() g.Result[*gotgbot.Message] {
	return sp.ctx.timers(sp.after, sp.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := sp.chatID.UnwrapOr(sp.ctx.EffectiveChat.Id)
		options := g.TransformSlice(sp.options, input.PollOption.Build)

		return g.ResultOf(sp.ctx.Bot.Raw().SendPoll(chatID, sp.question.Std(), options, sp.opts))
	})
}
