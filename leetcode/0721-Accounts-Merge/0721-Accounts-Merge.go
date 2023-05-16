package leetcode

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)
	uf := UnionFindStr{}
	uf.Init()

}

type UnionFindStr struct {
	father map[string]string
}

func (ufs *UnionFindStr) Init() {
	ufs.father = make(map[string]string)
}

func (ufs *UnionFindStr) Union(x string, y string) {
	fx := ufs.father[x]
	fy := ufs.father[y]
	if fx < fy {
		ufs.father[fy] = fx
	} else {
		ufs.father[fx] = fy
	}
}

func (ufs *UnionFindStr) Find(x string) string {
	if x != ufs.father[x] {
		ufs.father[x] = ufs.Find(ufs.father[x])
	}
	return ufs.father[x]
}

// My Solution
// {account index : {email : true}} --> UnionFind(account index) --> {root index : {email: true}} --> [root name: [emails]]
// 需要两两account比较，不是最优的解法
func accountsMerge_mysoln(accounts [][]string) [][]string {
	n := len(accounts)
	uf := UnionFind{}
	uf.Init()
	for i := 0; i < n; i++ {
		uf.father[i] = i
	}

	emails := make(map[int]map[string]bool) // {i : {email : true}}
	for i, account := range accounts {
		emails[i] = make(map[string]bool)
		for j := 1; j < len(account); j++ {
			emails[i][account[j]] = true
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := 1; k < len(accounts[j]); k++ {
				if emails[i][accounts[j][k]] {
					if uf.Find(i) != uf.Find(j) {
						uf.Union(i, j)
					}
				}
			}
		}
	}

	family := make(map[int]map[string]bool) // {root index : {email : true} }
	for i := 0; i < n; i++ {
		root := uf.Find(i)
		for _, email := range accounts[i][1:] {
			if _, ok := family[root]; !ok {
				family[root] = make(map[string]bool)
			}
			family[root][email] = true
		}
	}

	res := make([][]string, 0)
	for root, emailMap := range family {
		name := accounts[root][0]
		temp := make([]string, 0)
		for k, _ := range emailMap {
			temp = append(temp, k)
		}
		sort.Strings(temp)
		temp = append([]string{name}, temp...)
		res = append(res, temp)
	}

	return res
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
