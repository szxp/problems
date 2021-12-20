// Planting Trees
// /problems/plantingtrees/file/statement/en/img-0001.JPG
// Photo from Wikipedia, after U. S. Department of Agriculture
// Farmer Jon has recently bought n tree seedlings that he wants to plant in his yard. It takes 1 day for Jon to plant a seedling1, and for each tree Jon knows exactly in how many days after planting it grows to full maturity. Jon would also like to throw a party for his farmer friends, but in order to impress them he would like to organize the party only after all the trees have grown. More precisely, the party can be organized at earliest on the next day after the last tree has grown up.
//
// Help Jon to find out when is the earliest day when the party can take place. Jon can choose the order of planting the trees as he likes, so he wants to plant the trees in such a way that the party will be as soon as possible.
//
// Input
// The input consists of two lines. The first line contains a single integer N (1≤N≤100000) denoting the number of seedlings. Then a line with N integers ti follows (1≤ti≤1000000), where ti denotes the number of days it takes for the ith tree to grow.
//
// Output
// You program should output exactly one line containing one integer, denoting the earliest day when the party can be organized. The days are numbered 1,2,3,… beginning from the current moment.
//
// Sample Input 1
// 4
// 2 3 4 3
//
// Sample Output 1
// 7
//
// Sample Input 2
// 6
// 39 38 9 35 39 20
//
// Sample Output 2
// 42

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
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

	const maxCapacity = 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	s := scanner.Text()
	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	seeds := make([]int, 0, n)
	var ind, t int
	scanner.Scan()
	b := scanner.Bytes()
	for len(b) > 0 {
		ind = bytes.IndexByte(b, ' ')
		if ind != -1 {
			t, err = strconv.Atoi(string(b[:ind]))
			if err != nil {
				return err
			}
			b = b[ind+1:]
		} else {
			t, err = strconv.Atoi(string(b))
			if err != nil {
				return err
			}
			b = b[len(b):]
		}
		seeds = append(seeds, t)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	sort.Ints(seeds)
	var max, offset, v int
	for i := len(seeds) - 1; i >= 0; i-- {
		t = seeds[i]
		offset = len(seeds) - 1 - i
		v = offset + t + 1
		if max < v {
			max = v
		}
	}

	fmt.Fprint(w, max+1)

	return w.Flush()
}
