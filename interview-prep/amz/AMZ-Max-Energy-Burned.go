package amz

import "sort"

/*
Given an array of length n: height[].
Define the energy burned by going from i-th to j-th element as (height[i] - height[j])^2.
You start at value 0.
You can iterate through height[] in any order.
Find the max energy burned by visiting all elements in height[].

Example:
height = [5, 2, 5]

optimal path: 0 -> h[2] -> h[0] -> h[1]
max energy = (0-5)^2 + (5-2)^2 + (2-5)^2 = 43

Output: 43
*/

/*
Key Observation: squared distance is largest when the two unused elements are farthest apart.
Approach: Greedy (Sorting + Two Pointers in opposite directions)
Sort the array in ascending order, then goes: 0 -> h[n-1] -> h[0] -> h[n-2] -> h[1] -> ...
*/
func getMaxEnergyBurned(height []int) int {
	n := len(height)
	sort.Ints(height)

	left, right := 0, n-1
	count := 0
	res := 0
	for left <= right {
		if count == 0 {
			res += height[right] * height[right]
			right--
			count++
		} else {
			if count%2 == 1 {
				res += (height[left] - height[right+1]) * (height[left] - height[right+1])
				left++
				count++
			} else {
				res += (height[right] - height[left-1]) * (height[right] - height[left-1])
				right--
				count++
			}
		}
	}
	return res

}

// same greedy approach, different implementation
func getMaxEnergyBurned_v2(height []int) int {
	n := len(height)

	// Sort the array in increasing order
	sort.Ints(height)

	// Two pointers from opposite directions
	// 双指针相向而行
	sequence := make([]int, 0)
	left, right := 0, n-1
	for left <= right {
		if left != right {
			sequence = append(sequence, height[right])
			right--
			sequence = append(sequence, height[left])
			left++
		} else {
			sequence = append(sequence, height[left])
			left++
		}
	}

	// prepend 0
	sequence = append([]int{0}, sequence...)

	res := 0
	for i := 1; i < len(sequence); i++ {
		diff := sequence[i] - sequence[i-1]
		res += diff * diff
	}

	return res
}

// DFS brute-force: get all permutations with duplicates, then find the perm that gives the max energy burned.
func getMaxEnergyBurned_dfs(height []int) int {

	// Step 1: DFS to get all permutations with dups
	sort.Ints(height)

	allPerms := make([][]int, 0)
	visited := make(map[int]bool)
	var dfs func(path []int)
	dfs = func(path []int) {
		// base case
		if len(path) == len(height) {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			allPerms = append(allPerms, pathCopy)
			return
		}

		// recursive
		for i := 0; i < len(height); i++ {
			// skip visited
			if visited[i] {
				continue
			}
			// skip duplicates
			if i > 0 && height[i] == height[i-1] && !visited[i-1] {
				continue
			}

			visited[i] = true
			path = append(path, height[i])
			dfs(path)
			visited[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs([]int{})

	// Step 2: for each permuation, find the one that gives max energy
	var calEnergy func(nums []int) int
	calEnergy = func(nums []int) int {
		n := len(nums)

		energy := nums[0] * nums[0]
		for i := 1; i < n; i++ {
			energy += (nums[i] - nums[i-1]) * (nums[i] - nums[i-1])
		}
		return energy
	}

	res := 0
	for _, perm := range allPerms {
		res = max(res, calEnergy(perm))
	}

	return res

}
