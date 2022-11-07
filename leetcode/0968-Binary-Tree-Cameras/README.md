# [968. Binary Tree Cameras](https://leetcode.com/problems/binary-tree-cameras/description/)

## Solution idea
### Post-Order Traversal

* 当前节点根据左、右子节点返回的状态决定自己的状态.

#### 个人总结
* **难点1**: 明确划分每个node所有可能的状态
    1. empty - 空节点
    2. uncovered - 未被camera覆盖
    3. covered - 被camera覆盖
    4. camera - 本身set camera

* **难点2**: 分类讨论所有可能的情况，并划分出状态优先级 (优先级意味着一种状态的出现要优先于其他状态被处理)
    * 每个节点会有左、右两个子节点，所以如何处理当前节点就要根据左、右子节点左右可能状态的组合
    * 一共会有 $4*4 = 16$种cases
        1. 左/右子节点 = empty, 另一个子节点 = 
            1. empty: 看父节点是否存在
            2. uncovered: 优先看uncovered，当前节点放置camera
            3. covered: 看父节点是否存在
            4. camera: 优先看camera，当前节点covered
        2. 左/右子节点 = uncovered, 另一个子节点 =
            1. empty: 优先看uncovered，当前节点放置camera
            2. uncovered: 优先看uncovered，当前节点放置camera
            3. covered: 优先看uncovered，当前节点放置camera
            4. camera: 优先看uncovered，当前节点放置camera
        3. 左/右子节点 = covered, 另一个子节点 =
            1. empty: 看父节点是否存在
            2. uncovered: 优先看uncovered，当前节点放置camera
            3. covered: 看父节点是否存在
            4. camera: 优先看camera，当前节点covered
        4. 左/右子节点 = camera, 另一个子节点 =
            1. empty: 优先看camera，当前节点covered
            2. uncovered: 优先看uncovered，当前节点放置camera
            3. covered: 优先看camera，当前节点covered
            4. camera: 优先看camera，当前节点covered
    * 最简单的，可以在code中把16种cases的处理方法都写一遍，但，从上述的总结中，我们可以归纳出**优先级**
        * **优先级 I**: 子节点只要有一个 uncovered, 处理办法一律是：当前节点必放置camera
        * **优先级 II**: 子节点只要有一个 camera, 除了另一个是 uncovered 并且被**优先级 I**照顾了的情况, 其余处理办法一律是：当前节点 covered
        * **优先级 III**: 子节点只要有一个 empty, 除了: 1. 另一个是 uncovered 并且被**优先级 I**照顾了的情况; 2. 另一个是 camera 并且被**优先级 II**照顾了的情况, 其余处理办法一律是：看父节点是否存在决定是否在当前节点放置camera
        * **优先级 IV**:子节点只要有一个 covered, 除了: 1. 另一个是 uncovered 并且被**优先级 I**照顾了的情况; 2. 另一个是 camera 并且被**优先级 II**照顾了的情况; 3. 另一个是 empty 并且被**优先级 III**照顾了的情况, 剩下的唯一种情况(左、右节点都covered)的处理办法是：看父节点是否存在决定是否在当前节点放置camera

Time complexity = $O(n)$

## Resource

* labuladong:
    * 首先我们列举一下一个节点可能存在的几种状态：
        * 该节点不在监控区域内，称为 uncover 状态；该节点在附近节点的监控范围内，称为 cover 状态；该节点自己装了摄像头，称为 camera 状态；该节点为空节点，称为 empty 状态.
    * 如何保证安装的摄像头数量尽可能少呢？显然就是要尽可能分散，让每个摄像头物尽其用。
        * 具体来说就是自底向上安装摄像头，在叶子节点的父节点上安装摄像头，然后每隔两层再安装（因为每个摄像头都可以管三层）
    * 那么一个节点在什么情况下需要被安装摄像头呢？
        * 显然是当这个节点的子节点处于 uncover 的状态的时候必须安装摄像头，以便覆盖子节点。
    * 综上，我们需要利用后序位置自底向上遍历二叉树，同时要利用子节点的状态以及父节点的状态，判断当前节点是否需要安装摄像头。

[代码随想录-968.监控二叉树](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0968.%E7%9B%91%E6%8E%A7%E4%BA%8C%E5%8F%89%E6%A0%91.md)