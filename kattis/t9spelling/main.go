// T9 Spelling
// The Latin alphabet contains 26 characters and telephones only have ten digits on the keypad. We would like to make it easier to write a message to your friend using a sequence of keypresses to indicate the desired characters. The letters are mapped onto the digits as shown below. To insert the character ‘B’ for instance, the program would press “22”. In order to insert two characters in sequence from the same key, the user must pause before pressing the key a second time. The space character ‘ ’ should be printed to indicate a pause. For example, “2 2” indicates “AA” whereas “22” indicates “B”.
//
// \includegraphics[width=0.28\textwidth ]{keypad.png}
// Figure 1: Phone keypad
// Input
// The first line of input gives the number of cases, N,1≤N≤100. N test cases follow. Each case is a line of text containing the desired message, which will be at most 1000 characters long. Each message will consist of only lowercase characters ‘a’–‘z’ and space characters ‘ ’. Pressing zero emits a space.
//
// Output
// For each test case, output one line containing “Case #x: ” followed by the message translated into the sequence of key presses.
//
// Sample Input 1
// 4
// hi
// yes
// foo  bar
// hello world
//
// Sample Output 1
// Case #1: 44 444
// Case #2: 999337777
// Case #3: 333666 6660 022 2777
// Case #4: 4433555 555666096667775553

package main

import (
	"bufio"
	"fmt"
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

	// skip the first line
	scanner.Scan()

	var t int
	for scanner.Scan() {
		b := scanner.Bytes()
		t++
		w.WriteString("Case #")
		fmt.Fprint(w, t)
		w.WriteString(": ")

		var prev, curr string
		for i := 0; i < len(b); i++ {
			curr, _ = dict[b[i]]
			if prev != "" && prev[0] == curr[0] {
				w.WriteByte(' ')
			}
			w.WriteString(curr)
			prev = curr
		}
		w.WriteByte('\n')
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}

var dict = map[byte]string{
	'a': "2",
	'b': "22",
	'c': "222",
	'd': "3",
	'e': "33",
	'f': "333",
	'g': "4",
	'h': "44",
	'i': "444",
	'j': "5",
	'k': "55",
	'l': "555",
	'm': "6",
	'n': "66",
	'o': "666",
	'p': "7",
	'q': "77",
	'r': "777",
	's': "7777",
	't': "8",
	'u': "88",
	'v': "888",
	'w': "9",
	'x': "99",
	'y': "999",
	'z': "9999",
	' ': "0",
}
