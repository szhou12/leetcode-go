# [463. Island Perimeter](https://leetcode.com/problems/island-perimeter/)

## Solution idea

### My Solution

* 这一题乍一看好像要用 DFS/BFS, 但其实不用, 而且思路很直观, 主要是 **分类讨论**
* 要素察觉：计算边界个数时，遇到0和出界这两种情况是等同的，都是增加一条边界
* 思路：遍历每个cell, 遇到1, 就看他的上下左右:
    1. Case 1: 如果四边都是1，也就是没有0/出界，这个cell不增加边界, count += 0
    2. Case 2: 如果有任意三边是1, 也就是一个0/出界, 这个cell增加一条边界, count += 1
    3. Case 3: 如果有任意两边是1, 也就是两个0/出界, 这个cell增加二条边界, count += 2
    4. Case 4: 如果有任意一边是1, 也就是三个0/出界, 这个cell增加三条边界, count += 3

Time complexity = $(m*n)$