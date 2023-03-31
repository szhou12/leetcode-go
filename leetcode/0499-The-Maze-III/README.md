# [499. The Maze III](https://leetcode.ca/2017-04-12-499-The-Maze-III/)

## Solution idea

### Dijkstra

#### 要点总结
1. [【LeetCode】505. The Maze II](https://www.bilibili.com/video/BV1UX4y157i5/?spm_id_from=333.999.0.0&vd_source=0c02ef6f6e7a2b0959d7dd28e9e49da4) 姐妹题
2. 主要框架参考505, 但要注意与505的差别:
    1. 终点是一个"洞": 在滑动时如果遇到了就会掉洞里. 也就是说, Dijkstra在Expand时，一直走一直走，但是途中遇到了洞就要break
    2. 双重排序: 此题需要 a. 路径最短; b. 上下左右的操作指令的字母序最小
3. 具体实现的key points:
    * node 储存双状态: path cost (int) + node位置 (int) + 到达node的指令 (string)
    * Priority Queue 实现:
        * 内容物: 因为储存的变量type不同，需要 自定义的Tuple结构体
        * 排序: 因为要双排序, 先按照 path cost 排序, 再按照 指令的lexicographical order排序
    * Dijkstra expand 一直走的时候如果遇到 "洞"，直接break

Time complexity = $O(mn \log mn)$


## Resource
[【LeetCode】499. The Maze III](https://www.bilibili.com/video/BV1F44y1P7FN/?spm_id_from=333.999.0.0&vd_source=0c02ef6f6e7a2b0959d7dd28e9e49da4)