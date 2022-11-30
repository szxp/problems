// 10 Kinds of People
// /problems/10kindsofpeople/file/statement/en/img-0001.jpg
// Image by Christiaan Colen
// The world is made up of 10 kinds of people, those who understand binary and those who do not. These different kinds of people do not always get along so well. Bob might ask for a 10000 ounce coffee (meaning binary) and Alice might make misinterpret his request as being in decimal and give him a 10011100010000 ounce coffee (binary). After Sue explains that this much coffee costs 100 dollars (decimal), Bob might assume he only has to pay 4 dollars (interpreting the price as being in binary). In response to these differences that are difficult to resolve, these two groups have divided the world into two regions, the binary-friendly zones and the decimal-friendly zones. They have even published a map like the following to help people keep up with where the areas are (they have used ones and zeros so nobody would have trouble reading it).
//
// 1111100000
// 1111000000
// 1110000011
// 0111100111
// 0011111111
// Users of binary have to stay in the zones marked with a zero. Users of decimal have to stay in the zones marked with a one. You have to figure out if it is possible for either type of person to get between various locations of interest. People can move north, south, east or west, but cannot move diagonally.
//
// Input
// Input starts with a line containing two positive integers, 1≤r≤1000 and 1≤c≤1000. The next r input lines give the contents of the map, each line containing exactly c characters (which are all chosen from 0 or 1).
//
// The next line has an integer 0≤n≤1000. The following n lines each contain one query, given as four integers: r1,c1 and r2,c2. These two pairs indicate two locations on the map, and their limits are 1≤r1,r2≤r and 1≤c1,c2≤c.
//
// Output
// For each query, output binary if a binary user can start from location r1,c1 and move to location r2,c2. Output decimal if a decimal user can move between the two locations. Otherwise, output neither.
//
// Sample Input 1
// 1 4
// 1100
// 2
// 1 1 1 4
// 1 1 1 1
//
// Sample Output 1
// neither
// decimal
//
// Sample Input 2
// 10 20
// 11111111111111111111
// 11000000000000000101
// 11111111111111110000
// 11111111111111110000
// 11000000000000000111
// 00011111111111111111
// 00111111111111111111
// 10000000000000001111
// 11111111111111111111
// 11111111111111111111
// 3
// 2 3 8 16
// 8 1 7 3
// 1 1 10 20
//
// Sample Output 2
// binary
// decimal
// neither

package main

import (
	"bufio"
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
	row := scanner.Bytes()
	r, i := nextInt(row, 0)
	c, i := nextInt(row, i)

	brows := make([]byte, r*c)
	bitmap := make([][]byte, r)
	for i := 0; i < r; i++ {
		bitmap[i] = brows[i*c : (i+1)*c]
	}

	vrows := make([]int, r*c)
	visit := make([][]int, r)
	for i := 0; i < r; i++ {
		visit[i] = vrows[i*c : (i+1)*c]
	}

	for k := 0; k < r; k++ {
		scanner.Scan()
		row = scanner.Bytes()
		copy(bitmap[k], row)
	}

	// skip next line
	scanner.Scan()

	var r1, c1, r2, c2, j, areaID int
	var tile1, tile2 byte
	for scanner.Scan() {
		row = scanner.Bytes()
		r1, j = nextInt(row, 0)
		c1, j = nextInt(row, j)
		r2, j = nextInt(row, j)
		c2, j = nextInt(row, j)

		r1--
		c1--
		r2--
		c2--

		tile1 = bitmap[r1][c1]
		tile2 = bitmap[r2][c2]

		if tile1 != tile2 {
			w.WriteString("neither\n")
			continue
		}

		if visit[r1][c1] == 0 {
			areaID++
			step(r1, c1, bitmap, visit, areaID, r2, c2, tile1)
		}

		if visit[r1][c1] == visit[r2][c2] {
			if tile1 == '0' {
				w.WriteString("binary\n")
			} else {
				w.WriteString("decimal\n")
			}
		} else {
			w.WriteString("neither\n")
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}

func step(
	i int,
	j int,
	bitmap [][]byte,
	visit [][]int,
	areaID int,
	r2 int,
	c2 int,
	tile byte,
) {
	visit[i][j] = areaID

	// left
	if 0 <= j-1 && bitmap[i][j-1] == tile && visit[i][j-1] == 0 {
		step(i, j-1, bitmap, visit, areaID, r2, c2, tile)
	}

	// right
	if j+1 < len(bitmap[i]) && bitmap[i][j+1] == tile && visit[i][j+1] == 0 {
		step(i, j+1, bitmap, visit, areaID, r2, c2, tile)
	}

	// up
	if 0 <= i-1 && bitmap[i-1][j] == tile && visit[i-1][j] == 0 {
		step(i-1, j, bitmap, visit, areaID, r2, c2, tile)
	}

	// down
	if i+1 < len(bitmap) && bitmap[i+1][j] == tile && visit[i+1][j] == 0 {
		step(i+1, j, bitmap, visit, areaID, r2, c2, tile)
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
