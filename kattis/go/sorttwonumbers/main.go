// Sort Two Numbers
// In this problem, your program should read two whole numbers (also called integers) from the input, and print them out in increasing order.
//
// Input
// The input contains one line, which has two integers a and b, separated by a single space. The bounds on these values are 0≤a,b≤1000000.
//
// Output
// Output the smaller number first, and the larger number second.
//
// Sample Input 1
// 3 4
//
// Sample Output 1
// 3 4
//
// Sample Input 2
// 987 23
//
// Sample Output 2
// 23 987

package main

import (
	"errors"
	"fmt"
	"io"
	"log"
)

func main() {
	var v1, v2 int
	_, err := fmt.Scanln(&v1, &v2)
	if err != nil && !errors.Is(err, io.EOF) {
		log.Fatalln(err)
	}
	if v1 <= v2 {
		fmt.Printf("%d %d\n", v1, v2)
	} else {
		fmt.Printf("%d %d\n", v2, v1)
	}
}
