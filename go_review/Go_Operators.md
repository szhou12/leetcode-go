# [Operators](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

## Contents
* [Bitwise Operators](#bitwise-operators)
* [% Operator](#mod-operator)

## Bitwise Operators
[Go - Bitwise Operators](https://www.tutorialspoint.com/go/go_bitwise_operators.htm)

| Operator | Description | 个人理解 |
| :-: | - | - |
| `&` | Binary AND Operator | 对应bit位上"与"操作，任一有0即为0，两者都是1才是1 |
| `\|` | Binary OR Operator | 对应bit位上"或"操作，任一有1即为1 |
| `^` | Binary XOR Operator | 对应bit位上"异或"操作，两者相同即为0，必须一个0一个1才是1 |
| `<<` | Binary Left Shift Operator | 从右侧补n个0。`a = 0011 1100` :arrow_right: `a<<2 = 1111 0000` |
| `>>` | Binary Right Shift Operator | 从左侧补n个0。`a = 0011 1100` :arrow_right: `a>>2 = 0000 1111` |

[Go << and >> operators - Stack Overflow](https://stackoverflow.com/questions/5801008/go-and-operators)
* `n << t`: 另一层意思 - "n times 2, t times". i.e. $n \times 2^t$
* `n >> t`: 另一层意思 - "n divided by 2, y times". i.e. $n \times 2^{-t}$

## Mod Operator
Go 中的 % 与 Python 中的 %, 在对负数取mod时结果会不同

Go:
```go
-2 % 8 = -2
```
vs.
Python:
```python
-2 % 8 = 6 = 8 - 2
```

所以，在Go中如果希望 % 的行为像Python中一样，总是返回正数的话，需要自己定义一个函数:
```go
func mod(a, b int) int {
    return (a % b + b) % b
}
```