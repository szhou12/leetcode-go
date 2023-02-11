# [LeetCode 1650. Lowest Common Ancestor of a Binary Tree III](https://iamageek.medium.com/leetcode-1650-lowest-common-ancestor-of-a-binary-tree-iii-6d008b93376c)

## Solution idea

### LCA with parent pointer

#### Method 1: HashMap

* Traverse from one of the nodes all the way up to the root

* Save its all the ancestors in a hash map (including itself)

* Then, traverse from the other node all the way up to the root, return the first parent node that is already in the hash map.

Time complexity = $O(h)$, Space complexity = $O(h)$

#### Method 2: NOT easy-to-understand, skipped

## Resource

[Solution - LeetCode 1650. Lowest Common Ancestor of a Binary Tree III](https://iamageek.medium.com/leetcode-1650-lowest-common-ancestor-of-a-binary-tree-iii-6d008b93376c)