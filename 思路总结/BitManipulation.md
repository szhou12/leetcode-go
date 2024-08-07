# Bit Manipulation

## 目录
* [n 的二进制表示中1的个数](#n-的二进制表示中1的个数)
* [Logic Operation: AND, OR, XOR](#logic-operation-and-or-xor)
* [Digital Counting](#digital-counting)
* [Bit Array As State](#bit-array-as-state)

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
| | AND | XOR | OR |
|-| - | - | - |
|性质1|一串元素(非负数)连续 bitwise AND，所得值一定越变越小，不会大于其中的最小元素。|一串元素(非负数)连续 bitwise OR，所得值一定越变越大，不会小于其中的最大元素。| 对一个数, (XOR k)偶数次，值保持不变；(XOR k)奇数次，值发生变化。 |
|性质1拓展|bitwise AND 元素(非负数)越多，结果只会减小不会增加。因为越多bit位的1被0抹消。(单调性)|bitwise OR 元素(非负数)越多，结果只会增大不会减小。因为越多bit位的1被保留。(单调性)||
|性质2|`a AND math.MaxInt = a` where `a` non-negative number and `math.MaxInt = 01111...1111`|`a OR 0 = a` where `a` can be any number||

* :red_circle: 多少子数组的bitwise AND结果是k: [3209. Number of Subarrays With AND Value of K](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3209-Number-of-Subarrays-With-AND-Value-of-K)
    * AND性质: 连续 bitwise AND 所得值一定越变越小。
    * 常数个key (31种)

* :red_circle: 最小递增数列的最后一个元素值是多少: [3133. Minimum Array End](https://github.com/szhou12/leetcode-go/blob/main/leetcode/3133-Minimum-Array-End/README.md)
    * AND性质: 连续 bitwise AND 所得值一定越变越小，不会大于其中的最小元素。

* :red_circle: 树上节点的最大XOR和: [3068. Find the Maximum Sum of Node Values](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3068-Find-the-Maximum-Sum-of-Node-Values)
    * XOR性质: 对一个数, (XOR k)偶数次，值保持不变；(XOR k)奇数次，值发生变化

* :red_circle: 用AND和OR求最大平方和: [2897. Apply Operations on Array to Maximize Sum of Squares](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2897-Apply-Operations-on-Array-to-Maximize-Sum-of-Squares)
    * AND, OR对数值大小的影响

* :red_circle: 在AND结果为最小的情况下尽可能切多个subarrays: [2871. Split Array Into Maximum Number of Subarrays](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2871-Split-Array-Into-Maximum-Number-of-Subarrays)
    * AND性质: AND越多元素，结果的值越小

## Digital Counting

* :red_circle: 第k小的组合数: [3116. Kth Smallest Amount With Single Denomination Combination](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3116-Kth-Smallest-Amount-With-Single-Denomination-Combination)
    * Binary Search + Grosper's Hack + Exclusion-Inclusion Principle + Divisors

* :red_circle: 最大满足特定条件的自然数: [3007. Maximum Number That Sum of the Prices Is Less Than or Equal to K](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3007-Maximum-Number-That-Sum-of-the-Prices-Is-Less-Than-or-Equal-to-K)
    * Binary Search 猜答案 + Digital Counting

## Bit Array As State
- 用bit array表示n个元素使用与不使用的状态。

* :lock: :red_circle: 最大化数组中连续元素的总和: [2992. Number of Self-Divisible Permutations](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2992-Number-of-Self-Divisible-Permutations)
    * DP + Bit Manipulation
    * `DP[i][state]`