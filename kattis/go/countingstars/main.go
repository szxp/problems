// Counting Stars
// The field of astronomy has been significantly advanced through the use of computer technology. Algorithms can automatically survey digital images of the night sky, looking for new patterns.
//
// For this problem, you should write such an analysis program which counts the number of stars visible in an bitmap image. An image consists of pixels, and each pixel is either black or white (represented by the characters # and -, respectively). All black pixels are considered to be part of the sky, and each white pixel is considered to be part of a star. White pixels that are adjacent vertically or horizontally are part of the same star.
//
// Input
// Each test case begins with a line containing a pair of integers 1≤m,n≤100. This is followed by m lines, each of which contains exactly n pixels. Input contains at least one and at most 50 test cases, and input ends at the end of file.
//
// Output
// For each case, display the case number followed by the number of stars that are visible in the corresponding image. Follow the format of the sample output.
//
// Sample Input 1
// 10 20
// #################---
// ##-###############--
// #---################
// ##-#################
// ########---#########
// #######-----########
// ########---#########
// ##################--
// #################---
// ##################-#
// 3 10
// #-########
// ----------
// #-########
//
// Sample Output 1
// Case 1: 4
// Case 2: 1

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

	var m, n, i, from, count, cas int
	var row []byte
	for scanner.Scan() {
		row = scanner.Bytes()
		m, i = nextInt(row, 0)
		n, i = nextInt(row, i)
		cas++
		count = 0

		img := make([][]byte, m)
		for k := 0; k < m; k++ {
			scanner.Scan()
			row = scanner.Bytes()
			img[k] = append(img[k], row...)
		}

		visit := make([][]bool, m)
		for x := range visit {
			visit[x] = make([]bool, n)
		}

		for k := 0; k < m; k++ {
			row = img[k]
			for j := 0; j < len(row); {
				from, j = nextPart(row, j)
				if from == -1 {
					continue
				}

				if !visit[k][from] {
					count++
					step(k, from, img, visit)
				}
			}
		}

		fmt.Fprintf(w, "Case %d: %d\n", cas, count)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}

func step(
	i int,
	j int,
	img [][]byte,
	visit [][]bool,
) {
	visit[i][j] = true

	// left
	if 0 <= j-1 && img[i][j-1] == '-' && !visit[i][j-1] {
		step(i, j-1, img, visit)
	}

	// right
	if j+1 < len(img[i]) && img[i][j+1] == '-' && !visit[i][j+1] {
		step(i, j+1, img, visit)
	}

	// up
	if 0 <= i-1 && img[i-1][j] == '-' && !visit[i-1][j] {
		step(i-1, j, img, visit)
	}

	// down
	if i+1 < len(img) && img[i+1][j] == '-' && !visit[i+1][j] {
		step(i+1, j, img, visit)
	}
}

func nextPart(b []byte, i int) (from, to int) {
	for ; i < len(b) && b[i] == '#'; i++ {
	}
	if i == len(b) {
		return -1, i
	}
	from = i
	for ; i < len(b) && b[i] == '-'; i++ {
	}
	return from, i
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
