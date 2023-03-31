# [505. The Maze II](https://leetcode.ca/all/505.html)

## Solution idea

### Dijkstra

#### 要点总结
1. 题目给的条件像是以前宝可梦打冰系道馆: 一条路确定了方向会一直滑到头。难点: 如何模拟/维护这个"同一个方向一直走不能停下来"的状态
    * 一开始我的想法是: 每个node记录来时的方向, 如果当前node不是墙，则可以继续沿着来时的方向前进，否则，就改变方向。但是这样做有一个严重问题: 每走过一个node都记录了从起点到当前的距离(“步数”)，但是根据题目设置，那些“滑过去”的中间node不能记录步数，只有撞到墙的时候才记录步数。
    * 解决方法: "不能停下来 == 一直往前走到停下为止" - Dijkstra 每次expand的时候，一个方向一直走到撞墙或出界, 此时才更新走到当前位置node的步数, 中间“滑过”的node一律不记录

Time complexity = $O(mn\log mn)$

## Resource
[【LeetCode】505. The Maze II](https://www.bilibili.com/video/BV1UX4y157i5/?spm_id_from=333.999.0.0&vd_source=0c02ef6f6e7a2b0959d7dd28e9e49da4)