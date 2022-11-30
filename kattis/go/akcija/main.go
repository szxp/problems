// Akcija
// There is a promotional offer in a bookstore “Take 3, pay for the 2 more expensive ones”. So, each customer who picks 3 books gets the cheapest one for free. Of course, the customer can take even more books and, depending on the way the books are arranged into groups of three, get the cheapest one in each group for free.
//
// For example, let the prices of the books taken by the customer be: 2,3,4,4,6,9,10. If he arranges them into the groups (10,3,2), (4,6,4) and (9), he will get the book priced 2 from the first group for free and the book priced 4 from the second group. We can see that he will not get anything for free from the third group because it contains only one book.
//
// The lady working in the bookstore is well-intentioned and she always wants to lower the price for each customer as much as possible. For given book prices, help the lady arrange the books into groups in the best way possible, so that the total price the customer has to pay is minimal.
//
// Please note: The lady doesn’t have to arrange the books into groups so that each group contains exactly 3 books, but the number of books in a group needs to be between 1 and 3, inclusively.
//
// Input
// The first line of input contains the integer N (1≤N≤100 000), the number of books the customer bought. Each of the following N lines contains a single integer Ci (1≤Ci≤100 000), the price of each book.
//
// Output
// The first and only line of output must contain the required minimal price.
//
// Sample Input 1
// 4
// 3
// 2
// 3
// 2
//
// Sample Output 1
// 8
//
// Sample Input 2
// 6
// 6
// 4
// 5
// 5
// 5
// 5
//
// Sample Output 2
// 21

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

	scanner.Scan()
	b := scanner.Bytes()
	n, _ := nextInt(b, 0)

	prices := make([]int, n)

	var c int
	for i := 0; i < n; i++ {
		scanner.Scan()
		b = scanner.Bytes()
		c, _ = nextInt(b, 0)
		prices[i] = c
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	sort.Ints(prices)

	var total int
	for i := len(prices) - 1; 0 <= i; i-- {
		if i%3 != 2 {
			total += prices[len(prices)-1-i]
		}
	}

	fmt.Fprint(w, total)

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
