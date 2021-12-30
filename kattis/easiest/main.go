// The Easiest Problem Is This One
// Some people think this is the easiest problem in today’s problem set. Some people think otherwise, since it involves sums of digits of numbers and that’s difficult to grasp.
//
// If we multiply a number N with another number m, the sum of digits typically changes. For example, if m=26 and N=3029, then N⋅m=78754 and the sum of the digits is 31, while the sum of digits of N is 14.
//
// However, there are some numbers that if multiplied by N will result in the same sum of digits as the original number N. For example, consider m=37,N=3029, then N⋅m=112073, which has sum of digits 14, same as the sum of digits of N.
//
// Your task is to find the smallest positive integer p among those that will result in the same sum of the digits when multiplied by N. To make the task a little bit more challenging, the number must also be higher than ten.
//
// Input
// The input consists of several test cases. Each case is described by a single line containing one positive integer number N,1≤N≤100000. The last test case is followed by a line containing zero.
//
// Output
// For each test case, print one line with a single integer number p which is the minimal number such that N⋅p has the same sum of digits as N and p is bigger than 10.
//
// Sample Input 1
// 3029
// 4
// 5
// 42
// 0
//
// Sample Output 1
// 37
// 28
// 28
// 25

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

	var b []byte
	var n int
	for scanner.Scan() {
		b = scanner.Bytes()
		n, _ = nextInt(b, 0)
		if n == 0 {
			break
		}

		fmt.Fprintln(w, p(n))
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}

func p(n int) int {
	sum1 := sumDigits(n)
	for i := 11; ; i++ {
		sum2 := sumDigits(n * i)
		if sum2 == sum1 {
			return i
		}
	}
	return 0
}

func sumDigits(n int) int {
	var sum int
	for _, d := range fmt.Sprint(n) {
		sum += int(d) - '0'
	}
	return sum
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
