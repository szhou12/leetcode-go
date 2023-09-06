# [Pass by Value (值传递)](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

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