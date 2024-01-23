package leetcode

func findMaximumNumber(k int64, x int) int64 {
	left, right := 1, int(1e15)

	for left < right {
		mid := right - (right-left)/2
		if !validPrice(mid, k, x) {
			right = mid - 1
		} else {
			left = mid
		}
	}

	return int64(left)
}

func validPrice(num int, k int64, x int) bool {
	// convert num to binary representation
	arr := make([]int, 0)

	numBinary := num
	for numBinary > 0 {
		arr = append(arr, numBinary%2)
		numBinary /= 2
	}

	// 注意: arr中的binary representation是反过来的
	// e.g. 10 = '1010'        (右 -> 左 = 低位 -> 高位)
	//      arr = [0, 1, 0, 1] (左 -> 右 = 低位 -> 高位)
	// 这样做跟接下来的大loop i 也有关系: i从最低位符合i%x==0的地方开始，每次跳x个位

	count := 0
	// i starts from 0-index by x-1; i jumps by every x
	for i := x - 1; (1 << i) <= num; i += x {
		if arr[i] == 0 {
			// L   L   L   i   H    H    H
			//   2^m       0   000:(HHH-1)
			highbits := 0
			for j := len(arr) - 1; j > i; j-- { // 从最高位往低处走
				highbits = highbits*2 + arr[j]
			}
			lowbits := pow(2, i)
			count += highbits * lowbits
		} else {
			// Case 1
			// L   L   L   i   H    H    H
			//   2^m       1   000:(HHH-1)
			highbits := 0
			for j := len(arr) - 1; j > i; j-- {
				highbits = highbits*2 + arr[j]
			}
			lowbits := pow(2, i)
			count += highbits * lowbits

			// Case 2
			// L   L   L   i   H    H    H
			// 000:LLL     1      HHH (fixed)
			lowbits = 0
			for j := i - 1; j >= 0; j-- { // 000:(LLL-1)
				lowbits = lowbits*2 + arr[j]
			}
			lowbits += 1
			count += lowbits
		}

		if int64(count) > k {
			return false
		}
	}
	return true
}

func pow(x, n int) int {
	// base case
	if n == 0 {
		return 1
	}

	// recursion
	if n%2 == 0 {
		return pow(x, n/2) * pow(x, n/2)
	} else {
		return x * pow(x, n/2) * pow(x, n/2)
	}
}
