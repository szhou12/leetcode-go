package leetcode

func minimumCost(start []int, target []int, specialRoads [][]int) int {

}

func encode(x, y int) int {
	return (x << 32) + y
}

func decode(n int) (int, int) {
	return n >> 32, n % (1 << 32)
}

type PQ [][]int // [[shortest path, node]]

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}
