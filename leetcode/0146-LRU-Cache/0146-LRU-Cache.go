package leetcode

type LRUCache struct {
	head, tail *Node
	keys       map[int]*Node
	cap        int
}

// Doubly Linked-List
type Node struct {
	key, val   int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		keys: make(map[int]*Node),
		cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// case 1: key already exists: update value, move node to front
	// case 2: key not yet exits: create the node
	if node, ok := this.keys[key]; ok {
		node.val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node = &Node{
			key: key,
			val: value,
		}
		this.keys[key] = node
		this.Add(node)
	}
	// remove least recently used if exceeds capacity
	if len(this.keys) > this.cap {
		delete(this.keys, this.tail.key)
		this.Remove(this.tail)
	}
}

func (this *LRUCache) Add(node *Node) {
	node.prev = nil
	node.next = this.head

	if this.head != nil {
		this.head.prev = node
	}
	this.head = node

	if this.tail == nil {
		this.tail = node
		this.tail.next = nil
	}
}

func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		this.head = node.next
		if node.next != nil {
			node.next.prev = nil
		}
		node.prev = nil
		return
	}
	if node == this.tail {
		this.tail = node.prev
		node.prev.next = nil
		node.prev = nil
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
