package ctx

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

// Business creates a business API handler for the given connection ID.
func (ctx *Context) Business(connectionID String) *business.Account {
	return business.NewAccount(ctx.Bot, connectionID)
}
