// Datum
// Write a program that, given a date in 2009, determines the day of week on that date.
//
// Input
// The first line contains two positive integers D (day) and M (month) separated by a space. The numbers will be a valid date in 2009.
//
// Output
// Output the day of the week on day D of month M in 2009. The output should be one of the words "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday" or "Sunday".
//
// Sample Input 1
// 1 1
//
// Sample Output 1
// Thursday
//
// Sample Input 2
// 17 1
//
// Sample Output 2
// Saturday
//
// Sample Input 3
// 25 9
//
// Sample Output 3
// Friday

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

	scanner.Scan()
	b := scanner.Bytes()

	if err := scanner.Err(); err != nil {
		return err
	}

	d, i := nextInt(b, 0)
	m, i := nextInt(b, i)

	monthLengths := []int{
		31, 28, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31,
	}

	var days int
	for i, l := range monthLengths {
		if i < m-1 {
			days += l
		}
	}
	days += d

	dayNames := []string{
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
		"Monday",
		"Tuesday",
	}

	fmt.Fprint(w, dayNames[days%7])

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
