# Bit Manipulation

## n 的二进制表示中1的个数
* `n &= (n-1)`
* [1356. Sort Integers by The Number of 1 Bits](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1356-Sort-Integers-by-The-Number-of-1-Bits)

* :red_circle: :secret: 最长"好"子数组: [2401. Longest Nice Subarray](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2401-Longest-Nice-Subarray)
    * Sliding Window (Flex) + Bitwise AND

* :red_circle: :secret: 可以得到bitwise OR最大值的最短子数组长度: [2411. Smallest Subarrays With Maximum Bitwise OR](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2411-Smallest-Subarrays-With-Maximum-Bitwise-OR)
    * Sliding Window (Flex) + Bitwise OR
    * Trick 1: 左边界从右往左移动。固定左边界，右边界尽量缩短。
    * Trick 2: 判断一个整数 第k bit位 是否有1: `num>>k & 1 == 1`

## Logic Operation: AND, OR, XOR
* :red_circle: 用AND和OR求最大平方和: [2897. Apply Operations on Array to Maximize Sum of Squares](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2897-Apply-Operations-on-Array-to-Maximize-Sum-of-Squares)
    * AND, OR对数值大小的影响

* :red_circle: 在AND结果为最小的情况下尽可能切多个subarrays: [2871. Split Array Into Maximum Number of Subarrays](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2871-Split-Array-Into-Maximum-Number-of-Subarrays)
    * AND的性质: AND越多元素，结果的值越小

