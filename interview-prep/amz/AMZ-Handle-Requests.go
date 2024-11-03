package amz

import "container/heap"

/*
There is a queue that sequentially stores n requests. queue = [r1, r2, r3, ..., rn].
Given an array of length n: wait[] where wait[i] = the max survival time a request can wait in the queue to be processed before it expires.
At each unit of time, the first request in the queue is processed. At the next unit of time, the processed request and expired requests are removed from the queue.
Return an array of length n: res[] where res[i] = the number of requests present in the queue at each unit of time until the queue is empty.

Example:
Input: wait = [2, 2, 3, 1]

- at time=0: 1st request is being processed. The number of requests present in queue is 4 as queue = [1,2,3,4]
- at time=1: 2nd request is being processed. 1st request has been processed and removed. Also, request 4 (wait[3]=1) has expired and been removed. The number of requests present in queue is 2 as queue = [2,3]
- at time=2: 3rd request is being processed. 2nd request has been processed and removed. The number of requests present in queue is 1 as queue = [3]
- at time=3: 3rd request has been processed and removed. The number of requests present in queue is 0 as queue = []

Output: [4, 2, 1, 0]
*/

/*
Solution: HashMap (Frequency Count)

Key Observation:
For a request i, 它只可能被提前(time < i)处理 (因为前面有request expire了), 也就是说, 最晚到time=i时刻它一定会被处理 (i.e.最晚它成为 1st request in queue的时刻)。所以，它能被正常处理的情况发生在 当且仅当 wait[i] > time=i (严格大于。如果等于，就是它恰好在最晚时刻expire了)

Demo:
wait = [3, 1, 2, 1]
waitTimeFreq = {1: 2, 2: 1, 3: 1}
at t=0: numReqs=4.

	no expired requests, 4 requests in queue.
	res=[4].
	1st request (wait[0]=3) is being processed and will be reduced at t=1 (numReqs=4-1=3).

at t=1: numReqs=3.

	waitTimeFreq[1]=2 expired requests, 3-2=1 request in queue.
	res=[4,1].
	1st request (wait[2]=2) is being processed and will be reduced at t=2 (numReqs=1-1=0).

at t=2: numReqs=0.

	waitTimeFreq[2]=1 expired requests but time=2 > 0 request left.
	res=[4,1,0].
	queue empty (numReqs=0-1=-1).

at t=3: numReqs=-1 break loop
*/
func getRequestsInQueue(wait []int) []int {

	waitTimeFreq := make(map[int]int)
	for _, waitTime := range wait {
		waitTimeFreq[waitTime]++
	}

	res := make([]int, 0)
	time := 0
	numReqs := len(wait)
	for numReqs >= 0 {
		if cnt, ok := waitTimeFreq[time]; ok && time < numReqs {
			numReqs -= cnt
		}
		res = append(res, numReqs)

		// Prepare for next time
		time++
		numReqs-- // 1st request is being processed at cur time and will be done at next time
	}
	return res

}

func getRequestsInQueue_v2(wait []int) []int {
	minHeap := &PQ_req{}
	heap.Init(minHeap)

    process := make([][]int, 0)
    for i, waitTime := range wait {
        if waitTime > i {
            process = append(process, []int{waitTime, i})
        } else {
            heap.Push(minHeap, []int{waitTime, i})
        }
    }

    time := 0
    p1 := 0
    m := len(process)
    res := make([]int, 0)
    for p1 < m || minHeap.Len() > 0 {
        // at any time t, expire requests first
        for minHeap.Len() > 0 && (*minHeap)[0][0] <= time {
            heap.Pop(minHeap)
        }

        res = append(res, (m-1-p1+1)+minHeap.Len())

        var y []int
        if minHeap.Len() > 0 {
            y = heap.Pop(minHeap).([]int)
        }

        var x []int
        if p1 < m {
            x = process[p1]
        }

        if len(x) > 0 && len(y) > 0 && x[1] < y[1] {
            p1++
            heap.Push(minHeap, y)
        }
        
        time++

    }
    return res

}

type PQ_req [][]int // [[wait[i], i], ...]

func (pq PQ_req) Len() int {
	return len(pq)
}

func (pq PQ_req) Less(i, j int) bool {
	if pq[i][0] == pq[j][0] {
		return pq[i][1] < pq[j][1]
	}
	return pq[i][0] < pq[j][0]
}

func (pq PQ_req) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ_req) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PQ_req) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}


