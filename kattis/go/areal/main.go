// A Real Challenge
// /problems/areal/file/statement/en/img-0001.jpg
// Cow image from Pixabay
// Old MacDonald had a farm, and on that farm she had a square-shaped pasture, and on that pasture she had a cow that was prone to escape. So now Old MacDonald wants to set up a fence around the pasture. Given the area of the pasture, how long a fence does Old MacDonald need to buy?
// Input
// The input consists of a single integer a (1≤a≤1018), the area in square meters of Old MacDonald’s pasture.
//
// Output
// Output the total length of fence needed for the pasture, in meters. The length should be accurate to an absolute or relative error of at most 10−6.
//
// Sample Input 1
// 16
//
// Sample Output 1
// 16
//
// Sample Input 2
// 5
//
// Sample Output 2
// 8.94427190999915878564

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

	a, _ := nextInt(b, 0)
	fmt.Fprint(w, math.Sqrt(float64(a))*4)

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
