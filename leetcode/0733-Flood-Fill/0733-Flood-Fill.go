package leetcode

var dir = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	m, n := len(image), len(image[0])
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}

	srcColor := image[sr][sc]
	dfs(image, sr, sc, color, srcColor, visited)

	return image
}

func dfs(image [][]int, x int, y int, color int, srcColor int, visited [][]int) {
	// base cases
	// if out of bound
	if !inBound(image, x, y) {
		return
	}
	// if already visited
	if visited[x][y] == 1 {
		return
	}
	// if cell(x, y) does not have the same original color as srcColor
	if image[x][y] != srcColor {
		return
	}

	// update: color cell(x, y)
	image[x][y] = color
	visited[x][y] = 1

	// recursiom
	for _, d := range dir {
		dx, dy := x + d[0], y + d[1]
		dfs(image, dx, dy, color, srcColor, visited)
	}

}

func inBound(image [][]int, x, y int) bool {
	m, n := len(image), len(image[0])
	return (0 <= x && x < m) && (0 <= y && y < n)
}