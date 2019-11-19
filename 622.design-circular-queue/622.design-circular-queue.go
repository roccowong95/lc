/*
 * @lc app=leetcode id=622 lang=golang
 *
 * [622] Design Circular Queue

Design your implementation of the circular queue. The circular queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle and the last position is connected back to the first position to make a circle. It is also called "Ring Buffer".

One of the benefits of the circular queue is that we can make use of the spaces in front of the queue. In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space in front of the queue. But using the circular queue, we can use the space to store new values.

Your implementation should support following operations:

	MyCircularQueue(k): Constructor, set the size of the queue to be k.
	Front: Get the front item from the queue. If the queue is empty, return -1.
	Rear: Get the last item from the queue. If the queue is empty, return -1.
	enQueue(value): Insert an element into the circular queue. Return true if the operation is successful.
	deQueue(): Delete an element from the circular queue. Return true if the operation is successful.
	isEmpty(): Checks whether the circular queue is empty or not.
	isFull(): Checks whether the circular queue is full or not.

Example:
	MyCircularQueue circularQueue = new MyCircularQueue(3); // set the size to be 3
	circularQueue.enQueue(1); // return true
	circularQueue.enQueue(2); // return true
	circularQueue.enQueue(3); // return true
	circularQueue.enQueue(4); // return false, the queue is full
	circularQueue.Rear(); // return 3
	circularQueue.isFull(); // return true
	circularQueue.deQueue(); // return true
	circularQueue.enQueue(4); // return true
	circularQueue.Rear(); // return 4

Note:
	All values will be in the range of [0, 1000].
	The number of operations will be in the range of[1, 1000].
	Please do not use the built-in Queue library.

*/
package main

// @lc code=start

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
	if this.len == this.cap {
		return false
	}
	if this.len != 0 {
		this.tail++
	}
	this.len++
	if this.tail == this.cap {
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
	if this.len != 1 {
		this.head++
	}
	this.len--
	if this.head == this.cap {
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

// @lc code=end

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
