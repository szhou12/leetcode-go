package leetcode

func minIncrease(n int, edges [][]int, cost []int) int {
    // step 1: reconstruct edges to adj-list representation
    next := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        next[u] = append(next[u], v)
        next[v] = append(next[v], u)
    }

    _, res := dfs(0, -1, cost, next)

    return res
}

// return: 
// 1st arg: max cost of subtree from cur to leaf
// 2nd arg: total boosts to make every cur-to-leaf path equal to max cost
func dfs(cur int, parent int, cost []int, next [][]int) (int, int) {
    maxCost := 0
    totalBoosts := 0
    paths := make([]int, 0)

    for _, nei := range next[cur] {
        if nei == parent {
            continue
        }
        pathCost, pathBoosts := dfs(nei, cur, cost, next)
        maxCost = max(maxCost, pathCost)
        totalBoosts += pathBoosts
        paths = append(paths, pathCost)
    }

    count := 0
    for _, p := range paths {
        if p < maxCost {
            count++
        }
    }

    return maxCost + cost[cur], totalBoosts + count

}