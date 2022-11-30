// Pizza Crust
// /problems/pizza2/file/statement/en/img-0001.jpg
// Photo by cyclonebill.
// George has bought a pizza. George loves cheese. George thinks the pizza does not have enough cheese. George gets angry.
//
// George’s pizza is round, and has a radius of R cm. The outermost C cm is crust, and does not have cheese. What percent of George’s pizza has cheese?
//
// Input
// The input consists of a single line with two space separated integers, R and C.
//
// Output
// Output the percentage of the pizza that has cheese. Your answer must have an absolute or relative error of at most 10−6.
//
// Limits
// 1≤C≤R≤100
//
// Sample Input 1
// 1 1
//
// Sample Output 1
// 0.000000000
//
// Sample Input 2
// 2 1
//
// Sample Output 2
// 25.000000

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	err := solve()
	if err != nil {
		log.Fatalln(err)
	}
}

func solve() error {
	w := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	b := scanner.Bytes()

	if err := scanner.Err(); err != nil {
		return err
	}

	r, i := nextInt(b, 0)
	c, i := nextInt(b, i)

	a1 := math.Pi * float64(r*r)
	a2 := math.Pi * float64((r-c)*(r-c))

	fmt.Fprint(w, a2/a1*100)

	return w.Flush()
}

func nextInt(b []byte, i int) (val, ni int) {
	for ; i < len(b) && !('0' <= b[i] && b[i] <= '9' || b[i] == '-'); i++ {
	}
	x := 0
	sign := 1
	for ; i < len(b) && ('0' <= b[i] && b[i] <= '9' || b[i] == '-'); i++ {
		if b[i] == '-' {
			sign = -1
		} else {
			x = x*10 + int(b[i]) - '0'
		}
	}
	return x * sign, i
}
