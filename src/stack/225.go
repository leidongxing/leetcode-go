package stack

type MyStack struct {
	Items []int
}

/** Initialize your data structure here. */
func Constructor1() MyStack {
	return MyStack{Items: []int{}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this == nil {
		*this = Constructor1()
	}
	this.Items = append(this.Items, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this == nil {
		*this = Constructor1()
	}
	if !this.Empty() {
		top := this.Items[len(this.Items)-1]

		items := this.Items
		this.Items = items[0 : len(items)-1]

		return top
	}
	return -0
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this == nil {
		*this = Constructor1()
	}
	if !this.Empty() {
		return this.Items[len(this.Items)-1]
	}
	return 0
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this == nil {
		*this = Constructor1()
	}
	if len(this.Items) > 0 {
		return false
	} else {
		return true
	}
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
