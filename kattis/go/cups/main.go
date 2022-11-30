// Stacking Cups
// /problems/cups/file/statement/en/img-0001.jpg
// Photo by amy selleck cc by 2.0
// You are programming a cup stacking module for your robot. This robot is equiped with several sensors that can accurately determine the radius and color of a cup. The problem is that there is a glitch in the robot’s core routine that processess sensor inputs so the radius is doubled, if the result of the color sensor arrives to the routine after the radius.
// For instance, for a red cup with a radius of 5 units, your module will receive either “red 5” or “10 red” message.
//
// Given a list of messages from the core routine, each describing a different cup, can you put the cups in order of the smallest to the largest?
//
// Input
// The first line of the input file contains an integer N, the number of cups (1≤N≤20). Next N lines will contain two tokens each, either as “color radius” or “diameter color”. The radius of a cup R will be a positive integer less than 1000. The color of a cup C will be a non-empty string of lower case English letters of length at most 20. All cups will be different in both size and color.
//
// Output
// Output colors of cups, one color per line, in order of increasing radius.
//
// Sample Input 1
// 3
// red 10
// 10 blue
// green 7
//
// Sample Output 1
// blue
// green
// red
//
//
// Design
//
// We can detect if the first token is a diameter or a color by examining the first byte on the input line. If the first character is a digit then the forst token it is the diameter otherwise it is the color.
//
// We need to sort the radius values.
// And we need to remember which color belongs to which radius. There are to ways to achieve that.
//
// 1. Storing the radius values in an int array and sort that array in increasing order. And storing the colors in an radius->color map. In this case we just go through the radius array and output the corresponding color. It requires an array and a map, which is more than what is needed in the second approach.
//
// 2. Defining a custom struct type for a pair of the radius and the color. Store the pairs in an array and sort the array by the radius in ascending order. Then go through the array and output the colors. This approach requies less memory than the first one. We would use only one array.

package main

import (
	"bufio"
	//"fmt"
	"bytes"
	"log"
	"os"
	"sort"
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
	n, _ := nextInt(b, 0)

	cups := make([]cup, n)
	var r, k, ind int
	var c string
	for i := 0; i < n; i++ {
		scanner.Scan()
		b = scanner.Bytes()
		if 48 <= b[0] && b[0] <= 57 {
			r, k = nextInt(b, 0)
			r /= 2
			c = string(b[k+1:])
		} else {
			ind = bytes.IndexByte(b, ' ')
			c = string(b[:ind])
			r, _ = nextInt(b, ind+1)
		}
		cups[i] = cup{radius: r, color: c}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	sort.SliceStable(
		cups,
		func(i, j int) bool {
			return cups[i].radius < cups[j].radius
		},
	)

	for _, cup := range cups {
		w.WriteString(cup.color)
		w.WriteByte('\n')
	}

	return nil
}

type cup struct {
	radius int
	color  string
}

func nextInt(b []byte, i int) (val, ni int) {
	for ; i < len(b) && !('0' <= b[i] && b[i] <= '9' || b[i] == '-'); i++ {
	}
	x := 0
	sign := 1
	for ; i < len(b) && ('0' <= b[i] && b[i] <= '9' || b[i] == '-'); i++ {
		if b[i] == '-' {
			sign = -1
		} else {
			x = x*10 + int(b[i]) - '0'
		}
	}
	return x * sign, i
}
