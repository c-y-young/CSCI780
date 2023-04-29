package agent

import (
	"gitlab.com/akita/akita/sim"
)

type Order struct {
	sim.MsgMeta

	Quatity int
}

func (o *Order) Meta() *sim.MsgMeta {
	return &o.MsgMeta
}

type Shipment struct {
	sim.MsgMeta

	Quatity int
}

func (s *Shipment) Meta() *sim.MsgMeta {
	return &s.MsgMeta
}
