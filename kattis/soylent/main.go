// Soylent
// Yraglac recently decided to try out Soylent, a meal replacement drink designed to meet all nutritional requirements for an average adult. Soylent not only tastes great but is also low-cost, which is important for Yraglac as he is currently on a budget. Each bottle provides 400 calories, so it is recommended that an individual should consume 5 bottles a day for 2000 total calories. However, Yraglac is wondering how many bottles he should consume if his daily calorie requirement is not the same as an average adult. He can only consume an integer number of bottles, and needs to consume at least his daily calorie requirement.
//
// Input
// The first line contains a single integer T≤1000 giving the number of test cases. Each test case consists of a single line with an integer N (0≤N≤10000), the number of calories Yraglac needs in a day.
//
// Output
// For each test case, output a single line containing the number of bottles Yraglac needs to consume for the day.
//
// Sample Input 1
// 2
// 2000
// 1600
//
// Sample Output 1
// 5
// 4
//
//
// Design
//
// We need as many 400 kcal bottles that would cover the calories requirements per line.
//
//
// If the remainder of calories % 400 is zero then the answer is calories / 400. But if the remainder is not zero then the answer is (calories / 400) + 1.

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

	// skip the first line
	scanner.Scan()

	for scanner.Scan() {
		b := scanner.Bytes()
		n, _ := nextInt(b, 0)

		bottles := n / 400
		if n%400 != 0 {
			bottles++
		}
		fmt.Fprintln(w, bottles)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
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
