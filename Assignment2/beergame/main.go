package main

import (
	"agent"

	"gitlab.com/akita/akita/sim"
)

type NewCustomerEvent struct {
	time     sim.VTimeInSec
	handler  sim.Handler
	quantity int
}

func (event *NewCustomerEvent) Handler() sim.Handler {
	return event.handler
}

func (event *NewCustomerEvent) Time() sim.VTimeInSec {
	return event.time
}

func (event *NewCustomerEvent) IsSecondary() bool {
	return false
}

type NewCustomerEventHandler struct {
	Retailer *agent.Agent
}

func (handler *NewCustomerEventHandler) Handle(event sim.Event) {
	order := &agent.Order{Quatity: event.(*NewCustomerEvent).quantity}
	handler.Retailer.DownStream.Recv(order)
}

func main() {
	engine := sim.NewSerialEngine()

	retailer := agent.NewAgent(engine, "RETAILER")
	retailer.IsRetailer = true
	retailer.IsFactory = false
	retailer.InventoryCost = 4
	retailer.LostCustomerPenalty = 30

	wholesaler := agent.NewAgent(engine, "Wholesaler")
	wholesaler.IsRetailer = false
	wholesaler.IsFactory = false
	wholesaler.InventoryCost = 3

	distributor := agent.NewAgent(engine, "Distributor")
	distributor.IsRetailer = false
	distributor.IsFactory = false
	distributor.InventoryCost = 2

	factory := agent.NewAgent(engine, "Factory")
	factory.IsRetailer = false
	factory.IsFactory = true
	factory.InventoryCost = 1

	retailer.UpStreamAgent = wholesaler.DownStream

	wholesaler.UpStreamAgent = distributor.DownStream

	wholesaler.DownStreamAgent = retailer.UpStream

	distributor.UpStreamAgent = factory.DownStream

	distributor.DownStreamAgent = wholesaler.UpStream

	factory.DownStreamAgent = distributor.UpStream

	con_1 := sim.NewDirectConnection("connection 1", engine, 1*sim.GHz)
	con_2 := sim.NewDirectConnection("connection 2", engine, 1*sim.GHz)
	con_3 := sim.NewDirectConnection("connection 3", engine, 1*sim.GHz)

	con_1.PlugIn(retailer.UpStream, 1)
	con_1.PlugIn(wholesaler.DownStream, 1)

	con_2.PlugIn(wholesaler.UpStream, 1)
	con_2.PlugIn(distributor.DownStream, 1)

	con_3.PlugIn(distributor.UpStream, 1)
	con_3.PlugIn(factory.DownStream, 1)

	quantity := 0
	handler := &NewCustomerEventHandler{Retailer: retailer}
	for i := 0; i < 20; i++ {
		if i <= 8 {
			quantity = 4
		} else {
			quantity = 16
		}

		event := &NewCustomerEvent{time: sim.VTimeInSec(), quantity: quantity, handler: handler}
		engine.Schedule(event)
	}

	cost := &agent.EndDayHook{
		Cost: 0,
	}

	retailer.AcceptHook(cost)

	wholesaler.AcceptHook(cost)

	distributor.AcceptHook(cost)

	factory.AcceptHook(cost)

	engine.Run()

	print("Total cost:", cost.Cost)

}
