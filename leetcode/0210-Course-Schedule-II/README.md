# [210. Course Schedule II](https://leetcode.com/problems/course-schedule-ii/)

## Solution idea

### Topological Sort
Topological sort to find a topological order

Time Complexity = $O(V+E)$ = O(numCourses + len(prerequisites)). We pop each node exactly once from the zero in-degree queue and that gives us V. Also, for each vertex, we iterate over its adjacency list and in totality, we iterate over all the edges in the graph which gives us E. Hence, O(V+E).

Space Complexity = $O(V+E)$. The in-degree array requires O(V) space. We use an intermediate queue data structure to keep all the nodes with 0 in-degree. In the worst case, there won't be any prerequisite relationship and the queue will contain all the vertices initially since all of them will have 0 in-degree. That gives us O(V). Additionally, we also use the adjacency list to represent our graph initially. The space occupied is defined by the number of edges because for each node as the key, we have all its adjacent nodes in the form of a list as the value. Hence, O(E). So, the overall space complexity is O(V+E).

## Resource
[wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/BFS/210.Course-Schedule-II)

[环检测及拓扑排序算法](https://labuladong.github.io/algo/di-yi-zhan-da78c/shou-ba-sh-03a72/huan-jian--e36de/)