package leetcdode

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
    // step 1: reconstruct adj-list repr
    next := make([][]int, n)
    for i := 0; i < n; i++ {
        next[i] = make([]int, 0)
    }
    for i := 0; i < n-1; i++ {
        next[i] = append(next[i], i+1)
    }

    res := make([]int, 0)
    for _, q := range queries {
        from, to := q[0], q[1]
        next[from] = append(next[from], to)

        dist := bfs(next)
        res = append(res, dist)
    }

    return res
}

func bfs(next [][]int) int {
    n := len(next)
    queue := make([]int, 0)

    ds := make([]int, n) // ds[i] = source node 0 to node i
    for i := 0; i < n; i++ {
        ds[i] = -1
    }

    // start node
    queue = append(queue, 0)
    ds[0] = 0

    // loop
    for len(queue) > 0 {
        // cur
        cur := queue[0]
        queue = queue[1:]

        // make the next move
        for _, nei := range next[cur] {
            if ds[nei] != -1 {
                continue
            }
            queue = append(queue, nei)
            ds[nei] = ds[cur]+1
        }
    }

    return ds[n-1]


}