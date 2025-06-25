# [79. Word Search](https://leetcode.com/problems/word-search/description/)

## Solution idea
### 2-D Grid DFS - backtracking
1. Easy to think of the DFS solution.
2. Some reflections stemming from this implementation:
    1. There are two places to place **out-of-bound** check and **already-visited** check: 1. place them in DFS's edge cases; 2. place them before each DFS call.
    2. Case 1: either place them in DFS's edge cases; Case 2: Or place them before each DFS call; Case 3: Or BOTH places.
    3. However, there is a caveat of placing before each DFS call - if you have edge case checking `idx == len(word)` instead of `idx == len(word) - 1`, there will be a bug as such a pre-check before each DFS call prevents the DFS from reaching the last of last letter as it's technically out-of-bound so we won't reach this edge case.
    4. So there is a trade-off: if you want to go "no-brainer", i.e., place these checks in both places, you must use `idx == len(word) - 1` as the edge case checking.
    5. General rule: place these checks in edge cases.
3. Don't forget to backtrack! After all stuff done to the current cell, before we return, we must reset current cell's status as un-visited to allow other paths to visit.

Time complexity = $O(m \times n \times 3^{L})$ where $L$ is the length of the word. For the backtracking function, initially we could have at most 4 directions to explore, but further the choices are reduced into 3 (since we won't go back to where we come from). As a result, the execution trace after the first step could be visualized as a 3-nary tree, each of the branches represent a potential exploration in the corresponding direction. Therefore, in the worst case, the total number of invocation would be the number of nodes in a full 3-nary tree, which is about $3^L$.

Space complexity = $O(L)$. The main consumption of the memory lies in the recursion call of the backtracking function. The maximum length of the call stack would be the length of the word.