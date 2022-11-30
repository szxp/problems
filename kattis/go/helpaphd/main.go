// Help a PhD candidate out!
// /problems/helpaphd/file/statement/en/img-0001.jpg
// Jon Marius
// Jon Marius forgot how to add two numbers while doing research for his PhD. And now he has a long list of addition problems that he needs to solve, in addition to his computer science ones! Can you help him?
// On his current list Jon Marius has two kinds of problems: addition problems on the form “a+b” and the ever returning problem “P=NP”. Jon Marius is a quite distracted person, so he might have to solve this last problem several times, since he keeps forgetting the solution. Also, he would like to solve these problems by himself, so you should skip these.
//
// Input
// The first line of input will be a single integer N (1≤N≤1000) denoting the number of testcases. Then follow N lines with either “P=NP” or an addition problem on the form “a+b”, where a,b∈[0,1000] are integers.
//
// Output
// Output the result of each addition. For lines containing “P=NP”, output “skipped”.
//
// Sample Input 1
// 4
// 2+2
// 1+2
// P=NP
// 0+0
//
// Sample Output 1
// 4
// 3
// skipped
// 0
//
//
//
// Design
//
// If the first character is P, the answer is "skipped".
// Otherwise parse the expression, evaluate it.

package main

import (
	"bufio"
	"bytes"
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

	var b []byte
	for scanner.Scan() {
		b = scanner.Bytes()

		if b[0] == 'P' {
			fmt.Fprintln(w, "skipped")
		} else {
			fmt.Fprintln(w, eval(b))
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func eval(expr []byte) int {
	i := bytes.IndexByte(expr, '+')
	a, _ := nextInt(expr[:i], 0)
	b, _ := nextInt(expr[i+1:], 0)
	return a + b
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
