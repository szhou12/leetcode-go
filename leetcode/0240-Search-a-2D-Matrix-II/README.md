# [240. Search a 2D Matrix II](https://leetcode.com/problems/search-a-2d-matrix-ii/description/)

## Solution idea

### 小巧思：从 左下角/右上角 出发做 matrix traversal

这道题说 matrix 从上到下递增，从左到右递增，**显然左上角是最小元素，右下角是最大元素**。我们如果想高效在 matrix 中搜索一个元素，肯定需要从某个角开始，比如说从左上角开始，然后每次只能向右或向下移动，不要走回头路。

如果真从左上角开始的话，就会发现无论向右还是向下走，元素大小都会增加，那么到底向右还是向下？不确定，那只好用类似 动态规划算法 的思路穷举了。

但实际上不用这么麻烦，我们不要从左上角/右下角开始，而是从右上角 (top-right) /左下角 (bottom-left) 开始，规定只能向左或向下移动。

**Why This Works?**

注意，如果向左移动，元素在减小，如果向下移动，元素在增大，这样的话我们就可以根据当前位置的元素和 target 的相对大小来判断应该往哪移动，不断接近直至找到target 或者出界。

这个方法可以**用到所有基于这种sorted matrix的题目**，比如 [373. Find K Pairs with Smallest Sums](https://leetcode.com/problems/find-k-pairs-with-smallest-sums/description/) 和 [378. Kth Smallest Element in a Sorted Matrix](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/description/)，就是配合二分法可以得到更高效的解。

Time complexity = $O(m+n)$ where $m = $ number of rows, $n = $ number of cols.