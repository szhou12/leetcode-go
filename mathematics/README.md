# Math Tricks


## 目录
* [十进制](#十进制)
* [Gosper's Hack](#gospers-hack)
* [因子](#因子divisor)

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

### N能被a整除的个数
* `[1, ..., n]` 自然数区间内能被 `a` 整除的个数 = `n / a`

### N的因子Pair的总数
* 一个整数 `n` 的**因子Pair**的总数的**上限** = `sqrt(n)`
    * 因为 `sqrt(n)` 把 `n` 的所有因数分成左右两边，如果一个因子在`sqrt(n)`的右边，则必然有一个与之配对的因子在左边。所以，`n`的**因子Pair**的总数最多不超过 `[1, ..., sqrt(n)]` 的长度，也就是总数上限为`sqrt(n)`

### N的因子的总数
* 如果 `n` 刚好能开平方, `因子总数 = 2 * 因子Pair总数 - 1` (因为 `sqrt(n)`自己跟自己配对)
* 如果 `n` 不能开平方, `因子总数 = 2 * 因子Pair总数`