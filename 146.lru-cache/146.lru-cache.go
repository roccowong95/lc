/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */
type LRUCache struct {
	cap int
	len int
	m map[int]*Node
	head *Node
	tail *Node
}

type Node struct {
	k int
	v int
	next *Node
	prev *Node
}


func Constructor(capacity int) LRUCache {
	tail := &Node{}
	head := &Node{next: tail}
	tail.prev = head
	return LRUCache{
		cap: capacity,
		m: make(map[int]*Node),
		head: head,
		tail: tail,
	}
}


func (this *LRUCache) Get(key int) int {
	// fmt.Printf("get %d, que: ", key)
	// defer func(c *LRUCache) {
	// 	p := c.head.next
	// 	for p.next != nil {
	// 		fmt.Printf("%d:%d ", p.k, p.v)
	// 		p = p.next
	// 	}
	// 	fmt.Println()
	// }(this)

	v, ok := this.m[key]
	if !ok || nil==v {
		return -1
	}
	this.putHead(v)
	return v.v
}


func (this *LRUCache) Put(key int, value int)  {
	// fmt.Printf("put %d:%d, que: ", key, value)
	// defer func(c *LRUCache) {
	// 	p := c.head.next
	// 	for p.next != nil {
	// 		fmt.Printf("%d:%d ", p.k, p.v)
	// 		p = p.next
	// 	}
	// 	fmt.Println()
	// }(this)

	v, ok := this.m[key]
	if ok {
		this.putHead(v)
		this.m[key].v = value
		return
	}

	n := &Node{v: value, k: key}
	this.putHead(n)
	this.m[key] = n

	if this.len < this.cap {
		// n := &Node{v: value, prev: this.head, next: this.head.next}
		this.len++
		return
	}

	last := this.tail.prev
	last.prev.next = last.next
	last.next.prev = last.prev
	delete(this.m, last.k)
}

func (this *LRUCache) putHead(n *Node) {
	if nil != n.prev {
		n.prev.next = n.next
	}
	if nil != n.next {
		n.next.prev = n.prev
	}
	n.next = this.head.next
	n.prev = this.head
	n.next.prev = n
	n.prev.next = n
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

