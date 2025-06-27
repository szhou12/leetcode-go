# [323. Number of Connected Components in an Undirected Graph](https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/description/)

## Solution idea
### DFS

Time complexity = $O(V + E)$. Building the adjacency list will take $O(E)$ operations, as we iterate over the list of edges once, and insert each edge into two lists. During the DFS traversal, each vertex will only be visited once. This is because we mark each vertex as visited as soon as we see it, and then we only visit vertices that are not marked as visited. In addition, when we iterate over the edge list of each vertex, we look at each edge once. This has a total cost of $O(V + E)$.

Space complexity = $O(V + E)$. Building the adjacency list will take $O(E)$ space. To keep track of visited vertices, an array of size $O(V)$ is required. Also, the run-time stack for DFS will use $O(V)$ space.
