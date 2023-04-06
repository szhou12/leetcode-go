# [1201. Ugly Number III](https://leetcode.com/problems/ugly-number-iii/description/)

## Solution idea

### Binary Search

#### 要点总结
1. 对于一个`mid`, 最好想要有一个`fcn(mid, a, b, c)`可以帮我们判断在`[1, ..., mid]`这段区间的自然数里有多少个丑数. 这样就很容易实现Binary Search模版
    * if < n : left = mid + 1
    * else   : right = mid

2. 怎么实现`fcn(mid, a, b, c)`? 利用容斥原理 (inclusion–exclusion principle)
    * `[1, ..., mid]`中能被`a`整除的个数 `A = mid / a`
    * `[1, ..., mid]`中能被`b`整除的个数 `B = mid / b`
    * `[1, ..., mid]`中能被`c`整除的个数 `C = mid / c`
    * `[1, ..., mid]`中能被`a & b`整除的个数 `A ∩ B = mid / lcm(a, b)`
    * `[1, ..., mid]`中能被`a & c`整除的个数 `A ∩ C = mid / lcm(a, c)`
    * `[1, ..., mid]`中能被`b & c`整除的个数 `B ∩ C = mid / lcm(b, c)`
    * `[1, ..., mid]`中能被`a & b & c`整除的个数 `A ∩ B ∩ C = mid / lcm(a, lcm(b, c))`
    * `[1, ..., mid]`中能被 `a, b, c` 任意一个整除的个数 = `A + B + C - A ∩ B - A ∩ C - B ∩ C + A ∩ B ∩ C`

3. 需要直接记住的数学定理:
    * `[1, ..., n]`自然数区间内能被 `a` 整除的个数 = `n / a`
    * `[1, ..., n]`自然数区间内能被 `a & b` 同时整除的个数 = `n / lcm(a, b)`
    * lcm 可由 gcd 计算得到: `lcm(a, b) = a * b / gcd(a, b)`
    * gcd 可由 Euclidean algorithm (辗转相除法) 计算得到

Time complexity = $O(\log n)$

## Resource
[微软面试题解析：丑数系列算法](https://mp.weixin.qq.com/s/XXsWwDml_zHiTEFPZtbe3g)

[【每日一题】1201. Ugly Number III, 10/29/2019](https://www.youtube.com/watch?v=60PyXFrEf44&ab_channel=HuifengGuan)