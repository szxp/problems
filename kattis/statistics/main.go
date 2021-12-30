// Statistics
// Research often involves dealing with large quantities of data, and those data are often too massive to examine manually. Statistical descriptions of data can help humans understand their basic properties. Consider a sample of n numbers X=(x1,x2,…,xn). Of many statistics that can be computed on X, some of the most important are the following:
//
// min(X): the smallest value in X
//
// max(X): the largest value in X
//
// range(X): max(X)−min(X)
//
// Write a program that will analyze samples of data and report these values for each sample.
//
// Input
// The input contains between 1 and 10 test cases. Each test case is described by one line of input, which begins with an integer 1≤n≤30 and is followed by n integers which make up the sample to be analyzed. Each value in the sample will be in the range −1000000 to 1000000. Input ends at the end of file.
//
// Output
// For each case, display Case X:, where X is the case number, followed by the min, max, and range of the sample (in that order). Follow the format of the sample output.
//
// Sample Input 1
// 2 4 10
// 9 2 5 6 4 5 9 2 1 4
// 7 6 10 1 2 5 10 9
// 1 9
//
// Sample Output 1
// Case 1: 4 10 6
// Case 2: 1 9 8
// Case 3: 1 10 9
// Case 4: 9 9 0

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

	var cas, i, min, max, ran, v int
	var b []byte
	for scanner.Scan() {
		b = scanner.Bytes()
		_, i = nextInt(b, 0)
		cas++

		min = 1000000
		max = -1000000
		for i < len(b) {
			v, i = nextInt(b, i)
			if v < min {
				min = v
			}
			if max < v {
				max = v
			}
		}
		ran = max - min

		fmt.Fprintf(
			w,
			"Case %d: %d %d %d\n",
			cas,
			min,
			max,
			ran,
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
