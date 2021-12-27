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
	const maxCapacity = 2048 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	b := scanner.Bytes()
	n, _ := nextInt(b, 0)
	//fmt.Fprintln(w, n)

	nums := make([]int, n)
	indices := make([]int, 100001)
	var a int
	k := -1
	scanner.Scan()
	b = scanner.Bytes()
	for i := 0; i < len(b); {
		a, i = nextInt(b, i)
		k++
		nums[k] = a
		indices[a+50000] += 1
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	sort.SliceStable(
		nums,
		func(i, j int) bool {
			return nums[i] < nums[j]
		},
	)

	//fmt.Println(nums)

	var x, y, sum, count int
	var ii, jj, kk, c int
	for a := 0; a < len(nums)-1; {
		b := a + 1
		//fmt.Println("a", a)
		x, a = next(nums, a)
		for b < len(nums) {
			//fmt.Println("b", b)
			y, b = next(nums, b)
			sum = x + y
			//fmt.Println(x, "+", y, "=", sum)

			if -50000 <= sum && sum <= 50000 {
				kk = indices[sum+50000]
				if kk > 0 {
					ii = indices[x+50000]
					jj = indices[y+50000]

					//fmt.Println("ii", ii, "jj", jj, "kk", kk)

					if x == 0 && y == 0 {
						jj -= 1
						kk -= 2
						c = ii * jj * kk
					} else if x == 0 || y == 0 {
						kk -= 1
						c = ii * jj * 2 * kk
					} else if x == y {
						jj -= 1
						c = ii * jj * kk
					} else {
						c = ii * jj * 2 * kk
					}

					//fmt.Println("c", c)
					count += c
				}
			}
		}
	}

	fmt.Fprint(w, count)

	return w.Flush()
}

func next(nums []int, i int) (v, ni int) {
	if i == len(nums) {
		return 0, i
	}
	v = nums[i]
	for ; i < len(nums) && nums[i] == v; i++ {
	}
	return v, i
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
