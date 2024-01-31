# [3012. Minimize Length of Array Using Operations](https://leetcode.com/problems/minimize-length-of-array-using-operations/description/)

## Solution idea
### Constructive Algorithm + Modulo
1. 理解题意：本题需要通过取模操作来“消掉”元素。观察:
```
x < y ==> x % y = x (消掉了较大的y ==> 如果找到最小值，就可以消掉其他所有元素)
x > y ==> x % y = y2 < y (得到了一个更小于取模y的一个数 ==> 找到一种方法，得到比最小值更小的值z)
0 不参与取模：0 % x or x % 0 都不行
```
2. 通过观察，得出解题规律：
假设最小值为 `x`, 它的出现频次为`cnt`:
* Case 1: If `x` occurs once: The minimum length is 1. 因为其他所有元素都比`x`大，能被消掉。
* Case 2: If `x` occurs more than once: if there exists `y` s.t. `y % x != 0`, then the minimum length is 1. 因为我们可以得到通过`y % x < x`，得到一个比`x`更小的非零最小值，然后回到Case 1消掉其他所有元素。
* Case 3: If `x` occurs more than once AND `y % x == 0` for all `y` (i.e. Neither Case 1 nor Case 2 holds), then:
    * if `cnt == even`, the minimum length is `cnt / 2` (e.g. `4 4 4 4 => 0 0`)
    * if `cnt == odd`, the minimum length is `cnt / 2 + 1` (e.g. `3 3 3 => 3 0`)

Time complexity = $O(n)$

## Resource
[【每日一题】LeetCode 3012. Minimize Length of Array Using Operations](https://www.youtube.com/watch?v=BA_MBQU0Rew&ab_channel=HuifengGuan)