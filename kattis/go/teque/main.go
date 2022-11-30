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

	scanner.Scan()
	b := scanner.Bytes()
	n, _ := nextInt(b, 0)

	q := newTeque(n)
	var x, ind int
	for scanner.Scan() {
		b = scanner.Bytes()
		ind = bytes.IndexByte(b, ' ')
		x, _ = nextInt(b, ind+1)

		switch {
		case b[0] == 'g': // get
			fmt.Fprint(w, q.get(x))
			w.WriteByte('\n')
		case b[5] == 'f': // push_front
			q.pushFront(x)
		case b[5] == 'b': // push_back
			q.pushBack(x)
		case b[5] == 'm': // push_middle
			q.pushMiddle(x)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}

type teque struct {
	a1  []int
	fp1 int
	lp1 int
	a2  []int
	fp2 int
	lp2 int
}

func newTeque(n int) *teque {
	return &teque{
		a1:  make([]int, 2*n),
		fp1: 2 * n / 2,
		lp1: 2 * n / 2,
		a2:  make([]int, 2*n),
		fp2: 2 * n / 2,
		lp2: 2 * n / 2,
	}
}

func (t *teque) pushFront(x int) {
	t.fp1--
	t.a1[t.fp1] = x
	if (t.lp1-t.fp1)-(t.lp2-t.fp2) == 2 {
		t.fp2--
		t.a2[t.fp2] = t.a1[t.lp1-1]
		t.lp1--
	}
}

func (t *teque) pushBack(x int) {
	t.lp2++
	t.a2[t.lp2-1] = x
	if (t.lp1 - t.fp1) < (t.lp2 - t.fp2) {
		t.lp1++
		t.a1[t.lp1-1] = t.a2[t.fp2]
		t.fp2++
	}
}

func (t *teque) pushMiddle(x int) {
	if t.lp1-t.fp1 == t.lp2-t.fp2 {
		t.lp1++
		t.a1[t.lp1-1] = x
	} else {
		t.fp2--
		t.a2[t.fp2] = x
	}
}

func (t *teque) get(i int) int {
	if i < t.lp1-t.fp1 {
		return t.a1[t.fp1+i]
	}
	return t.a2[t.fp2+(i-(t.lp1-t.fp1))]
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
