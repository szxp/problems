// Line Them Up
// An eccentric coach asks players on the team to line up alphabetically at the start of practice. The coach does not tell the players whether they need to line up in increasing or decreasing order, so they guess. If they guess wrong, the coach makes them run laps before practice. Given a list of names, you are to determine if the list is in increasing alphabetical order, decreasing alphabetical order or neither.
//
// Input
// The input consists of a single test case. The first line will contain the number N of people on the team (2≤N≤20). Following that are N lines, each containing the name of one person. A name will be at least two characters and at most 12 characters in length and will consist only of capital letters, and with no white spaces (sorry BILLY BOB and MARY JOE). Duplicates names will not be allowed on a team.
//
// Output
// Output a single word: INCREASING if the list is in increasing alphabetical order, DECREASING if it is in decreasing alphabetical order, and otherwise NEITHER.
//
// Sample Input 1
// 5
// JOE
// BOB
// ANDY
// AL
// ADAM
//
// Sample Output 1
// DECREASING
//
// Sample Input 2
// 11
// HOPE
// ALI
// BECKY
// JULIE
// MEGHAN
// LAUREN
// MORGAN
// CARLI
// MEGAN
// ALEX
// TOBIN
//
// Sample Output 2
// NEITHER
//
// Sample Input 3
// 4
// GEORGE
// JOHN
// PAUL
// RINGO
//
// Sample Output 3
// INCREASING

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

	// skip the first line
	scanner.Scan()

	var ord, ordPrev order
	var s, prev string
	for scanner.Scan() {
		s = scanner.Text()

		if prev != "" {
			if prev < s {
				ord = inc
			} else {
				ord = dec
			}
			if ordPrev != 0 && ord != ordPrev {
				ord = nei
				break
			}
			ordPrev = ord
		}

		prev = s
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if ord == nei {
		fmt.Fprint(w, "NEITHER")
	} else if ord == inc {
		fmt.Fprint(w, "INCREASING")
	} else if ord == dec {
		fmt.Fprint(w, "DECREASING")
	}

	return w.Flush()
}

type order int

const nei order = 1
const inc order = 2
const dec order = 3
