// Above Average
// /problems/aboveaverage/file/statement/en/img-0001.jpg
// It is said that 90% of frosh expect to be above average in their class. You are to provide a reality check.
// Input
// The first line of standard input contains an integer 1≤C≤50, the number of test cases. C data sets follow. Each data set begins with an integer, N, the number of people in the class (1≤N≤1000). N integers follow, separated by spaces or newlines, each giving the final grade (an integer between 0 and 100) of a student in the class.
//
// Output
// For each case you are to output a line giving the percentage of students whose grade is above average, rounded to exactly 3 decimal places.
//
// Sample Input 1
// 5
// 5 50 50 70 80 100
// 7 100 95 90 80 70 60 50
// 3 70 90 80
// 3 70 90 81
// 9 100 99 98 97 96 95 94 93 91
//
// Sample Output 1
// 40.000%
// 57.143%
// 33.333%
// 66.667%
// 55.556%

package main

import (
	"bufio"
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
	scanner.Split(bufio.ScanWords)

	var b []byte
	scanner.Scan()
	b = scanner.Bytes()
	c, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	var n int
	for i := 0; i < c; i++ {
		scanner.Scan()
		b = scanner.Bytes()
		n, err = strconv.Atoi(string(b))
		if err != nil {
			return err
		}

		var g, sum int
		grades := make([]int, n)
		for k := 0; k < n; k++ {
			scanner.Scan()
			b = scanner.Bytes()
			g, err = strconv.Atoi(string(b))
			if err != nil {
				return err
			}
			grades[k] = g
			sum += g
		}

		var above int
		avg := float64(sum) / float64(len(grades))
		for _, g := range grades {
			if avg < float64(g) {
				above++
			}
		}

		ratio := (float64(above) / float64(len(grades))) * 100.0
		fmt.Fprintf(w, "%.3f%%\n", ratio)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}
