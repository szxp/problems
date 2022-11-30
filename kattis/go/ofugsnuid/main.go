// Reverse
// Little Jóna needs a program. The program should read integers and print them in reverse order. Jóna asks for your help.
//
// Input
// The first line of input contains the integer n. Then comes a list of n integers, each on its own line.
//
// Output
// The list should be printed in reverse order of input.
//
// Points
// Group
//
// Points
//
// Input size
//
// 1
//
// 25
//
// n=1
//
// 2
//
// 25
//
// 1≤n≤5
//
// 3
//
// 25
//
// 1≤n≤103
//
// 4
//
// 25
//
// 1≤n≤105
//
// Sample Input 1
// 5
// 1
// 2
// 3
// 4
// 5
//
// Sample Output 1
// 5
// 4
// 3
// 2
// 1
//
// Sample Input 2
// 3
// 10
// 12
// 9
//
// Sample Output 2
// 9
// 12
// 10

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
	n, _ := nextInt(b, 0)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		b = scanner.Bytes()
		nums[i], _ = nextInt(b, 0)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	for i := n - 1; 0 <= i; i-- {
		fmt.Fprintln(w, nums[i])
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
