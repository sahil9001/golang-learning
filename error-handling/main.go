package main

import (
    "errors"  // Used for creating and comparing errors
    "fmt"
)

// ErrDivideByZero is a **sentinel error**.
// Defined at package level so it can be compared reliably (errors.Is).
var ErrDivideByZero = errors.New("divide by zero")

// Divide performs integer division and returns an error if b is zero.
func Divide(a, b int) (int, error) {
    if b == 0 {
        // Return zero result and the sentinel error
        return 0, ErrDivideByZero
    }
    // Valid division: return result and nil error
    return a / b, nil
}

func main() {
    a, b := 10, 0

    // Call Divide and check both result and error
    result, err := Divide(a, b)
    if err != nil {
        // Switch lets us handle specific error types cleanly
        switch {
        case errors.Is(err, ErrDivideByZero):
            // Compare using errors.Is() – works even if wrapped errors are introduced later
            fmt.Println("divide by zero error")

        default:
            // Unexpected error: print it out
            fmt.Printf("unexpected division error: %s\n", err)
        }
        return
    }

    // No error: safely print the result
    fmt.Printf("%d / %d = %d\n", a, b, result)
}