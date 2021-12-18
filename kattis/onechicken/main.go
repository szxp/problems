// One Chicken Per Person!
// Dr. Chaz is hosting a programming contest wrap up dinner. Dr. Chaz has severe OCD and is very strict on rules during dinner, specifically, he needs to be sure that everyone take exactly 1 piece of chicken at his buffet, even if that will result in an enormous amount of leftovers. This is why every year before the dinner, Dr. Chaz would give a powerful speech: “Everyone, one chicken per person!”
//
// However, Dr. Chaz does not always have an idea how many pieces of chicken he needs, he believes if there are N people at the buffet and everyone takes exactly 1 piece of chicken, providing M pieces of chicken will be perfect, i.e., is enough and will have no leftovers. Help Dr. Chaz find out whether his decision is good or not!
//
// Input
// The first line contain integers 0≤N≤1000, 0≤M≤1000, N≠M , the number of people at the buffet and the number of pieces of chicken Dr. Chaz is providing.
//
// Output
// Output a single line of the form “Dr. Chaz will have P piece[s] of chicken left over!”, if Dr. Chaz has enough chicken and P pieces of chicken will be left over, or “Dr. Chaz needs Q more piece[s] of chicken!” if Dr. Chaz does not have enough pieces of chicken and needs Q more.
//
// Sample Input 1
// 20 100
//
// Sample Output 1
// Dr. Chaz will have 80 pieces of chicken left over!
//
// Sample Input 2
// 2 3
//
// Sample Output 2
// Dr. Chaz will have 1 piece of chicken left over!
//
// Sample Input 3
// 10 1
//
// Sample Output 3
// Dr. Chaz needs 9 more pieces of chicken!

package main

import (
	"fmt"
	"log"
)

const unit1 = "piece"
const unit2 = "pieces"

func main() {
	var n, m int
	_, err := fmt.Scanln(&n, &m)
	if err != nil {
		log.Fatalln(err)
	}

	d := n - m
	unit := unit2
	if d == 1 || d == -1 {
		unit = unit1
	}
	if d < 0 {
		fmt.Printf(
			"Dr. Chaz will have %d "+
				"%s of chicken left over!",
			-1*d,
			unit,
		)
	} else {
		fmt.Printf(
			"Dr. Chaz needs %d "+
				"more %s of chicken!",
			d,
			unit,
		)
	}
}

