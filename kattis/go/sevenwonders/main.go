// Seven Wonders
// /problems/sevenwonders/file/statement/en/img-0001.jpg
// Seven Wonders is a card drafting game in which players build structures to earn points. The player who ends with the most points wins. One winning strategy is to focus on building scientific structures. There are three types of scientific structure cards: Tablet (‘T’), Compass (‘C’), and Gear (‘G’). For each type of cards, a player earns a number of points that is equal to the squared number of that type of cards played. Additionally, for each set of three different scientific cards, a player scores 7 points.
// For example, if a player plays 3 Tablet cards, 2 Compass cards and 1 Gear card, she gets 32+22+12+7=21 points.
//
// It might be tedious to calculate how many scientific points a player gets by the end of each game. Therefore, you are here to help write a program for the calculation to save everyone’s time.
//
// Input
// The input has a single string with no more than 50 characters. The string contains only letters ‘T’, ‘C’ or ‘G’, which denote the scientific cards a player has played in a Seven Wonders game.
//
// Output
// Output the number of scientific points the player earns.
//
// Note
// Seven Wonders was created by Antoine Bauza, and published by Repos Production. Antoine Bauza and Repos Production do not endorse and have no involvement with the ProgNova contest.
//
// Sample Input 1
// TCGTTC
//
// Sample Output 1
// 21
//
// Sample Input 2
// CCC
//
// Sample Output 2
// 9
//
// Sample Input 3
// TTCCGG
//
// Sample Output 3
// 26

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

	scanner.Scan()
	b := scanner.Bytes()

	var t, c, g, min, sum int
	for _, v := range b {
		if v == 'T' {
			t++
		}
		if v == 'C' {
			c++
		}
		if v == 'G' {
			g++
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	min = minInt(t, c)
	min = minInt(min, g)
	sum = t*t + c*c + g*g + min*7
	fmt.Fprint(w, sum)
	return w.Flush()
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
