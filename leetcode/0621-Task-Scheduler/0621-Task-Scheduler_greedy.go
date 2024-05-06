package leetcode

// Solution 2: Greedu
func leastInterval_greedy(tasks []byte, n int) int {
	// map{task: frequency}
	count := make(map[byte]int)
	for _, task := range tasks {
		count[task]++
	}

	// get hightest frequency
	maxFreq := 0
	for _, freq := range count {
		maxFreq = max(maxFreq, freq)
	}

	// 计算“零头” i.e., 最高频次的任务种类
	tail := 0
	for _, freq := range count {
		if freq == maxFreq {
			tail++
		}
	}

	// 两种填充情况取最大值
	return max((maxFreq-1)*(n+1)+tail, len(tasks))
}
