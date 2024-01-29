# [3007. Maximum Number That Sum of the Prices Is Less Than or Equal to K](https://leetcode.com/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/description/)

## Solution idea
1. 理解题意：注意一个点，对一个自然数 `x <= num`，它每一个满足要求`i%x == 0` 的bit位，如果是1 (set bit)，那么就要计算一次。也就是说，是需要有double counting的。

### Binary Search + Digital Counting
1. Binary Search: 首先要有直觉，题目具有单调性：num越大，price越高。所以可以使用 Binary Search 来猜答案。
2. Digital Counting: 本题难点在于对于任意一个猜的`mid`，折半的条件是什么？或者说，对应的`price`怎么计算？这里使用 Digital Counting。有两种方法，下面会分别讨论。

#### Digital Counting - Method 1
1. 每猜一个`num`，从 1～`num` 每一个自然数都分别计算 bit 1 的个数比较笨拙。Method 1的方法是：与其每一个数计算有多少个 bit 1，逆向思维，对于`num`的每一个要求的bit位`i`，计算有多少个组合数满足：1. bit位`i=1`; 2. 这个组合的十进制表示数值`<= num`。
2. 假设，`num`二进制表示为：(低位) `L   L   L   i   H    H    H` (高位)，以`i`位计算组合数。

**Case 1**: 如果bit位`i`为0
```
L   L   L   i   H    H    H
  2^m       0   000:(HHH-1)
```
* 把`i`位改设为1。
* 高位 (相对于`i`) 允许的组合数范围：`000 ~ (HHH-1)`，上界到`HHH-1`， 因为需要`<= num`。组合数总共有 `HHH` 个。
* 低位 (相对于`i`) 允许的组合数范围：可以任意组合。组合数总共有 `2^m` 个。
* 总组合数 = `HHH * 2^m`

**Case 2**: 如果bit位`i`为1
```
L   L   L   i   H    H    H
  2^m       1   000:(HHH-1)
```
* Case 2a: 与Case 1一样
  * 高位 (相对于`i`) 允许的组合数范围：`000 ~ (HHH-1)`，上界到`HHH-1`。组合数总共有 `HHH` 个。
  * 低位 (相对于`i`) 允许的组合数范围：可以任意组合。组合数总共有 `2^m` 个。
```
L   L   L   i   H    H    H
  000:LLL     1     HHH (fixed)
```
* Case 2b:
  * 高位 (相对于`i`) 允许的组合数范围：固定为`HHH`。组合数总共有1个。
  * 低位 (相对于`i`) 允许的组合数范围：`000 ~ LLL`，上界到`LLL` (此时就是`num`自身)。组合数总共有 `LLL+1` 个。
* 总组合数 = `HHH * 2^m + 1 * (LLL+1)`

3. 实现细节
  * 二进制表示默认的位数读取(从低位到高位)是从右到左。实现上，从左到右表示低位到高位。用一个array来存每一位的bit。
  * 题目中二进制表示的位数是1-indexed。但是上述的实现的array是0-indexed。`i`位每`x`步一跳，如果是1-indexed，`x=2`，则`i=2, 4, 6, ...`。在0-indexed的array中对应的位置则是`i=1, 3, 5, ...`。所以，`i`的loop可以写成：`for i := x-1; (1<<i) <= num; i += x { ... }`。`i=x-1`左移调整到0-indexed。`(1<<i) <= num`最高位`i`不超过`num`。`i += x`每`x`步一跳。

