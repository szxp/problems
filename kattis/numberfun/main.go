// Number Fun
// /problems/numberfun/file/statement/en/img-0001.png
// Source: WikiMedia Commons
// Ms. Greene is trying to design a game for her third-grade class to practice their addition, subtraction, multiplication, and division. She would like for every student in her class to be able to “think mathematically” and determine if any two given numbers can be added, subtracted, multiplied, or divided in any way to produce a third given number. However, she would like to be able to check her students’ work quickly without having to think about it herself.
//
// Help Ms. Greene by writing a program that inputs two given numbers and determines whether or not it is possible to add, subtract, multiply, or divide those two numbers in any order to produce the third number. Only one operation may be used to produce the third number.
//
// Input
// Each input file will start with a single integer N (1≤N≤10000) denoting the number of test cases. The following N lines will contain sets of 3 numbers a,b,c (1≤a,b,c≤10000).
//
// Output
// For each test case, determine whether or not it is possible to produce the third number, c, using the first two numbers, a and b, using addition, subtraction, multiplication, or division.
//
// Sample Input 1
// 6
// 1 2 3
// 2 24 12
// 5 3 1
// 9 16 7
// 7 2 14
// 12 4 2
//
// Sample Output 1
// Possible
// Possible
// Impossible
// Possible
// Possible
// Impossible

package main

import (
	"bufio"
	//"fmt"
	"log"
	"math"
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
	const maxCapacity = 512 * 1024
	sbuf := make([]byte, maxCapacity)
	scanner.Buffer(sbuf, maxCapacity)

	// skip the first line
	scanner.Scan()

	var buf []byte
	var a, b, c, i int
	var x, y, z float64
	for scanner.Scan() {
		buf = scanner.Bytes()
		a, i = nextInt(buf, 0)
		b, i = nextInt(buf, i)
		c, i = nextInt(buf, i)

		x = float64(a)
		y = float64(b)
		z = float64(c)

		switch {
		case math.Abs(x+y-z) < 10e-9, math.Abs(x-y-z) < 10e-9 || math.Abs(y-x-z) < 10e-9, math.Abs(x*y-z) < 10e-9, math.Abs(x/y-z) < 10e-9 || math.Abs(y/x-z) < 10e-9:
			w.WriteString("Possible\n")
		default:
			w.WriteString("Impossible\n")
		}
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
