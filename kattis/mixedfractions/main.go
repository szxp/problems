// Mixed Fractions
// You are part of a team developing software to help students learn basic mathematics. You are to write one part of that software, which is to display possibly improper fractions as mixed fractions. A proper fraction is one where the numerator is less than the denominator; a mixed fraction is a whole number followed by a proper fraction. For example the improper fraction 27/12 is equivalent to the mixed fraction 2 3/12. You should not reduce the fraction (i.e. don’t change 3/12 to 1/4).
//
// Input
// Input has one test case per line. Each test case contains two integers in the range [1,231−1]. The first number is the numerator and the second is the denominator. A line containing 0 0 will follow the last test case.
//
// Output
// For each test case, display the resulting mixed fraction as a whole number followed by a proper fraction, using whitespace to separate the output tokens.
//
// Sample Input 1
// 27 12
// 2460000 98400
// 3 4000
// 0 0
//
// Sample Output 1
// 2 3 / 12
// 25 0 / 98400
// 0 3 / 4000

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	var err error
	var s string
	var nums []string
	var n, d, whole, frac int
	for scanner.Scan() {
		s = scanner.Text()
		if s == "0 0" {
			break
		}

		nums = strings.Split(s, " ")
		n, err = strconv.Atoi(nums[0])
		if err != nil {
			return err
		}
		d, err = strconv.Atoi(nums[1])
		if err != nil {
			return err
		}

		whole = n / d
		frac = n - (whole * d)
		fmt.Fprintf(
			w,
			"%d %d / %d\n",
			whole,
			frac,
			d,
		)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}
