package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// cadence 用於控制節奏的條件變數
// 用於控制所有動作之間的節奏。通過定期廣播，它模擬了一個固定的步伐，使所有goroutine以相同的節奏進行。
var cadence = sync.NewCond(&sync.Mutex{})

// 啟動一个 goroutine，每隔 1 毫秒廣播一次條件變數
// 從而喚醒所有等待在 cadence 上的 goroutine。
func init() {
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()
}

// takeStep 模擬每一步操作
// 用於等待和釋放 cadence
// goroutine 在需要同步的地方會調用 cadence.Wait()，進行等待，直到收到廣播信號。
func takeStep() {
	cadence.L.Lock()
	cadence.Wait()
	cadence.L.Unlock()
}

// tryDir 允许一個人嘗試向某個方向移動並返回它們是否成功
// 使用 atomic 來保證對共享變數的操作是原子的
func tryDir(dirName string, dir *int32, out *bytes.Buffer) bool {
	fmt.Fprintf(out, " %v", dirName)
	atomic.AddInt32(dir, 1)
	takeStep()
	if atomic.LoadInt32(dir) == 1 {
		fmt.Fprint(out, ". Success!")
		return true
	}
	takeStep()
	atomic.AddInt32(dir, -1)
	return false
}

var left, right int32

// tryLeft 嘗試向左移動
func tryLeft(out *bytes.Buffer) bool { return tryDir("left", &left, out) }

// tryRight 嘗試向右移動
func tryRight(out *bytes.Buffer) bool { return tryDir("right", &right, out) }

// walk 函數模擬一個人在走廊中行走的過程。
// 它會嘗試多次向左或向右移動，如果成功則返回，否則在所有嘗試失敗後放棄。
func walk(walking *sync.WaitGroup, name string) {
	var out bytes.Buffer
	defer func() { fmt.Println(out.String()) }()
	defer walking.Done()
	fmt.Fprintf(&out, "%v is trying to scoot:", name)
	for i := 0; i < 500; i++ {
		if tryLeft(&out) || tryRight(&out) {
			return
		}
	}
	fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
}

func main() {
	// 創建了兩個 goroutine 分別模擬 "Alice" 和 "Barbara" 的行走過程，並等待它們完成
	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
}
