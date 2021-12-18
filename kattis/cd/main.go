// CD
// Jack and Jill have decided to sell some of their Compact Discs, while they still have some value. They have decided to sell one of each of the CD titles that they both own. How many CDs can Jack and Jill sell?
// 
// Neither Jack nor Jill owns more than one copy of each CD.
// 
// Input
// The input consists of a sequence of test cases. The first line of each test case contains two positive integers N and M, each at most one million, specifying the number of CDs owned by Jack and by Jill, respectively. This line is followed by N lines listing the catalog numbers of the CDs owned by Jack in increasing order, and M more lines listing the catalog numbers of the CDs owned by Jill in increasing order. Each catalog number is a positive integer no greater than one billion. The input is terminated by a line containing two zeros. This last line is not a test case and should not be processed.
// 
// Output
// For each test case, output a line containing one integer, the number of CDs that Jack and Jill both own.
// 
// Sample Input 1
// 3 3
// 1
// 2
// 3
// 1
// 2
// 4
// 0 0
// 
// Sample Output 1
// 2

package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    const maxCapacity = 1024*1024
    buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(buf, maxCapacity)

    var jackCDs []int

	for {
        scanner.Scan()
        s := scanner.Text()
        ind := strings.IndexByte(s, ' ')

        n, err := strconv.Atoi(s[:ind])
        if err != nil {
            log.Fatalln(err)
        }

        m, err := strconv.Atoi(s[ind+1:])
        if err != nil {
            log.Fatalln(err)
        }

        if n == 0 && m == 0 {
            break
        }

        jackCDs = make([]int, 0, n)
        for i:=0; i<n; i++ {
            scanner.Scan()
            s = scanner.Text()
            id, err := strconv.Atoi(s)
            if err != nil {
                log.Fatalln(err)
            }
            jackCDs = append(jackCDs, id)
        }

        count := 0
        i := 0
        for k:=0; k<m; k++ {
            scanner.Scan()
            s = scanner.Text()
            id, err := strconv.Atoi(s)
            if err != nil {
                log.Fatalln(err)
            }

            for ; i < n && jackCDs[i] < id; i++ {}
            if i < n && jackCDs[i] == id {
                count++
                i++
            }
        }
        fmt.Println(count)
    }
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(
            os.Stderr,
            "reading standard input:",
            err,
        )
	}
}
