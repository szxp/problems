// Detailed Differences
// One of the most basic problems in information processing is identifying differences between data. This is useful when comparing files, for example. For this problem, write a program which identifies the differences between pairs of strings to make it easier for humans to see the differences.
//
// Your program should identify those characters which differ between the two given strings in a visually striking way. Output the two input strings on two lines, and then identify the differences on the line below using periods (for identical characters) and asterisks (for differing characters). For example:
//
// ATCCGCTTAGAGGGATT
// GTCCGTTTAGAAGGTTT
// *....*.....*..*..
// Input
// The first line of input contains an integer 1≤n≤500, indicating the number of test cases that follow. Each test case is a pair of lines of the same length, 1 to 50 characters. Each string contains only letters (a-z, A-Z) or digits (0-9).
//
// Output
// For each test case, output the two lines in the order they appear in the input. Output a third line indicating similarities and differences as described above. Finally, output a blank line after each test case.
//
// Sample Input 1
// 3
// ATCCGCTTAGAGGGATT
// GTCCGTTTAGAAGGTTT
// abcdefghijklmnopqrstuvwxyz
// bcdefghijklmnopqrstuvwxyza
// abcdefghijklmnopqrstuvwxyz0123456789
// abcdefghijklmnopqrstuvwxyz0123456789
//
// Sample Output 1
// ATCCGCTTAGAGGGATT
// GTCCGTTTAGAAGGTTT
// *....*.....*..*..
//
// abcdefghijklmnopqrstuvwxyz
// bcdefghijklmnopqrstuvwxyza
// **************************
//
// abcdefghijklmnopqrstuvwxyz0123456789
// abcdefghijklmnopqrstuvwxyz0123456789
// ....................................

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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
	w := bufio.NewWriter(os.Stdout)
	b := make([]byte, 5+500*2*52)

	n, err := readAll(os.Stdin, b)
	if err != nil {
		return err
	}
	b = b[:n]

	var m, k int
	i := bytes.IndexByte(b, '\n')
	b = b[i+1:] // skip the first line

	for {
		i = bytes.IndexByte(b, '\n')
		if i == -1 {
			break
		}
		m = i
		if b[i-1] == '\r' {
			m -= 1
		}

		// first string
		w.Write(b[:m])
		w.WriteByte('\n')

		// second string
		w.Write(b[i+1 : i+1+m])
		w.WriteByte('\n')

		// diff
		for k = 0; k < m; k++ {
			if b[k] == b[i+1+k] {
				w.WriteByte('.')
			} else {
				w.WriteByte('*')
			}
		}
		w.WriteByte('\n')
		w.WriteByte('\n')

		b = b[2*(i+1):]
	}

	return w.Flush()
}

func readAll(r io.Reader, b []byte) (int, error) {
	var n, nr int
	var err error
	for {
		if n == len(b) {
			return 0, fmt.Errorf("zero length")
		}
		nr, err = r.Read(b[n:])
		n += nr
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
	}
	return n, err

}
