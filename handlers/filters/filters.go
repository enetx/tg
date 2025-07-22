package filters

import "github.com/PaulSonOfLars/gotgbot/v2"

// BusinessMessagesDeleted is a filter function that determines whether to handle deleted business messages.
type BusinessMessagesDeleted func(*gotgbot.BusinessMessagesDeleted) bool
