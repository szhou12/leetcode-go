# Handbook

## 常用性质
### Subarray
1. subarray总个数: 一个长度为n的array，subarray的总个数 = `n*(n-1)/2 + n` = `n*(n+1)/2` = $O(n^2)$
    * `n*(n-1)/2`: n个位置可以选择作为subarray的起始位置，剩下n-1个位置可以选择作为subarray的终止位置。起始位置和终止位置可以互换，n*(n-1)对于同一个subarray算了两遍，故除以2。
    * `n`: 每个元素自己单独作为subarray。
    * 另一种计算方式: n个长度=1的subarray, n-1个长度=2的subarray, n-2个长度=3的subarray, ..., 1个长度=n的subarray。所以总和为 n + (n-1) + (n-2) + ... + 1 = n*(n+1)/2
### Bitwise Operations
1. bitwise AND: `X1 & X2 & X3 & ... & Xn <= min{X1, X2, X3, ..., Xn}`
    * 一串元素连续 bitwise AND, 所得值一定越变越小，不会大于其中的最小元素。
2. bitwise XOR: `X (^k) (^k) ... (^k) = X` if `(^k)` even number of times; `X (^k) (^k) ... (^k) != X ^ k` if `(^k)` odd number of times
    * 对一个数, (XOR k)偶数次，值保持不变；(XOR k)奇数次，值发生变化。

