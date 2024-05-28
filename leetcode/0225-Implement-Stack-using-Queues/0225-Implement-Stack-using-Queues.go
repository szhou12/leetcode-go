package leetcode

// 左、右两个stack背靠背
// Push: 把元素push进右stack
// Pop:
// 	case 1 - 如果左stack非空，直接pop左stack栈顶元素；
//  case 2 - 如果左stack空，把右stack元素全部导入左stack，再pop左stack栈顶元素
type MyStack struct {
}

func Constructor() MyStack {
	
}

func (this *MyStack) Push(x int) {

}

func (this *MyStack) Pop() int {

}

func (this *MyStack) Top() int {

}

func (this *MyStack) Empty() bool {

}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
