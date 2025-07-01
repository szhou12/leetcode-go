package leetcode

// Use dummy head and dummy tail to simplify the node insertion and removal

// Doubly Linked-List
type Node struct {
	key int
	val int
	prev *Node
	next *Node
}

type LRUCache struct {
	nodes map[int]*Node
	cap int
	head *Node
	tail *Node
}

func Constructor(capacity int) LRUCache {
	// order: head -> least recently used -> most recently used -> tail
	head := &Node{key: -1, val:-1}
	tail := &Node{key: -1, val: -1}
	// Remember to connect head and tail first!
	head.next = tail
	tail.prev = head

	return LRUCache{
		nodes: make(map[int]*Node),
		cap: capacity,
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
	// if key exists, change node to tail and return its value
	if node, ok := this.nodes[key]; ok {
		this.remove(node)
		this.insert(node)
		return node.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int)  {
	// if key exists, update value and change node to tail as recently used
	if node, ok := this.nodes[key]; ok {
		node.val = value
		this.remove(node)
		this.insert(node)
		return 
	}

	// if not yet exists, create a new node
	newNode := &Node{key: key, val: value}
	this.nodes[key] = newNode
	this.insert(newNode)
	
	// if full, evict the least recently used node
	if len(this.nodes) > this.cap {
		this.evict()
	}
}

// <---> [tail.prev] <---> [tail]
func (this *LRUCache) insert(node *Node) {
	// connect the node to the list
	node.prev = this.tail.prev
	node.next = this.tail

	// redirect existing nodes' connections
	this.tail.prev.next = node
	this.tail.prev = node
}

// <---> [node.prev] <---> [node] <---> [node.next] <--->
func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev	
}

func (this *LRUCache) evict() {
	// remove the least recently used node aka head.next
	lruNode := this.head.next
	this.remove(lruNode)
	delete(this.nodes, lruNode.key)
}


