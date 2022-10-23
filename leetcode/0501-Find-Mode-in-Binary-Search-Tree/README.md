# [501. Find Mode in Binary Search Tree](https://leetcode.com/problems/find-mode-in-binary-search-tree/)

## Solution idea

### Map - naive solution
* Use a map to count each distinct node value
* Loop the map to get the most frequent (mode) node value(s)
* **Pros**: 适用于任何tree structure, 不局限于 BST
* **Cons**: 额外的space 开销, 没有利用到BST的性质 (ie. 中序遍历是 increasing order --> node value就有序了)

Time complexity = $O(n)$, Space complexity = $O(n)$

### Better Solution - 利用BST的性质
* *NOTE:* 用Go非常不好写, 而且in-order traversal要写成 anonymous func 才能对, 写成helper func无法pass
* BST性质: 中序遍历是 increasing order
* 通过 in-order traversal, 我们就可以在树上模拟 在一个increasing order的array上如何通过一个loop找到 身为众数的 元素(s)
* 在一个increasing order的array上 判断当前元素是否属于众数, 只需要拿它与前一个元素比较 (所以需要 `prev node`)
    * 遍历有序数组的元素出现频率，从头遍历，那么一定是相邻两个元素作比较，然后就把出现频率最高的元素输出就可以了
* 形象化:
```
[..., B, A, C, ...]

In-order(root.Left) // 相当于完成了 [..., B] 部分
当前层 do something // 相当于现在是拿A与B比较, 并且完成 [..., B, A] 部分
In-order(root.Rightyeft) // 相当于是在完成 [..., B, A, C] 部分

```
* 当前层的两层嵌套逻辑
    1. 首先，判断 当前元素 是否 与prev node 相同
    2. 然后，相同与不同这两种case 要分别比较 curCount 与 maxCount - curCount > maxCount 和 curCount == maxCount 两种情况

Time complexity = $O(n)$, Space complexity = $O(1)$

## Resource
[代码随想录-501.二叉搜索树中的众数](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0501.%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E4%B8%AD%E7%9A%84%E4%BC%97%E6%95%B0.md)
