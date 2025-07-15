package handlers

import "github.com/enetx/tg/core"

type Handlers struct {
	bot                core.BotAPI
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

func NewHandlers(bot core.BotAPI) *Handlers {
	return &Handlers{
		bot:                bot,
		Message:            &MessageHandlers{bot},
		Callback:           &CallbackHandlers{bot},
		Inline:             &InlineQueryHandlers{bot},
		Poll:               &PollHandlers{bot},
		PollAnswer:         &PollAnswerHandlers{bot},
		ChatMember:         &ChatMemberHandlers{bot},
		MyChatMember:       &MyChatMemberHandlers{bot},
		ChatJoinRequest:    &ChatJoinRequestHandlers{bot},
		ChosenInlineResult: &ChosenInlineResultHandlers{bot},
		Shipping:           &ShippingHandlers{bot},
		PreCheckout:        &PreCheckoutHandlers{bot},
		Reaction:           &ReactionHandlers{bot},
		PaidMedia:          &PaidMediaHandlers{bot},
	}
}

func (h *Handlers) Any(fn Handler) core.BotAPI {
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
