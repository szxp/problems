// 4 thought
// Write a program which, given an integer n as input, will produce a mathematical expression whose solution is n. The solution is restricted to using exactly four 4’s and exactly three of the binary operations selected from the set {∗,+,−,/}. The number 4 is the ONLY number you can use. You are not allowed to concatenate fours to generate other numbers, such as 44 or 444.
//
// For example given n=0, a solution is 4∗4−4∗4=0. Given n=7, a solution is 4+4−4 / 4=7. Division is considered truncating integer division, so that 1/4 is 0 (instead of 0.25). Assume the usual precedence of operations so that 4+4∗4=20, not 32. Not all integer inputs have solutions using four 4’s with the aforementioned restrictions (consider n=11).
//
// Hint: Using your forehead and some forethought should make an answer forthcoming. When in doubt use the fourth.
//
// Input
// Input begins with an integer 1≤m≤1000, indicating the number of test cases that follow. Each of the next m lines contain exactly one integer value for n in the range −1000000≤n≤1000000.
//
// Output
// For each test case print one line of output containing either an equation using four 4’s to reach the target number or the phrase no solution. Print the equation following the format of the sample output; use spaces to separate the numbers and symbols printed. If there is more than one such equation which evaluates to the target integer, print any one of them.
//
// Sample Input 1
// 5
// 9
// 0
// 7
// 11
// 24
//
// Sample Output 1
// 4 + 4 + 4 / 4 = 9
// 4 * 4 - 4 * 4 = 0
// 4 + 4 - 4 / 4 = 7
// no solution
// 4 * 4 + 4 + 4 = 24

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

	cache := make([]byte, (60+1+256)*3)
	c := newCalculator()
	ops := []byte{'+', '-', '*', '/'}
	e := []byte{'4', ' ', '+', ' ', '4', ' ', '+', ' ', '4', ' ', '+', ' ', '4'}

	var res int
	var op1, op2, op3 byte
	for k := 0; k < len(ops); k++ {
		for l := 0; l < len(ops); l++ {
			for m := 0; m < len(ops); m++ {
				op1 = ops[k]
				op2 = ops[l]
				op3 = ops[m]

				e[2] = op1
				e[6] = op2
				e[10] = op3
				res = c.evaluate(e)

				cache[(res+60)*3] = op1
				cache[(res+60)*3+1] = op2
				cache[(res+60)*3+2] = op3
			}
		}
	}

	const maxCapacity = 512 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)

	// skip the first line
	scanner.Scan()

	var s string
	var n int
	var err error
	for scanner.Scan() {
		s = scanner.Text()
		n, err = strconv.Atoi(s)
		if err != nil {
			return err
		}

		if n < -60 || 256 < n {
			w.WriteString("no solution\n")
			continue
		}

		op1 = cache[(n+60)*3]
		if op1 == '+' || op1 == '-' || op1 == '*' || op1 == '/' {
			op2 = cache[(n+60)*3+1]
			op3 = cache[(n+60)*3+2]

			fmt.Fprintf(
				w,
				"4 %c 4 %c 4 %c 4 = %d\n",
				op1,
				op2,
				op3,
				n,
			)
		} else {
			w.WriteString("no solution\n")
		}
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}

type calculator struct {
	q   []int
	ops []byte
}

func newCalculator() *calculator {
	return &calculator{
		q:   make([]int, 0, 4),
		ops: make([]byte, 0, 3),
	}
}

func (c *calculator) evaluate(e []byte) int {
	c.q = c.q[:0]
	c.ops = c.ops[:0]

	var op byte
	for _, b := range e {
		switch b {
		case '4':
			c.appendQueue(4)
		case '+', '-', '*', '/':
			for len(c.ops) > 0 {
				if b == '+' || b == '-' {
					c.popOp()
					continue
				}

				op = c.ops[len(c.ops)-1]
				if op == '*' || op == '/' {
					c.popOp()
					continue
				}

				break
			}
			c.pushOp(b)
		}
	}

	for len(c.ops) > 0 {
		c.popOp()
	}

	return c.q[0]
}

func (c *calculator) appendQueue(arg int) {
	c.q = append(c.q, arg)
}

func (c *calculator) pushOp(op byte) {
	c.ops = append(c.ops, op)
}

func (c *calculator) popOp() {
	q2 := len(c.q) - 2
	q1 := len(c.q) - 1
	o1 := len(c.ops) - 1

	switch c.ops[o1] {
	case '+':
		c.q[q2] = c.q[q2] + c.q[q1]
	case '-':
		c.q[q2] = c.q[q2] - c.q[q1]
	case '*':
		c.q[q2] = c.q[q2] * c.q[q1]
	case '/':
		c.q[q2] = c.q[q2] / c.q[q1]
	}

	c.q = c.q[:q1]
	c.ops = c.ops[:o1]
}
