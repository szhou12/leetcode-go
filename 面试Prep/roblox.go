package interviews

import (
	"math"
	"strconv"
	"strings"
)

// greater/smaller/tie v.s. pivot
func pivotCompare(nums []int, pivot int) string {
	greater := 0
	smaller := 0
	for _, val := range nums {
		if val < pivot {
			smaller++
		} else if val > pivot {
			greater++
		}
	}

	if greater > smaller {
		return "greater"
	} else if greater < smaller {
		return "smaller"
	} else {
		return "tie"
	}
}

/*********************************************************************/
// lowest in range
func lowestInRange(nums []int, nRange []int) int {
	res := math.MaxInt
	found := false

	for _, num := range nums {
		if nRange[0] < num && num < nRange[1] {
			res = min(res, num)
			found = true
		}
	}

	if found {
		return res
	}
	return 0
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

/*********************************************************************/
// league-style competition
func compete(wins []int, draws []int, scored []int, conceded []int) int {
	highest := math.MinInt
	tops := []int{}
	n := len(wins)
	for i := 0; i < n; i++ {
		score := 3*wins[i] + draws[i]
		if score == highest {
			tops = append(tops, i)
		} else if score > highest {
			tops = []int{i}
			highest = score
		} else {
			continue
		}
	}

	if len(tops) == 1 {
		return tops[0]
	}

	var res int
	netGoals := math.MinInt
	for _, id := range tops {
		curGoals := scored[id] - conceded[id]
		if curGoals > netGoals {
			netGoals = curGoals
			res = id
		}
	}
	return res
}

/*********************************************************************/
// BFS / DFS
// Candy Crush - burst bubbles
type coord struct {
	x int
	y int
}

func bfs(bubbles [][]int, x int, y int, visited map[coord]bool, m int, n int) []coord {
	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	queue := make([]coord, 0)
	start := coord{x: x, y: y}
	queue = append(queue, start)
	target := bubbles[x][y] // target color

	res := []coord{start}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[:len(queue)-1]
		for _, dir := range dirs {
			newX, newY := cur.x+dir[0], cur.y+dir[1]
			newCoord := coord{x: newX, y: newY}
			if newX < 0 || newX >= m { // newX out of bound
				continue
			}
			if newY < 0 || newY >= n { // newY out of bound
				continue
			}
			if _, ok := visited[newCoord]; ok { // already visited
				continue
			}
			if bubbles[newX][newY] != target { // not the same color
				continue
			}

			res = append(res, newCoord)
			visited[newCoord] = true
			queue = append(queue, newCoord)
		}
	}
	return res

}

func candyCrush(bubbles [][]int) [][]int {
	m, n := len(bubbles), len(bubbles[0])
	visited := make(map[coord]bool)
	removes := make([]coord, 0)

	// find same colors
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			curCoord := coord{x: i, y: j}
			visited[curCoord] = true
			sameColorCells := bfs(bubbles, i, j, visited, m, n)
			if len(sameColorCells) >= 3 {
				removes = append(removes, sameColorCells...)
			}
		}
	}

	// change them to 0
	for _, c := range removes {
		bubbles[c.x][c.y] = 0
	}

	// Move Zeros (leetcode 283) for every col
	for j := 0; j < n; j++ { // move zeros for every col
		left := m - 1
		for right := m - 1; right >= 0; right-- {
			if bubbles[right][j] != 0 {
				bubbles[right][j], bubbles[left][j] = bubbles[left][j], bubbles[right][j]
				left--
			}
		}
	}

	return bubbles

}

/*********************************************************************/
// blocks & obstacles
func buildBlocks(operations [][]int) string {
	res := ""
	obstacles := make(map[int]bool)

	for _, op := range operations {
		if op[0] == 1 {
			obstacles[op[1]] = true
		} else {
			size := op[2]
			x := op[1]
			possible := true
			for i := x - size; i < x; i++ {
				if obstacles[i] {
					possible = false
				}
			}
			if possible {
				res += "1"
			} else {
				res += "0"
			}
		}
	}
	return res

}

/*********************************************************************/
// monotonic
func monotonic(nums []int) []int {
	var res []int
	for i := 1; i < len(nums)-1; i++ {
		if (nums[i-1] < nums[i] && nums[i] < nums[i+1]) || (nums[i-1] > nums[i] && nums[i] > nums[i+1]) {
			res = append(res, 1)
		} else {
			res = append(res, 0)
		}
	}

	return res
}

/*********************************************************************/
// Binary Search
// bus schedule
func lastDepart(schedule []string, time string) int {
	// convert time from string to int
	target := convert(time)
	var timeSeq []int
	for _, t := range schedule {
		timeSeq = append(timeSeq, convert(t))
	}

	// Binary Search: find the closet time (smaller) to target

}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid
		} else {
			right = mid - 1
		}
	}
}

func convert(time string) int {
	time2 := strings.Split(time, ":")
	hr, _ := strconv.Atoi(time2[0])
	min, _ := strconv.Atoi(time2[1])

	return hr*60 + min
}
