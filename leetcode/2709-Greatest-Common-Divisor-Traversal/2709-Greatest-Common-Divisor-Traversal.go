package leetcode

func canTraverseAllPairs(nums []int) bool {
	maxNum := maxElement(nums)
	primes := eratosthenes(maxNum)
	m := len(primes)
	n := len(nums)
	primeIdx := make(map[int]int)
	for i, p := range primes {
		primeIdx[p] = i
	}

	uf := UnionFind{}
	uf.Init()

	// Note: all primes' index are assigned after n
	for i := 0; i < n+m; i++ {
		uf.father[i] = i
	}

	for i, x := range nums {
		// try whether each prime is factor of x
		for j, p := range primes {
			if p > x {
				break
			}
			// leftover x itself is the prime factor
			if p*p > x {
				if uf.Find(i) != uf.Find(n+primeIdx[x]) {
					uf.Union(i, n+primeIdx[x])
				}
				break
			}
			// if p divides x: union p + factorize p out of x
			if x%p == 0 {
				if uf.Find(i) != uf.Find(n+j) {
					uf.Union(i, n+j)
				}
				for x%p == 0 {
					x /= p
				}
			}
		}
	}

	// check whether all nums are connected
	for i := 1; i < n; i++ {
		if uf.Find(0) != uf.Find(i) {
			return false
		}
	}
	return true
}

func maxElement(nums []int) int {
	res := -1
	for _, val := range nums {
		if val > res {
			res = val
		}
	}
	return res
}

func eratosthenes(n int) []int {
	var primes []int
	sieve := make([]int, n+1)
	for i := 2; i*i <= n; i++ {
		if sieve[i] == 1 {
			continue
		}
		j := i * 2
		for j <= n {
			sieve[j] = 1
			j += i
		}
	}

	for i := 2; i <= n; i++ {
		if sieve[i] == 0 {
			primes = append(primes, i)
		}
	}
	return primes
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
