// Sibice
// Young Mirko threw matches all over the floor of his room.
//
// His mom did not like that and ordered him to put all the matches in a box. Mirko soon noticed that not all of the matches on the floor fit in the box, so he decided to take the matches that don’t fit and throw them in the neighbour’s garbage, where his mom (hopefully) won’t find them.
//
// Help Mirko determine which of the matches fit in the box his mom gave him. A match fits in the box if its entire length can lie on the bottom of the box. Mirko examines the matches one by one.
//
// Input
// The first line of input contains an integer N (1≤N≤50), the number of matches on the floor, and two integers W and H, the dimensions of the box (1≤W≤100, 1≤H≤100).
//
// Each of the following N lines contains a single integer between 1 and 1000 (inclusive), the length of one match.
//
// Output
// For each match, in the order they were given in the input, output on a separate line “DA” if the match fits in the box or “NE” if it does not.
//
// Sample Input 1
// 5 3 4
// 3
// 4
// 5
// 6
// 7
//
// Sample Output 1
// DA
// DA
// DA
// NE
// NE
//
// Sample Input 2
// 2 12 17
// 21
// 20
//
// Sample Output 2
// NE
// DA

package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s := scanner.Text()
	nums := strings.Split(s, " ")

	wi, err := strconv.Atoi(nums[1])
	if err != nil {
		log.Fatalln(err)
	}
	he, err := strconv.Atoi(nums[2])
	if err != nil {
		log.Fatalln(err)
	}

	r := math.Sqrt(float64(wi*wi + he*he))
	for scanner.Scan() {
		s := scanner.Text()
		m, err := strconv.ParseFloat(s, 64)
		if err != nil {
			log.Fatalln(err)
		}

		if m <= r {
			w.Write([]byte("DA\n"))
		} else {
			w.Write([]byte("NE\n"))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	w.Flush()
}
