package agent

import (
	"gitlab.com/akita/akita/sim"
)

var EndDayHookPos = &sim.HookPos{Name: "End Day Hook Pos"}

type EndDayHook struct {
	Cost int
}

func (hook *EndDayHook) Func(ctx sim.HookCtx) {
	//The Func method of the hook should first check if the position is EndDayHookPos. If it is, the function should accumulate the cost of the agent of the day
	if ctx.Pos == EndDayHookPos {
		agent := ctx.Domain.(*Agent)
		hook.Cost = agent.Inventory * agent.InventoryCost

		// If the agent is a retailer, also add the backlog * LostCustomerPenalty to the cost and set the backlog to 0.
		if agent.IsRetailer {
			hook.Cost = hook.Cost + agent.Backlog*agent.LostCustomerPenalty
			agent.Backlog = 0
		}
	}
}
