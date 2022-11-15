# [1365. How Many Numbers Are Smaller Than the Current Number](https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/description/)

## Solution idea

### Brute Force

Time complexity = $O(n^2)$

### Optimal: Sort + HashMap

* 首先要找小于当前数字的数字，那么从小到大排序之后，该数字之前的数字就都是比它小的了。
* 所以可以定义一个新数组，将数组排个序。
* **排序之后，其实每一个数值的下标就代表这前面有几个比它小的了。**
* 用一个Hash Map（本题可以就用一个数组）来做数值和下标的映射。这样就可以通过数值快速知道下标（也就是前面有几个比它小的）。
* 数值相同怎么办？
    * 技巧: 在构造数组hash的时候，从后向前遍历，这样hash里存放的就是相同元素最左面的数值和下标了

Time complexity = $O(n\log n)$

## Resource
[代码随想录-1365.有多少小于当前数字的数字](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/1365.%E6%9C%89%E5%A4%9A%E5%B0%91%E5%B0%8F%E4%BA%8E%E5%BD%93%E5%89%8D%E6%95%B0%E5%AD%97%E7%9A%84%E6%95%B0%E5%AD%97.md)
