// 2048
// /problems/2048/file/statement/en/img-0001.jpg
// screenshot taken from http://gabrielecirulli.github.io/2048/
// 2048 is a single-player puzzle game created by Gabriele Cirulli1. It is played on a 4×4 grid that contains integers ≥2 that are powers of 2. The player can use a keyboard arrow key (left/up/right/down) to move all the tiles simultaneously. Tiles slide as far as possible in the chosen direction until they are stopped by either another tile or the edge of the grid. If two tiles of the same number collide while moving, they will merge into a tile with the total value of the two tiles that collided. The resulting tile cannot merge with another tile again in the same move. Please observe this merging behavior carefully in all Sample Inputs and Outputs.
//
// Input
// The input is always a valid game state of a 2048 puzzle. The first four lines of input, that each contains four integers, describe the 16 integers in the 4×4 grid of 2048 puzzle. The j-th integer in the i-th line denotes the content of the cell located at the i-th row and the j-th cell. For this problem, all integers in the input will be either {0, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}. Integer 0 means an empty cell.
//
// The fifth line of input contains an integer 0, 1, 2, or 3 that denotes a left, up, right, or down move executed by the player, respectively.
//
// Output
// Output four lines with four integers each. Two integers in a line must be separated by a single space. This describes the new state of the 4×4 grid of 2048 puzzle. Again, integer 0 means an empty cell. Note that in this problem, you can ignore the part from the 2048 puzzle where it introduces a new random tile with a value of either 2 or 4 in an empty spot of the board at the start of a new turn.
//
// Sample Input 1
// 2 0 0 2
// 4 16 8 2
// 2 64 32 4
// 1024 1024 64 0
// 0
//
// Sample Output 1
// 4 0 0 0
// 4 16 8 2
// 2 64 32 4
// 2048 64 0 0
//
// Sample Input 2
// 2 0 0 2
// 4 16 8 2
// 2 64 32 4
// 1024 1024 64 0
// 1
//
// Sample Output 2
// 2 16 8 4
// 4 64 32 4
// 2 1024 64 0
// 1024 0 0 0
//
// Sample Input 3
// 2 0 0 2
// 4 16 8 2
// 2 64 32 4
// 1024 1024 64 0
// 2
//
// Sample Output 3
// 0 0 0 4
// 4 16 8 2
// 2 64 32 4
// 0 0 2048 64
//
// Sample Input 4
// 2 0 0 2
// 4 16 8 2
// 2 64 32 4
// 1024 1024 64 0
// 3
//
// Sample Output 4
// 2 0 0 0
// 4 16 8 0
// 2 64 32 4
// 1024 1024 64 4
//
// Sample Input 5
// 2 2 4 8
// 4 0 4 4
// 16 16 16 16
// 32 16 16 32
// 0
//
// Sample Output 5
// 4 4 8 0
// 8 4 0 0
// 32 32 0 0
// 32 32 32 0
//
// Sample Input 6
// 2 2 4 8
// 4 0 4 4
// 16 16 16 16
// 32 16 16 32
// 2
//
// Sample Output 6
// 0 4 4 8
// 0 0 4 8
// 0 0 32 32
// 0 32 32 32
//
// Footnotes
//
// Based on 1024 by Veewo Studio and conceptually similar to Threes by Asher Vollmer.

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

	puzzle := make([][]int, 4)
	for i := 0; i < 4; i++ {
		puzzle[i] = make([]int, 4)
	}

	var b []byte
	for i := 0; i < 4; i++ {
		scanner.Scan()
		b = scanner.Bytes()
		var n, j int
		for k := 0; k < len(b); {
			n, k = nextInt(b, k)
			puzzle[i][j] = n
			j++
		}
	}

	scanner.Scan()
	b = scanner.Bytes()
	dir, _ := nextInt(b, 0)

	if err := scanner.Err(); err != nil {
		return err
	}

	switch dir {
	case 0:
		moveLeft(puzzle)
		mergeLeft(puzzle)
		moveLeft(puzzle)
	case 1:
		moveUp(puzzle)
		mergeUp(puzzle)
		moveUp(puzzle)
	case 2:
		moveRight(puzzle)
		mergeRight(puzzle)
		moveRight(puzzle)
	case 3:
		moveDown(puzzle)
		mergeDown(puzzle)
		moveDown(puzzle)
	}

	dump(puzzle)

	return w.Flush()
}

func moveLeft(p [][]int) {
	for i := 0; i < len(p); i++ {
		moveLineLeft(p, i)
	}
}

