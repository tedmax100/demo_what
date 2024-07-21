package main

import (
	"bytes"
	"fmt"
	"math/rand"
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

// backoff 等待一段隨機時間
func backoff() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
}

// walk 函數模擬一個人在走廊中行走的過程
func walk(walking *sync.WaitGroup, name string) {
	var out bytes.Buffer
	defer func() { fmt.Println(out.String()) }()
	defer walking.Done()
	fmt.Fprintf(&out, "%v is trying to scoot:", name)
	for i := 0; i < 5; i++ {
		if tryLeft(&out) || tryRight(&out) {
			return
		}
		backoff() // 添加backoff time
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
