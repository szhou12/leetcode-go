package leetcode

/*
PrefixProduct[] and SuffixProduct[]
*/
func productExceptSelf(nums []int) []int {
    n := len(nums)
    prefixProduct := make([]int, n)
    prefixProduct[0] = nums[0]
    for i := 1; i < n; i++ {
        prefixProduct[i] = nums[i] * prefixProduct[i-1]
    }
    suffixProduct := make([]int, n)
    suffixProduct[n-1] = nums[n-1]
    for i := n-2; i>=0; i-- {
        suffixProduct[i] = suffixProduct[i+1] *nums[i]
    }

    res := make([]int, n)
    for i:= 0; i < n; i++ {
        cur := 1
        if i-1 >= 0 {
            cur *= prefixProduct[i-1]
        }
        if i+1 < n {
            cur *= suffixProduct[i+1]
        }
        res[i] = cur
    }

    return res
}

/*
Optimized PrefixProduct[] and SuffixProduct[]
Directly calculate the prefix product and suffix product on res[]
NOTE: res[i] definition (only up to i-1 and i+1) is a bit different from prefixProduct[i] and suffixProduct[i] (up to i)
*/
func productExceptSelf_opt(nums []int) []int {
	n := len(nums)

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}

	// prefix product [0...-1] for res[i]
	cur := 1
	for i := 0; i < n; i++ {
		res[i] *= cur
		cur *= nums[i]
	}

	// suffix product [i+1...n-1] for res[i]
	cur = 1
	for i := n-1; i >= 0; i-- {
		res[i] *= cur
		cur *= nums[i]
	}

	return res

}