// Teque
// You have probably heard about the deque (double-ended queue) data structure, which allows for efficient pushing and popping of elements from both the front and back of the queue. Depending on the implementation, it also allows for efficient random access to any index element of the queue as well. Now, we want you to bring this data structure up to the next level, the teque (triple-ended queue)!
//
// The teque supports the following four operations:
//
// push_back x: insert the element x into the back of the teque.
//
// push_front x: insert the element x into the front of the teque.
//
// push_middle x: insert the element x into the middle of the teque. The inserted element x now becomes the new middle element of the teque. If k is the size of the teque before the insertion, the insertion index for x is (k+1)/2 (using 0-based indexing).
//
// get i: prints out the ith index element (0-based) of the teque.
//
// Input
// The first line contains an integer N (1≤N≤106) denoting the number of operations for the teque. Each of the next N lines contains a string S, denoting one of the above commands, followed by one integer x. If S is a push_back, push_front, or push_middle command, x (1≤x≤109), else for a get command, i (0≤i≤(size of teque)−1). We guarantee that the teque is not empty when any get command is given.
//
// Output
// For each get i command, print the value inside the ith index element of the teque in a new line.
//
// Warning
// The I/O files are large. Please use fast I/O methods.
//
// Sample Input 1
// 9
// push_back 9
// push_front 3
// push_middle 5
// get 0
// get 1
// get 2
// push_middle 1
// get 1
// get 2
//
// Sample Output 1
// 3
// 5
// 9
// 5
// 1

package main

import (
	"bufio"
	"bytes"
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
	const maxCapacity = 2 * 1024 * 1024
	w := bufio.NewWriterSize(os.Stdout, maxCapacity)

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	// skip the first number
	scanner.Scan()
	s := scanner.Text()
	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	q := newTeque(n)
	var op string
	var x, ind int
	var b []byte
	for scanner.Scan() {
		b = scanner.Bytes()
		ind = bytes.IndexByte(b, ' ')
		op = string(b[:ind])
		x, err = strconv.Atoi(string(b[ind+1:]))
		if err != nil {
			return err
		}

		switch op {
		case "push_front":
			q.pushFront(x)
		case "push_back":
			q.pushBack(x)
		case "push_middle":
			q.pushMiddle(x)
		case "get":
			fmt.Fprint(w, q.get(x))
			w.WriteByte('\n')
		}
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}

type teque struct {
	a     []int
	first int
	last  int
}

func newTeque(n int) *teque {
	return &teque{
		a:     make([]int, 2*n),
		first: 2 * n / 2,
		last:  2 * n / 2,
	}
}

func (t *teque) pushFront(x int) {
	t.first--
	t.a[t.first] = x
//	log.Println("front", t.first, t.last)
//	log.Println(t.a)
}

func (t *teque) pushBack(x int) {
	t.last++
	t.a[t.last-1] = x
//	log.Println("back", t.first, t.last)
//	log.Println(t.a)
}

func (t *teque) pushMiddle(x int) {
	k := t.last - t.first
	i := t.first + ((k + 1) / 2)
	copy(t.a[i+1:], t.a[i:])
	t.a[i] = x
	t.last++
//	log.Println("middle", t.first, t.last)
//	log.Println(t.a)
}

func (t *teque) get(i int) int {
	return t.a[t.first+i]
}
