package leetcode

import "math"

func minimumCost(target string, words []string, costs []int) int {
	n := len(target)

	// Step 1: Build Trie
	root := newTrieNode()
	for i, word := range words {
		node := root
		for _, char := range word {
			idx := int(char - 'a')
			if node.children[idx] == nil {
				node.children[idx] = newTrieNode()
			}
			node = node.children[idx]
		}

		// NOTE: there is a test case where the same word is associated with different costs, keep the smallest
		if node.cost == -1 {
			node.cost = costs[i]
		} else {
			node.cost = min(node.cost, costs[i])
		}
	}

	// Step 2: DFS + Memo

	// memo[i] := already seen min cost to chop target[i:n-1]
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = -1
	}

	res := dfs(target, 0, root, memo)
	if res != math.MaxInt/2 {
		return res
	}

	return -1
}


// [X X X X] {X  X  X | X  X}
//           cur    i
func dfs(target string, cur int, root *TrieNode, memo []int) int {
    // base case
    if cur == len(target) {
        // if nothing left to chop, return 0 cost
        return 0
    }

    // memoization / pruning
    if memo[cur] != -1 {
        return memo[cur]
    }

    res := math.MaxInt / 2
    node := root
    for i := cur; i < len(target); i++ {
        idx := int(target[i] - 'a')
        if node.children[idx] == nil {
            break
        }
        node = node.children[idx]
        if node.cost != -1 {
            res = min(res, dfs(target, i+1, root, memo) + node.cost)
        }

    }

    memo[cur] = res
    
    return res
}

type TrieNode struct {
	children [26]*TrieNode
	cost     int // also work as "isEnd"
}

func newTrieNode() *TrieNode {
	node := &TrieNode{
		cost: -1,
	}

	for i := 0; i < 26; i++ {
		node.children[i] = nil
	}

	return node
}
