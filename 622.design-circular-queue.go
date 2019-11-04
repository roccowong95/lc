/*
 * @lc app=leetcode id=622 lang=golang
 *
 * [622] Design Circular Queue
 */
 type MyCircularQueue struct {
	cap  int
	len  int
	head int
	tail int
	dat  []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		cap: k,
		dat: make([]int, k),
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	// make sure dat doesn't overlap
	if this.cap == this.len {
		return false
	}
	this.len++
	// if buffer is empty, do not incr this.tail.
	if this.len == 1 {
		this.dat[this.tail] = value
		return true
	}
	// move tail forward
	this.tail++
	if this.tail == this.cap {
		// if tail is at the end of array,
		// move it back to zero.
		this.tail = 0
	}
	this.dat[this.tail] = value
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.len == 0 {
		return false
	}
	this.len--
	if this.len != 0 {
		this.head++
	}
	if this.head == this.cap {
		// move head back to zero if it's beyond boundary.
		this.head = 0
	}
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.dat[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.dat[this.tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.len == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.len == this.cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */