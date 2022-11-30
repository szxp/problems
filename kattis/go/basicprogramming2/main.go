// Basic Programming 2
// You think you can code?
// This problem will test you on various basic programming techniques.
// You are given two integers N and t; and then an array A of N integers (0-based indexing).
// Based on the value of t, you will perform an action on A.
//
// t
//
// Action Needed
//
// 1
//
// Print “Yes” if there are two integers x∈A and y∈A such that x≠y and x+y=7777,
//
//
// or “No” otherwise (without the quotes)
//
// 2
//
// Print “Unique” if all integers in A are different;
//
//
// or print “Contains duplicate” otherwise (without the quotes)
//
// 3
//
// Find and print the integer that appears >N2 times in A;
//
//
// or print −1 if such integer cannot be found
//
// 4
//
// Find and print the median integer of A if N is odd;
//
//
// or print both median integers of A if N is even (separate them with a single space)
//
// 5
//
// Print integers in A that fall between a range [100…999] in sorted order;
//
//
// (print a single space between two integers)
//
// Input
// The first line of the input contains an integer N and t (3≤N≤200000; 1≤t≤5).
// The second line of the input contains N non-negative 32-bit signed integers.
//
// Output
// For each test case, output the required answer based on the value of t.
//
// Scoring
// There are 20 hidden test cases that test various requirements of this problem.
// All 20 test cases will be tested.
// Each hidden test case worth 5 points (the 5 sample test cases below worth 0 point).
//
// Sample Input 1
// 7 1
// 1 7770 3 4 5 6 7
//
// Sample Output 1
// Yes
//
// Sample Input 2
// 7 2
// 1 2 3 4 5 6 7
//
// Sample Output 2
// Unique
//
// Sample Input 3
// 7 3
// 1 1 1 1 2 2 2
//
// Sample Output 3
// 1
//
// Sample Input 4
// 8 4
// 8 1 4 3 6 7 5 2
//
// Sample Output 4
// 4 5
//
// Sample Input 5
// 7 5
// 210 999 1000 543 321 99 777
//
// Sample Output 5
// 210 321 543 777 999

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	const maxCapacity = 2048 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	b := scanner.Bytes()
	n, i := nextInt(b, 0)
	t, i := nextInt(b, i)
	scanner.Scan()
	b = scanner.Bytes()
	a := make([]int, n)
	parseNums(a, b)
	sort.Ints(a)

	if err := scanner.Err(); err != nil {
		return err
	}

	switch t {
	case 1:
		action1(w, n, a)
	case 2:
		action2(w, n, a)
	case 3:
		action3(w, n, a)
	case 4:
		action4(w, n, a)
	case 5:
		action5(w, n, a)
	}

	return w.Flush()
}

func action1(
	w *bufio.Writer,
	n int,
	a []int,
) {
	i := len(a) - 1
	for ; 0 <= i && 7777 < a[i]; i-- {
	}

	j := 0
	for ; j < i; i-- {
		x := 7777 - a[i]

		for ; j < i && a[j] < x; j++ {
		}

		if a[j]+a[i] == 7777 {
			w.WriteString("Yes")
			return
		}
	}
	w.WriteString("No")
}

func action2(
	w *bufio.Writer,
	n int,
	a []int,
) {
	prev := -1
	for _, v := range a {
		if v != prev {
			prev = v
		} else {
			w.WriteString("Contains duplicate")
			return
		}
	}
	w.WriteString("Unique")
}

func action3(
	w *bufio.Writer,
	n int,
	a []int,
) {
	threshold := n / 2
	prev := -1
	var count int
	for _, v := range a {
		if v != prev {
			prev = v
			count = 1
		} else {
			count++
		}
		if threshold < count {
			fmt.Fprint(w, v)
			return
		}
	}
	w.WriteString("-1")
}

func action4(
	w *bufio.Writer,
	n int,
	a []int,
) {
	if n%2 == 0 {
		fmt.Fprint(w, a[n/2-1])
		w.WriteByte(' ')
		fmt.Fprint(w, a[n/2])
	} else {
		fmt.Fprint(w, a[n/2])
	}
}

func action5(
	w *bufio.Writer,
	n int,
	a []int,
) {
	var found bool
	for _, v := range a {
		if found {
			w.WriteByte(' ')
		}
		if 100 <= v && v <= 999 {
			found = true
			fmt.Fprint(w, v)
		}
	}
}

func parseNums(a []int, b []byte) {
	var v int
	k := -1
	for i := 0; i < len(b); {
		v, i = nextInt(b, i)
		k++
		a[k] = v
	}
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
