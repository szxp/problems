// Secret Message
// Jack and Jill developed a special encryption method, so they can enjoy conversations without worrrying about eavesdroppers. Here is how: let L be the length of the original message, and M be the smallest square number greater than or equal to L. Add (M−L) asterisks to the message, giving a padded message with length M. Use the padded message to fill a table of size K×K, where K2=M. Fill the table in row-major order (top to bottom row, left to right column in each row). Rotate the table 90 degrees clockwise. The encrypted message comes from reading the message in row-major order from the rotated table, omitting any asterisks.
//
// For example, given the original message ‘iloveyouJack’, the message length is L=12. Thus the padded message is ‘iloveyouJack****’, with length M=16. Below are the two tables before and after rotation.
//
// ilov
// eyou
// Jack
// ****
//
// *Jei
// *ayl
// *coo
// *kuv
//
// Then we read the secret message as ‘Jeiaylcookuv’.
//
// Input
// The first line of input is the number of original messages, 1≤N≤100. The following N lines each have a message to encrypt. Each message contains only characters a–z (lower and upper case), and has length 1≤L≤10000.
//
// Output
// For each original message, output the secret message.
//
// Sample Input 1
// 2
// iloveyoutooJill
// TheContestisOver
//
// Sample Output 1
// iteiloylloooJuv
// OsoTvtnheiterseC
//
//
//
// Design
//
// First let's find the number K,
// which the square root of M.
// If the square root of L is a whole number
// then it is K itself.
// But if it's not a whole number
// then we just need to round that number
// up to the nearest integer to get K.
//
// So we have a KxK table.
//
// The output is basically the same as
// reading the table before rotating
// in column major order
// starting from the bottom left corner
// from bottom to up
// and then left to right.
//
// The bottom left corner is k(k-1).
// From bottom to up
// read the characters at indexes
// k(k-1), k(k-1)-k, k(k-1)-2k, k(k-1)-3k, ... k(k-1)-(k-1)k.
// Then add one to them
// and read the characters again.
// Do this for a total of k times.
// Output characters only
// if the actual index
// is greater than or equal to zero
// and less than L.
//
// In this approach we do not have to
// pad the text with asterisks
// and build a two dimensional table.

package main

import (
	"bufio"
	"log"
	"math"
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
	defer w.Flush()

	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 10002
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	// skip the first line
	scanner.Scan()

	for scanner.Scan() {
		encrypt(w, scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func encrypt(w *bufio.Writer, b []byte) {
	l := len(b)
	k := rowLen(l)
	m := k * k
	for i := k * (k - 1); i < m; i++ {
		for j := i; 0 <= j; j -= k {
			if j < l {
				w.WriteByte(b[j])
			}
		}
	}
	w.WriteByte('\n')
}

func rowLen(l int) int {
	kf := math.Sqrt(float64(l))
	ki := int(kf)
	if kf == float64(ki) {
		return ki
	}
	return int(math.Ceil(kf))
}
