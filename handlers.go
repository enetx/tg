package tg

type Handlers struct {
	bot                *Bot
	Message            *MessageHandlers
	Callback           *CallbackHandlers
	Inline             *InlineQueryHandlers
	Poll               *PollHandlers
	PollAnswer         *PollAnswerHandlers
	ChatMember         *ChatMemberHandlers
	MyChatMember       *MyChatMemberHandlers
	ChatJoinRequest    *ChatJoinRequestHandlers
	ChosenInlineResult *ChosenInlineResultHandlers
	Shipping           *ShippingHandlers
	PreCheckout        *PreCheckoutHandlers
	Reaction           *ReactionHandlers
	PaidMedia          *PaidMediaHandlers
}

func newHandlers(b *Bot) *Handlers {
	return &Handlers{
		bot:                b,
		Message:            &MessageHandlers{b},
		Callback:           &CallbackHandlers{b},
		Inline:             &InlineQueryHandlers{b},
		Poll:               &PollHandlers{b},
		PollAnswer:         &PollAnswerHandlers{b},
		ChatMember:         &ChatMemberHandlers{b},
		MyChatMember:       &MyChatMemberHandlers{b},
		ChatJoinRequest:    &ChatJoinRequestHandlers{b},
		ChosenInlineResult: &ChosenInlineResultHandlers{b},
		Shipping:           &ShippingHandlers{b},
		PreCheckout:        &PreCheckoutHandlers{b},
		Reaction:           &ReactionHandlers{b},
		PaidMedia:          &PaidMediaHandlers{b},
	}
}

func (h *Handlers) Any(fn Handler) *Bot {
	h.Message.Any(fn)
	h.Callback.Any(fn)
	h.Inline.Any(fn)
	h.Poll.Any(fn)
	h.PollAnswer.Any(fn)
	h.ChatMember.Any(fn)
	h.MyChatMember.Any(fn)
	h.ChatJoinRequest.Any(fn)
	h.ChosenInlineResult.Any(fn)
	h.Shipping.Any(fn)
	h.PreCheckout.Any(fn)
	h.Reaction.Any(fn)
	h.PaidMedia.Any(fn)

	return h.bot
}
