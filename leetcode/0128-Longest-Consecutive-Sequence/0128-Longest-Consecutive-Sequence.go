package leetcode

func longestConsecutive(nums []int) int {
	uf := UnionFind{}
	uf.Init()
	for _, num := range nums {
		uf.father[num] = num
	}

	for _, num := range nums {
		// if num - 1, num + 1 exist in uf.father, union
		if _, ok := uf.father[num-1]; ok {
			if uf.Find(num-1) != uf.Find(num) {
				uf.Union(num-1, num)
			}
			if uf.Find(num+1) != uf.Find(num) {
				uf.Union(num+1, num)
			}
		}
	}

	// count the number of elements in each component
	family := make(map[int]map[int]bool) // NOTE: value must use a set instead of []int to avoid duplicate family members
	for _, num := range nums {
		root := uf.Find(num)
		if _, ok := family[root]; !ok {
			family[root] = make(map[int]bool)
		}
		family[root][num] = true
	}

	// find the largest component
	res := 0
	for _, members := range family {
		res = max(res, len(members))
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type UnionFind struct {
	father map[int]int
}

func (uf *UnionFind) Init() {
	uf.father = make(map[int]int)
}

func (uf *UnionFind) Union(x int, y int) {
	fx := uf.father[x]
	fy := uf.father[y]
	if fx < fy {
		uf.father[fy] = fx
	} else {
		uf.father[fx] = fy
	}
}

func (uf *UnionFind) Find(x int) int {
	if x != uf.father[x] {
		uf.father[x] = uf.Find(uf.father[x])
	}
	return uf.father[x]
}
