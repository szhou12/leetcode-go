# Golang Review
本文主要进行针对在LeetCode刷题过程中所遇到的有关Go语言知识点的总结和补充

## Contents
* [Print](#print)
    * [Print byte as character](#print-byte-as-character-not-value)
    * [Print variable type](#print-variable-type)
* [Pass by Value (值传递)](#pass-by-value-值传递)
* [Slice](#slice)
    * [Slice vs. Array](#slice-vs-array)
    * [Pass into a function](#pass-into-a-function)
    * [How to index a slice pointer](#how-to-index-a-slice-pointer)
    * [Create slice of slices](#create-slice-of-slices)
    * [Sort slice of objects](#sort-slice-of-objects)
    * [Sort slice of slice](#sort-slice-of-slice)
    * [Slice of Map](#slice-of-map)
    * [Slice as a key in map](#slice-as-a-key-in-map)
    * [Prepend elements](#prepend-elements)

## Print
### Print byte as character, not value
```go
fmt.Printf("char: %c ", char)
```

### Print variable type
[Print type of variable in Go (Golang)](https://gosamples.dev/print-type/)
```go
fmt.Printf("type: %T ", var)
```

## Pass by Value (值传递)
[go语言参数传递到底是传值还是传引用？](https://segmentfault.com/a/1190000037763005)
* 个人总结:
    * Go就是值传递，可以确认的是Go语言中所有的传参都是值传递 (传值)，都是一个副本，一个拷贝。因为拷贝的内容有时候是非引用类型（int、string、struct等这些），这样在函数中就无法修改原内容数据；有的是引用类型（pointer、map、slice、chan等这些），这样就可以修改原内容数据 (ie. 拷贝的是原内容数据的指针)。
    * 但是要注意！如果函数里要对slice进行 `append()` 操作，则必须传引用 `fcn(*nums)`，直接传slice `fcn(nums)` 没用。 因为，`append()` 和 `copy()` 一样会深拷贝，即生成新的一个底层数组 (backing array)，所以，指针地址也会改变。也就是说，function内部一旦发生 `append()`，它赋值给的slice所指向的backing array 就已经不再是输入时候的slice所指向的backing array了。
    * 所以，一个function的argument如果需要传入 pointer、map、slice、chan，除非function内部会发生深拷贝 (eg. slice `append()`, `copy()`). 可以直接传入该数据结构，不用传入其指针. e.g. `helper(map[int]bool)` 而不用 `helper(*map[int]bool)`，在main function中再操作 原map 就会看到被helper修改过内容的map。
```go
func main() {
    var args = [][]int64{{1}, {1, 2, 3}}
	fmt.Printf("切片args的地址： %p \n", args)
	fmt.Printf("切片args第一个元素的地址： %p \n", &args[0])
	modifiedNumber(args)
	fmt.Println(args)
}

func modifiedNumber(args [][]int64) {
	fmt.Printf("形参切片的地址 %p \n", args)
	fmt.Printf("形参切片args第一个元素的地址： %p \n", &args[0])
	fmt.Printf("直接对形参切片args取地址%v \n", &args)
	args = append(args, []int64{5, 6}) //args地址会改变
	fmt.Printf("更改后，形参切片的地址 %p \n", args) 
	//args[0] = []int64{10} //args地址不会改变
}
```


## Slice
### Slice vs. Array
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

### Pass into a function

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

### How to index a slice pointer
[why is indexing on the slice pointer not allowed in golang - Stack Overflow](https://stackoverflow.com/questions/38468258/why-is-indexing-on-the-slice-pointer-not-allowed-in-golang)
* TLDR:
    * To give some practical reason, this is likely due to the fact that the pointer doesn't point to the beginning of an array (like the block of memory.)
```go
(*nums)[i]
```

### Create slice of slices
[Go slice - working with slices in Golang](https://zetcode.com/golang/slice/)
```go
inputSlice := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}} //都要用花括号！
```

### Sort slice of objects
[sort package](https://pkg.go.dev/sort#Slice)

[The 3 ways to sort in Go](https://yourbasic.org/golang/how-to-sort-in-go/)
* 个人总结：
    * `[]int`, `[]string`, `[]float64` 分别有专属的sort API: `sort.Ints()`, `sort.Strings()`, `sort.Float64s()`. 都是增序排列。
    * 常用的可自定义sort function: `sort.Slice()`
```go
sort.Slice(people, func(i, j int) bool { 
    return people[i].Name < people[j].Name
})
```

### Sort slice of slice
[golang sort slices of slice by first element - Stack Overflow](https://stackoverflow.com/questions/55360091/golang-sort-slices-of-slice-by-first-element)

[Checking the equality of two slices - Stack Overflow](https://stackoverflow.com/questions/15311969/checking-the-equality-of-two-slices)


### Slice of Map
[Slice of Map in Go (Golang)](https://golangbyexample.com/slice-map-golang/)
```go
next := make([]map[int]bool, n)
```

### Slice as a key in map
[Slice as a key in map - Stack Overflow](https://stackoverflow.com/questions/20297503/slice-as-a-key-in-map)
* 个人总结：
    * 只能用定长的 array 作为 key (e.g. `[2]int{1,2}`), 不能用 slice 作为 key (e.g. `[]int{1,2}`)
```go
m := make(map[[2]int]bool)
```

### Prepend elements
[How to prepend int to slice - Stack Overflow](https://stackoverflow.com/questions/53737435/how-to-prepend-int-to-slice)
```go
nums = append([]int{1}, nums...)
```

## Priority Queue
### Max/Min Heap
[Is there a more generic way to implement 2 kinds of Heap (Max and Min) in Go Lang](https://stackoverflow.com/questions/23580285/is-there-a-more-generic-way-to-implement-2-kinds-of-heap-max-and-min-in-go-lan)
```go
type PQ []int

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i] < pq[j] } // MinHeap
func (pq PQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

// Push and Pop use pointer receivers because they modify the slice's length, not just its contents.
func (pq *PQ) Push(x interface{}) {
    // NOTE: 这里为什么要手动解引用 *pq？ 这是因为append()只作用于slice而不是slice 的 pointer; 同时, append()也对slice进行了长度改变
    *pq = append(*pq, x.(int))
}

func (pq *PQ) Pop() interface{} {
    // NOTE: 这里为什么要手动解引用 *pq？ 这是因为len()和slice indexing只作用于slice而不是slice pointer; 同时, *pq = (*pq)[:n-1]也对slice进行了长度改变
    n := len(*pq)
    temp := (*pq)[n-1]
    *pq = (*pq)[:n-1]
    return temp
}

// Top()
(*minHeap)[0]
```