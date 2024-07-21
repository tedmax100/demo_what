package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// cadence 用於控制節奏的條件變數
var cadence = sync.NewCond(&sync.Mutex{})

func init() {
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()
}

// takeStep 模擬每一步操作
func takeStep() {
	cadence.L.Lock()
	cadence.Wait()
	cadence.L.Unlock()
}

// tryDir 允许一個人嘗試向某個方向移動並返回它們是否成功
// 使用 priority 值來增加和減少方向計數器，確保高優先級的處理程序優先獲取資源。這樣可以有效地減少碰撞和活鎖情況。
func tryDir(dirName string, dir *int32, out *bytes.Buffer, priority int) bool {
	fmt.Fprintf(out, " %v (priority %d)", dirName, priority)
	atomic.AddInt32(dir, int32(priority))
	takeStep()
	if atomic.LoadInt32(dir) == int32(priority) {
		fmt.Fprint(out, ". Success!")
		return true
	}
	takeStep()
	atomic.AddInt32(dir, -int32(priority))
	return false
}

var left, right int32

// tryLeft 嘗試向左移動
func tryLeft(out *bytes.Buffer, priority int) bool { return tryDir("left", &left, out, priority) }

// tryRight 嘗試向右移動
func tryRight(out *bytes.Buffer, priority int) bool { return tryDir("right", &right, out, priority) }

// walk 函數模擬一個人在走廊中行走的過程
func walk(walking *sync.WaitGroup, name string, priority int) {
	var out bytes.Buffer
	defer func() { fmt.Println(out.String()) }()
	defer walking.Done()
	fmt.Fprintf(&out, "%v is trying to scoot:", name)
	for i := 0; i < 5; i++ {
		if tryLeft(&out, priority) || tryRight(&out, priority) {
			return
		}
	}
	fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
}

func main() {
	// 創建了兩個 goroutine 分別模擬 "Alice" 和 "Barbara" 的行走過程，並等待它們完成
	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice", 2)   // Alice 具有較高優先級
	go walk(&peopleInHallway, "Barbara", 1) // Barbara 具有較低優先級
	peopleInHallway.Wait()
}
