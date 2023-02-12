# [1676. Lowest Common Ancestor of a Binary Tree IV](https://zhenchaogan.gitbook.io/leetcode-solution/leetcode-1676-lowest-common-ancestor-of-a-binary-tree-iv)

## Solution idea

### LCA

* input nodes **一定**存在于tree中

* 逻辑与LCA I: [236. Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)一致

* 因为input nodes 是一个array，所以使用一个 **HashMap** 来当作input nodes的set
    * Base case comparison: LCA I vs. LCA IV
    * `if root == p || root == q {return root}` vs. `if set[root] == true {return root}`

Time complexity = $O(h)$, Space complexity = $O(n)$ where $n$ is number of input nodes