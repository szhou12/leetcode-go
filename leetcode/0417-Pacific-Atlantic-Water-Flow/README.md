# [417. Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/)

## Solution idea

### DFS

**Key 1**: 用DFS暴力搜索每个cell是否可以通向 pacific 和 atlantis

**Key 2**: 逆向思维 - 从棋盘的四个边界分别出发回流，看哪些cell可以回流到 (下一个cell比当前cell更高才能回流到)

时间复杂度分析：

根据 Early Stopping Case 3, 其实 有visited函数记录 走过的节点，而走过的节点是不会再走第二次的。

所以 调用dfs函数，**只要参数传入的是 数组pacific，那么地图中 每一个节点其实就遍历一次，无论你调用多少次。**

同理，调用 dfs函数，**只要参数传入的是 数组atlantic，地图中每个节点也只会遍历一次。**

这段代码，pacific遍历花费 $O(m\times n)$, atlantic遍历花费 $O(m\times n)$; 所以，总共花费 $O(2\times m\times n)$
```
	for i := 0; i < rows; i++ {
		dfs(heights, i, 0, &pacific, math.MinInt)       // 左边界回流
		dfs(heights, i, cols-1, &atlantis, math.MinInt) // 右边界回流
	}
	for j := 0; j < cols; j++ {
		dfs(heights, 0, j, &pacific, math.MinInt)       // 上边界回流
		dfs(heights, rows-1, j, &atlantis, math.MinInt) // 下边界回流
	}
```

Time complexity = $O(2\times m \times n) + O(m\times n)$ = $O(m\times n)$

## Resrouces

[参考答案1](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0417.%E5%A4%AA%E5%B9%B3%E6%B4%8B%E5%A4%A7%E8%A5%BF%E6%B4%8B%E6%B0%B4%E6%B5%81%E9%97%AE%E9%A2%98.md)

[参考答案2](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0417.Pacific-Atlantic-Water-Flow)

[参考答案3](https://github.com/wisdompeak/LeetCode/tree/master/DFS/417.Pacific-Atlantic-Water-Flow)