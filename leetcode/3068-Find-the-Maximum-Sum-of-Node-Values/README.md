# [3068. Find the Maximum Sum of Node Values](https://leetcode.com/problems/find-the-maximum-sum-of-node-values/description/)

## Solution idea
### Bit Manipulation (XOR)
1. 突破点: 所有突破点都来自于本题中对operation的描述
    ```
    Alice can perform the following operation any number of times (including zero) on the tree:

    Choose any edge [u, v] connecting the nodes u and v, and update their values as follows:
    - nums[u] = nums[u] XOR k
    - nums[v] = nums[v] XOR k
    ```
    1. Why any number of times?
        - 先做一个实验：对于一个数(XOR k)无数次会发生什么?
        ```
        3: 0 0 1 1
        7: 0 1 1 1
        3xor7 = 0100 = 4
        3xor7xor7 = 3
        3xor7xor7xor7 = 0100
        3 xor^even 7 = 3
        3 xor^odd 7 = 0100
        ```
        - 通过实验，可以发现：(XOR k)偶数次，原数不变；(XOR k)奇数次，原数发生变化。并且任意奇数次的变化结果相同。
        - 所以可以得出结论：本题中对一个节点进行(XOR k)操作，操作一次就够了。
    2. Why tree? 
        - 迷惑我的是这道题为什么给出一个tree？
        - 解答要重新回到对operation的描述：每次选一个边，边两头儿的node都要进行(XOR k)
        - tree中任意两个node都是连通的。如果我们对两个非直接连通的node (i.e., 这两个node不是由一条边直接相连) 同时进行 (XOR k)操作一次会发生什么？可以发现，路径上的每一个中间node自然地进行了(XOR k)操作两次 (可以随便举个例子画一画)，值保持不变。也就是说，这道题中tree的作用就是：我们可以任意选择一对儿(即使不是直接连通的)node进行(XOR k)操作。
    3. 好了，现在我们可以自由选择一对儿一对儿的node进行(XOR k)操作，怎么选可以达到最大值？
        - 优先选择进行(XOR k)操作后，增量较大的node
        - 实现上：使用一个array`diff`来记录每个node进行(XOR k)操作后的增量。然后对`diff`进行从大到小排序。
        ```
        diff [+4, +3, +2, +2, +1, -2, -1]
              ------
                      ------
        ```
        - 遍历`diff`：每次选一对儿，累计增量和。记录遍历途中得到的最大增量和。

Time complexity = $O(n\log n)$

## Resource
[【每日一题】LeetCode 3068. Find the Maximum Sum of Node Values](https://www.youtube.com/watch?v=CF7rTxasbow&t=613s&ab_channel=HuifengGuan)