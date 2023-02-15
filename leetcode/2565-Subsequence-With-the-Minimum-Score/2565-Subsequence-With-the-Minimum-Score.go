package leetcode

/**
1. 翻译题目: 删除t的一个 subarray = t[left:right] (甭管内部怎么删的), 使得剩下的chars能组成s的subseq
2. 根据翻译, 删除的这个subarray越长, 剩下的chars越容易组成s的subseq (eg. 最极端的例子是t删除到什么都不剩了，一个空串肯定是s的subseq)
	--> 单调性 --> 暗示可以用 binary search 猜答案
3. 根据单调性, t[left:right]之间的chars不删白不删, 删的越多t就越容易成为s的subseq. 所以不用管那么多, 不需要在它内部额外操作, 就给全删了吧

t:   x x x  {y y y}  z z z
            -------
*/
func minimumScore(s string, t string) int {

}
