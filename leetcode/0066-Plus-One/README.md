# [66. Plus One](https://leetcode.com/problems/plus-one/)

## Solution ideas

### Recursion

My first idea is to use **Recursion**:

Starting from the last position:

Case 1: if the last digit `< 9`, increment this digit, return array

Case 2.1: Otherwise, if current digit being pointed to is the first digit, need to return `[1, 0]`

Case 2.2: Otherwise, current digit becomes $0$ and do recursion on subarray

Time complexity = $O(n)$

**CAVEAT**: Case 2.1 can take some time to implement correctly.

### Iterative
从数组尾部开始往前扫，逐位进位即可。最高位如果还有进位需要在数组里面第 $0$ 位再插入一个 $1$.

[Resource](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0066.Plus-One)