package leetcode

func groupStrings(words []string) []int {
	uf := UnionFind{}
	uf.Init()

	n := len(words)
	// Step 0: every word is set itself
	for i := 0; i < n; i++ {
		uf.father[i] = i
	}

	encode := make([]int, n)    // encode[i] = word[i] encoding form
	family := make(map[int]int) // {encode: word index that represents this encode family}

	// Step 1: words with same encoding belong to same family
	for i, word := range words {
		curEncode := 0

		for j := 0; j < len(word); j++ {
			ch := word[j] - 'a'
			curEncode = curEncode | (1 << ch)
		}

		encode[i] = curEncode

		// if current word's encode has no family yet, create one
		if _, ok := family[curEncode]; !ok {
			family[curEncode] = i
		} else { // if family exists, union family representative (index) with cur index
			repr := family[curEncode]
			if uf.Find(i) != uf.Find(repr) {
				uf.Union(i, repr)
			}
		}
	}

	// Step 2: connect words that belong to same family by operations
	// Only need to check by deleting every char in each word
	for i := 0; i < n; i++ {
		curEncode := encode[i]

		// delete a encoded letter
		for l := 0; l < 26; l++ {
			// skip if has no such letter
			if curEncode&(1<<l) == 0 {
				continue
			}

			delEncode := curEncode - (1 << l)

			// if encode with deleted letter has no family yet, create one
			if _, ok := family[delEncode]; !ok {
				family[delEncode] = i
			} else { // if family exists, union family representative (index) with cur index
				repr := family[delEncode]
				if uf.Find(i) != uf.Find(repr) {
					uf.Union(i, repr)
				}
			}
		}
	}

	// Step 3: count # of connected families
	group := make(map[int]int) // {word index : # of members}
	for i := 0; i < n; i++ {
		group[uf.Find(i)]++
	}
	maxGroupSize := 0
	for _, v := range group {
		maxGroupSize = max(maxGroupSize, v)
	}

	return []int{len(group), maxGroupSize}

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
