// FizzBuzz
// /problems/fizzbuzz/file/statement/en/img-0001.png
// Image by chris morgan cc by
// According to Wikipedia, FizzBuzz is a group word game for children to teach them about division. This may or may not be true, but this question is generally used to torture screen young computer science graduates during programming interviews.
// Basically, this is how it works: you print the integers from 1 to N, replacing any of them divisible by X with Fizz or, if they are divisible by Y, with Buzz. If the number is divisible by both X and Y, you print FizzBuzz instead.
// 
// Check the samples for further clarification.
// 
// Input
// Input contains a single test case. Each test case contains three integers on a single line, X, Y and N (1≤X<Y≤N≤100).
// 
// Output
// Print integers from 1 to N in order, each on its own line, replacing the ones divisible by X with Fizz, the ones divisible by Y with Buzz and ones divisible by both X and Y with FizzBuzz.
// 
// Sample Input 1	
// 2 3 7
// 
// Sample Output 1
// 1
// Fizz
// Buzz
// Fizz
// 5
// FizzBuzz
// 7
// 
// Sample Input 2	
// 2 4 7
// 
// Sample Output 2
// 1
// Fizz
// 3
// FizzBuzz
// 5
// Fizz
// 7
// 
// Sample Input 3	
// 3 5 7
// 
// Sample Output 3
// 1
// 2
// Fizz
// 4
// Buzz
// Fizz
// 7

package main

import (
    "os"
    "log"
    "bufio"
    "fmt"
)

func main() {
    var x, y, n int

    _, err := fmt.Scanln(&x, &y, &n)
    if err != nil {
        log.Fatalln(err)
    }

    w := bufio.NewWriter(os.Stdout)

    var f, b bool
    for i:=1; i<=n; i++ {
        f = i%x == 0
        b = i%y == 0
        if f {
            fmt.Fprint(w, "Fizz")
        }
        if b {
            fmt.Fprint(w, "Buzz")
        }
        if !f && !b {
            fmt.Fprint(w, i)
        }
        fmt.Fprintln(w)
    }
    w.Flush()
}


