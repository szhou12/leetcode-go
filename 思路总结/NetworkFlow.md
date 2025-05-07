# Network Flow

## 目录
- [Terminology](#terminology)
- [Ford Fulkerson Algorithm](#ford-fulkerson)
- [Tricks](#tricks)
- [Resources](#resources)

## Terminology
1. **Augmenting Path**: A path $P$ from node `s` to node `t` where all edges (can be original or reverse edges) have remaining capacity > 0.
2. **Residual Graph/Network**: $G' = G + \overline{G} = (V, \overrightarrow{E} \cup \overleftarrow{E})$. Augmented from original graph $G$: for each original edge $A \rightarrow B$, add a reverse edge $A \leftarrow B$.
3. **Residual/Remaining Capacity** of edge: $e.capacity - e.flow$ 
4. **Residual Edges**: All edges (including original and reverse edges) in the Residual Graph $G'$.
    - Definition: edge with remaining capacity > 0.
    - Reason to exist: to "undo" bad augmenting paths that do not lead to max flow.
5. **Bottleneck Edge**: The edge on an Augmenting Path with the minimum Remaining Capacity among all edges on this path.
6. Max Flow = sum of all found bottleneck edges' flow.

```
Ford_Fulkerson(G, s, t):
    // Step 1: Construct Residual Graph G'
    For each original edge (A -> B) in G:
        add a reverse edge (A <- B)
    
    // Step 2: Assign initial values to all edges in G'
    orignal edge: e(flow, capacity, cost) = e(0, cap, cost)
    reverse edge: e(flow, capacity, cost) = e(0, 0, -cost)

    // Step 3: As long as an augmenting path exists, take this path, 
    // find the bottleneck edge and its remaining capacity c, 
    // update all edges flow on this path by this remaining capacity c
    while (P = findAugmentingPath(G', s, t) && P != null):
        c = findMinRemainingCapacity(P)
        all forward edges flow += c
        all backward edges flow -= c
        update all edges' remaining capacity

    return maxFlow
```

NOTE: "Foward" only means leading the direction from `s -> t` on the path P. Thus, a foward edge can be an original edge as well as a reverse edge. "Backward" only means leading the opposite direction of its forward edge. Thus, it can be an original edge as well as a reverse edge.

## Tricks
1. How to find Augmenting Path: Always pick the "shortest" path
    - Edmonds-Karp: BFS. shortest = least number of edges.
    - Min Cost Max Flow: Dijkstra's. shortest = least cost.
2. SuperSource `S` & SuperSink `T`
    - On a graph with "Supply Nodes" and "Demand Nodes", add `S` and a directed edge `S -> u` for all supply nodes `u`; add `T` and a directed edge `v -> T` for all demand nodes `v`.
    - capacity of `S -> u` = |supply of `u`|. cost = 0.
    - capacity of `v -> T` = |demand of `v`|. cost = 0.
3. Undirected Graph
4. Johnson's Reweighting / Potential

## Resources
- [Youtube | Max Flow Ford Fulkerson | Network Flow | Graph Theory](https://www.youtube.com/watch?v=LdOnanfc5TM&ab_channel=WilliamFiset)
- [Youtube | Edmonds Karp Algorithm | Network Flow | Graph Theory](https://www.youtube.com/watch?v=RppuJYwlcI8&ab_channel=WilliamFiset)
- [Blog | Minimum-cost flow - Successive shortest path algorithm](https://cp-algorithms.com/graph/min_cost_flow.html)
- [Github | williamfiset/Algorithms/networkflow](https://github.com/williamfiset/Algorithms/tree/master/src/main/java/com/williamfiset/algorithms/graphtheory/networkflow)