// Last Factorial Digit
// /problems/lastfactorialdigit/file/statement/en/img-0001.jpg
// Factorials on the complex plane, by Dmitrii Kouznetsov
// The factorial of N, written as N!, is defined as the product of all the integers from 1 to N. For example, 3!=1×2×3=6.
// 
// This number can be very large, so instead of computing the entire product, just compute the last digit of N! (when N! is written in base 10).
// 
// Input
// The first line of input contains a positive integer 1≤T≤10, the number of test cases. Each of the next T lines contains a single positive integer N. N is at most 10.
// 
// Output
// For each value of N, print the last digit of N!.
// 
// Sample Input 1	
// 3
// 1
// 2
// 3
// 
// Sample Output 1
// 1
// 2
// 6
// 
// Sample Input 2	
// 2
// 5
// 2
// 
// Sample Output 2
// 0
// 2

package main

import (
    "bufio"
    "log"
    "os"
    "strconv"
)

func main() {
    w := bufio.NewWriter(os.Stdout)
    scanner := bufio.NewScanner(os.Stdin)

    // skip the first line
    scanner.Scan()

    for scanner.Scan() {
        s := scanner.Text()
        n, err := strconv.Atoi(s)
        if err != nil {
            log.Fatalln(err)
        }

        f := 1
        for i:=1; i<=n; i++ {
            f *= i
        }

        last := int64(f % 10)
        w.Write([]byte(strconv.FormatInt(last, 10)))
        w.Write([]byte("\n"))
    }
    w.Flush()
}
