package rogue

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type handler struct {
	player.NopHandler
	p *player.Player
}

func (h *handler) HandleAttackEntity(ctx *event.Context, e world.Entity, force, height *float64, critical *bool) {
h.p.

}
