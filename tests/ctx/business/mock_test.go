package business_test

import "github.com/PaulSonOfLars/gotgbot/v2"

type mockBot struct{}

func (m *mockBot) Raw() *gotgbot.Bot {
	return &gotgbot.Bot{}
}
