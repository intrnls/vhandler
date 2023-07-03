package vhandler

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
)

type InventoryHandler struct {
	inventory.Handler

	inv *Inventory
}

func newInventoryHandler(inv *Inventory) *InventoryHandler {
	return &InventoryHandler{inv: inv}
}

type InventoryTakeHandler func(ctx *event.Context, slot int, it item.Stack)

func (i *InventoryHandler) HandleTake(ctx *event.Context, slot int, it item.Stack) {
	i.inv.handlers[InventoryTakeId].handle(func(h Handler) {
		h.(InventoryTakeHandler)(ctx, slot, it)
	})
}

type InventoryPlaceHandler func(ctx *event.Context, slot int, it item.Stack)

func (i *InventoryHandler) HandlePlace(ctx *event.Context, slot int, it item.Stack) {
	i.inv.handlers[InventoryPlaceId].handle(func(h Handler) {
		h.(InventoryPlaceHandler)(ctx, slot, it)
	})
}

type InventoryDropHandler func(ctx *event.Context, slot int, it item.Stack)

func (i *InventoryHandler) HandleDrop(ctx *event.Context, slot int, it item.Stack) {
	i.inv.handlers[InventoryDropId].handle(func(h Handler) {
		h.(InventoryDropHandler)(ctx, slot, it)
	})
}
