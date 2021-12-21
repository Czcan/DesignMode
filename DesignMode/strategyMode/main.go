package main

func main() {
	lfu := &lfu{}
	cache := initCache(lfu)

	cache.add("a", "0")

	cache.add("b", "1")

	cache.add("c", "2")

	lru := &lru{}
	cache.setEvictionAlgo(lru)
	cache.add("d", "3")

	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("d", "4")

}
