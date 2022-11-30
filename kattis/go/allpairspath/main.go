// All Pairs Shortest Path
// Input
// The input consists of several test cases. Each test case starts with a line with three non-negative integers, 1≤n≤150, 0≤m≤5000 and 1≤q≤1000, separated by single single spaces, where n is the numbers of nodes in the graph, m the number of edges and q the number of queries. Nodes are numbered from 0 to n−1. Then follow m lines, each line consisting of three (space-separated) integers u, v and w indicating that there is an edge from u to v in the graph with weight −1000≤w≤1000. Then follow q lines of queries, each consisting of two node numbers u and v (separated by a space), asking for the minimum distance from node u to node v.
//
// Input will be terminated by a line containing 0 0 0, this line should not be processed.
//
// Output
// For each query, output a single line containing the minimum distance from node u to v, or the word Impossible if there is no path from u to v, or -Infinity if there are arbitrarily short paths from u to v. Print a blank line after each test case.
//
// Sample Input 1
// 4 3 4
// 0 1 2
// 1 2 2
// 3 3 1
// 0 2
// 1 2
// 3 0
// 3 3
// 2 1 2
// 0 1 100
// 0 1
// 1 0
// 0 0 0
//
// Sample Output 1
// 4
// 2
// Impossible
// 0
//
// 100
// Impossible

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
	w := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)

	var b []byte
	var n, m, q, t, ind, u, v, we int
	var err error
	for scanner.Scan() {
		b = scanner.Bytes()
		if len(b) == 5 && b[0] == 48 && b[2] == 48 && b[4] == 48 {
			break
		}
		t++

		ind = bytes.IndexByte(b, ' ')
		n, err = strconv.Atoi(string(b[:ind]))
		if err != nil {
			return err
		}
		b = b[ind+1:]
		ind = bytes.IndexByte(b, ' ')
		m, err = strconv.Atoi(string(b[:ind]))
		if err != nil {
			return err
		}
		q, err = strconv.Atoi(string(b[ind+1:]))
		if err != nil {
			return err
		}

		if t > 1 {
			w.WriteByte('\n')
		}
		g := make([][]*edge, n)
		for i := 0; i < m; i++ {
			scanner.Scan()
			b = scanner.Bytes()
			ind = bytes.IndexByte(b, ' ')
			u, err = strconv.Atoi(string(b[:ind]))
			if err != nil {
				return err
			}
			b = b[ind+1:]
			ind = bytes.IndexByte(b, ' ')
			v, err = strconv.Atoi(string(b[:ind]))
			if err != nil {
				return err
			}
			we, err = strconv.Atoi(string(b[ind+1:]))
			if err != nil {
				return err
			}

			g[u] = append(g[u], &edge{
				to:     v,
				weight: we,
			})
		}

		p := make([]int, 0, n)
		var q1, q2 int
		for i := 0; i < q; i++ {
			scanner.Scan()
			b = scanner.Bytes()
			ind = bytes.IndexByte(b, ' ')
			q1, err = strconv.Atoi(string(b[:ind]))
			if err != nil {
				return err
			}
			q2, err = strconv.Atoi(string(b[ind+1:]))
			if err != nil {
				return err
			}

			maxInt := 1<<31 - 1
			weMin := maxInt
			p = p[:0]
			step(q1, 0, 0, 0, p, g, q2, &weMin)
			if weMin == maxInt {
				w.WriteString("Impossible\n")
			} else {
				fmt.Fprint(w, weMin)
				w.WriteByte('\n')
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}

func step(
	k int,
	we int,
	weCum int,
	depth int,
	p []int,
	g [][]*edge,
	exit int,
	weMin *int,
) {
	//	log.Println("step into", k, "dep", depth)
	p = append(p, k)

	if k == exit {
		//		log.Println("exit", k, p, weCum)
		if weCum < *weMin {
			*weMin = weCum
			//			log.Println("weMin", *weMin)
		}
		return
	}

	for _, n := range p[:depth] {
		if n == k {
			//			log.Println("circle", k, p)
			return
		}
	}

	for _, e := range g[k] {
		step(
			e.to,
			e.weight,
			weCum+e.weight,
			depth+1,
			p,
			g,
			exit,
			weMin,
		)
	}
}

type edge struct {
	to     int
	weight int
}
