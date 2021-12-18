// Bank Queue
// /problems/bank/file/statement/en/img-0001.jpg
// Oliver is a manager of a bank near KTH and wants to close soon. There are many people standing in the queue wanting to put cash into their accounts after they heard that the bank increased the interest rates by 42% (from 0.01% per year to 0.0142% per year).
// However, there are too many people and only one counter is open which can serve one person per minute. Greedy as Oliver is, he would like to select some people in the queue, so that the total amount of cash stored by these people is as big as possible and that money then can work for the bank overnight.
//
// There is a problem, though. Some people don’t have the time to wait until the bank closes because they have to run somewhere else, so they have to be served before a certain time, after which they just leave. Oliver also turned off the infrared door sensor outside the bank, so that no more people can enter, because it’s already too crowded in the hall.
//
// Task
// Help Oliver calculate how much cash he can get from the people currently standing in the queue before the bank closes by serving at most one person per minute.
//
// Input
// The first line of input contains two integers N (1≤N≤10000) and T (1≤T≤47), the number of people in the queue and the time in minutes until Oliver closes the bank. Then follow N lines, each with 2 integers ci and ti, denoting the amount of cash in Swedish crowns person i has and the time in minutes from now after which person i leaves if not served. Note that it takes one minute to serve a person and you must begin serving a person at time ti at the latest. You can assume that 1≤ci≤100000 and 0≤ti<T.
//
// Output
// Output one line with the maximum amount of money you can get from the people in the queue before the bank closes.
//
// Sample Input 1
// 4 4
// 1000 1
// 2000 2
// 500 2
// 1200 0
//
// Sample Output 1
// 4200
//
// Sample Input 2
// 3 4
// 1000 0
// 2000 1
// 500 1
//
// Sample Output 2
// 3000

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s := scanner.Text()
	nums := strings.Split(s, " ")
	n, err := strconv.Atoi(nums[0])
	if err != nil {
		log.Fatalln(err)
	}
	ct, err := strconv.Atoi(nums[1])
	if err != nil {
		log.Fatalln(err)
	}

	q := newQueue(ct)
	var sek, min int
	for i := 0; i < n; i++ {
		scanner.Scan()
		s = scanner.Text()
		nums = strings.Split(s, " ")
		sek, err = strconv.Atoi(nums[0])
		if err != nil {
			log.Fatalln(err)
		}
		min, err = strconv.Atoi(nums[1])
		if err != nil {
			log.Fatalln(err)
		}
		q.Add(sek, min)
	}

	//q.Dump()

	var sum int
	for i := ct - 1; i >= 0; i-- {
		sek = q.Next(i)
		//q.Dump()
		sum += sek
	}
	fmt.Println(sum)
}

type amount struct {
	SEK  int
	Next *amount
}

type queue struct {
	Wtimes []*amount
}

func newQueue(size int) *queue {
	return &queue{
		Wtimes: make([]*amount, size),
	}
}

func (q *queue) Add(sek, min int) {
	head := q.Wtimes[min]
	var prev *amount
	curr := head
	for curr != nil && sek < curr.SEK {
		prev = curr
		curr = curr.Next
	}

	a := &amount{
		SEK:  sek,
		Next: curr,
	}

	if prev != nil {
		prev.Next = a
	}

	if curr == head {
		q.Wtimes[min] = a
	}
}

func (q *queue) Next(i int) int {
	maxk := -1
	var curr, max *amount
	for k := i; k < len(q.Wtimes); k++ {
		curr = q.Wtimes[k]
		if curr != nil && (max == nil || curr.SEK > max.SEK) {
			max = curr
			maxk = k
		}
	}
	if maxk == -1 {
		return 0
	}
	q.Wtimes[maxk] = max.Next
	return max.SEK
}

func (q *queue) Dump() {
	for min, a := range q.Wtimes {
		fmt.Print(min, ": ")
		for a != nil {
			fmt.Print(a.SEK, ", ")
			a = a.Next
		}
		fmt.Println()
	}
}
