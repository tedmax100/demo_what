package main

// hash table
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkNode
	Head, Tail *DLinkNode
}

// double linked list
type DLinkNode struct {
	key, value int
	Pre, Next  *DLinkNode
}

// InitDlinkNode
func InitDlinkNode(key, value int) *DLinkNode {
	return &DLinkNode{key, value, nil, nil}
}

// init hash table and double linked list
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		0,
		capacity,
		map[int]*DLinkNode{},
		InitDlinkNode(0, 0),
		InitDlinkNode(0, 0),
	}
	l.Head.Next = l.Tail
	l.Tail.Pre = l.Head
	return l
}

func (lru *LRUCache) Get(key int) int {
	//not hit，return -1
	if _, ok := lru.cache[key]; !ok {
		return -1
	}
	//cache hit，update access summary and return value
	node := lru.cache[key]
	lru.UpdateToHead(node)
	return node.value
}

func (lru *LRUCache) Put(key int, value int) {
	if _, ok := lru.cache[key]; !ok {
		node := InitDlinkNode(key, value)
		for lru.size >= lru.capacity {
			lru.DeleteLast()
		}
		lru.cache[key] = node
		lru.InsertNewHead(node)
	} else {
		//update access summary
		node := lru.cache[key]
		node.value = value
		lru.UpdateToHead(node)
	}
}

// update access summary
func (lru *LRUCache) UpdateToHead(node *DLinkNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	temp := lru.Head.Next
	lru.Head.Next = node
	node.Pre = lru.Head
	node.Next = temp
	temp.Pre = node

}

// Remove the least recently used element
func (lru *LRUCache) DeleteLast() {
	node := lru.Tail.Pre
	lru.Tail.Pre = node.Pre
	node.Pre.Next = node.Next
	node.Pre = nil
	node.Next = nil
	lru.size--
	delete(lru.cache, node.key)
}

// add new element and size + 1
func (lru *LRUCache) InsertNewHead(node *DLinkNode) {
	temp := lru.Head.Next
	lru.Head.Next = node
	node.Pre = lru.Head
	temp.Pre = node
	node.Next = temp
	lru.size++
}
