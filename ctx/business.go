package ctx

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

// Business creates a business API handler for the given connection ID.
func (ctx *Context) Business(connectionID g.String) *business.Account {
	return business.NewAccount(ctx.Bot, connectionID)
}
