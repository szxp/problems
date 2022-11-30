// Zamka
// The impossible has happened. Bear G. has fallen into his own trap. Lured by a delicious box of Domaćica, without even thinking, he rushed and fell into his trap. In order to get out of the trap, he must solve the following task with your help. You are given three integers L, D and X.
//
// determine the minimal integer N such that L≤N≤D and the sum of its digits is X
//
// determine the maximal integer M such that L≤M≤D and the sum of its digits is X
//
// Bear will be able to escape from the trap if he correctly determines numbers N and M. The numbers N and M will always exist.
//
// Input
// The first line of input contains the integer L (1≤L≤10000), the number from the task. The second line of input contains the integer D (1≤D≤10000, L≤D), the number from the task. The third line of input contains the integer X (1≤X≤36), the number from the task.
//
// Output
// The first line of output must contain the integer N from the task. The second line of output must contain the integer M from the task.
//
// Sample Input 1
// 1
// 100
// 4
//
// Sample Output 1
// 4
// 40
//
// Sample Input 2
// 100
// 500
// 12
//
// Sample Output 2
// 129
// 480
//
// Sample Input 3
// 1
// 10000
// 1
//
// Sample Output 3
// 1
// 10000

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

	scanner.Scan()
	s := scanner.Text()
	l, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	scanner.Scan()
	s = scanner.Text()
	d, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	scanner.Scan()
	s = scanner.Text()
	x, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	min := find(l, d, 1, x)
	max := find(d, l, -1, x)

	fmt.Fprint(w, min)
	w.WriteByte('\n')
	fmt.Fprint(w, max)
	w.WriteByte('\n')

	return w.Flush()
}

func find(from, to, step, sum int) int {
	i := from
	for {
		if sumDigits(i) == sum {
			return i
		}
		if i == to {
			break
		}
		i += step
	}
	return from
}

func sumDigits(v int) int {
	sum := 0
	for _, c := range fmt.Sprint(v) {
		sum += int(c) - 48
	}
	return sum
}
