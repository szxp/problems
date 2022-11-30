// Cold-puter Science
// We’re not going to sugar-coat it: Chicago’s winters can be rough. The temperatures sometimes dip to uncomfortable levels and, after last year’s “polar vortex”, the University of Chicago Weather Service wants to find out exactly how bad the winter was. More specifically, they are interested in knowing the total number of days in which the temperature was below zero degrees Celsius.
//
// Input
// The input is composed of two lines. The first line contains a single positive integer n (1≤n≤100) that specifies the number of temperatures collected by the University of Chicago Weather Service. The second line contains n temperatures, each separated by a single space. Each temperature is represented by an integer t (−1000000≤t≤1000000)
//
// Output
// You must print a single integer: the number of temperatures strictly less than zero.
//
// Sample Input 1
// 3
// 5 -10 15
//
// Sample Output 1
// 1
//
// Sample Input 2
// 5
// -14 -5 -39 -5 -7
//
// Sample Output 2
// 5

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	buf := make([]byte, 1024)

	k := 0
	for {
		n, err := os.Stdin.Read(buf)
		for _, c := range buf[:n] {
			if c == '-' {
				k++
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
	}

	fmt.Println(k)
}
