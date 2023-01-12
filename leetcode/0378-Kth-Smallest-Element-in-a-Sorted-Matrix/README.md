# [378. Kth Smallest Element in a Sorted Matrix](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/description/)

## Solution idea

### Binary Search (Optimal Solution)

**Binary Search 猜答案**

* 确定搜索空间：根据题目给出的矩阵性质，每一行每一列都是 non-decreasing，可以确定，min value一定是左上角 `matrix[0][0]` (作为BS的下界/左边界), max value一定是右下角 `matrix[m-1][n-1]` (作为BS的上界/右边界)。在这个解空间里面二分搜索所有值，找到第 K 小的元素。
* 确定收敛条件：每次得到一个 `mid`, 找出 `matrix` 中比 `mid` 小于等于的元素个数 `count`
    1. `if count < k`: 下界/左边界太小了，不可能是解，收敛下界/左边界 （不包括下界/左边界）
    2. `if count >= k`: 上界/右边界有可能太大了，有可能是解，收敛上界/右边界 （包括上界/右边界）
* 注意！因为题目中说了，一定会存在第 K 小元素，所以二分搜索到一个元素的时候，一定会得出结果
* 如何找出 `matrix` 中比 `mid` 小的元素个数 `count`？
    * 这时就要利用`matrix`中元素排列的性质：从左下角出发往右上角找(向右走或向上走) 或者 从右上角出发往左下角找(向下走或向左走)
    * 假设选择：从左下角出发往右上角找
    1. `if matrix[i][j] <= mid`: `(i, j)`位置上的元素符合条件，并且根据矩阵性质，`(i, j)`上面的所有元素也都符合条件，所以都算入符合条件的个数中 (`i+1`)；向右走, 以达到增大下一个元素的值的目的
    2. `if matrix[i][j] > mid`: `(i, j)`位置上的元素不符合条件，向上走，以达到减小下一个元素的值的目的

Time complexity = $O((m+n)*\log D)$ where $m = $ number of rows, $n = $ number of cols, $D = $ difference between the max value in the matrix and the min value in the matrix.

### Priority Queue - Max Heap

 容易想到的解法是Priority Queue:
 * 维护一个max heap
 * 遍历每一个元素：如果max heap没有满员 或者 当前元素小于栈顶元素，则入队
 * 最后max heap的栈顶元素就是第k小的元素

Time complexity = $O(m*n*\log k)$

## Resources
[【每日一题】LeetCode 378. Kth Smallest Element in a Sorted Matrix, 10/23/2021](https://www.youtube.com/watch?v=JJUv4DDLSB4&ab_channel=HuifengGuan)

[halfrost/LeetCode-Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0378.Kth-Smallest-Element-in-a-Sorted-Matrix)

[Leetcode Solution discussion](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/solutions/1322101/c-java-python-maxheap-minheap-binary-search-picture-explain-clean-concise/?orderBy=most_votes)

## 姐妹题
* Leetcode 373
* Leetcode 240