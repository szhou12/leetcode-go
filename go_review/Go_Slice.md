# [Slice](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

## Contents
* [Slice vs. Array](#slice-vs-array)
* [Pass into a function](#pass-into-a-function)
* [How to index a slice pointer](#how-to-index-a-slice-pointer)
* [Create slice of slices](#create-slice-of-slices)
* [Sort slice of objects](#sort-slice-of-objects)
* [Sort slice of slice](#sort-slice-of-slice)
* [Slice of Map](#slice-of-map)
* [Slice as a key in map](#slice-as-a-key-in-map)
* [Prepend elements](#prepend-elements)
* [Compare Two Slices Equality](#compare-two-slices-equality)

## Slice vs. Array
[Revisiting Arrays and Slices in Go](https://www.developer.com/languages/arrays-slices-golang/)
* TLDR:
    * The basic difference between a slice and an array is that a slice is a reference to a contiguous segment of an array. 
    * Unlike an **array, which is a value-type**, **slice is a reference type**. 
    * A slice can be a complete array or a part of an array, indicated by the start and end index.
    * A slice, therefore, is also an array that pours a context of dynamism onto the underlying array, which otherwise is a static contiguous memory allocation.
* 个人总结：
    * 简而言之，声明的时候标明元素个数的是array(=定长的slice)，没标明元素个数是slice。
```go
nums := [3]int{1, 2, 3} // array
nums := []int{1, 2, 3} // slice
```

## Pass into a function

[Are slices passed by value?](https://stackoverflow.com/questions/39993688/are-slices-passed-by-value)
* TLDR:
    * Everything in Go is passed by value, slices too. **But a slice value is a header, describing a contiguous section of a backing array, and a slice value only contains a pointer to the array where the elements are actually stored.** The slice value does not include its elements (unlike arrays).
    * So when you pass a slice to a function, a copy will be made from this header, including the pointer, which will point to the same backing array. Modifying the elements of the slice implies modifying the elements of the backing array, and so all slices which share the same backing array will "observe" the change.
* 个人总结：
    * 把一个slice传入function的确是pass by value, 也就是进行了一次copy。
    * slice每个元素不是存的值本身，而是一个pointer。底层有一个开辟的连续物理空间存放slice的所有元素的值 (backing array)。每一个pointer指向物理空间的一个值。
    * 所以，虽然slice传入function是pass by value，但是copy的实际上是每一个pointer，对应的pointer指向底层物理空间中同一个值。所以，function内对slice的CRUD操作也会反映在原来的slice上。
    * 相反的，array每个元素是存的值本身，所以，array传入function后，就生成了另一个array，function内对array的CRUD操作不会反映在原来的array上。

![slice_into_fcn](https://user-images.githubusercontent.com/35708194/231896890-7f5ba00d-6b58-4d9f-993d-aadc601082e4.png)
```go
// slice 传入 function后所作修改会反映在原来的slice上
func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := change(s1)
	fmt.Println(s2) // [10, 2, 3, 4]
	fmt.Println(s1) // [10, 2, 3, 4] 原slice也变了
}

func change(nums []int) []int {
	nums[0] = 10
	return nums
}

// array 传入 function后所作修改不会反映在原来的array上
func main() {
	a1 := [4]int{1, 2, 3, 4}
	a2 := change(a1)
	fmt.Println(a2) // [10, 2, 3, 4]
	fmt.Println(a1) // [1, 2, 3, 4] 原array没变
}

func change(nums [4]int) [4]int {
	nums[0] = 10
	return nums
}
```

## How to index a slice pointer
[why is indexing on the slice pointer not allowed in golang - Stack Overflow](https://stackoverflow.com/questions/38468258/why-is-indexing-on-the-slice-pointer-not-allowed-in-golang)
* TLDR:
    * To give some practical reason, this is likely due to the fact that the pointer doesn't point to the beginning of an array (like the block of memory.)
```go
(*nums)[i]
```

## Create slice of slices
[Go slice - working with slices in Golang](https://zetcode.com/golang/slice/)
```go
inputSlice := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}} //都要用花括号！
```

## Sort slice of objects
[sort package](https://pkg.go.dev/sort#Slice)

[The 3 ways to sort in Go](https://yourbasic.org/golang/how-to-sort-in-go/)
* 个人总结：
    * `[]int`, `[]string`, `[]float64` 分别有专属的sort API: `sort.Ints()`, `sort.Strings()`, `sort.Float64s()`. 都是增序排列。
    * 自 go 1.21 起, 新增 API `slices.Sort()` [Go Slices Sort](https://pkg.go.dev/slices#Sort)。更快！
    * 常用的可自定义sort function: `sort.Slice()`
```go
sort.Slice(people, func(i, j int) bool { 
    return people[i].Name < people[j].Name
})
```

## Sort slice of slice
[golang sort slices of slice by first element - Stack Overflow](https://stackoverflow.com/questions/55360091/golang-sort-slices-of-slice-by-first-element)

[Checking the equality of two slices - Stack Overflow](https://stackoverflow.com/questions/15311969/checking-the-equality-of-two-slices)


## Slice of Map
[Slice of Map in Go (Golang)](https://golangbyexample.com/slice-map-golang/)
```go
next := make([]map[int]bool, n)
```

## Slice as a key in map
[Slice as a key in map - Stack Overflow](https://stackoverflow.com/questions/20297503/slice-as-a-key-in-map)
* 个人总结：
    * 只能用定长的 array 作为 key (e.g. `[2]int{1,2}`), 不能用 slice 作为 key (e.g. `[]int{1,2}`)
```go
m := make(map[[2]int]bool)
```

## Prepend elements
[How to prepend int to slice - Stack Overflow](https://stackoverflow.com/questions/53737435/how-to-prepend-int-to-slice)
```go
nums = append([]int{1}, nums...)
```

## Compare Two Slices Equality
[Checking the equality of two slices](https://stackoverflow.com/questions/15311969/checking-the-equality-of-two-slices)
1. 调用 `reflect.DeepEqual()`
```go
a := []int{4,5,6}
b := []int{4,5,6}
reflect.DeepEqual(a, b)
```
2. 自己写一个helper function
```go
func testEq(a, b []int) bool {
    // Check 长度
    if len(a) != len(b) {
        return false
    }
    // Check 元素
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
```