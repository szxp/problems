// I've Been Everywhere, Man
// /problems/everywhere/file/statement/en/img-0001.png
// Alice travels a lot for her work. Each time she travels, she visits a single city before returning home.
// Someone recently asked her “how many different cities have you visited for work?” Thankfully Alice has kept a log of her trips. Help Alice figure out the number of cities she has visited at least once.
//
// Input
// The first line of input contains a single positive integer T≤50 indicating the number of test cases. The first line of each test case also contains a single positive integer n indicating the number of work trips Alice has taken so far. The following n lines describe these trips. The ith such line simply contains the name of the city Alice visited on her ith trip.
//
// Alice’s work only sends her to cities with simple names: city names only contain lowercase letters, have at least one letter, and do not contain spaces.
//
// The number of trips is at most 100 and no city name contains more than 20 characters.
//
// Output
// For each test case, simply output a single line containing a single integer that is the number of distinct cities that Alice has visited on her work trips.
//
// Sample Input 1
// 2
// 7
// saskatoon
// toronto
// winnipeg
// toronto
// vancouver
// saskatoon
// toronto
// 3
// edmonton
// edmonton
// edmonton
//
// Sample Output 1
// 4
// 1

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s := scanner.Text()
	t, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < t; i++ {
		scanner.Scan()
		s := scanner.Text()
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalln(err)
		}

		cities := make(map[string]struct{}, n)
		for k := 0; k < n; k++ {
			scanner.Scan()
			s := scanner.Text()
			if _, ok := cities[s]; !ok {
				cities[s] = struct{}{}
			}
		}

		w.Write([]byte(fmt.Sprint(len(cities))))
		w.Write([]byte("\n"))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	w.Flush()
}
