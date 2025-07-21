package handlers

import "github.com/PaulSonOfLars/gotgbot/v2/ext"

type namedHandler struct {
	name string
	ext.Handler
}

// Name returns the name of the handler.
func (n namedHandler) Name() string {
	return n.name
}
