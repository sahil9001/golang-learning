package main

import "fmt"

func main() {
    // 1. Create a slice with make (length 2, capacity 5)
    nums := make([]int, 2, 5)
    fmt.Printf("slice: %v, len=%d, cap=%d\n", nums, len(nums), cap(nums))
    // Initially: slice: [0 0], len=2, cap=5

    // 2. Create a map with make (string → int)
    counts := make(map[string]int)
    fmt.Println("map before:", counts)
    counts["apple"] = 3
    fmt.Println("map after:", counts)
    // map before: map[], map after: map[apple:3]

    // 3. Create a buffered channel with make (int channel with buffer size 3)
    ch := make(chan int, 3)
    ch <- 10
    ch <- 20
    fmt.Printf("channel buffered size: %d (len)\n", len(ch))
    // Channel length: 2 of capacity 3
}