package main

import "fmt"

func main() {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	fmt.Println(lfu.Get(1)) // 输出 1
	lfu.Put(3, 3)           // 移除 key 2
	fmt.Println(lfu.Get(2)) // 输出 -1
	fmt.Println(lfu.Get(3)) // 输出 3
	lfu.Put(4, 4)           // 移除 key 1
	fmt.Println(lfu.Get(1)) // 输出 -1
	fmt.Println(lfu.Get(3)) // 输出 3
	fmt.Println(lfu.Get(4)) // 输出 4
}
