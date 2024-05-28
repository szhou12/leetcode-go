package leetcode

type MyQueue struct {
	// imports elements from outside world
	// exports elements ONLY to master
	buffer []int
	// imports elements ONLY from buffer
	// exports elemnts to outside world
	master []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.buffer = append(this.buffer, x)
}

func (this *MyQueue) Pop() int {
	res := this.Peek() // import elements from buffer to master if necessary
	this.master = this.master[:len(this.master)-1]
	return res
}

func (this *MyQueue) Peek() int {
	// Prioritize emptying elements already in master
	// Only import elements from buffer when master has no elements
	// because elements already in master are ALWAYS older than elements in buffer
	if len(this.master) == 0 {
		for len(this.buffer) > 0 {
			this.master = append(this.master, this.buffer[len(this.buffer)-1])
			this.buffer = this.buffer[:len(this.buffer)-1]
		}
	}
	return this.master[len(this.master)-1]

}

func (this *MyQueue) Empty() bool {
	return len(this.master) == 0 && len(this.buffer) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
