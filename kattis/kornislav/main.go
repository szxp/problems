// Kornislav
// Kornislav the turtle never has anything interesting to do. Since he’s going to live for three hundred years, he keeps trying to find way to kill time. This weekend he started playing "enclose the largest rectangle".
//
// To start with, Kornislav needs four positive integers. He tries to enclose a rectangle by moving in one direction, then turning 90 degrees, then walking in the new direction etc. Kornislav makes a total of three 90-degree turns and walks four segments.
//
// When walking in some direction, the number of steps he takes must be equal to one of the four chosen integers and each integer must be used exactly once. Depending on the order in which Kornislav uses the integers, his walk will make various shapes, some of which don’t contain enclosed rectangles. Write a program that calculates the largest rectangle the turtle can enclose with its walk.
//
// Input
// The first line contains four positive integers A,B,C and D (0<A,B,C,D<100), the four chosen integers.
//
// Output
// Output the largest area.
//
// Sample Input 1
// 1 2 3 4
//
// Sample Output 1
// 3
//
// Sample Input 2
// 4 4 3 4
//
// Sample Output 2
// 12
//
//
// Design
//
// How long can the longest side of the rectangle be?
// Take the minimum of the two largest numbers.
//
// What is the maximum length of the other side?
// Repeat the previous step for the remaining two numbers.
//
// The answer is the product of these minimum numbers.

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
	defer w.Flush()

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	b := scanner.Bytes()

	if err := scanner.Err(); err != nil {
		return err
	}

	var i int
	nums := make([]int, 4)
	nums[0], i = nextInt(b, 0)
	nums[1], i = nextInt(b, i)
	nums[2], i = nextInt(b, i)
	nums[3], i = nextInt(b, i)

	sort.Ints(nums)

	area := minInt(nums[3], nums[2]) * minInt(nums[1], nums[0])
	fmt.Fprint(w, area)

	return nil
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
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
