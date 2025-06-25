# [779. K-th Symbol in Grammar](https://leetcode.com/problems/k-th-symbol-in-grammar/description/)

## Solution idea
### Math + Recursion
1. Relationship between k-th child and parent in a tree: if k is even, k-th child is right child, its parent is at k/2; if k is odd, k-th child is left child, its parent is at (k+1)/2.
2. So if the parent node's value decides the child node's value.

Time complexity = $O(n)$ because each recursion call reduces n by 1.

Space complexity = $O(n)$

## Resource
[My 3 lines C++ recursive solution](https://leetcode.com/problems/k-th-symbol-in-grammar/solutions/113697/my-3-lines-c-recursive-solution/)