// Another Candies
// A class went to a school trip. And, as usually, all N kids have got their backpacks stuffed with candy. But soon quarrels started all over the place, as some of the kids had more candies than others. Soon, the teacher realized that he has to step in: "Everybody, listen! Put all the candies you have on this table here!"
//
// Soon, there was quite a large heap of candies on the teacher’s table. "Now, I will divide the candies into N equal heaps and everyone will get one of them." announced the teacher.
//
// "Wait, is this really possible?" wondered some of the smarter kids.
//
// Task
// You are given the number of candies each child brought. Find out whether the teacher can divide the candies into N exactly equal heaps. (For the purpose of this task, all candies are of the same type.)
//
// Input
// The first line of the input file contains an integer T,1≤T≤100 specifying the number of test cases. Each test case is preceded by a blank line.
//
// Each test case looks as follows: The first line contains N,1≤N≤20000 – the number of children. Each of the next N lines contains the number of candies one child brought. Each child has less than 263.
//
// Output
// For each of the test cases output a single line with a single word "YES" if the candies can be distributed equally, or "NO" otherwise.
//
// Sample Input 1
// 2
//
// 5
// 5
// 2
// 7
// 3
// 8
//
// 6
// 7
// 11
// 2
// 7
// 3
// 4
//
// Sample Output 1
// YES
// NO

package main

import (
	"bufio"
	//"fmt"
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
	t, _ := nextInt(b, 0)

	for i := 0; i < t; i++ {
		scanner.Scan()
		scanner.Scan()
		b = scanner.Bytes()
		n, _ := nextInt(b, 0)

		var sum uint64
		for k := 0; k < n; k++ {
			scanner.Scan()
			b = scanner.Bytes()
			c, _ := nextInt(b, 0)
			sum += uint64(c)
			sum = sum % uint64(n)
		}

		if sum == 0 {
			w.WriteString("YES\n")
		} else {
			w.WriteString("NO\n")
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
