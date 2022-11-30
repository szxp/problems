// Where's My Internet??
// /problems/wheresmyinternet/file/statement/en/img-0001.jpg
// Photo by Jerry John from Flickr
// A new town is being built far out in the country, and currently there are N houses. People have already started moving in. However, some of the houses aren’t connected to the internet yet, and naturally residents are outraged.
//
// The houses are numbered 1 to N. House number 1 has already been connected to the internet via a long network cable to a neighboring town. The plan is to provide internet to other houses by connecting pairs of houses with separate network cables. A house is connected to the internet if it has a network cable to another house that’s already connected to the internet.
//
// Given a list of which pairs of houses are already connected by a network cable, determine which houses are not yet connected to the internet.
//
// Input
// The first line of input contains two integers 1≤N,M≤200000, where N is the number of houses and M is the number of network cables already deployed. Then follow M lines, each containing a pair of distinct house numbers 1≤a,b≤N meaning that house number a and house number b are already connected by a network cable. Each house pair is listed at most once in the input.
//
// Output
// If all the houses are already connected to the internet, output one line containing the string Connected. Otherwise, output a list of house numbers in increasing order, one per line, representing the houses that are not yet connected to the internet.
//
// Sample Input 1
// 6 4
// 1 2
// 2 3
// 3 4
// 5 6
//
// Sample Output 1
// 5
// 6
//
// Sample Input 2
// 2 1
// 2 1
//
// Sample Output 2
// Connected
//
// Sample Input 3
// 4 3
// 2 3
// 4 2
// 3 4
//
// Sample Output 3
// 2
// 3
// 4

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

	n, err := strconv.Atoi(string(b[:ind]))
	if err != nil {
		return err
	}
	m, err := strconv.Atoi(string(b[ind+1:]))
	if err != nil {
		return err
	}

	var h1, h2 int
	cables := make([][]int, n)
	for i := 0; i < m; i++ {
		scanner.Scan()
		b = scanner.Bytes()
		ind = bytes.IndexByte(b, ' ')

		h1, err = strconv.Atoi(string(b[:ind]))
		if err != nil {
			return err
		}
		h2, err = strconv.Atoi(string(b[ind+1:]))
		if err != nil {
			return err
		}

		cables[h1-1] = append(cables[h1-1], h2-1)
		cables[h2-1] = append(cables[h2-1], h1-1)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	conn := make([]bool, n)
	var count int
	step(0, cables, conn, &count)

	if len(conn) == count {
		w.WriteString("Connected")
	} else {
		for i, c := range conn {
			if !c {
				fmt.Fprint(w, i+1)
				w.WriteByte('\n')
			}
		}
	}

	return w.Flush()
}

func step(
	i int,
	cables [][]int,
	conn []bool,
	count *int,
) {
	conn[i] = true
	*count += 1
	for _, to := range cables[i] {
		if !conn[to] {
			step(to, cables, conn, count)
		}
	}
}
