// Filip
// Mirko has a younger brother, Filip, who just started going to school and is having trouble with numbers. To help him get the hang of the number scale, his teacher writes two three-digit numbers. She asks Filip to compare those numbers, but instead of interpreting them with the leftmost most significant digit, he needs to interpret them the other way around, with the most significant digit being the rightmost one. He then has to tell the teacher the larger of the two numbers.
//
// Write a program that will check Filip’s answers.
//
// Input
// The first and only line of input contains two three-digit numbers, A and B. A and B will not be equal and will not contain any zeroes.
//
// Output
// The first and only line of output should contain the larger of the numbers in the input, compared as described in the task. The number should be written reversed, to display to Filip how he should read it.
//
// Sample Input 1
// 734 893
//
// Sample Output 1
// 437
//
// Sample Input 2
// 221 231
//
// Sample Output 2
// 132
//
// Sample Input 3
// 839 237
//
// Sample Output 3
// 938

package main

import (
	"fmt"
	"log"
)

func main() {
	var n1, n2 string
	_, err := fmt.Scanln(&n1, &n2)
	if err != nil {
		log.Fatalln(err)
	}

	var larger string
	for i := len(n1) - 1; i >= 0; i-- {
		if larger != "" {
			fmt.Printf("%c", larger[i])
		} else if n1[i] > n2[i] {
			larger = n1
			fmt.Printf("%c", n1[i])
		} else if n2[i] > n1[i] {
			larger = n2
			fmt.Printf("%c", n2[i])
		} else {
			fmt.Printf("%c", n1[i])
		}
	}
}
