// Cetvrta
// Mirko needs to choose four points in the plane so that they form a rectangle with sides parallel to the axes. He has already chosen three points and is confident that he hasn’t made a mistake, but is having trouble locating the last point. Help him.
//
// Input
// Each of the three points already chosen will be given on a separate line. All coordinates will be integers between 1 and 1000.
//
// Output
// Output the coordinates of the fourth vertex of the rectangle.
//
// Sample Input 1
// 5 5
// 5 7
// 7 5
//
// Sample Output 1
// 7 7
//
// Sample Input 2
// 30 20
// 10 10
// 10 20
//
// Sample Output 2
// 30 10

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
	x1, i := nextInt(b, 0)
	y1, i := nextInt(b, i)

	scanner.Scan()
	b = scanner.Bytes()
	x2, i := nextInt(b, 0)
	y2, i := nextInt(b, i)

	scanner.Scan()
	b = scanner.Bytes()
	x3, i := nextInt(b, 0)
	y3, i := nextInt(b, i)

	if err := scanner.Err(); err != nil {
		return err
	}

	var x, y int
	switch {
	case x1 == x2:
		x = x3
	case x1 == x3:
		x = x2
	case x2 == x3:
		x = x1
	}

	switch {
	case y1 == y2:
		y = y3
	case y1 == y3:
		y = y2
	case y2 == y3:
		y = y1
	}

	fmt.Fprint(w, x, y)

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
