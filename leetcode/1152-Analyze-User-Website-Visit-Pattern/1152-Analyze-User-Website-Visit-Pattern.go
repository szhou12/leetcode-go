package leetcode

import "sort"

type Tuple struct {
	username  string
	timestamp int
	website   string
}

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	n := len(username)

	// step 1: create a slice of Tuple
	tupleArr := make([]Tuple, 0)
	for i := 0; i < n; i++ {
		tupleArr = append(tupleArr, Tuple{username: username[i], timestamp: timestamp[i], website: website[i]})
	}

	// step 2: sort tuples by timestamp
	sort.Slice(tupleArr, func(i, j int) bool {
		return tupleArr[i].timestamp < tupleArr[j].timestamp
	})

	// step 3: map to track websites visited by each user in chronological order
	webs := make(map[string][]string) // {username: [web1, web2, ...]}
	for _, tuple := range tupleArr {
		webs[tuple.username] = append(webs[tuple.username], tuple.website)
	}

	// step 4: map to count occurrences of each 3-length pattern
	patternCount := make(map[[3]string]int) // {[web1, web2, web3]: 1}

	// for each user, put all their 3-length pattern into patternCount
	for _, theirWebs := range webs {
		visitedPattern := make(map[[3]string]bool) // dedup same pattern by same user
		// Note: j must depend on i, k must depend on j. CANT set k=i+2!
		for i := 0; i < len(theirWebs)-2; i++ {
			for j := i + 1; j < len(theirWebs)-1; j++ {
				for k := j + 1; k < len(theirWebs); k++ {
					cur := [3]string{theirWebs[i], theirWebs[j], theirWebs[k]}
					if visitedPattern[cur] {
						continue
					}
					visitedPattern[cur] = true
					patternCount[cur]++
				}
			}
		}
	}

	// step 5: find the lexi smallest pattern with the max count
	var res [3]string
	maxCount := 0
	for pattern, count := range patternCount {
		if count > maxCount || (count == maxCount && isLexiSmaller(pattern[:], res[:])) {
			res = pattern
			maxCount = count
		}
	}

	return res[:]

}

func isLexiSmaller(cur, res []string) bool {
	for i := 0; i < len(cur); i++ {
		if cur[i] < res[i] {
			return true
		}
		if cur[i] > res[i] {
			return false
		}
	}

	return false
}
