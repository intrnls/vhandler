package vhandler

import (
	"errors"

	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/imlighty/vhandler/priority"
)

type Inventory struct {
	handlers map[handlerId]*subHandler

	h *InventoryHandler
}

func NewInventory() *Inventory {
	v := &Inventory{handlers: map[handlerId]*subHandler{}}

	v.handlers[InventoryTakeId] = newSubHandler()
	v.handlers[InventoryPlaceId] = newSubHandler()
	v.handlers[InventoryDropId] = newSubHandler()

	return v
}

func (v *Inventory) OnTake(p priority.Priority, h InventoryTakeHandler) {
	v.handlers[InventoryTakeId].add(p, h)
}

func (v *Inventory) OnPlace(p priority.Priority, h InventoryPlaceHandler) {
	v.handlers[InventoryPlaceId].add(p, h)
}

func (v *Inventory) OnDrop(p priority.Priority, h InventoryDropHandler) {
	v.handlers[InventoryDropId].add(p, h)
}

func (v *Inventory) Set(i *inventory.Inventory) {
	i.Handle(v.h)
}

func (v *Inventory) Attach(p priority.Priority, ih inventory.Handler) {
	nop := inventory.NopHandler{}
	nopHandlers := v.getHandlers(nop)

	handlers := v.getHandlers(ih)
	for hId, handler := range handlers {
		if handler == nopHandlers[hId] {
			continue
		}
		v.handlers[hId].add(p, handler)
	}
}

func (v *Inventory) Detach(ih inventory.Handler) error {
	handlers := v.getHandlers(ih)
	for hId, handler := range handlers {
		if err := v.handlers[hId].remove(handler); err != nil {
			return err
		}
	}
	return nil
}

func (v *Inventory) Remove(h Handler) error {
	hId, ok := v.getHandlerId(h)
	if !ok {
		return errors.New("invalid handler")
	}
	return v.handlers[hId].remove(h)
}

func (*Inventory) getHandlers(h inventory.Handler) map[handlerId]Handler {
	var handlers map[handlerId]Handler

	handlers[InventoryTakeId] = h.HandleTake
	handlers[InventoryPlaceId] = h.HandlePlace
	handlers[InventoryDropId] = h.HandleDrop

	return handlers
}

func (v *Inventory) getHandlerId(h Handler) (handlerId, bool) {
	switch h.(type) {
	case InventoryTakeHandler:
		return InventoryTakeId, true
	case InventoryPlaceHandler:
		return InventoryPlaceId, true
	case InventoryDropHandler:
		return InventoryDropId, true
	default:
		return 0, false
	}
}
