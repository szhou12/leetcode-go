# [269. Alien Dictionary](https://leetcode.ca/all/269.html)
# [269. Alien Dictionary](https://leetcode.com/problems/alien-dictionary/description/)

## Solution idea
### Topological Sort
#### 要点总结
1. 用 Topological Sort 解题不难想到：如果在input给的顺序上 字母a 先于 字母b，则可以总结出1组方向: `a -> b`
2. 难点：在于实现上，如何利用 lexicographical order 来构建 adjacency-list representation
```
1. 按input的顺序，两两单词比较相同位置上的字母
2. `i`位置上字母如果相同，则跳过，比较`i+1`位置上的字母; 如果不同，则找到一组方向，确保无重复后，添加进图，然后直接break，比较下一组两两单词比较
```
* 简言之，每比较两两一组的单词，只考虑相同位置上首次出现不同字母的情况，之后位置有没有不同字母不再考虑，这是因为之后位置的不同字母已经不再包含先后顺序的信息了 (因为它们前面的字母就已经不同了)

## Resource
- [【每日一题】269. Alien Dictionary, 7/21/2020](https://www.youtube.com/watch?v=yfGJFDkyEmE&ab_channel=HuifengGuan)
- [wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/BFS/269.Alien-Dictionary)
