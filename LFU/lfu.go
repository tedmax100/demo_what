package main

// double linked list node
type DLinkNode struct {
	key, value, freq int
	prev, next       *DLinkNode
}

// init double linked list
func InitDLinkNode(key, value int) *DLinkNode {
	return &DLinkNode{
		key:   key,
		value: value,
		freq:  1, // 初始频率为 1
	}
}

// LFU cache
type LFUCache struct {
	capacity int
	size     int
	minFreq  int
	cache    map[int]*DLinkNode
	freqMap  map[int]*DLinkList
}

// double linked list
type DLinkList struct {
	head, tail *DLinkNode
}

// init double linked list
func InitDLinkList() *DLinkList {
	head := &DLinkNode{}
	tail := &DLinkNode{}
	head.next = tail
	tail.prev = head
	return &DLinkList{
		head: head,
		tail: tail,
	}
}

// remove node from list
func (list *DLinkList) removeNode(node *DLinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// add node into head of list
func (list *DLinkList) addToHead(node *DLinkNode) {
	node.next = list.head.next
	node.prev = list.head
	list.head.next.prev = node
	list.head.next = node
}

// init LFU cache
func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		size:     0,
		minFreq:  0,
		cache:    make(map[int]*DLinkNode),
		freqMap:  make(map[int]*DLinkList),
	}
}

// get value
func (this *LFUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.increaseFreq(node)
		return node.value
	}
	return -1
}

// put element
func (this *LFUCache) Put(key, value int) {
	if this.capacity == 0 {
		return
	}

	if node, ok := this.cache[key]; ok {
		node.value = value
		this.increaseFreq(node)
	} else {
		if this.size == this.capacity {
			this.removeMinFreqNode()
		}

		node := InitDLinkNode(key, value)
		this.cache[key] = node
		if _, ok := this.freqMap[1]; !ok {
			this.freqMap[1] = InitDLinkList()
		}
		this.freqMap[1].addToHead(node)
		this.minFreq = 1
		this.size++
	}
}

// increate freq for node
func (this *LFUCache) increaseFreq(node *DLinkNode) {
	freq := node.freq
	this.freqMap[freq].removeNode(node)
	if this.freqMap[freq].head.next == this.freqMap[freq].tail {
		delete(this.freqMap, freq)
		if this.minFreq == freq {
			this.minFreq++
		}
	}

	node.freq++
	if _, ok := this.freqMap[node.freq]; !ok {
		this.freqMap[node.freq] = InitDLinkList()
	}
	this.freqMap[node.freq].addToHead(node)
}

// remove the least freq node
func (this *LFUCache) removeMinFreqNode() {
	list := this.freqMap[this.minFreq]
	node := list.tail.prev
	list.removeNode(node)
	delete(this.cache, node.key)
	this.size--
}
