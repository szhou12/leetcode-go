package amz

/*
There are n machines on a 2D-grid.
Given two arrays of length n: x[], y[], where (x[i], y[i]) is the coordinate of the i-th machine.
A machine is considered idle if it is surrounded by other machines in its all four directions (i.e., up, down, left, and right). Note that they do not necessarily have to be adjacent.
e.g. The machine at (2, 3) is idle if it has machines at left(1, 3), right(3, 3), up(2, 5), and down(2, 2) respectively.
Find the total number of idle machines.

Constraints:
1 ≤ n ≤ 105
-10^9 ≤ x[i], y[i] ≤ 10^9

Example:
x = [1, 1, 1, 2, 2, 2, 2, 3, 3, 3]
y = [1, 2, 3, 1, 2, 3, 5, 1, 2, 3]
Output: 2 - machines at (2, 2) and (2, 3) are idle

x = [0, 0, 0, 0, 0, 1, 1, 1, 2, -1, -1, -2, -1]
y = [-1, 0, 1, 2, -2, 0, 1, -1, 0, 1, -1, 0, 0]
Output: 5 - machines at (0, 0), (1, 0), (-1, -0), (0, 1), and (0, -1) are idle
*/


/*
Idea: HashMap
Use 2 hashmaps: 
one hashmap stores, for each available row, [min col value, max col value].
another hashmap stores, for each available col, [min row value, max row value].

Then, for each machine, check if it is idle by checking if it is surrounded by other machines in all four directions.

Time complexity := O(n)
*/
func NumNotWorkingMachinese(x, y []int) int {
	n := len(x)

	// for each row, find (y_min, y_max)
	// for each col, find (x_min, x_max)
	rows := make(map[int][]int)
	cols := make(map[int][]int)
	for i := 0; i < n; i++ {
		curX, curY := x[i], y[i]
		if tempY, ok := rows[curX]; ok {
			rows[curX][0] = min(tempY[0], curY)
			rows[curX][1] = max(tempY[1], curY)
		} else {
			rows[curX] = []int{curY, curY}
		}

		if tempX, ok := cols[curY]; ok {
			cols[curY][0] = min(tempX[0], curX)
			cols[curY][1] = max(tempX[1], curX)
		} else {
			cols[curY] = []int{curX, curX}
		}

	}

	isIdle := func(dx, dy int) bool {
		return rows[dx][0] < dy && dy < rows[dx][1] && cols[dy][0] < dx && dx < cols[dy][1]
	}

	res := 0
	// for each machine (dx, dy), check if rows[dx][0] < dy < rows[dx][1] && cols[dy][0] < dx < cols[dy][1]
	for i := 0; i < n; i++ {
		if isIdle(x[i], y[i]) {
			res++
		}
	}

	return res

}
