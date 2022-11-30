// Basic Programming 1
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
// Print 7, regardless of the content of A
//
// 2
//
// Print “Bigger” if A[0]>A[1], or
//
//
// Print “Equal” if A[0]==A[1], or
//
//
// Print “Smaller” otherwise (without the quotes);
//
//
// Ignore other indices of A, if any
//
// 3
//
// Print the median of three integers {A[0], A[1], and A[2]};
//
//
// Ignore other indices of A, if any
//
// 4
//
// Print the sum of all integers in A
//
// 5
//
// Print the sum of all even integers in A
//
// 6
//
// Apply modulo (%) 26 to each integer in A,
//
//
// Map integer 0/1/…/25 to character ‘a’/‘b’/…/‘z’,
//
//
// Finally, print the sequence of characters as a string (without the spaces)
//
// 7
//
// a. Start from index i=0;
//
//
// b. Jump to index i=A[i];
//
//
// c. If the current index i is outside the valid bound of [0..N-1], print “Out” and stop;
//
//
// d. Else if the current index i is index N-1, print “Done” and stop;
//
//
// e1. Otherwise, repeat step b;
//
//
// e2. If doing this leads to an infinite loop, print “Cyclic” and stop;
//
//
// (all output are without the quotes)
//
// Input
// The first line of the input contains an integer N and t (3≤N≤200000; 1≤t≤7).
// The second line of the input contains N non-negative 32-bit signed integers.
//
// Output
// For each test case, output the required answer based on the value of t.
//
// Scoring
// There are 20 hidden test cases that test various requirements of this problem.
// All 20 test cases will be tested.
// Each hidden test case worth 5 points (the 7 sample test cases below worth 0 point).
//
// Sample Input 1
// 7 1
// 1 2 3 4 5 6 7
//
// Sample Output 1
// 7
//
// Sample Input 2
// 7 2
// 1 2 3 4 5 6 7
//
// Sample Output 2
// Smaller
//
// Sample Input 3
// 7 3
// 1 2 3 4 5 6 7
//
// Sample Output 3
// 2
//
// Sample Input 4
// 7 4
// 1 2 3 4 5 6 7
//
// Sample Output 4
// 28
//
// Sample Input 5
// 7 5
// 1 2 3 4 5 6 7
//
// Sample Output 5
// 12
//
// Sample Input 6
// 10 6
// 7 4 11 37 14 22 40 17 11 3
//
// Sample Output 6
// helloworld
//
// Sample Input 7
// 3 7
// 1 0 2
//
// Sample Output 7
// Cyclic

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

	if err := scanner.Err(); err != nil {
		return err
	}

	switch t {
	case 1:
		w.WriteString("7")
	case 2:
		action2(w, n, a)
	case 3:
		sort.Ints(a[:3])
		fmt.Fprint(w, a[1])
	case 4:
		action4(w, n, a)
	case 5:
		action5(w, n, a)
	case 6:
		action6(w, n, a)
	case 7:
		action7(w, n, a)
	}

	return w.Flush()
}

func action2(
	w *bufio.Writer,
	n int,
	a []int,
) {
	if a[0] > a[1] {
		w.WriteString("Bigger")
	} else if a[0] == a[1] {
		w.WriteString("Equal")
	} else {
		w.WriteString("Smaller")
	}
}

func action4(
	w *bufio.Writer,
	n int,
	a []int,
) {
	var sum int
	for _, v := range a {
		sum += v
	}
	fmt.Fprint(w, sum)
}

func action5(
	w *bufio.Writer,
	n int,
	a []int,
) {
	var sum int
	for _, v := range a {
		if v%2 == 0 {
			sum += v
		}

	}
	fmt.Fprint(w, sum)
}

func action6(
	w *bufio.Writer,
	n int,
	a []int,
) {
	for _, v := range a {
		w.WriteByte(byte(v%26 + 97))
	}
}

func action7(
	w *bufio.Writer,
	n int,
	a []int,
) {
	used := make([]bool, n)

	i := 0
	for {
		i = a[i]
		if i < 0 || len(a)-1 < i {
			w.WriteString("Out")
			return
		}
		if i == len(a)-1 {
			w.WriteString("Done")
			return
		}
		if used[i] {
			w.WriteString("Cyclic")
			return
		}
		used[i] = true
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
