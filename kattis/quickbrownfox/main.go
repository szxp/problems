// Quick Brown Fox
// /problems/quickbrownfox/file/statement/en/img-0001.jpg
// Photo by Neil McIntosh
// A pangram is a phrase that includes at least one occurrence of each of the 26 letters, ‘a’…‘z’. You’re probably familiar with this one: “The quick brown fox jumps over the lazy dog.”
//
// Your job is to recognize pangrams. For phrases that don’t contain every letter, report what letters are missing. We’ll say that a particular letter occurs in the phrase if it occurs as either upper case or lower case.
//
// Input
// Input starts with a line containing an integer 1≤N≤50. The next N lines are each a single phrase, possibly containing upper and lower case letters, spaces, decimal digits and punctuation characters ‘.’, ‘,’, ‘?’, ‘!’, ‘’’ and ‘"’. Each phrase contains at least one and no more than 100 characters.
//
// Output
// For each input phrase, output “pangram” if it qualifies as a pangram. Otherwise, output the word “missing” followed by a space and then the list of letters that didn’t occur in the phrase. The list of missing letters should be reported in lower case and should be sorted alphabetically.
//
// Sample Input 1
// 3
// The quick brown fox jumps over the lazy dog.
// ZYXW, vu TSR Ponm lkj ihgfd CBA.
// .,?!'" 92384 abcde FGHIJ
//
// Sample Output 1
// pangram
// missing eq
// missing klmnopqrstuvwxyz

package main

import (
	"bufio"
	//"fmt"
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
	scanner := bufio.NewScanner(os.Stdin)

	// skip the first ine
	scanner.Scan()

	var pangram bool
	var s string
	letters := make([]bool, 26)
	for scanner.Scan() {
		s = scanner.Text()

		for _, c := range s {
			if 65 <= c && c <= 90 {
				letters[c-65] = true
			}
			if 97 <= c && c <= 122 {
				letters[c-97] = true
			}
		}

		pangram = true
		for i, l := range letters {
			if !l {
				if pangram {
					w.WriteString("missing ")
					pangram = false
				}
				w.WriteByte(byte(i + 97))
			}
		}
		if !pangram {
			w.WriteByte('\n')
		} else {
			w.WriteString("pangram\n")
		}

		reset(letters)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}

func reset(a []bool) {
	for i, _ := range a {
		a[i] = false
	}
}
