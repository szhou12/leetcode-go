# [String](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

## Contents
* [Strings are Immutable](#strings-are-immutable)
* [Sort a String](#sort-a-string)
* [Reverse a String](#reverse-a-string)
* [rune vs. byte](#rune-vs-byte)
* [Substract Two Bytes](#substract-two-bytes)
* [String Concatenation](#string-concatenation)

## Strings are Immutable
[Swap Characters of a string in Go (Golang)](https://golangbyexample.com/swap-characters-string-golang/)
* TLDR:
    * Once a string varibale is declared, it’s immutable. That is, you CANNOT swap characters by index. You need to convert `string` to `[]rune` beforehand (i.e., char[] in Java).

## Sort a String
[golang - how to sort string or []byte?](https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte)
```go
word := "1BCagM9"
// Step 1: convert string to []rune
s := []rune(word)
// Step 2: sort in lexicographical order
sort.Slice(s, func(i int, j int) bool {
    return s[i] < s[j]
})
// Step 3: convert []rune back to string
word = string(s)
```

## Reverse a String
[How to reverse a string in Golang](https://www.educative.io/answers/how-to-reverse-a-string-in-golang)

[How to reverse a string in Go?](https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go)

* Formal Template:
```go
func reverse(s string) string {
    // Step 1: Convert string to []rune
    runes := []rune(s)
    // Step 2: Swap the runes
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    // Step 3: Convert back to string
    return string(runes)
}
```

* Variation (Not Recommended):
```go
func reverse(word string) string {
    // Prepend each character to the front of return value
    reversedWord := ""
    for _, char := range word {
        reversedWord = string(char) + reversedWord
    }
    return reversedWord
}
```

## rune vs. byte
[Rune vs byte ranging over string](https://stackoverflow.com/questions/58635507/rune-vs-byte-ranging-over-string)

* In Go, Characters are expressed by enclosing them in single quotes like this: **'a'**
* A character is either type `byte` or type `rune`.
* `byte` is an alias for `uint8`, an unsigned 8-bit integer. Governed by the rules of unsigned arithmetic, meaning (smaller byte - larger byte) wraps aournd, never giving negative value. 
```go
// Indexing a string always gives a byte
// Even if it's a Unicode character, it's still a byte. MAY only be accessing just a part of it, not the whole character.
word := "ig"
x := word[1] // x is byte
```
* `rune` is an alias for `int32`, a 32-bit integer. NOT governed by the rules of unsigned arithmetic, meaning (smaller rune - larger rune) gives negative value.
```go
// Go handles character literal as rune
x = 'i' // x is rune
```
* `byte` type represents **ASCII** characters; `rune` type represents a more broader set of **Unicode** characters.
* Extract characters from a string:
    1. if we `range` a string, the characters we get are `rune` type
    ```go
    s := "hello"
    for _, c := range s {
        fmt.Printf("%c\n", c) // c is rune type
    }
    ```
    2. if we **index** a string, the characters we get are `byte` type
    ```go
    s := "hello"
    for i := 0; i < len(s); i++ {
        fmt.Printf("%c\n", s[i]) // s[i] is byte type
    }
    s[2] // byte type
    ```

## Substract Two Bytes
* In Go, `byte` is an alias for `uint8`, which is an **unsigned** 8-bit integer.
* This means: When you subtract one byte from another, if the result is **negative**, it wraps around under the rules of unsigned arithmetic, producing a large positive number instead of a negative one.
```go
// e.g.
word := "ig"
fmt.Println(word[1] - word[0]) // answer = 254
// Explanation: 'g' (103) - 'i' (105) in unsigned arithmetic wraps around, giving 256 - 2 = 254.

// Solution: to properly get negative value, convert each byte to int before subtraction
fmt.Println(int(word[1]) - int(word[0])) // answer = -2

// Note: Go handles character literal (i.e. not indexing into string) as rune
fmt.Println('g' - 'i') // answer = -2
```

## String Concatenation
### `+` Operator For Small Concatenations
- For a small number of strings, use `+` to concatenate strings.
    - Pros: Simple and convenient
    - Cons: Inefficient. Avoid using `+` in a loop because it can lead to excessive memory allocation and copying. Each concatenation with `+` creates a new string because strings in Go are immutable.
```go
concat := "a" + "b" + "c"
```
### `strings.Builder` for Efficient Large Concatenations
- For a **large** number of strings, use `strings.Builder` to efficiently concatenate strings.
    - Pros: Efficient. Using `strings.Builder` in a loop. It avoids the immutability overhead by working with a mutable buffer underneath. It minimizes memory copying and allocations, which makes it more efficient.
    - Cons: More complex than using `+`.
```go
var builder strings.Builder
for _, str := range []string{"a", "b", "c", "d"} {
    builder.WriteString(str)
}
concat := builder.String()
```

## Convert Single Byte to String
- Directly call `string()`
```go
x := 'a'
str_x := string(x)
```

## Conver Multiple Bytes to String
- Put all bytes into a `[]byte` and call `string()`
```go
x := []byte{'a', 'b', 'c'}
str_x := string(x)
```