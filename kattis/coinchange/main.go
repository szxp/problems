package main

import (
    "fmt"
)

func main() {
    coins := []int{20000, 10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 5}
    amount := 32465

    counts := change(coins, amount)
    fmt.Println(counts)
}

func change(coins []int, amount int) []int {
    counts := make([]int, len(coins))

    for i, c := range coins {
        if amount == 0 {
            break
        }
        counts[i] = amount / c
        amount = amount % c
    }

    if amount != 0 {
        return nil
    }
    return counts
}
