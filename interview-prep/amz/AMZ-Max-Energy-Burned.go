package amz

/*
Given an array of length n - height[].
Define the energy burned by going from i-th to j-th element as (height[i] - height[j])^2.
You start at 0.
You can iterate through height[] in any order.
Find the max energy burned by visiting all elements in height[].

Example:
height = [5, 2, 5]

optimal path: 0 -> h[2] -> h[0] -> h[1]
max energy = (0-5)^2 + (5-2)^2 + (2-5)^2 = 43

Output: 43
*/

func getMaxEnergyBurned_dfs(height []int) int {

}