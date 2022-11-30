// R2
//
// The number S is called the mean of two numbers R1 and R2 if S is equal to (R1+R2)/2. Mirko’s birthday present for Slavko was two integers R1 and R2. Slavko promptly calculated their mean which also happened to be an integer but then lost R2! Help Slavko restore R2.
//
// Input
// The first and only line of input contains two integers R1 and S, both between −1000 and 1000.
//
// Output
// Output R2 on a single line.
//
// Sample Input 1
// 11 15
//
// Sample Output 1
// 19
//
// Sample Input 2
// 4 3
//
// Sample Output 2
// 2

package main

import (
	"errors"
	"fmt"
	"io"
	"log"
)

func main() {
	var r1, s int64
	_, err := fmt.Scanln(&r1, &s)
	if err != nil && !errors.Is(err, io.EOF) {
		log.Fatalln(err)
	}
	fmt.Println(2*s - r1)
}
