# Math Tricks


## 目录
* [十进制](#十进制)
    * [从 integer N 取 digits](#从-integer-n-取-digits)
    * [把 digits 组合成 integer N](#把-digits-组合成-integer-n)
* [Gosper's Hack](#gospers-hack)
* [因子](#因子divisor)
    * [GCD](#gcd-greatest-common-divisor)
    * [LCM](#lcm-least-common-multiple)
    * [N能被因子d整除的个数](#n能被因子d整除的个数)
    * [N的因子对的总数](#n的因子对的总数)
    * [N的因子的总数](#n的因子的总数)
* [组合数](#组合数)
    * [杨辉三角 (Pascal's Triangle) 求C(n, k)](#杨辉三角-pascals-triangle)

## 十进制
### 从 integer N 取 digits
* `n % 10` 得到末位
* `n / 10` 砍掉末位

### 把 digits 组合成 integer N
```
n := 0
for each digit {
    n = (n * 10) + digit
}
```

## Gosper's Hack
* 通过位运算，用二进制数模拟生成所有组合数 ${n \choose k}$.
* 注意！只适合 $n$ 比较小的情况。$n \leq 32$.
```go
func GospersHack(n int, k int) {

	state := (1 << k) - 1

	for state < (1 << n) {

		// do something with current combination
		doSomething(state)

		c := state & -state
		r := state + c
		state = (((r ^ state) >> 2) / c) | r
	}
}
```
* Time complexity = $O({n \choose k})$
* Resources:
    1. [Gosper's Hack Explained](https://programmingforinsomniacs.blogspot.com/2018/03/gospers-hack-explained.html)
        - Explain the mechanism of Gosper's Hack pretty well!!!
    2. [算法学习笔记(75): Gosper's Hack](https://zhuanlan.zhihu.com/p/360512296)

## 因子/Divisor
### GCD (Greatest Common Divisor)
* Euclidean Algorithm (辗转相除法)
```go
func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a % b)
}
```

### LCM (Least Common Multiple)
```go
lcm(a, b) = a * b / gcd(a, b)
```

### N能被因子d整除的个数
* `[1, ..., N]` 自然数区间内能被 `d` 整除的个数 = `N / d`

### N的因子对的总数
* 一个整数 `n` 的**因子对 (Divisor Pair)**的总数的**上限** = `sqrt(n)`
    * 因为 `sqrt(n)` 把 `n` 的所有因数分成左右两边，如果一个因子在`sqrt(n)`的右边，则必然有一个与之配对的因子在左边。所以，`n`的**因子对**的总数最多不超过 `[1, ..., sqrt(n)]` 的长度，也就是总数上限为`sqrt(n)`。
        * 注意！`sqrt(n)`只能是一个上限估计，不完全准确！
    * `n`的所有**因子对**中偏小的divisor，一定在`[1, ..., sqrt(n)]`范围内。
    * e.g. `n = 12`: divisors = 1, 2, 3, 4, 6, 12; divisor pairs = (1, 12), (2, 6), (3, 4)
        * $\sqrt{12} = 2\sqrt{3} = 2\cdot 1.732 \approx 3.4$. 因子对有3对.
    * e.g. `n = 16`: divisors = 1, 2, 4, 8, 16; divisor pairs = (1, 16), (2, 8), (4, 4)
        * $\sqrt{16} = 4$. 因子对有3对, 不超过4对.

### N的因子的总数
* 如果 `n` 刚好能开平方, `因子总数 = 2 * 因子对总数 - 1` (因为 `sqrt(n)`自己跟自己配对)
    * `n = 16`: divisors = 2 * 3 - 1 = 5
* 如果 `n` 不能开平方, `因子总数 = 2 * 因子对总数`
    * `n = 12`: divisors = 2 * 3 = 6
* 规律: 如果正整数`n`有奇数个因子,则`n`为完全平方数
* 因子个数求解公式: 将整数`n`分解为质因子乘积形式, 然后将每个质因子的幂分别加一相乘. `n=(a*a*a)*(b*b)*(c)`, 则因子个数=(3+1)(2+1)(1+1)
    * 200=(2*2*2) * (5*5). 因子个数=(3+1)(2+1)=12个

## 组合数
### 杨辉三角 (Pascal's Triangle)
1. 利用杨辉三角求任意组合数 ${n \choose k}$: `C(n, k) = `杨辉三角中第n行的k-th元素。
2. 杨辉三角第`i`行值的物理意义：二项式`(a+b)^i` (`i` = 0, 1, 2, ...) 展开后所有项的系数。
3. [代码实现]()