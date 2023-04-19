package leetcode

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	// Step 1: Reconstruct adj-list repr + calc degree
	degree := make(map[string]int)
	// Not necessary to have this: all supplies have degree = 0
	for _, s := range supplies {
		degree[s] = 0
	}
	next := make(map[string]map[string]bool) // {from: {to: true}, {...}}
	for i := 0; i < len(recipes); i++ {
		to := recipes[i]
		froms := ingredients[i]
		for _, from := range froms {
			if _, ok := next[from]; !ok { // if key doesn't exist, create it
				next[from] = make(map[string]bool)
			}
			if next[from][to] { // bc we use map, need to avoid duplicates
				continue
			}
			next[from][to] = true
			degree[to]++
		}
	}

	// Step 2: Topological Sort
	recipeMap := make(map[string]bool) // Mark recipes that can be made
	for _, recipe := range recipes {
		recipeMap[recipe] = false
	}

	queue := make([]string, 0)

	// Start nodes: supplies
	queue = append(queue, supplies...)

	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]
		// update results
		if _, ok := recipeMap[cur]; ok { // if cur is a recipe we want
			recipeMap[cur] = true
		}

		// Make the next move
		for nei, _ := range next[cur] {
			degree[nei]--
			if degree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	// Step 3: collect recipes that can be made
	res := make([]string, 0)
	for k, v := range recipeMap {
		if v {
			res = append(res, k)
		}
	}

	return res
}
