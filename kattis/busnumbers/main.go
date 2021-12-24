// Bus Numbers
// Your favourite public transport company LS (we cannot use their real name here, so we permuted the letters) wants to change signs on all bus stops. Some bus stops have quite a few buses that stop there and listing all the buses takes space. However, if for example buses 141, 142, 143 stop there, we can write 141–143 instead and this saves space. Note that just for two buses this does not save space.
//
// You are given a list of buses that stop at a bus stop. What is the shortest representation of this list?
//
// Input
// The first line of the input contains one integer N,1≤N≤1000, the number of buses that stop at a bus stop. Then next line contains a list of N space separated integers between 1 and 1000, which denote the bus numbers. All numbers are distinct.
//
// Output
// Print the shortest representation of the list of bus numbers. Use the format as in the example, separate numbers with single spaces and output them in sorted order.
//
// Sample Input 1
// 6
// 180 141 174 143 142 175
//
// Sample Output 1
// 141-143 174 175 180

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

	// skip the first line
	scanner.Scan()

	scanner.Scan()
	b := scanner.Bytes()

	if err := scanner.Err(); err != nil {
		return err
	}

	buses := make([]bool, 1001)
	var x int
	for i := 0; i < len(b); {
		x, i = nextInt(b, i)
		buses[x] = true
	}

	var from, to, k int
	for i := 0; i < len(buses); {
		from, to, i = nextRange(buses, i)
		if from > 0 && from < len(buses) {
			k++
			if k > 1 {
				w.WriteByte(' ')
			}
			if from == to {
				fmt.Fprint(w, from)
			} else if to-from == 1 {
				fmt.Fprintf(w, "%d %d", from, to)
			} else {
				fmt.Fprintf(w, "%d-%d", from, to)
			}
		}
	}

	return w.Flush()
}

func nextRange(b []bool, i int) (from, to, ni int) {
	for ; i < len(b) && !b[i]; i++ {
	}
	from = i
	for ; i < len(b) && b[i]; i++ {
	}
	to = i - 1
	return from, to, i
}

func nextInt(b []byte, i int) (val, ni int) {
	for ; i < len(b) && !isDigit(b[i]); i++ {
	}
	x := 0
	for ; i < len(b) && isDigit(b[i]); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}

func isDigit(c byte) bool {
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}
	return false
}
