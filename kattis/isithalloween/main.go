// IsItHalloween.com
// /problems/isithalloween/file/statement/en/img-0001.jpg
// Happy Halloween! Author: Petar Milošević, cc by-sa
// HiQ recently got an assignment from a client to create a clone of the immensely popular website https://IsItHalloween.com. The website is a very simple one. People will visit the site occasionally to see if it is Halloween. Whenever it is, the website should print out yup, otherwise it should print out nope on the screen.
//
// Since HiQ is such a popular firm, they don’t have time to complete this assignment right now. Their frontend engineers have already programmed the frontend of the website that prints out yup or nope, but not the backend microservice that determines whether it is indeed Halloween or not. Do you have time to help them?
//
// The behaviour of the server should be as follows: it gets as input the current date in the format FEB 9, where FEB is the month given in three letters (JAN, FEB, MAR, APR, MAY, JUN, JUL, AUG, SEP, OCT, NOV, DEC) and 9 is the day of the month starting at 1. It should then determine if this date represents October 31 or December 25 (since 318=2510).
//
// Input
// The input consists of a single line containing a date of the format FEB 9, with the month and date separated by a single space.
//
// Output
// If the date is October 31 or December 25, output yup. Otherwise, output nope.
//
// Sample Input 1
// OCT 31
//
// Sample Output 1
// yup
//
// Sample Input 2
// JUN 24
//
// Sample Output 2
// nope

package main

import (
	"bufio"
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
	b := make([]byte, 8)

	n, err := readAll(os.Stdin, b)
	if err != nil {
		return err
	}
	b = b[:n]

	if isOct31(b) || isDec25(b) {
		w.WriteString("yup")
	} else {
		w.WriteString("nope")
	}

	return w.Flush()
}

func isOct31(b []byte) bool {
	if len(b) < 6 {
		return false
	}
	return b[0] == 'O' && b[1] == 'C' && b[2] == 'T' && b[3] == ' ' && b[4] == '3' && b[5] == '1'
}

func isDec25(b []byte) bool {
	if len(b) < 6 {
		return false
	}
	return b[0] == 'D' && b[1] == 'E' && b[2] == 'C' && b[3] == ' ' && b[4] == '2' && b[5] == '5'
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
