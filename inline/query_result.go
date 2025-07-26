package inline

import "github.com/PaulSonOfLars/gotgbot/v2"

// QueryResult represents an interface for all inline query result builders.
type QueryResult interface {
	Build() gotgbot.InlineQueryResult
}
