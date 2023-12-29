# [String](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

## Contents
* [Strings are Immutable](#strings-are-immutable)
* [Sort a String](#sort-a-string)
* [Reverse a String](#reverse-a-string)
* [rune vs. byte](#rune-vs-byte)

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
    ruens := []rune(s)
    // Step 2: Swap the runes
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1{
        runes[i], runes[j] = runes[j], runes[i]
    }
    // Step 3: Convert back to string
    return string(ruens)
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
