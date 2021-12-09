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
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	const maxCapacity = 2 * 1024 * 1024
	w := bufio.NewWriterSize(os.Stdout, maxCapacity)

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	// skip the first number
	scanner.Scan()

	q := &teque{}
	var s, op, arg string
	var i, ind int
	var err error
    var e *element
	for scanner.Scan() {
		s = scanner.Text()
		ind = strings.IndexByte(s, ' ')
		op = s[:ind]
		arg = s[ind+1:]

		switch op {
		case "push_front":
			q.pushFront(arg)
		case "push_back":
			q.pushBack(arg)
		case "push_middle":
			q.pushMiddle(arg)
		case "get":
			i, err = strconv.Atoi(arg)
			if err != nil {
				log.Fatalln(err)
			}
            e = q.get(i)
            //fmt.Println("i", i, e)
            if e != nil {
			    w.Write([]byte(e.s))
			    w.Write([]byte("\n"))
            }
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	//q.dump()
	w.Flush()
}

type teque struct {
	size   int
	front  *element
	back   *element
	middle *element
}

func (t *teque) pushFront(s string) {
    e := t.push(s, t.front, nil)
	t.front = e
    if t.size == 1 {
        t.back = e
        t.middle = e
    } else if t.size%2 == 1 {
        t.middle = t.middle.prev
    }
    //t.dump()
}

func (t *teque) pushBack(s string) {
    e := t.push(s, nil, t.back)
	t.back = e
    if t.size == 1 {
        t.front = e
        t.middle = e
    } else if t.size%2 == 0 {
        t.middle = t.middle.next
    }
    //t.dump()
}

func (t *teque) pushMiddle(s string) {
    var e *element
    if t.middle != nil {
        if t.size%2 == 1 {
            e = t.push(s, t.middle.next, t.middle)
        } else {
            e = t.push(s, t.middle, t.middle.prev)
        }
    } else {
        e = t.push(s, nil, nil)
    }

    t.middle = e
    if t.size == 2 {
        t.back = e
    } else if t.size == 1 {
        t.front = e
        t.back = e
    }
    //t.dump()
}

func (t *teque) push(
    s string,
    next *element,
    prev *element,
) *element {
	e := &element{
		s:    s,
        next: next,
        prev: prev,
	}
    if prev != nil {
        prev.next = e
    }
    if next != nil {
        next.prev = e
    }
    t.size += 1
    return e
}

func (t *teque) get(i int) *element {
    fi := 0
    mi := t.size / 2
	bi := t.size - 1

    if i <= mi {
        hi := (mi - fi) / 2
        if i <= hi {
            return t.find(i, t.front, fi, true)
        } else {
            return t.find(i, t.middle, mi, false)
        }
    } else {
        hi := mi + ((bi - mi) / 2)
        if i <= hi {
            return t.find(i, t.middle, mi, true)
        } else {
            return t.find(i, t.back, bi, false)
        }
    }
}

func (t *teque) find(
    i int,
    from *element,
    fromi int,
    forward bool,
) *element {
    e := from
	for e != nil {
        if fromi == i {
            return e
        }
        if forward {
		    e = e.next
            fromi++
        } else {
		    e = e.prev
            fromi--
        }
	}
    return nil
}

func (t *teque) dump() {
	e := t.front
	for e != nil {
		fmt.Print(e.s, ",")
		e = e.next
	}
    fmt.Println()
    fmt.Println(
        "F", t.front.s,
        "M", t.middle.s,
        "B", t.back.s,
    )
}

type element struct {
	s    string
	next *element
	prev *element
}