## Algorithms
| 算法 | 类型题 |
| :-: | - |
|[3-Pass](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/3-Pass.md)|求总数/Count类型题 <br>求Max or Min/优化类型题<br>|
|[BFS & Dijkstra](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/BFS_Dijkstra.md)|基础题: 层级遍历<br>常规题<br>综合题<br>岛屿沉没类/连通图个数<br>Dijkstra<br>BFS + PQ<br>Topological Sort<br>Floyd-Warshall<br>|
|[Binary Search](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/BinarySearch.md)|经典题<br>Lower Bound & Upper Bound<br>First/Last Occurrence<br>Closet Element Greater/Less<br>BS猜答案<br>非常规<br>|
|[Bit Manipulation](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/BitManipulation.md)|n 的二进制表示中1的个数<br>Logic Operation: AND, OR, XOR<br>Digital Counting<br>Bit Array As State<br>|
|[DFS](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/DFS.md)|DFS暴力穷举<br>N-Queens<br>岛屿沉没<br>类DP<br>Permutation<br>All Subsets<br>图的遍历<br>Min-Max Strategy<br>DFS + Memoization<br>|
|[Dynamic Programming](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/DP.md)|基础题<br>买卖股票<br>抢劫房子<br>背包问题<br>回文串<br>Subarray<br>Subsequnce<br>二维/高维DP<br>|
|[Greedy](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/Greedy.md)|Scheduling<br>找规律<br>分情况讨论<br>|
|[KMP](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/KMP.md)|字符串中寻找子串<br>|
|[Linked List](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/LinkedList.md)|Dummy Head<br>递归<br>双指针同向而行，相隔k步，pace一样<br>双指针同向而行, 快慢指针<br>综合题<br>不等长两条链<br>|
|[Prefix Sum & Diff Array](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/PrefixSum_DiffArray.md)|一维差分<br>二维差分<br>扫描线<br>整体区间的增减<br>|
|[Recursion & Tree](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/Recursion_Tree.md)|Traversal Order 构造 Binary Tree<br> 构造 BST<br> 找到所有 Unique BST <br>Complete Binary Tree<br> 找Tree的所有路径<br>Post-Order Traversal<br>In-Order Traversal<br>BST<br>LCA<br>Find Paths in Tree<br>|
|[Trie](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/Trie.md)|经典题<br>Design题<br>|
|[Stack](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/Stack.md)|经典题<br>单调栈模板<br>|
|[Two Pointers](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/TwoPointers.md)|X-Sum<br>Sliding Window<br>|
|[Union Find](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/UnionFind.md)|基础题<br>MST<br>|
|[Math](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/Math.md)|String上模拟加减运算<br>中位数定理<br>Constructive<br>Degree of Freedom<br>|
|[找规律](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/%E6%89%BE%E8%A7%84%E5%BE%8B%E9%A2%98.md)|螺旋格子<br>DP/单调栈 + 中心开花<br>翻转String<br>Rotate/Circular Array/String<br>Compress String<br>打擂台<br>Overlapped Intervals 的充分必要条件<br>|
|[常用技巧](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/%E5%B8%B8%E7%94%A8%E6%8A%80%E5%B7%A7.md)|Graph: 2D坐标转换为1D节点<br>Bit: 用位操作遍历所有子集<br>|
|[Segment Tree](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/SegmentTree.md)||
|[String](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/String.md)||
|[Sorting](https://github.com/szhou12/leetcode-go/blob/main/%E6%80%9D%E8%B7%AF%E6%80%BB%E7%BB%93/Sorting.md)||

## 按题目特征分类
| 特征 | 题号 |
| :-: | - |
|Run any number of times |3012|
|Grid / Matrix |130, 200, 407, 417, 499, 505, 778, 1020, 1368, 1632, 2258, 2290, 2392, 2503, 2577, 2812, 3148, 3195, 3197, 3202|
|2-D坐标 |2101|
|K Sorted Lists |632|

* K Sorted Lists 切入点: 每个list**最小的元素**先放一块排个序看看有没有什么规律？跟其他元素的关系是什么？以它们为subgroup中如何找到符合题意的答案？可否往更大的元素推演转换？

## Practice First
### :red_circle: Hard
* [76. Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/description/)
* [297. Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/)
* [42. Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/description/)
* [295. Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/description/)
* [127. Word Ladder](https://leetcode.com/problems/word-ladder/description/)
* [224. Basic Calculator](https://leetcode.com/problems/basic-calculator/description/)
* [1235. Maximum Profit in Job Scheduling](https://leetcode.com/problems/maximum-profit-in-job-scheduling/description/)
    * Greedy (Sort by end date) + DP + Binary Search (UpperBound)
    * 不使用Greedy解，但是沿用sort的思路
* [23. Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/description/)
* [84. Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/description/)
* [124. Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/description/)
* [895. Maximum Frequency Stack](https://leetcode.com/problems/maximum-frequency-stack/description/)
* [4. Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/description/)
* [329. Longest Increasing Path in a Matrix](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/description/)
* [32. Longest Valid Parentheses](https://leetcode.com/problems/longest-valid-parentheses/description/)
* :lock: [588. Design In-Memory File System](https://leetcode.ca/all/588.html)
* :lock: [759. Employee Free Time](https://leetcode.ca/all/759.html)
* [212. Word Search II](https://leetcode.com/problems/word-search-ii/description/)
* :lock: [269. Alien Dictionary](https://leetcode.ca/all/269.html)
* [815. Bus Routes](https://leetcode.com/problems/bus-routes/description/)
* [239. Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/description/)
* [336. Palindrome Pairs](https://leetcode.com/problems/palindrome-pairs/description/)
* [25. Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/description/)
* [37. Sudoku Solver](https://leetcode.com/problems/sudoku-solver/description/)
    * DFS
* [41. First Missing Positive](https://leetcode.com/problems/first-missing-positive/description/)
    * 找规律
* [51. N-Queens](https://leetcode.com/problems/n-queens/description/)
    * DFS
* [632. Smallest Range Covering Elements from K Lists](https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/description/)
    *  K sorted lists + Greedy + minHeap


## Emojis

| ico | meaning |
| :-: | - |
| :bell: | 优先反复练习, 解题思路代表一类思想! |
| :purple_circle: | 极难, 完全没有思路, 即使参考了答案, 之后再做也很可能做不出来 |
| :red_circle: | 困难, 完全没有思路, 参考了答案 |
| :orange_circle: | 中等偏难, 有正确的思路但是不完整, 缺少解题的关键, 无法直接出答案, 需要参考版本答案 |
| :yellow_circle: | 中等, 有思路但是实现上有问题, 最终实现版本参考了答案 |
| :green_circle: | 简单, 可直接出bug-free code |
| :lock: | 加锁题, 尚未提交过leetcode |
| :secret: | 题解使用了值得记住的技巧 |
| :star: | 类型题的解题思路总结 |
| :arrow_right: | 一般用于总结思路时候逻辑的递进 |
| :bulb: | 解题思路总结中，idea层面的总结 |
| :mag: | 解题思路总结中，tricks/小技巧/应用层面的总结 |
| :rotating_light: | 注意！注意！注意！重要的事情说三遍！！！ |

## Readme content structure

每道题的Readme结构统一如下:

* Solution idea
    * [solution type]
        * 要点总结
        * 物理意义
        * 代码结构
        * Time complexity & Space complexity
* Resource

## Resource
* [emoji-cheat-sheet](https://github.com/ikatyang/emoji-cheat-sheet/blob/master/README.md#table-of-contents)
* [How to Write a Good README File for Your GitHub Project](https://www.freecodecamp.org/news/how-to-write-a-good-readme-file/)