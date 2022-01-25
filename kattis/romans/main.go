// Roaming Romans
// /problems/romans/file/statement/en/img-0001.png
// The English word “mile” derives from the Latin “mille passus”, meaning “a thousand paces”. A Roman mile was the distance a soldier would walk in 1000 paces (a pace being two steps, one with each foot).
// Over time, the actual distance referred to as a “mile” has changed. The modern English mile is 5280 (modern) feet. The Roman mile is believed to have been about 4854 (modern) feet. Therefore a distance of x English miles would correspond to 1000⋅52804854 Roman paces.
//
// Write a program to convert distances in English miles into Roman paces.
//
// Input
// Input will consist of a single line containing a single real number 0≤X≤1000 denoting a distance in English miles. The number X has at most 3 digits of precision after the decimal point.
//
// Output
// Print an integer denoting the closest number of Roman paces equivalent to X. Your answer should be rounded to the closest integer (with an exact .5 decimal part rounded up).
//
// Sample Input 1	Sample Output 1
// 1.0
// 1088
// Sample Input 2	Sample Output 2
// 20.267
// 22046
//
//
//
// Design
//
// English miles (X) can be converted into Roman paces (R) as follows:
//
// R = Round( x * 1000 (5280 / 4854) )

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	defer w.Flush()

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	b := scanner.Bytes()

	if err := scanner.Err(); err != nil {
		return err
	}

	x, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return err
	}

	fmt.Fprint(w, int64(math.Round(x*1000.0*(5280.0/4854.0))))
	return nil
}
