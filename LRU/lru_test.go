package main

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)

	// 插入兩個值
	cache.Put(1, 1)
	cache.Put(2, 2)

	// 測試 Get 方法
	if cache.Get(1) != 1 {
		t.Errorf("Expected 1, got %d", cache.Get(1))
	}

	// 超過容量，插入新值，1應被保留，2被移除
	cache.Put(3, 3)
	if cache.Get(2) != -1 {
		t.Errorf("Expected -1, got %d", cache.Get(2))
	}

	// 測試新插入的值是否存在
	if cache.Get(3) != 3 {
		t.Errorf("Expected 3, got %d", cache.Get(3))
	}

	// 插入另一個新值，1應被移除
	cache.Put(4, 4)
	if cache.Get(1) != -1 {
		t.Errorf("Expected -1, got %d", cache.Get(1))
	}

	// 測試其他值是否正常
	if cache.Get(3) != 3 {
		t.Errorf("Expected 3, got %d", cache.Get(3))
	}
	if cache.Get(4) != 4 {
		t.Errorf("Expected 4, got %d", cache.Get(4))
	}
}

func TestLRUCacheUpdate(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	// 更新已存在的值
	cache.Put(1, 10)
	if cache.Get(1) != 10 {
		t.Errorf("Expected 10, got %d", cache.Get(1))
	}

	// 更新後仍在最前端
	cache.Put(3, 3)
	if cache.Get(2) != -1 {
		t.Errorf("Expected -1, got %d", cache.Get(2))
	}
}

func TestLRUCacheEdgeCases(t *testing.T) {
	cache := Constructor(1)
	cache.Put(1, 1)
	cache.Put(2, 2)

	// 容量1時，測試是否移除正確
	if cache.Get(1) != -1 {
		t.Errorf("Expected -1, got %d", cache.Get(1))
	}
	if cache.Get(2) != 2 {
		t.Errorf("Expected 2, got %d", cache.Get(2))
	}

	// 重複插入相同的鍵
	cache.Put(2, 20)
	if cache.Get(2) != 20 {
		t.Errorf("Expected 20, got %d", cache.Get(2))
	}
}

func TestLRUCacheLargeCapacity(t *testing.T) {
	cache := Constructor(1000)
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}
	for i := 0; i < 1000; i++ {
		if cache.Get(i) != i {
			t.Errorf("Expected %d, got %d", i, cache.Get(i))
		}
	}

	// 插入超過容量的元素
	cache.Put(1000, 1000)
	if cache.Get(0) != -1 {
		t.Errorf("Expected -1, got %d", cache.Get(0))
	}
	if cache.Get(1000) != 1000 {
		t.Errorf("Expected 1000, got %d", cache.Get(1000))
	}
}
