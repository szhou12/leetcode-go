package networkflow

import (
	"testing"
)

/*
Terminal Commands:

	Run all tests that contains func_name:
	/networkflow $ go test -run TestMinCostMaxFlowJohnson

	Print out result (t.Logf) in console:
	/networkflow $ go test -v -run TestMinCostMaxFlowJohnson

	Run the only test named exactly as func_name:
	/networkflow $ go test -run ^TestMinCostMaxFlowJohnson$

*/

func TestMinCostMaxFlowJohnson(t *testing.T) {
	n := 6
	s := n - 1
	target := n - 2
	edges := [][]int{
		{s, 1, 4, 10},
		{s, 2, 2, 30},
		{1, 2, 2, 10},
		{1, target, 0, 9999},
		{2, target, 4, 10},
	}

	maxFlow, minCost := MinCostMaxFlowJohnson(n, edges, s, target)

	if maxFlow != 4 || minCost != 140 {
		t.Errorf("(maxFlow, minCost) Expected (4, 140), Got (%d, %d)", maxFlow, minCost)
	}
}
