package leetcode

/*

   [ X X X X X X X X X ]
-1   0       j       i

sum[j:i] = prefixSum[i] - prefixSum[j-1] ([j:i] inclusive)
sum[i] = prefixSum[i] - prefixSum[-1]
                            0

if sum[j:i] divisble by k <=> (prefixSum[i] % k) and (prefixSum[j-1] % k) must have the same remainder

(sum[j:i] % k) == 0 iff r1 == r2 where r1 = (prefixSum[i] % k) and r2 = (prefixSum[j-1] % k)

So we define the hash map where key is the remainder value and value is the count already added.

So as of nums[:i] where the r = (prefixSum[i] % k), the total number of subarrays that are divisible by k is the value of the hash map at key r (ie. how many 0 <= j < i s.t. (prefixSum[j] % k) == r)

*/

func subarraysDivByK(nums []int, k int) int {
	dict := make(map[int]int) // { (prefixSum % k) : count }
	dict[0] = 1 // prefixSum 结合 HashMap 时的通用初始化技巧: key = 0 表示 prefixSum % k = 0, 也就是 prefixSum = 0, 这个0指代index=-1时的prefixSum[-1], value=1表示j=-1算一个count，虚拟地认为index=-1是存在的。

	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		r := (sum % k + k) % k // to avoid sum % k < 0, we add an extra k and then mod k again.
		res += dict[r]
		dict[r]++
	}

	return res
}