# [Binary Form](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

## Contents
* [Binary Representation of Integer](#binary-representation-of-integer)



## Binary Representation of Integer
### Integer :arrow_right: Binary String
[Converting from an integer to its binary representation](https://stackoverflow.com/questions/13870845/converting-from-an-integer-to-its-binary-representation)

[Int -> Binary: func FormatInt](https://pkg.go.dev/strconv#FormatInt)
```go
func FormatInt(i int64, base int) string
```
e.g.
```go
n := 123
fmt.Println(strconv.FormatInt(int64(n), 2)) // "1111011"
```

### Binary String :arrow_right: Integer
[Go - convert string which represent binary number into int](https://stackoverflow.com/questions/9271469/go-convert-string-which-represent-binary-number-into-int)

[Binary -> Int: func ParseInt](https://pkg.go.dev/strconv#ParseInt)
```go
func ParseInt(s string, base int, bitSize int) (i int64, err error)
// The bitSize argument specifies the integer type that the result must fit into. 
// Bit sizes 0, 8, 16, 32, and 64 correspond to int, int8, int16, int32, and int64. 
// If bitSize is below 0 or above 64, an error is returned.
```
e.g.
```go
s := "101"
if i, err := strconv.ParseInt(s, 2, 0); err != nil {
    fmt.Println(err)
} else {
    fmt.Println(int(i)) // 5, type: int
    fmt.Printf("type: %T ", i) // type: int64
}
```