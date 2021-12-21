// Almost Perfect
// A positive integer p is called a perfect number if all the proper divisors of p sum to p exactly. Integer d is a proper divisor of p if 1≤d≤p−1 and p is evenly divisible by d. For example, the number 28 is a perfect number, since its proper divisors (which are 1, 2, 4, 7 and 14) add up to 28.
//
// Perfect numbers are rare; only 50 of them are known (as of 2017). Perhaps the definition of perfection is a little too strict. Instead, we will consider numbers that we’ll call almost perfect. Positive integer p is almost perfect if the proper divisors of p sum to a value that differs from p by no more than two.
//
// Input
// Input consists of a sequence of up to 500 integers, one per line. Each integer is in the range 2 to 109 (inclusive). Input ends at end of file.
//
// Output
// For each input value, output the same value and then one of the following: “perfect” (if the number is perfect), “almost perfect” (if it is almost perfect but not perfect), or “not perfect” (otherwise).
//
// Sample Input 1
// 6
// 65
// 650
//
// Sample Output 1
// 6 perfect
// 65 not perfect
// 650 almost perfect

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	var s string
	var err error
	var p, sum int
	for scanner.Scan() {
		s = scanner.Text()
		p, err = strconv.Atoi(s)
		if err != nil {
			return err
		}

		if _, ok := perfectNums[p]; ok {
			fmt.Fprintf(w, "%s perfect\n", s)
			continue
		}

		if p != 3 && p%2 == 1 {
			fmt.Fprintf(w, "%s not perfect\n", s)
			continue
		}

		sum = sumProperDivisors(p)
		if p-2 <= sum && sum <= p+2 {
			fmt.Fprintf(w, "%s almost perfect\n", s)
		} else {
			fmt.Fprintf(w, "%s not perfect\n", s)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}

func sumProperDivisors(n int) int {
	if n < 2 {
		return 1
	}
	sum := 0
	var other int
	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			other = n / i
			if i != other {
				sum += i + other
			} else {
				sum += i
			}
		}
	}
	return sum + 1
}

var perfectNums = map[int]struct{}{
	6:          struct{}{},
	28:         struct{}{},
	496:        struct{}{},
	8128:       struct{}{},
	33550336:   struct{}{},
	8589869056: struct{}{},
}
