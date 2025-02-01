# [2454. Next Greater Element IV](https://leetcode.com/problems/next-greater-element-iv/description/)

## Solution idea
### Stack - nextGreater
1. 我们已经熟悉基础的 find the next greater element 怎么做。那么，有什么办法我们可以利用上 nextGreater ?
2. 重新审视基础算法的机制，发现：为了维护递减stack，每次栈顶元素遇到更大的来者，需要出栈时，这是它的“第一次”first next greater element。注意到重点了吗？没错，就是这个“第一次”。本题中，我们需要找到“第二次”，那么，需要将这些有过“第一次”的元素和没有过的元素做标记区分。
3. 怎么标记有过“第一次”的元素？不出栈？显然不行，因为栈的递减性质就无法维护。那么，暂时把它们拿出来，存在某个地方？好像可以。拿什么数据结构存比较合适呢？怎么存，我们才能在接下来遇到“第二次”时，高效地更新结果？既然在一个递减stack里，我们就可以遇到“第一次”，那么，照猫画虎，再来一个递减stack，用来遇到“第二次”？好像可以哎，来试试看。
4. 创建两个递减栈: stack1, stack2。 stack1存没有过“第一次”的元素，stack2存有过“第一次”的元素。每次来一个来者，先跟stack2中的栈顶元素比较，如果来者更大，好了，栈顶元素遇到它的“第二次”了，出栈，我们更新结果。然后这个来者不能进入stack2，因为它也还没有“第一次”。所以，这个来者紧接着要跟stack1中的栈顶元素比较，如果来者更大，好了，栈顶元素遇到它的“第一次”了，出栈，暂时存到一个缓存里(数组)。直到stack1中元素都比来者大了，那么这个来者可以进入stack1。而缓存里的这些有了“第一次”的元素要以逆序 (先大后小) 倒入stack2. 注意，这时可能stack2中还有元素，但这没关系，因为这些留守的元素一定比来者大，而来者又大于这些新来的，所以留守的一定大于新来的。stack2的递减性质没有破坏。

Time complexity = $O(4n)$

## Resource
[【每日一题】LeetCode 2454. Next Greater Element IV](https://www.youtube.com/watch?v=vrNFhKKHkP0&ab_channel=HuifengGuan)