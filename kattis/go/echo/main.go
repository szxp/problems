// Echo Echo Echo
// If you have ever visited a canyon or a cave, you may have yelled and heard the echo of your own voice. In this problem, you should simulate that effect:
//
// Hello! Hello! Hello!
// Your program will be given as input a single word, and it should print out that word three times, separating the words with spaces.
//
// Input
// Input is a single line, containing exactly one word. The only characters used are upper and/or lowercase letters (a–z). The word contains at least one and at most 15 characters.
//
// Output
// Output the given word three times, separated by spaces.
//
// Sample Input 1
// Hello
//
// Sample Output 1
// Hello Hello Hello
//
// Sample Input 2
// ECHO
//
// Sample Output 2
// ECHO ECHO ECHO

package main

import (
	"errors"
	"fmt"
	"io"
	"log"
)

func main() {
	var w string
	_, err := fmt.Scanln(&w)
	if err != nil && !errors.Is(err, io.EOF) {
		log.Fatalln(err)
	}
	fmt.Printf("%s %s %s\n", w, w, w)
}
