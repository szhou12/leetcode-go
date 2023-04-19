# [2115. Find All Possible Recipes from Given Supplies](https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/description/)

## Solution idea

### Topological Sort
#### 要点总结
1. 本题思路不算难。`supplies -> ingredients -> recipes` 有一个明显的有方向的逻辑链条，很自然能想到 Topological Sort
2. 本题难点在于：构造 Adj-list representation - 只有`ingredients`和`recipes`中的元素需要参与构建 Adj-list；选用那种data structure合适？
    * 本人解题的时候使用的是map of maps `map[string]map[string]bool`。 也可以用map of slices `map[string][]string`。区别只在于，前者需要避免重复edge来防止degree多+1，后者不需要考虑、可直接append

#### 代码结构
```
Step 1: Reconstruct adj-list representation + Calculate degree
Step 2: Topological Sort - recipe whose degree can become 0 is a possible one
Step 3: Collect all possible recipes
```

Time complexity = $O(n)$ where `n = |recipes| + |ingredients| + |supplies|`

## Resource
[【每日一题】LeetCode 2115. Find All Possible Recipes from Given Supplies](https://www.youtube.com/watch?v=FuuGI4yg1q0)