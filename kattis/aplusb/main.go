// A+B Problem
// Given N integers in the range [−50000,50000], how many ways are there to pick three integers ai, aj, ak, such that i, j, k are pairwise distinct and ai+aj=ak? Two ways are different if their ordered triples (i,j,k) of indices are different.
//
// Input
// The first line of input consists of a single integer N (1≤N≤200000). The next line consists of N space-separated integers a1,a2,…,aN.
//
// Output
// Output an integer representing the number of ways.
//
// Sample Input 1
// 4
// 1 2 3 4
//
// Sample Output 1
// 4
//
// Sample Input 2
// 6
// 1 1 3 3 4 6
//
// Sample Output 2
// 10

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
	const maxCapacity = 512 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	b := scanner.Bytes()
	n, _ := nextInt(b, 0)
	//fmt.Fprintln(w, n)

	nums := make([]num, n)
	ord := make([]*num, n)
	var a int
	k := -1
	scanner.Scan()
	b = scanner.Bytes()
	for i := 0; i < len(b); {
		a, i = nextInt(b, i)
		k++
		nums[k].val = a
		nums[k].ind = k
		ord[k] = &nums[k]
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	sort.SliceStable(
		ord,
		func(i, j int) bool {
			return ord[i].val < ord[j].val
		},
	)

	var count int
	for u, v := range ord {
		if v.val > 0 {
			count += pairsLeft(ord, u)
			count += pairsSides(ord, u)
		} else if v.val < 0 {
			count += pairsRight(ord, u)
			count += pairsSides(ord, u)
		} else {
			count += pairsLeft(ord, u)
			count += pairsRight(ord, u)
			count += pairsSides(ord, u)
		}
	}

	fmt.Fprint(w, count)

	return w.Flush()
}

func pairsLeft(ord []*num, u int) int {
	var count, sum int
	v := ord[u]
	var va, vb *num

	for a := u - 1; a > 0; a-- {
		b := a - 1
		for ; b >= 0; b-- {
			va = ord[a]
			vb = ord[b]
			sum = va.val + vb.val

			if sum < v.val {
				break
			}

			if sum == v.val {
				count += 2
			}
		}

		if b == a-1 {
			break
		}
	}
	return count
}

func pairsRight(ord []*num, u int) int {
	var count, sum int
	v := ord[u]
	var va, vb *num

	for a := u + 1; a < len(ord)-1; a++ {
		b := a + 1
		for ; b < len(ord); b++ {
			va = ord[a]
			vb = ord[b]
			sum = va.val + vb.val

			if sum > v.val {
				break
			}

			if sum == v.val {
				count += 2
			}
		}

		if b == a+1 {
			break
		}
	}
	return count
}

func pairsSides(ord []*num, u int) int {
	var count, sum int
	v := ord[u]
	var va, vb *num

	if u == 0 || u == len(ord)-1 {
		return 0
	}

	for a := u - 1; a >= 0; a-- {
		for b := u + 1; b < len(ord); b++ {
			va = ord[a]
			vb = ord[b]
			sum = va.val + vb.val

			if sum > v.val {
				break
			}

			if sum == v.val {
				count += 2
			}
		}
	}
	return count
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

type num struct {
	val int
	ind int
}
