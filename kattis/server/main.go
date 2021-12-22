// Server
// /problems/server/file/statement/en/img-0001.jpg
// Picture from Wikimedia Commons
// You are in charge of a server that needs to run some submitted tasks on a first-come, first-served basis. Each day, you can dedicate the server to run these tasks for at most T minutes. Given the time each task takes, you want to know how many of them will be finished today.
//
// Consider the following example. Assume T=180 and the tasks take 45,30,55,20,80, and 20 minutes (in order they are submitted). Then, only four tasks can be completed. The first four tasks can be completed because they take 150 minutes, but not the first five, because they take 230 minutes which is greater than 180. Notice that although there is enough time to perform the sixth task (which takes 20 minutes) after completing the fourth task, you cannot do that because the fifth task is not done yet.
//
// Input
// The input consists of a single test case. The first line contains two integers n and T where 1≤n≤50 is the number of tasks and 1≤T≤500. The next line contains n positive integers no more than 100 indicating how long each task takes in order they are submitted.
//
// Output
// Display the number of tasks that can be completed in T minutes on a first-come, first-served basis.
//
// Sample Input 1
// 6 180
// 45 30 55 20 80 20
//
// Sample Output 1
// 4
//
// Sample Input 2
// 10 60
// 20 7 10 8 10 27 2 3 10 5
//
// Sample Output 2
// 5

package main

import (
	"bufio"
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
	scanner.Split(bufio.ScanWords)

	// skip the first word
	scanner.Scan()

	scanner.Scan()
	s := scanner.Text()
	t, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	var x, sum, c int
	for scanner.Scan() {
		s = scanner.Text()
		x, err = strconv.Atoi(s)
		if err != nil {
			return err
		}
		sum += x
		if sum <= t {
			c++
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w, c)

	return w.Flush()
}
