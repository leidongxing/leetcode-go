package stack

type MyQueue struct {
	stackPush []int
	stackPop  []int
}

func (this *MyQueue) pushToPop() {
	if len(this.stackPop) <= 0 {
		for _, val := range this.stackPush {
			this.stackPop = append(this.stackPop, val)
		}
		this.stackPush = nil
	}
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stackPush = append(this.stackPush, x)
	this.pushToPop()
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.stackPop) <= 0 && len(this.stackPush) <= 0 {
		return 0
	}
	this.pushToPop()
	ret := this.stackPop[0]
	this.stackPop = this.stackPop[1:]
	return ret
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.stackPop) <= 0 && len(this.stackPush) <= 0 {
		return 0
	}
	this.pushToPop()
	return this.stackPop[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return (len(this.stackPop) <= 0 && len(this.stackPush) <= 0)
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
