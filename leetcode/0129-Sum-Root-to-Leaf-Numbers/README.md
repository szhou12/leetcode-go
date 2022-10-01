# [129. Sum Root to Leaf Numbers](https://leetcode.com/problems/sum-root-to-leaf-numbers/)

## Solution idea

### DFS
从根节点出发，进行深搜，每到一个叶子节点就累加路径上求的值
```
# levels = root to leaf node
# branches = left / right
```

Time complexity = $O(2^n)$