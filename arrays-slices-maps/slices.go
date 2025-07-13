package main

import "fmt"

func main() {
    // 1. Create a slice literal with 3 elements
    nums := []int{10, 20, 30}
    fmt.Printf("nums: %v, len=%d, cap=%d\n", nums, len(nums), cap(nums))

    // 2. Create a slice using make (length 2, capacity 4)
    s := make([]int, 2, 4)
    s[0], s[1] = 1, 2
    fmt.Printf("s: %v, len=%d, cap=%d\n", s, len(s), cap(s))

    // 3. Slice a portion of nums
    sub := nums[1:3] // elements at index 1 and 2
    fmt.Printf("sub: %v, len=%d, cap=%d (shares backing array)\n", sub, len(sub), cap(sub))

    // 4. Modify sub; it affects nums
    sub[0] = 99
    fmt.Println("after sub[0]=99 → nums:", nums)

    // 5. Append a new element to s
    s = append(s, 3, 4)
    fmt.Printf("after append, s: %v, len=%d, cap=%d\n", s, len(s), cap(s))

    // 6. Appending beyond capacity creates a new backing array
    s = append(s, 5)
    fmt.Printf("after exceed cap, s: %v, len=%d, cap=%d\n", s, len(s), cap(s))
}