// Judging Moose
// /problems/judgingmoose/file/statement/en/img-0001.png
// Picture by Ryan Hagerty/US Fish and Wildlife Service, public domain
// When determining the age of a bull moose, the number of tines (sharp points), extending from the main antlers, can be used. An older bull moose tends to have more tines than a younger moose. However, just counting the number of tines can be misleading, as a moose can break off the tines, for example when fighting with other moose. Therefore, a point system is used when describing the antlers of a bull moose.
// The point system works like this: If the number of tines on the left side and the right side match, the moose is said to have the even sum of the number of points. So, “an even 6-point moose”, would have three tines on each side. If the moose has a different number of tines on the left and right side, the moose is said to have twice the highest number of tines, but it is odd. So “an odd 10-point moose” would have 5 tines on one side, and 4 or less tines on the other side.
//
// Can you figure out how many points a moose has, given the number of tines on the left and right side?
//
// Input
// The input contains a single line with two integers ℓ and r, where 0≤ℓ≤20 is the number of tines on the left, and 0≤r≤20 is the number of tines on the right.
//
// Output
// Output a single line describing the moose. For even pointed moose, output “Even x” where x is the points of the moose. For odd pointed moose, output “Odd x” where x is the points of the moose. If the moose has no tines, output “Not a moose”
//
// Sample Input 1
// 2 3
//
// Sample Output 1
// Odd 6
//
// Sample Input 2
// 3 3
//
// Sample Output 2
// Even 6
//
// Sample Input 3
// 0 0
//
// Sample Output 3
// Not a moose

package main

import (
	"fmt"
	"log"
)

func main() {
	var l, r int
	_, err := fmt.Scanln(&l, &r)
	if err != nil {
		log.Fatalln(err)
	}

	if l == 0 && r == 0 {
		fmt.Print("Not a moose")
	} else if l > r {
		fmt.Printf("Odd %d", 2*l)
	} else if l < r {
		fmt.Printf("Odd %d", 2*r)
	} else {
		fmt.Printf("Even %d", 2*l)
	}
}

