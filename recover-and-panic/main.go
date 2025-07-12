package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    // 1. This deferred anonymous function will run if f panics.
    //    It checks recover(), and if there was a panic, prints a message.
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")

    // 2. Call g starting at 0.
    g(0)

    // 3. If g panics and is recovered, we would continue here.
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))  // triggers panic with value of i
    }
    // 4. Defer printing happens when g returns (or panics).
    defer fmt.Println("Defer in g", i)

    fmt.Println("Printing in g", i)

    // 5. Recursive call to g with i+1
    g(i + 1)
}