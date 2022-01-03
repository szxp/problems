// Symmetric Order
// In your job at Albatross Circus Management (yes, it’s run by a bunch of clowns), you have just finished writing a program whose output is a list of names in nondescending order by length (so that each name is at least as long as the one preceding it). However, your boss does not like the way the output looks, and instead wants the output to appear more symmetric, with the shorter strings at the top and bottom and the longer strings in the middle. His rule is that each pair of names belongs on opposite ends of the list, and the first name in the pair is always in the top part of the list. In the first example set below, Bo and Pat are the first pair, Jean and Kevin the second pair, etc.
//
// Input
// The input consists of one or more sets of strings, followed by a final line containing only the value 0. Each set starts with a line containing an integer, n, which is the number of strings in the set, followed by n strings, one per line, sorted in nondescending order by length. None of the strings contain spaces. There is at least one and no more than 15 strings per set. Each string is at most 25 characters long.
//
// Output
// For each input set print “SET n” on a line, where n starts at 1, followed by the output set as shown in the sample output.
//
// Sample Input 1
// 7
// Bo
// Pat
// Jean
// Kevin
// Claude
// William
// Marybeth
// 6
// Jim
// Ben
// Zoe
// Joey
// Frederick
// Annabelle
// 5
// John
// Bill
// Fran
// Stan
// Cece
// 0
//
// Sample Output 1
// SET 1
// Bo
// Jean
// Claude
// Marybeth
// William
// Kevin
// Pat
// SET 2
// Jim
// Zoe
// Frederick
// Annabelle
// Joey
// Ben
// SET 3
// John
// Fran
// Cece
// Stan
// Bill

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

	maxn := 15
	names := make([]string, maxn)
	var b []byte
	var set, fp, lp int
	for scanner.Scan() {
		b = scanner.Bytes()
		n, _ := nextInt(b, 0)
		if n == 0 {
			break
		}
		set++
		fp = 0
		lp = maxn

		for i := 0; i < n; i++ {
			scanner.Scan()
			if i%2 == 0 {
				fp++
				names[fp-1] = scanner.Text()
			} else {
				lp--
				names[lp] = scanner.Text()
			}
		}

		fmt.Fprintf(w, "SET %d\n", set)
		for i := 0; i < fp; i++ {
			fmt.Fprintln(w, names[i])
		}
		for i := lp; i < maxn; i++ {
			fmt.Fprintln(w, names[i])
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

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
