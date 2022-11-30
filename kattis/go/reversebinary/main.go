// Reversed Binary Numbers
// Yi has moved to Sweden and now goes to school here. The first years of schooling she got in China, and the curricula do not match completely in the two countries. Yi likes mathematics, but now... The teacher explains the algorithm for subtraction on the board, and Yi is bored. Maybe it is possible to perform the same calculations on the numbers corresponding to the reversed binary representations of the numbers on the board? Yi dreams away and starts constructing a program that reverses the binary representation, in her mind. As soon as the lecture ends, she will go home and write it on her computer.
//
// Task
// Your task will be to write a program for reversing numbers in binary. For instance, the binary representation of 13 is 1101, and reversing it gives 1011, which corresponds to number 11.
//
// Input
// The input contains a single line with an integer N, 1≤N≤1000000000.
//
// Output
// Output one line with one integer, the number we get by reversing the binary representation of N.
//
// Sample Input 1
// 13
//
// Sample Output 1
// 11
//
// Sample Input 2
// 47
//
// Sample Output 2
// 61

package main

import (
	"fmt"
	"log"
	//"math"
)

func main() {
	var n int32
	_, err := fmt.Scanln(&n)
	if err != nil {
		log.Fatalln(err)
	}

	var r int32
	s := fmt.Sprintf("%b", n)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 49 { // digit one
			r += int32(1) << i
		}
	}
	fmt.Print(r)
}
