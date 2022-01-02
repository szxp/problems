// Piece of Cake!
// It is Greg’s birthday! To celebrate, his friend Sam invites Greg and two other friends for a small party. Of course, every birthday party must have cake.
//
// Sam ordered a square cake. She makes a single horizontal cut and a single vertical cut. In her excitement to eat cake, Sam forgot to make these cuts through the middle of the cake.
//
// Of course, the biggest piece of cake should go to Greg since it is his birthday. Help Sam determine the volume of the biggest piece of cake that resulted from these two cuts.
//
// \includegraphics{cake.png}
// Input
// The input consists of a single line containing three integers n (2≤n≤10000), the length of the sides of the square cake in centimeters, h (0<h<n), the distance of the horizontal cut from the top edge of the cake in centimeters, and v (0<v<n), the distance of the vertical cut from the left edge of the cake in centimeters. This is illustrated in the figure above.
//
// Each cake is 4 centimeters thick.
//
// Output
// Display the volume (in cubic centimeters) of the largest of the four pieces of cake after the horizontal and vertical cuts are made.
//
// Sample Input 1
// 10 4 7
//
// Sample Output 1
// 168
//
// Sample Input 2
// 5 2 2
//
// Sample Output 2
// 36
//
// Sample Input 3
// 4 2 1
//
// Sample Output 3
// 24

package main

import (
	"bufio"
	"fmt"
	"log"
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

	n, i := nextInt(b, 0)
	h, i := nextInt(b, i)
	v, i := nextInt(b, i)

	h = maxInt(h, n-h)
	v = maxInt(v, n-v)

	fmt.Fprint(w, h*v*4)

	return w.Flush()
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
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
