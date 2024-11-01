package amz

import "container/heap"

/*
Same as Leetcode 0621. Task Scheduler
Given a string "requests" of length n where each character represents a request from a certain region waiting to be processed.
Given an integer "gap".
Every two requests from the same region cannot be scheduled within the same gap units of time.
How would you re-arrange the requests properly in order to find the minimum time to finish processing all requests?

Example:
requests = "AAABBB", gap = 2
arrangement: "AB_AB_AB" (where "_" represents idle time)
Output: 8
*/

/*
Key Observations:
1. smallest interval unit = gap+1. Within a smallest interval unit, there must be different requests + idles if not enough requests diversity
2. Min time to finish all requests == make each interval unit as compact as possible == use idles as least as possible == fill up with each interval unit with diverse requests as many as possible
3. How to fill up compactly? Prioritize arranging requests with high frequency. This way, we keep the diversity of requests as long as possible and delay the time to use idles. If otherwise, we arrange requests with low frequency first, we will lose diversity very quickly and have to use idles very soon.

Solution: Greedy + MaxHeap + Frequency Count
Time Complexity: O(nlogn)
*/
func leastInterval(requests string, gap int) int {
	// smallest interval unit = gap+1
	// within a smallest interval unit, we need to arrange different requests. If not enough requests, we supplement with idle
	unit := gap + 1

	// count the frequency of each request. HashMap = {req : count}
	countMap := make(map[rune]int)
	for _, req := range requests {
		countMap[req]++
	}

	maxHeap := &PQReqScheduler{}
	heap.Init(maxHeap)
	for req, freq := range countMap {
		heap.Push(maxHeap, RequestPair{req: req, freq: freq})
	}

	res := 0

	for maxHeap.Len() > 0 {
		k := min(unit, maxHeap.Len()) // maxHeap may not have "unit" many of different requests. If not enough, we supplement with idle
		queue := make([]RequestPair, 0) // temporarily store RequestPair whose freq has not been used up

		for i := 0; i < k; i++ {
			cur := heap.Pop(maxHeap).(RequestPair)
			cur.freq--
			if cur.freq > 0 {
				queue = append(queue, cur)
			}
		}

		// update res
		if len(queue) > 0 {
			res += unit // supplement with idel
		} else { // len(queue) == 0 means last interval, no supplement with idle
			res += k
		}

		// throw elements in queue back to maxHeap
		for _, reqPair := range queue {
			heap.Push(maxHeap, reqPair)
		}
	}

	return res
}

type RequestPair struct {
	req  rune
	freq int
}

type PQReqScheduler []RequestPair

func (pq PQReqScheduler) Len() int {
	return len(pq)
}

func (pq PQReqScheduler) Less(i, j int) bool {
	return pq[i].freq > pq[j].freq
}

func (pq PQReqScheduler) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQReqScheduler) Push(x interface{}) {
	*pq = append(*pq, x.(RequestPair))
}

func (pq *PQReqScheduler) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}
