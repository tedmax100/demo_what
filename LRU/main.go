package main

func main() {
	cache := Constructor(5)

	for i := 0; i < 10; i++ {
		cache.Put(i, i)
	}
}
