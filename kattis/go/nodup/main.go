// No Duplicates
// There is a game in which you try not to repeat a word while your opponent tries to see if you have repeated one.
//
// "THE RAIN IN SPAIN" has no repeats.
//
// "IN THE RAIN AND THE SNOW" repeats THE.
//
// "THE RAIN IN SPAIN IN THE PLAIN" repeats THE and IN.
//
// Write a program to test a phrase.
//
// Input
// Input is a line containing words separated by single spaces, where a word consists of one or more uppercase letters. A line contains no more than 80 characters.
//
// Output
// The output is "yes" if no word is repeated, and "no" if one or more words repeat.
//
// Sample Input 1
// THE RAIN IN SPAIN
//
// Sample Output 1
// yes
//
// Sample Input 2
// IN THE RAIN AND THE SNOW
//
// Sample Output 2
// no
//
// Sample Input 3
// THE RAIN IN SPAIN IN THE PLAIN
//
// Sample Output 3
// no

package main

import (
	"bufio"
	//"fmt"
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

	var found bool
	var from int
	var tok []byte
	for i := 0; i < len(b); {
		from, i = nextToken(b, i)
		if from > -1 {
			tok = b[from:i]
			if exists(tok, b[i:]) {
				found = true
				break
			}
		}
	}

	if found {
		w.WriteString("no")
	} else {
		w.WriteString("yes")
	}

	return w.Flush()
}

func nextToken(b []byte, i int) (from, to int) {
	for ; i < len(b) && b[i] == ' '; i++ {
	}
	if i == len(b) {
		return -1, i
	}
	from = i
	for ; i < len(b) && b[i] != ' '; i++ {
	}
	to = i
	return
}

func exists(tok, b []byte) bool {
	var s int
	var found bool
	for i := 0; i < len(b); {
		for ; i < len(b) && b[i] == ' '; i++ {
		}

		s = i
		found = true
		for k := 0; found && k < len(tok) && s+k < len(b) && b[s+k] != ' '; k++ {
			if tok[k] != b[s+k] {
				found = false
			}
		}

		for ; i < len(b) && b[i] != ' '; i++ {
		}

		if found && i-s == len(tok) {
			return true
		}
	}
	return false
}
