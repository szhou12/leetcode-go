# Math

## 目录
* [String上模拟加减运算](#string上模拟加减运算)
* [中位数定理](#中位数定理)
* [Constructive Algorithm (构造题)](#constructive-algorithm)
* [Degree of Freedom (自由度)](#degree-of-freedom)
* [Frequency](#frequency)
* [Combinatorics](#combinatorics)
* [Geometry](#geometry)

## String上模拟加减运算
* :red_circle: 寻找最接近的回文串: [564. Find the Closest Palindrome](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0564-Find-the-Closest-Palindrome)
    * 镜像复制形成回文串
    * String上进行十进制的加减操作
    * Go中，`string`, `[]rune`, `rune`, `int`之间的转换
    * 可以作为很多涉及string的加减运算，回文串等题目的母题

* :red_circle: 最小代价将数组化为Equalindromic: [2967. Minimum Cost to Make Array Equalindromic](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2967-Minimum-Cost-to-Make-Array-Equalindromic)
    

## 中位数定理
* :red_circle: 最小代价将数组化为Equalindromic: [2967. Minimum Cost to Make Array Equalindromic](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2967-Minimum-Cost-to-Make-Array-Equalindromic)
    * 中位数定理：确定最小代价的目标点
    * Closest Palindrome: 找到与string表示的目标点的最接近的回文串

## Constructive Algorithm
* :red_circle: 模运算缩短array长度: [3012. Minimize Length of Array Using Operations](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3012-Minimize-Length-of-Array-Using-Operations)

## Degree of Freedom (自由度)
* :red_circle: [3132. Find the Integer Added to Array II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3132-Find-the-Integer-Added-to-Array-II)
    * Degree of Freedom + Sorting

## Frequency
* :yellow_circle: digit不同的pair总数: [3153. Sum of Digit Differences of All Pairs](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3153-Sum-of-Digit-Differences-of-All-Pairs)
    * digit counting

* :yellow_circle: 最少更改数使字符串以k长重复: [3137. Minimum Number of Operations to Make Word K-Periodic](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3137-Minimum-Number-of-Operations-to-Make-Word-K-Periodic)
    * map to record frequency

## Combinatorics
* Grosper's Hack: Use bit operation to generate all combinations (states) of $\binom{n}{k}$
* Pascal's Triangle = Binomial Coefficients: In Pascal's Triangle, the value of the element at row $n$ and position $k$ (starting from 0) = $\binom{n}{k}$

* :red_circle: k秒后n号位的值: [3179. Find the N-th Value After K Seconds](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3179-Find-the-N-th-Value-After-K-Seconds)
    * Method 2: resembles Pascal's Triangle to calculate $\binom{n+k-1}{k}$


## Geometry
* :red_circle: 覆盖所有1的最小矩形II: [3197. Find the Minimum Area to Cover All Ones II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3197-Find-the-Minimum-Area-to-Cover-All-Ones-II)

* :yellow_circle: 覆盖所有1的最小矩形I: [3195. Find the Minimum Area to Cover All Ones I](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3195-Find-the-Minimum-Area-to-Cover-All-Ones-I)

## Modulo
## Definition
- if `a = k * n + r`, then `a % n = r`.
### Key Properties
- Symmetry: if `a % n = r`, then `(a+n) % n = r`.
- Transitivity: if `a % n = b % n` and `b % n = c % n`, then `a % n = c % n`.
- Distributivity: (e.g., a = 13, b = 11, n = 2)
    - `(a + b) % n = (a % n + b % n) % n` 
    - `(a * b) % n = ((a % n) * (b % n)) % n`
- Congruences: `a % n = b % n` $\Longleftrightarrow $ $a \equiv b(\mod n)$ $\Longleftrightarrow $ a和b在取模n时"相等" (aka.同余)