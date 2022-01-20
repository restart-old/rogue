package rogue

import (
	"github.com/RestartFU/tickerFunc"
	"github.com/df-HCF/class"
	"github.com/df-mc/dragonfly/server/item/armour"
	"github.com/df-mc/dragonfly/server/player"
)

type Rogue struct {
	tickers []*tickerFunc.Ticker
}

func (*Rogue) New(p *player.Player) class.Class {
	rogue := &Rogue{}
	return rogue
}

func (*Rogue) Armour() class.Armour {
	return class.Armour{
		armour.TierChain,
		armour.TierChain,
		armour.TierChain,
		armour.TierChain,
	}
}

func (r *Rogue) Handler(p *player.Player) player.Handler {
	return &handler{p: p}
}

func (r *Rogue) Tickers() []*tickerFunc.Ticker {
	return r.tickers
}