func moveLineLeft(p [][]int, i int) {
	var j int
	for ; j < len(p[i]) && p[i][j] != 0; j++ {
	}
	k := findNonZeroRight(p[i], j+1)

	for j < len(p[i]) && k < len(p[i]) {
		p[i][j] = p[i][k]
		p[i][k] = 0
		j = k
		k = findNonZeroRight(p[i], j+1)
	}
}

func findNonZeroRight(row []int, from int) int {
	for ; from < len(row) && row[from] == 0; from++ {
	}
	return from
}

func mergeLeft(p [][]int) {
	for i := 0; i < len(p); i++ {
		mergeLineLeft(p, i)
	}
}

func mergeLineLeft(p [][]int, i int) {
	for j := 0; j < len(p[i])-1; {
		if p[i][j] > 0 && p[i][j+1] > 0 && p[i][j] == p[i][j+1] {
			p[i][j] *= 2
			p[i][j+1] = 0
			j += 2
		} else {
			j++
		}
	}
}

func moveUp(p [][]int) {
	for j := 0; j < len(p[0]); j++ {
		moveLineUp(p, j)
	}
}

func moveLineUp(p [][]int, j int) {
	var i int
	for ; i < len(p) && p[i][j] != 0; i++ {
	}
	k := findNonZeroDown(p, j, i+1)

	for i < len(p) && k < len(p) {
		p[i][j] = p[k][j]
		p[k][j] = 0
		i = k
		k = findNonZeroDown(p, j, i+1)
	}
}

func findNonZeroDown(p [][]int, j, i int) int {
	for ; i < len(p) && p[i][j] == 0; i++ {
	}
	return i
}

func mergeUp(p [][]int) {
	for j := 0; j < len(p[0]); j++ {
		mergeLineUp(p, j)
	}
}

func mergeLineUp(p [][]int, j int) {
	for i := 0; i < len(p)-1; {
		if p[i][j] > 0 && p[i+1][j] > 0 && p[i][j] == p[i+1][j] {
			p[i][j] *= 2
			p[i+1][j] = 0
			i += 2
		} else {
			i++
		}
	}
}

func moveRight(p [][]int) {
	for i := 0; i < len(p); i++ {
		moveLineRight(p, i)
	}
}

func moveLineRight(p [][]int, i int) {
	j := len(p[i]) - 1
	for ; 0 <= j && p[i][j] != 0; j-- {
	}
	k := findNonZeroLeft(p[i], j-1)

	for 0 <= j && 0 <= k {
		p[i][j] = p[i][k]
		p[i][k] = 0
		j = k
		k = findNonZeroLeft(p[i], j-1)
	}
}

func findNonZeroLeft(row []int, from int) int {
	for ; 0 <= from && row[from] == 0; from-- {
	}
	return from
}

func mergeRight(p [][]int) {
	for i := 0; i < len(p); i++ {
		mergeLineRight(p, i)
	}
}

func mergeLineRight(p [][]int, i int) {
	for j := len(p[i]) - 1; 0 < j; {
		if p[i][j] > 0 && p[i][j-1] > 0 && p[i][j] == p[i][j-1] {
			p[i][j] *= 2
			p[i][j-1] = 0
			j -= 2
		} else {
			j--
		}
	}
}

func moveDown(p [][]int) {
	for j := 0; j < len(p[0]); j++ {
		moveLineDown(p, j)
	}
}

func moveLineDown(p [][]int, j int) {
	i := len(p) - 1
	for ; 0 <= i && p[i][j] != 0; i-- {
	}
	k := findNonZeroUp(p, j, i-1)

	for 0 <= i && 0 <= k {
		p[i][j] = p[k][j]
		p[k][j] = 0
		i = k
		k = findNonZeroUp(p, j, i-1)
	}
}

func findNonZeroUp(p [][]int, j, i int) int {
	for ; 0 <= i && p[i][j] == 0; i-- {
	}
	return i
}

func mergeDown(p [][]int) {
	for j := 0; j < len(p[0]); j++ {
		mergeLineDown(p, j)
	}
}

func mergeLineDown(p [][]int, j int) {
	for i := len(p) - 1; 0 < i; {
		if p[i][j] > 0 && p[i-1][j] > 0 && p[i][j] == p[i-1][j] {
			p[i][j] *= 2
			p[i-1][j] = 0
			i -= 2
		} else {
			i--
		}
	}
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

func dump(p [][]int) {
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[i]); j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(p[i][j])
		}
		fmt.Println()
	}
}
