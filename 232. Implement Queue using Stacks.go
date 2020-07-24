// Time complexity: Push - O(n) per operation, Pop - O(1) per operation.
// Push can optimize to O(1) per operation
// Space complexity: O(n)


type MyQueue struct {
    stack []int
    back []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{stack:make([]int, 0), back:make([]int, 0)}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	// make sure all number is push back into stack
    for len(this.back)!=0 {
        val := this.back[len(this.back)-1]
        this.back = this.back[:len(this.back)-1]
        this.stack = append(this.stack, val)
    }
    this.stack = append(this.stack, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	// push all stack number into queue
    for len(this.stack)!=0 {
        val := this.stack[len(this.stack)-1]
        this.stack = this.stack[:len(this.stack)-1]
        this.back = append(this.back, val)
    }
    if len(this.back)==0 {
        return 0
    }
    val := this.back[len(this.back)-1]
    this.back = this.back[:len(this.back)-1]
    return val
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	// push all stack number into queue
    for len(this.stack)!=0 {
        val := this.stack[len(this.stack)-1]
        this.stack = this.stack[:len(this.stack)-1]
        this.back = append(this.back, val)
    }
    if len(this.back)==0 {
        return 0
    }
    return this.back[len(this.back)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    if len(this.stack) == 0 && len(this.back) == 0 {
        return true
    }
    return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */