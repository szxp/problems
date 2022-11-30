// Dice Cup
// \includegraphics[width=0.7\textwidth ]{BluePlatonicDice.png}
// In many table-top games it is common to use different dice to simulate random events. A “d” or “D” is used to indicate a die with a specific number of faces, d4 indicating a four-sided die, for example. If several dice of the same type are to be rolled, this is indicated by a leading number specifying the number of dice. Hence, 2d6 means the player should roll two six-sided dice and sum the result face values.
//
// Task
// Write a program to compute the most likely outcomes for the sum of two dice rolls. Assume each die has numbered faces starting at 1 and that each face has equal roll probability.
//
// Input
// The input consists of a single line with two integer numbers, N,M, specifying the number of faces of the two dice.
//
// Constraints
// 4≤N,M≤20 Number of faces.
//
// Output
// A line with the most likely outcome for the sum; in case of several outcomes with the same probability, they must be listed from lowest to highest value in separate lines.
//
// Sample Input 1
// 6 6
//
// Sample Output 1
// 7
//
// Sample Input 2
// 6 4
//
// Sample Output 2
// 5
// 6
// 7
//
// Sample Input 3
// 12 20
//
// Sample Output 3
// 13
// 14
// 15
// 16
// 17
// 18
// 19
// 20
// 21

package main

import (
	"bufio"
	"bytes"
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

	scanner.Scan()
	b := scanner.Bytes()
	ind := bytes.IndexByte(b, ' ')
	d1, err := strconv.Atoi(string(b[:ind]))
	if err != nil {
		return err
	}
	d2, err := strconv.Atoi(string(b[ind+1:]))
	if err != nil {
		return err
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	probpair := 100.0 / float64(d1) / float64(d2)
	var probmax float64
	sums := make([]float64, d1+d2+1)
	var sum int
	for i := 1; i <= d1; i++ {
		for k := 1; k <= d2; k++ {
			sum = i + k
			sums[sum] += probpair
			if probmax < sums[sum] {
				probmax = sums[sum]
			}
		}
	}

	var diff float64
	for sum, p := range sums {
		diff = p - probmax
		if -1e-9 < diff && diff < 1e-9 {
			fmt.Fprintln(w, sum)
		}
	}

	return w.Flush()
}
