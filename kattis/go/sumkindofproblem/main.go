// Sum Kind of Problem
// For this problem you will compute various running sums of values for positive integers.
//
// Input
// The first line of input contains a single integer P, (1≤P≤10000), which is the number of data sets that follow. Each data set should be processed identically and independently. Each data set consists of a single line of input. It contains the data set number, K, followed by an integer N, (1≤N≤10000).
//
// Output
// For each data set there is one line of output. The single output line consists of the data set number, K, followed by a single space followed by three space separated integers S1, S2 and S3 such that:
//
// S1 = The sum of the first N positive integers.
//
// S2 = The sum of the first N odd integers.
//
// S3 = The sum of the first N even integers.
//
// Sample Input 1
// 3
// 1 1
// 2 10
// 3 1001
//
// Sample Output 1
// 1 1 1 2
// 2 55 100 110
// 3 501501 1002001 1003002

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

	// skip the first line
	scanner.Scan()

	for scanner.Scan() {
		b := scanner.Bytes()
		cas, i := nextInt(b, 0)
		n, i := nextInt(b, i)
		fmt.Fprintln(
			w,
			cas,
			(n*(n+1))/2,
			n*n,
			n*(n+1),
		)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
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
