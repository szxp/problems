// Soda Slurper
// /problems/sodaslurper/file/statement/en/img-0001.jpg
// Tim is an absolutely obsessive soda drinker, he simply cannot get enough. Most annoyingly though, he almost never has any money, so his only obvious legal way to obtain more soda is to take the money he gets when he recycles empty soda bottles to buy new ones. In addition to the empty bottles resulting from his own consumption he sometimes find empty bottles in the street. One day he was extra thirsty, so he actually drank sodas until he couldn’t afford a new one.
//
// Input
// Three non-negative integers e,f,c, where e<1000 equals the number of empty soda bottles in Tim’s possession at the start of the day, f<1000 the number of empty soda bottles found during the day, and 2≤c<2000 the number of empty bottles required to buy a new soda.
//
// Output
// How many sodas did Tim drink on his extra thirsty day?
//
// Sample Input 1
// 9 0 3
//
// Sample Output 1
// 4
//
// Sample Input 2
// 5 5 2
//
// Sample Output 2
// 9

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

	if err := scanner.Err(); err != nil {
		return err
	}

	e, i := nextInt(b, 0)
	f, i := nextInt(b, i)
	c, i := nextInt(b, i)

	emptyBottles := e + f
	var drank int
	for {
		bought := emptyBottles / c
		if bought == 0 {
			break
		}

		drank += bought
		emptyBottles -= bought * c
		emptyBottles += bought
	}

	fmt.Fprint(w, drank)

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
