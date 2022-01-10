// Booking a Room
// /problems/bookingaroom/file/statement/en/img-0001.jpg
// Photo by Igor523 via Wikimedia Commons, cc by-sa
// Going to a contest such as NWERC is not all fun and games, there are also some worldly matters to tend to. One of these is to book hotel rooms in time, before all the rooms in town are booked.
// In this problem, you should write a program to search for available rooms in a given hotel. The hotel has r rooms, numbered from 1 to r, and you will be given a list describing which of these rooms are already booked.
//
// Input
// The input consists of:
//
// one line with two integers r and n (1≤r≤100, 0≤n≤r), the number of rooms in the hotel and the number of rooms that are already booked, respectively;
//
// n lines, each with an integer between 1 and r (inclusive), a room number that is already booked;
//
// All n room numbers of the already booked rooms are distinct.
//
// Output
// If there are available rooms, output the room number of any such room. Otherwise, output “too late”.
//
// Sample Input 1
// 100 5
// 42
// 3
// 2
// 99
// 1
//
// Sample Output 1
// 23
//
// Sample Input 2
// 3 3
// 2
// 3
// 1
//
// Sample Output 2
// too late
//
//
// Design
//
// The simplest case is when the number of rooms (r) and the number of booked rooms (n) is equal. In this case the asnwer is always "too late", regardless of the following n lines of input.
// 
// The remaining case is when the number of booked rooms (n) is less than the number of rooms (r).
// 
// Room numbers can be used as array indexes.
// Let's create an array of size r+1 to remember if a room is booked or not. The index 0 will never be used as a book number.
// 
// While reading the input we are going to set the value of the corresponding room in the array to true.
// 
// After processing the input we just need to find the first non-booked room in the array and output its index. It is the first index where the corresponding value is false.

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
	defer w.Flush()

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	b := scanner.Bytes()
	r, i := nextInt(b, 0)
	n, i := nextInt(b, i)

	if r == n {
		w.WriteString("too late")
		return nil
	}

	rooms := make([]bool, r+1)
	for scanner.Scan() {
		b = scanner.Bytes()
		m, _ := nextInt(b, 0)
		rooms[m] = true
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	var free int
	for i, v := range rooms {
		if 0 < i && !v {
			free = i
			break
		}
	}

	fmt.Fprint(w, free)
	return nil
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
