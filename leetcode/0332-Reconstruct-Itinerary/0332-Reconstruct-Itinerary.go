package leetcode

import "sort"

type destPair struct {
	dest    string
	visited bool
}

func findItinerary(tickets [][]string) []string {
	// for each 出发地，map its 到达地list
	dests := make(map[string][]*destPair)
	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		if _, ok := dests[from]; !ok {
			dests[from] = make([]*destPair, 0)
		}
		dests[from] = append(dests[from], &destPair{dest: to, visited: false})
	}
	// sort each 到达地list by lexical order
	for from, _ := range dests {
		sort.Slice(dests[from], func(i, j int) bool {
			return dests[from][i].dest < dests[from][j].dest
		})
	}

	var res []string
	res = append(res, "JFK")
	dfs(dests, len(tickets), &res)
	return res

}

func dfs(dests map[string][]*destPair, ticketNum int, res *[]string) bool {
	// Base Case: all tickets have been used once and only once
	if ticketNum+1 == len(*res) {
		return true
	}

	depart := (*res)[len(*res)-1]
	for _, destPair := range dests[depart] {
		if !destPair.visited {
			*res = append(*res, destPair.dest)
			destPair.visited = true
			if dfs(dests, ticketNum, res) {
				return true
			}
			*res = (*res)[:len(*res)-1]
			destPair.visited = false
		}
	}
	return false
}