#### Digital Counting - Method 2
1. Method 2是通过观察bit位重复的规律进行总结归纳。
假设`num=14`，列出前十五行`0 ~ num` (注意！Method 2要从0算起才能找到规律) 所有数字的二进制表示：
```
       4   3   2   1   --> bit位
     +---+---+---+---+
     | 0 | 0 | 0 | 0 | --> 0
     +---+---+---+---+
     | 0 | 0 | 0 | 1 | --> 1
     +---+---+---+---+ 
     | 0 | 0 | 1 | 0 | --> 2
     +---+---+---+---+ 
     | 0 | 0 | 1 | 1 | --> 3
     +---+---+---+---+ 
     | 0 | 1 | 0 | 0 | --> 4
     +---+---+---+---+
     | 0 | 1 | 0 | 1 | --> 5
     +---+---+---+---+ 
     | 0 | 1 | 1 | 0 | --> 6
     +---+---+---+---+ 
     | 0 | 1 | 1 | 1 | --> 7
     +---+---+---+---+ 
     | 1 | 0 | 0 | 0 | --> 8
     +---+---+---+---+ 
     | 1 | 0 | 0 | 1 | --> 9
     +---+---+---+---+ 
     | 1 | 0 | 1 | 0 | --> 10
     +---+---+---+---+ 
     | 1 | 0 | 1 | 1 | --> 11
     +---+---+---+---+ 
     | 1 | 1 | 0 | 0 | --> 12
     +---+---+---+---+
     | 1 | 1 | 0 | 1 | --> 13         
   +-------------------+
   | +---+---+---+---+ |
   | | 1 | 1 | 1 | 0 | | --> 14
   | +---+---+---+---+ |
   +-------------------+
     | 1 | 1 | 1 | 1 | --> 15
     +---+---+---+---+
```
2. 注意：Method 2使用1-indexed，低位到高位是从右往左。
3. 观察得出规律：for `ith column`, the bit toggles after every $2^{i-1}$ row
  * `1st column`: the bit toggles after every `1 = 2^0` row
  * `2nd column`: the bit toggles after every `2 = 2^1` rows
  * `3rd column`: the bit toggles after every `4 = 2^2` rows
  * `4th column`: the bit toggles after every `8 = 2^3` rows
4. 根据上述规律，把一次`0/1`变换视作一组group，则发现：for `ith column`, group size $=2^i$ (number of bits)。
  * `1st column`: group size `2 = 2^1`. `01 01 01 01 ...`
  * `2nd column`: group size `4 = 2^2`. `0011 0011 0011 0011 ...`
  * `3rd column`: group size `8 = 2^3`. `00001111 00001111 00001111 00001111 ...`
  * `4th column`: group size `16 = 2^4`. `0000000011111111 0000000011111111 ...`
5. 根据上述规律，我们发现：
    1. 每一个group一半是0，一半是1。bit 1数量表示为：$\frac{2^i}{2} = 2^{i-1}$
    2. bit 0总是先于bit 1出现。
6. 根据总结的规律，归纳统计`ith column`的bit 1总数的计算公式：
  * current column = `i`, group size = `2^i`, number of bit 1 per group = `2^(i-1)`, 用`rows = num+1`表示从`0`到`num`依次列出的总行数
  * 计算公式 Part 1: `(rows / 2^i) * 2^(i-1)` 
    * 注意！除法向下取整，所以这里的group数只计算了能被整除的部分
  * 计算公式 Part 2: `max{0, (rows % 2^i) - 2^(i-1)}`
    * remainder group size = `rows % 2^i`
    * 0优先出现 且 一半是0，所以`(rows % 2^i) - 2^(i-1)`表示剩余的group中bit 1的数量
    * 显然，remainder group size不足一半的时候，不会有1，所以取`max{0, ...}`

Time complexity = $\log n \times k$ where $\log n$ is binary search and `k` is iterating over all bit columns. 

## Resource
[【每日一题】LeetCode 3007. Maximum Number That Sum of the Prices Is Less Than or Equal to K](https://www.youtube.com/watch?v=tw6jJCIq0lU&t=20s&ab_channel=HuifengGuan)

[Leetcode Top Solution](https://leetcode.com/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/solutions/4563689/c-solution-bit-manipulation-binary-search-comprehensive-visualization)