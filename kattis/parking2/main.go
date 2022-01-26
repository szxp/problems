// Parking
// /problems/parking2/file/statement/en/img-0001.png
// When shopping on Long Street, Michael usually parks his car at some random location, and then walks to the stores he needs. Can you help Michael choose a place to park which minimises the distance he needs to walk on his shopping round?
// Long Street is a straight line, where all positions are integer. You pay for parking in a specific slot, which is an integer position on Long Street. Michael does not want to pay for more than one parking though. He is very strong, and does not mind carrying all the bags around.
//
// Input
// The first line of input gives the number of test cases, 1≤t≤100. There are two lines for each test case. The first gives the number of stores Michael wants to visit, 1≤n≤20, and the second gives their n integer positions on Long Street, 0≤xi≤99.
//
// Output
// Output for each test case a line with the minimal distance Michael must walk given optimal parking.
//
// Sample Input 1	Sample Output 1
// 2
// 4
// 24 13 89 37
// 6
// 7 30 41 14 39 42
// 152
// 70
//
//
//
// Design
//
// The optimal parking slot is between the stores
// with the minimum and maximum store number (inclusive).
//
// But within that parking range
// it does not matter where Micheal will park,
// because wherever he parks
// he have to go past the stores twice.
// So the minimum distance he needs to walk
// is the double of the distance
// between the minimum and the maximum store number.

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
	defer w.Flush()

	scanner := bufio.NewScanner(os.Stdin)

	// skip the first line
	scanner.Scan()

	var b []byte
	for scanner.Scan() {
		scanner.Scan()
		b = scanner.Bytes()

		var x, max int
		min := 99
		for i := 0; i < len(b); {
			x, i = nextInt(b, i)
			if x < min {
				min = x
			}
			if max < x {
				max = x
			}
		}
		fmt.Fprintln(w, (max-min)*2)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
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
