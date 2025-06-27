package leetcode

type Node struct {
    key int
    value int
    freq int
    prev *Node
    next *Node
}

func newNode(key, value, freq int) *Node {
    return &Node{key: key, value: value, freq: freq}
}

func dummyNode() *Node {
    return newNode(-1, -1, -1)
}

type LinkedList struct {
    head *Node
    tail *Node
    length int
}

func newLinkedList() *LinkedList {
	// order: least recently used --> most recently used
    head, tail := dummyNode(), dummyNode()
    head.next = tail
    tail.prev = head
    return &LinkedList{head: head, tail: tail}
}

// <---> [node.prev] <---> [node] <---> [node.next] <--->
func (l *LinkedList) remove(node *Node) {
    node.prev.next = node.next
    node.next.prev = node.prev
    l.length--
}

// <---> [tail.prev] <---> [tail]
func (l *LinkedList) insert(node *Node) {
    tail := l.tail
    // connects new node
    node.prev = tail.prev
    node.next = tail
    // redirect connections of existing nodes
    tail.prev.next = node
    tail.prev = node
    
    l.length++
}

type LFUCache struct {
	cap int
    minFreq int // least frequently used counter
    nodes map[int]*Node // {key: (key, value, freq, prev, next)}
    lists map[int]*LinkedList // {freq: key1 <--> key2 <--> key3 }
}


func Constructor(capacity int) LFUCache {
	return LFUCache{
        cap: capacity,
        nodes: make(map[int]*Node),
        lists: make(map[int]*LinkedList),
    }
}


func (this *LFUCache) Get(key int) int {
	node, exist := this.nodes[key]
    if !exist {
        return -1
    }
    this.updateFreq(node)
    return node.value
}


func (this *LFUCache) Put(key int, value int)  {
	// Update the value of the key if present, or inserts the key if not already present.
	if node, ok := this.nodes[key]; ok {
        node.value = value
        this.updateFreq(node)
        return
    }

	// when cap is reached, 
	//remove the least recently used node from the least frequently used linked list
	if this.isFull() {
        this.evict()
    }

    this.minFreq = 1
    node := newNode(key, value, 1)
    this.insert(node)
}

func (this *LFUCache) updateFreq(node *Node) {
	list := this.lists[node.freq]
	list.remove(node) // take out the node from the linked list and add it to the last to become most recently used in its group

	// if the input node is from the least frequently used (minFreq) group and is the only node in the group
	// then we need to increment the minFreq counter
	if list.length == 0 && node.freq == this.minFreq {
		this.minFreq++
	}

	node.freq++
	this.insert(node)
}

func (this *LFUCache) insert(node *Node) {
    list, exist := this.lists[node.freq]
    if !exist {
        list = newLinkedList()
    }

    list.insert(node) // append to the last as most recently used
    this.lists[node.freq] = list
    this.nodes[node.key] = node
}

func (this *LFUCache) evict() {
    minList := this.lists[this.minFreq]
    lfuNode := minList.head.next
    minList.remove(lfuNode)
    delete(this.nodes, lfuNode.key)
}

func (this *LFUCache) isFull() bool {
    return len(this.nodes) == this.cap
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */