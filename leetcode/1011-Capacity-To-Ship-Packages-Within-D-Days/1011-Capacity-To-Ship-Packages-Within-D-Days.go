package main

func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0

	for _, weight := range weights {
		left = max(left, weight)
		right += weight
	}

	// binary search
	for left < right {
		mid := left + (right-left)/2
        if isOk(weights, days, mid) {
            right = mid
        } else {
            left = mid + 1
        }
	}

    return left
}

func isOk(weights []int, days int, cap int) bool {
    needDays := 1 // start with 1 day

    curLoads := 0
    for _, w := range weights {
        if curLoads + w > cap {
            curLoads = w
            needDays++
            continue
        }
        curLoads += w
    }

    return needDays <= days
}

// Two Pointers implementation
func isOk_twoPointers(weights []int, days int, cap int) bool {
    needDays := 1
    n := len(weights)
    l := 0
    
    for l < n {
        loads := 0
        r := l
        
        // Load as many packages as possible for current day
        for r < n && loads + weights[r] <= cap {
            loads += weights[r]
            r++
        }
        
        // Move to next day
        l = r
        if l < n {  // Only increment days if there are more packages
            needDays++
        }
    }
    
    return needDays <= days
}
