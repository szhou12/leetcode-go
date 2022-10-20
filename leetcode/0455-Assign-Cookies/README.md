# [455. Assign Cookies](https://leetcode.com/problems/assign-cookies/)

## Solution idea

### My Solution - Greedy Algorithm 基础题
* sort each array in increasing order

* assign a pointer to each array
    1. 满足条件，两个指针都移动一格，count+=1
    2. 不满足，只移动 cookie-size 指针

Time complexity = $O(n\log n)$

## Resource
[代码随想录-455.分发饼干](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0455.%E5%88%86%E5%8F%91%E9%A5%BC%E5%B9%B2.md)