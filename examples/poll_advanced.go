package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

// Advanced polls (Bot API 9.5 / 9.6): descriptions, shuffled/members-only voting,
// hidden results, revoting, user-added options, country restrictions, media in the
// question and in individual options, plus multi-answer quizzes.
func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// /survey — a regular poll using the new 9.5/9.6 options.
	b.Command("survey", func(ctx *ctx.Context) error {
		return ctx.SendPoll(g.String("How do you commute to work?")).
			Description(g.String("Pick the option you use most often")).
			Option(input.Choice("Car")).
			Option(input.Choice("Bike")).
			Option(input.Choice("Public transport")).
			MultipleAnswers().
			AllowRevoting().
			ShuffleOptions().
			AllowAddingOptions().
			HideResultsUntilClosed().
			MembersOnly().
			CountryCodes(g.String("US"), g.String("GB")).
			Send().Err()
	})

	// /mediapoll — attach media to the question and to individual options.
	b.Command("mediapoll", func(ctx *ctx.Context) error {
		return ctx.SendPoll(g.String("Where should we meet?")).
			Media(input.Photo(file.Input(g.String("map.jpg")).Ok())).
			Option(input.Choice("The office").
				Media(input.VenueMedia(40.7128, -74.0060, g.String("HQ"), g.String("New York, NY")))).
			Option(input.Choice("The park").
				Media(input.LocationMedia(40.7829, -73.9654))).
			Send().Err()
	})

	// /quiz — a quiz with several correct options and a media-rich explanation.
	b.Command("quiz", func(ctx *ctx.Context) error {
		return ctx.SendPoll(g.String("Which of these are prime numbers?")).
			Option(input.Choice("2")).
			Option(input.Choice("3")).
			Option(input.Choice("4")).
			Option(input.Choice("5")).
			Quiz(0, 1, 3). // 2, 3 and 5 are prime
			Explanation(g.String("4 = 2 × 2 is the only composite number here")).
			ExplanationMedia(input.LocationMedia(0, 0)).
			Send().Err()
	})

	b.Polling().Start()
}
