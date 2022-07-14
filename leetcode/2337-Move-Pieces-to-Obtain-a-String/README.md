# [2337. Move Pieces to Obtain a String](https://leetcode.com/problems/move-pieces-to-obtain-a-string/)

## Solution idea
**Two Pointers**: one pointer for `start`, one pointer for `target`

The mechanism is similar to `merge` part in merge sort.

**Key Observation 1**

Ignoring `_`, the relative position of any `L` and `R` must remain through their starting and ending positions.

For example: if starting `LRR`, then ending must be `LRR`. It **can't** be other permutations like `RLR` or `RRL`.

**Key Observation 2**

`L` can only walk towards left, thus, the starting index must $>$ the ending index. In other words, if the starting index $<$ the ending index, then impossible to move like this.


`R` can only walk towards right, thus, the starting index must $<$ the ending index. In other words, if the starting index $>$ the ending index, then impossible to move like this.

**Simpler Solution**

Based on above two key observations, a simpler approach is to do 3 passes on `start` and `target`:

1st pass: check relative positions of all `L` and `R`

2nd pass: check starting (start) and ending (target) index of each `L`

3rd pass: check starting (start) and ending (target) index of each `R`

Time complexity = $O(n)$

## Resources
[【每日一题】LeetCode 2337. Move Pieces to Obtain a String](https://www.youtube.com/watch?v=b7jH5-CFnAM)
