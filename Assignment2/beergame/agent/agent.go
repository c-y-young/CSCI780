package agent

import (
	"gitlab.com/akita/akita/sim"
)

type Agent struct {
	*sim.TickingComponent
	Upstream        sim.Port
	DownStream      sim.Port
	UpstreamAgent   sim.Port
	DownstreamAgent sim.Port

	Inventory           int
	Backlog             int
	OnOrder             int
	InventoryCost       int
	LostCustomerPenalty int
	IsRetailer          bool
	IsFactory           bool
}

func NewAgent(engine sim.Engine, name string) *Agent {

	agent := new(Agent)

	agent.TickingComponent = sim.NewTickingComponent(name, engine, 1, agent)

	agent.Upstream = sim.NewLimitNumMsgPort(agent, 4, name+"'s up stream")

	agent.DownStream = sim.NewLimitNumMsgPort(agent, 4, name+"'s down stream")

	return agent

}

func (agent *Agent) Tick(currentTime sim.VTimeInSec) bool {
	// receive order
	// check the Downstream port for any incoming messages
	incomingMessagesOrder := agent.DownStream.Retrieve(sim.VTimeInSec(currentTime))

	// If there is a message, cast it to Order and add the quantity to the backlog.
	order, status := incomingMessagesOrder.(*Order)
	//
	if status == 1 {
		agent.Backlog = agent.Backlog + order.Quatity
	}

	// receive shipments
	//check the Upstream port for any incoming messages
	incomingMessagesShipments := agent.UpStream.Retrive(sim.VTimeInSec(currentTime))
	//If there is a message, cast it to Shipment and add the quantity to the inventory.
	shipment, status := incomingMessagesShipments.(*Shipment)
	//
	if status == 1 {
		agent.Inventory == agent.nventory+shipment.Quatity
	}

	// send shipments
	// If either the inventory or the backlog is 0, do not create a shipment.
	if agent.Inventory > 0 && agent.Backlog > 0 {
		if agent.Inventory < agent.Backlog {
			quatity := float64(agent.Inventory)
		} else {
			quatity := float64(agent.Backlog)
		}
	}
	// If the shipment is sent, reduce the backlog and the inventory by the quantity of the shipment.
	// If the agent is retailer, simply update the variables but do not create a shipment.
	if agent.IsRetailer {
		agent.Inventory = agent.Inventory - quatity
		agent.Backlog = agent.Backlog - quatity
	} else {
		shipment := &Shipment{
			MsgMeta: sim.MsgMeta{Src: agent.DownStream, Dst: agent.DownstreamAgent, SendTime: currentTime},
			Quatity: quatity,
		}
	}
	agent.DownStream.Send(shipment)

	agent.Inventory = agent.Inventory - quatity
	agent.Backlog = agent.Backlog - quatity

	// send order
	// quantity = backlog - inventory - onorder.
	// If the order is sent, update the onorder variable
	// If the agent is the factory, simply update the inventory variable but do not create an order.
	if agent.IsFactory {
		agent.Inventory = agent.Inventory + 50
	} else {
		quatity := agent.Backlog - agent.Inventory - agent.OnOrder
		if quatity {
			order := &Order{
				MsgMeta: sim.MsgMeta{Src: agent.UpStream, Dst: agent.UpstreamAgent, SendTime: currentTime},
				Quatity: quatity,
			}
			agent.Upstream.Send(order)
			agent.OnOrder = agent.OnOrder + quatity
		}
	}

	if currentTime > 1 {
		return false
	} else {
		hookCtx := sim.HookCtx{
			Domain: agent,
			Pos:    EndDayHookPos,
		}
		agent.InvokeHook(hookCtx)
		return true
	}
}
