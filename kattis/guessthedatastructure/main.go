// I Can Guess the Data Structure!
// There is a bag-like data structure, supporting two operations:
// 
// 1 x: Throw an element x into the bag.
// 
// 2: Take out an element from the bag.
// 
// Given a sequence of operations with return values, you’re going to guess the data structure. It is a stack (Last-In, First-Out), a queue (First-In, First-Out), a priority-queue (Always take out larger elements first) or something else that you can hardly imagine!
// 
// Input
// There are several test cases. Each test case begins with a line containing a single integer n (1≤n≤1000). Each of the next n lines is either a type-1 command, or an integer 2 followed by an integer x. This means that executing the type-2 command returned the element x. The value of x is always a positive integer not larger than 100. The input is terminated by end-of-file (EOF). The size of input file does not exceed 1MB.
// 
// Output
// For each test case, output one of the following:
// 
// stack
// It’s definitely a stack.
// 
// queue
// It’s definitely a queue.
// 
// priority queue
// It’s definitely a priority queue.
// 
// impossible
// It can’t be a stack, a queue or a priority queue.
// 
// not sure
// It can be more than one of the three data structures mentioned above.
// 
// Sample Input 1	
// 6
// 1 1
// 1 2
// 1 3
// 2 1
// 2 2
// 2 3
// 6
// 1 1
// 1 2
// 1 3
// 2 3
// 2 2
// 2 1
// 2
// 1 1
// 2 2
// 4
// 1 2
// 1 1
// 2 1
// 2 2
// 7
// 1 2
// 1 5
// 1 1
// 1 3
// 2 5
// 1 4
// 2 4
// 1
// 2 1
// 
// Sample Output 1
// queue
// not sure
// impossible
// stack
// priority queue
// impossible

package main

import (
    "bufio"
    "os"
    "log"
    "strconv"
    "strings"
    "fmt"
)

func main() {
    w := bufio.NewWriter(os.Stdout)
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        s1 := scanner.Text()
        n, err := strconv.Atoi(s1)
        if err != nil {
            log.Fatalln(err)
        }

        stack := &list{}
        queue := &list{}
        queuePri := &list{}
        okStack := true
        okQueue := true
        okQueuePri := true
        for k:=0; k<n; k++ {
            scanner.Scan()
            s2 := scanner.Text()
            nums := strings.Split(s2, " ")
            op := nums[0]
            x, err := strconv.Atoi(nums[1])
            if err != nil {
                log.Fatalln(err)
            }

            if op == "1" {
                stack.Push(x)
                queue.Append(x)
                queuePri.AppendPriority(x)
            } else {
                rs := stack.Pop()
                okStack = okStack && rs != nil && *rs == x
                rq := queue.Pop()
                okQueue = okQueue && rq != nil && *rq == x
                rp := queuePri.Pop()
                okQueuePri = okQueuePri && rp != nil && *rp == x
            }
        }

        if (okStack && okQueue) ||
            (okQueue && okQueuePri) ||
            (okQueuePri && okStack) {
            w.Write([]byte("not sure\n"))
        } else if okStack {
            w.Write([]byte("stack\n"))
        } else if okQueue {
            w.Write([]byte("queue\n"))
        } else if okQueuePri {
            w.Write([]byte("priority queue\n"))
        } else {
            w.Write([]byte("impossible\n"))
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatalln(err)
    }

    w.Flush()
}

type list struct {
    size int
    front *entry
    back *entry
}

type entry struct {
    x int
    next *entry
    prev *entry
}

func (l *list) Push(x int) {
    e := &entry{
        x: x,
        next: l.front,
    }
    if l.front != nil {
        l.front.prev = e
    }
    l.front = e
    l.size += 1
    if l.size == 1 {
        l.back = e
    }
}

func (l *list) Append(x int) {
    e := &entry{
        x: x,
        prev: l.back,
    }
    if l.back != nil {
        l.back.next = e
    }
    l.back = e
    l.size += 1
    if l.size == 1 {
        l.front = e
    }
}

func (l *list) AppendPriority(x int) {
    curr := l.front
    for curr != nil && x < curr.x {
        curr = curr.next
    }
    e := &entry{
        x: x,
        next: curr,
    }
    if curr != nil {
        e.prev = curr.prev
        if e.prev != nil {
            e.prev.next = e
        }
        curr.prev = e

        if l.front == curr {
            l.front = e
        }
    } else {
        if l.back != nil {
            l.back.next = e
            e.prev = l.back
        }
        l.back = e
    }
    l.size += 1
    if l.size == 1 {
        l.front = e
        l.back = e
    }
}

func (l *list) Pop() *int {
    if l.front != nil {
        x := l.front.x
        if l.front.next != nil {
            l.front.next.prev = nil
        }
        l.front =  l.front.next
        l.size -= 1
        if l.size == 0 {
            l.back = nil
        }
        return &x
    }
    return nil
}

func (l *list) Dump() {
    curr := l.front
    for curr != nil {
        fmt.Println(curr.x)
        curr = curr.next
    }
}


