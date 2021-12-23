// Odd Man Out
// You are hosting a party with G guests and notice that there is an odd number of guests! When planning the party you deliberately invited only couples and gave each couple a unique number C on their invitation. You would like to single out whoever came alone by asking all of the guests for their invitation numbers.
//
// Input
// The first line of input gives the number of cases, N. N test cases follow. For each test case there will be:
//
// One line containing the value G the number of guests.
//
// One line containing a space-separated list of G integers. Each integer C indicates the invitation code of a guest.
//
// You may assume that 1≤N≤50,0<C<231,3≤G<1000.
//
// Output
// For each test case, output one line containing “Case #x: ” followed by the number C of the guest who is alone.
//
// Sample Input 1
// 3
// 3
// 1 2147483647 2147483647
// 5
// 3 4 7 4 3
// 5
// 2 10 2 10 5
//
// Sample Output 1
// Case #1: 1
// Case #2: 7
// Case #3: 5

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return err
	}

	guests := make(map[int]struct{}, 1000)
	var g, c, m int
	for t := 0; t < n; t++ {
		scanner.Scan()
		g, err = strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}

		for i := 0; i < g; i++ {
			scanner.Scan()
			c, err = strconv.Atoi(scanner.Text())
			if err != nil {
				return err
			}

			_, ok := guests[c]
			if ok {
				delete(guests, c)
			} else {
				guests[c] = struct{}{}
			}
		}

		m = missing(guests)
		fmt.Fprintf(w, "Case #%d: %d\n", t+1, m)
		reset(guests)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}

func reset(m map[int]struct{}) {
	for k := range m {
		delete(m, k)
	}
}

func missing(m map[int]struct{}) int {
	for k := range m {
		return k
	}
	return 0
}
