// Assigning Workstations
// /problems/workstations/file/statement/en/img-0001.jpg
// Picture by NASA via WikiMedia Commons
// Penelope is part of the admin team of the newly built supercomputer. Her job is to assign workstations to the researchers who come here to run their computations at the supercomputer.
// Penelope is very lazy and hates unlocking machines for the arriving researchers. She can unlock the machines remotely from her desk, but does not feel that this menial task matches her qualifications. Should she decide to ignore the security guidelines she could simply ask the researchers not to lock their workstations when they leave, and then assign new researchers to workstations that are not used any more but that are still unlocked. That way, she only needs to unlock each workstation for the first researcher using it, which would be a huge improvement for Penelope.
//
// Unfortunately, unused workstations lock themselves automatically if they are unused for more than m minutes. After a workstation has locked itself, Penelope has to unlock it again for the next researcher using it. Given the exact schedule of arriving and leaving researchers, can you tell Penelope how many unlockings she may save by asking the researchers not to lock their workstations when they leave and assigning arriving researchers to workstations in an optimal way? You may assume that there are always enough workstations available.
//
// Input
// The input consists of:
//
// one line with two integers n (1≤n≤300000), the number of researchers, and m (1≤m≤108), the number of minutes of inactivity after which a workstation locks itself;
//
// n lines each with two integers a and s (1≤a,s≤108), representing a researcher that arrives after a minutes and stays for exactly s minutes.
//
// Output
// Output the maximum number of unlockings Penelope may save herself.
//
// Sample Input 1
// 3 5
// 1 5
// 6 3
// 14 6
//
// Sample Output 1
// 2
//
// Sample Input 2
// 5 10
// 2 6
// 1 2
// 17 7
// 3 9
// 15 6
//
// Sample Output 2
// 3

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

	scanner.Scan()
	b := scanner.Bytes()
	n, i := nextInt(b, 0)
	m, i := nextInt(b, i)

	var a, s int
	arrive := make([]int, 0, n)
	finish := make([]int, 0, n)
	for scanner.Scan() {
		b = scanner.Bytes()
		a, i = nextInt(b, 0)
		s, i = nextInt(b, i)
		arrive = append(arrive, a)
		finish = append(finish, a+s)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	sort.Ints(arrive)
	sort.Ints(finish)

	var save int
	for _, a := range arrive {
		// skip closed free machines
		i := 0
		for ; i < len(finish) && finish[i]+m < a; i++ {
		}

		if i < len(finish) && finish[i] <= a && a <= finish[i]+m {
			save += 1
			i++
		}

		finish = finish[i:]
	}

	fmt.Fprint(w, save)

	return w.Flush()
}

func nextInt(b []byte, i int) (val, ni int) {
	for ; i < len(b) && !('0' <= b[i] && b[i] <= '9'); i++ {
	}
	x := 0
	for ; i < len(b) && '0' <= b[i] && b[i] <= '9'; i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}
