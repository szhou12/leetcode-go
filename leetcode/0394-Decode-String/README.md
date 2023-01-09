# [394. Decode String](https://leetcode.com/problems/decode-string/description/)

## Soltion idea - Stack
* 看到题目input带括号，首先考虑用 **Stack** 方法来解
* **Key 1**: 本题使用两个 **Stack**, 一个存 numbers, 一个存 letters
* **Key 2**: 注意括号会有嵌套的情况 e.g. `3[a2[c]]`
* **Key 3**: 多位数的number要一次拿完 e.g. `100[leetcode]`

Case 1: 遇见 `number`, 进栈 `numbers stack`

Case 2: 遇见 `[`, 进栈 `letters stack`

Case 3: 遇见 `letter`, 进栈 `letters stack`

Case 4: 遇见 `]`, pop `letters stack` 栈顶元素直到 `[`; pop `numbers stack` 栈顶元素; 重复 `n` 遍再倒回 `letters stack`

Time complexity = $O(n)$

## Resrouces
[贾考博 LeetCode 394. Decode String - 如果你愿意一层一层一层地剥开我的心](https://www.youtube.com/watch?v=JXlosO-4BSI&ab_channel=%E8%B4%BE%E8%80%83%E5%8D%9A)