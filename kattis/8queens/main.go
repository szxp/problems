// Eight Queens
// In the game of chess, the queen is a powerful piece. It can attack by moving any number of spaces in its current row, in its column or diagonally.
//
// In the eight queens puzzle, eight queens must be placed on a standard 8×8 chess board so that no queen can attack another. The center figure below shows an invalid solution; two queens can attack each other diagonally. The figure on the right shows a valid solution. Given a description of a chess board, your job is to determine whether or not it represents a valid solution to the eight queens puzzle.
//
// \includegraphics[width=0.7\textwidth ]{chess}
// Figure 1: Queen movement (left), invalid solution (center), valid solution (right).
// Input
// Input will contain a description of a single chess board, given as eight lines of eight characters each. Input lines will consist of only the characters ‘.’ and ‘*’. The ‘.’ character represents an empty space on the board, and the ‘*’ character represents a queen.
//
// Output
// Print a single line of output. Print the word “valid” if the given chess board is a valid solution to the eight queens problem. Otherwise, print “invalid”.
//
// Sample Input 1
// *.......
// ..*.....
// ....*...
// ......*.
// .*......
// .......*
// .....*..
// ...*....
//
// Sample Output 1
// invalid
//
// Sample Input 2
// *.......
// ......*.
// ....*...
// .......*
// .*......
// ...*....
// .....*..
// ..*.....
//
// Sample Output 2
// valid

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	w := os.Stdout
	scanner := bufio.NewScanner(os.Stdin)

	valid := true
	queens := make([]*pos, 0, 8)
	var s string
	for r := 0; r < 8; r++ {
		scanner.Scan()
		s = scanner.Text()
		for i, c := range s {
			if c == '*' {
				queens = append(
					queens,
					&pos{row: r, col: i},
				)
			}
		}
	}

	//log.Println(queens)

	if len(queens) < 8 {
		w.Write([]byte("invalid"))
		return
	}

	var rdiff, cdiff int
	var p1, p2 *pos
LOOP:
	for i1 := 0; i1 < 7; i1++ {
		p1 = queens[i1]
		for i2 := i1 + 1; i2 < 8; i2++ {
			p2 = queens[i2]
			if p1.row == p2.row {
				valid = false
				break LOOP
			}
			if p1.col == p2.col {
				valid = false
				break LOOP
			}

			rdiff = int(math.Abs(
				float64(p1.row - p2.row),
			))
			cdiff = int(math.Abs(
				float64(p1.col - p2.col),
			))
			if rdiff == cdiff {
				valid = false
				break LOOP
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	if valid {
		w.Write([]byte("valid"))
	} else {
		w.Write([]byte("invalid"))
	}
}

type pos struct {
	row int
	col int
}

func (p *pos) String() string {
	return fmt.Sprintf("%d:%d", p.row, p.col)
}
