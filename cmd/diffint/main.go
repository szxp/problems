// A Different Problem
// Write a program that computes the difference between non-negative integers.
// 
// Input
// Each line of the input consists of a pair of integers. Each integer is between 0 and 1015 (inclusive). The input is terminated by end of file.
// 
// Output
// For each pair of integers in the input, output one line, containing the absolute value of their difference.
// 
// Sample Input 1
// 10 12
// 71293781758123 72784
// 1 12345677654321
// 
// Sample Output 1
// 2
// 71293781685339
// 12345677654320

package main

import (
    "fmt"
    "errors"
    "io"
    "math"
    "log"
)

func main() {
    var v1, v2 int64
    for {
        _, err := fmt.Scanln(&v1, &v2)
        if err != nil {
            if errors.Is(err, io.EOF) {
                break
            } else {
               log.Fatalln(err)
            }
        }

        d := int64(math.Abs(float64(v1 - v2)))
        fmt.Println(d)
    }
}
