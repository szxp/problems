// Poker Hand
// You are given a five-card hand drawn from a standard 52-card deck. The strength of your hand is the maximum value k such that there are k cards in your hand that have the same rank.
//
// Compute the strength of your hand.
//
// Input
// The input will consist of a single line, with five two-character strings separated by spaces.
//
// The first character in each string will be the rank of the card, and will be one of A23456789TJQK. The second character in the string will be the suit of the card, and will be one of CDHS.
//
// You may assume all the strings are distinct.
//
// Output
// Output, on a single line, the strength of your hand.
//
// Sample Input 1
// AC AD AH AS KD
//
// Sample Output 1
// 4
//
// Sample Input 2
// 2C 4D 4H 2D 2H
//
// Sample Output 2
// 3
//
// Sample Input 3
// AH 2H 3H 4H 5H
//
// Sample Output 3
// 1
//
//
//
// Design
//
// Map the possible ranks (A23456789TJQK) into an int array.
// We can count the number of ranks in the array.
// While processing the input we can remember the actual maximum count in another int variable.

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

	scanner.Scan()
	b := scanner.Bytes()

	if err := scanner.Err(); err != nil {
		return err
	}

	var strength int
	ranks := make([]int, 13)
	for i := 0; i <= 12; i += 3 {
		k := index(b[i])
		ranks[k]++
		if strength < ranks[k] {
			strength = ranks[k]
		}
	}

	fmt.Fprint(w, strength)
	return nil
}

func index(rank byte) int {
	switch rank {
	case 'A':
		return 0
	case 'T':
		return 9
	case 'J':
		return 10
	case 'Q':
		return 11
	case 'K':
		return 12
	}
	return int(rank - '0') - 1
}
