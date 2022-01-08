// Dice Game
// /problems/dicegame/file/statement/en/img-0001.jpg
// Photo by JD Hancock
// Gunnar and Emma play a lot of board games at home, so they own many dice that are not normal 6-sided dice. For example they own a die that has 10 sides with numbers 47,48,…,56 on it.
//
// There has been a big storm in Stockholm, so Gunnar and Emma have been stuck at home without electricity for a couple of hours. They have finished playing all the games they have, so they came up with a new one. Each player has 2 dice which he or she rolls. The player with a bigger sum wins. If both sums are the same, the game ends in a tie.
//
// Task
// Given the description of Gunnar’s and Emma’s dice, which player has higher chances of winning?
//
// All of their dice have the following property: each die contains numbers a,a+1,…,b, where a and b are the lowest and highest numbers respectively on the die. Each number appears exactly on one side, so the die has b−a+1 sides.
//
// Input
// The first line contains four integers a1,b1,a2,b2 that describe Gunnar’s dice. Die number i contains numbers ai,ai+1,…,bi on its sides. You may assume that 1≤ai≤bi≤100. You can further assume that each die has at least four sides, so ai+3≤bi.
//
// The second line contains the description of Emma’s dice in the same format.
//
// Output
// Output the name of the player that has higher probability of winning. Output “Tie” if both players have same probability of winning.
//
// Sample Input 1
// 1 4 1 4
// 1 6 1 6
//
// Sample Output 1
// Emma
//
// Sample Input 2
// 1 8 1 8
// 1 10 2 5
//
// Sample Output 2
// Tie
//
// Sample Input 3
// 2 5 2 7
// 1 5 2 5
//
// Sample Output 3
// Gunnar

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

	scanner.Scan()
	b := scanner.Bytes()
	a1, i := nextInt(b, 0)
	b1, i := nextInt(b, i)
	a2, i := nextInt(b, i)
	b2, i := nextInt(b, i)
	gMin := a1 + a2
	gMax := b1 + b2
	gunnarSums := make([]float64, b1+b2+1)
	calcSumProbs(gunnarSums, a1, b1, a2, b2)

	scanner.Scan()
	b = scanner.Bytes()
	a1, i = nextInt(b, 0)
	b1, i = nextInt(b, i)
	a2, i = nextInt(b, i)
	b2, i = nextInt(b, i)
	eMin := a1 + a2
	eMax := b1 + b2
	emmaSums := make([]float64, b1+b2+1)
	calcSumProbs(emmaSums, a1, b1, a2, b2)

	if err := scanner.Err(); err != nil {
		return err
	}

	var gunnarAdv, emmaAdv float64
	sumMin := minInts(gMin, eMin)
	sumMax := maxInts(gMax, eMax)
	for sum := sumMin; sum <= sumMax; sum++ {
		gp := probHigherSum(sum, gunnarSums)
		ep := probHigherSum(sum, emmaSums)
		if ep < gp {
			gunnarAdv += gp - ep
		} else {
			emmaAdv += ep - gp
		}
	}

	diff := gunnarAdv - emmaAdv

	if -10e-6 < diff && diff < 10e-6 {
		w.WriteString("Tie")
	} else if 0 < diff {
		w.WriteString("Gunnar")
	} else {
		w.WriteString("Emma")
	}

	return w.Flush()
}

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInts(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func calcSumProbs(sums []float64, a1, b1, a2, b2 int) {
	p := (1.0 / float64(b1-a1+1)) * (1.0 / float64(b2-a2+1))

	for i := a1; i <= b1; i++ {
		for j := a2; j <= b2; j++ {
			sums[i+j] += p
		}
	}
}

func probHigherSum(sum int, sums []float64) float64 {
	var p float64
	for i := sum + 1; i < len(sums); i++ {
		if sum < i {
			p += sums[i]
		}
	}
	return p
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
