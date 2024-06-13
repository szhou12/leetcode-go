# [272. Closest Binary Search Tree Value II](https://leetcode.ca/all/272.html)

## Solution idea
### Recursion: In-Order Traversal + Queue
1. Binary Search Tree的一个重要性质：中序遍历可以得到递增数列。
2. 在一个递增数列上，怎么找出k个最接近target的数？用一个queue撸一遍。中序遍历从小到大遍历元素的时候，每遍历一个元素，就把它加入queue中，直到queue中已经有k个元素。下一个元素要进queue的之前，这个元素与target的差值 和 queue首元素与target的差值做比较，如果下一个元素的差值更小，那么扔掉queue首元素，加入下一个元素；如果更大，那么可以就此打住，不再往后看了，因为递增的关系，之后的元素带来的差值只会更大。
3. 为什么使用queue可以保证正确性？因为queue中的元素也是按照递增排列的，所以queue首的元素一定是queue里所有元素里与target最远的元素。

Time complexity = $O(n)$