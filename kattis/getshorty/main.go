// Get Shorty
// Nils and Mikael are intergalaxial fighters as well as lethal enemies. Now Nils has managed to capture the poor Mikael in his dark dungeons, and it is up to you to help Mikael escape with as much of his pride intact as possible.
//
// The dungeons can be viewed as a set of corridors and intersections. Each corridor joins two intersections. There are no guards, traps, or locked doors in Nils’ dungeon. However, there is one obstacle which makes escaping from the dungeon a perilious project: in each corridor there is a sentry, armed with a factor weapon. (As is commonly known, a factor weapon with factor f reduces the size of its target to a factor f of its original size, e.g. if Mikael is 8 gobs large and is hit by a factor weapon with factor f=0.25 his size will be reduced to 2 gobs.)
//
// Mikael will not be able to pass through a corridor without being hit by the factor weapon (but luckily enough, reloading the factor weapon takes enough time that the sentry will only have time to shoot him once as he passes through). It seems inevitable that Mikael will come out of this adventure a smaller man, but since the sentries have different factors in their factor weapons, his final size depends very much on the route he takes to the exit of the dungeons. Naturally, he would like to lose as little size as possible, and has asked you to help him accomplish that.
//
// Input
// Input consists of a series of test cases (at most 20). Each test case begins with a line consisting of two integers n, m separated by a single space, where 2≤n≤10000 is the number of intersections and 1≤m≤15000 is the number of corridors in Nils’ dungeon. Then follow m lines, each containing two integers x, y and a real number f (with at most four decimals), indicating that there is a corridor between intersections x and y, and that the factor weapon of the sentry in that corridor has factor 0≤f≤1. Intersections are numbered from 0 to n−1. Mikael always starts in intersection 0, and the exit is located in intersection n−1.
//
// The last case will be followed by a case where n=m=0, which should not be processed.
//
// Output
// For each test case, output a single line containing a real number (with exactly four decimals) indicating how big a fraction of Mikael will be left when he reaches the exit, assuming he chooses the best possible route through the dungeon. You may assume that it is always possible for Mikael to reach the exit.
//
// Sample Input 1
// 3 3
// 0 1 0.9
// 1 2 0.9
// 0 2 0.8
// 2 1
// 1 0 1
// 0 0
//
// Sample Output 1
// 0.8100
// 1.0000

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		s := scanner.Text()
		ind := strings.IndexByte(s, ' ')

		n, err := strconv.Atoi(s[:ind])
		if err != nil {
			log.Fatalln(err)
		}
		m, err := strconv.Atoi(s[ind+1:])
		if err != nil {
			log.Fatalln(err)
		}

		if n == 0 && m == 0 {
			break
		}

		dungeon := make([][]*dir, n)
		for i := 0; i < m; i++ {
			scanner.Scan()
			s := scanner.Text()
			nums := strings.Split(s, " ")

			x, err := strconv.Atoi(nums[0])
			if err != nil {
				log.Fatalln(err)
			}
			y, err := strconv.Atoi(nums[1])
			if err != nil {
				log.Fatalln(err)
			}
			factor, err := strconv.ParseFloat(
				nums[2],
				64,
			)
			if err != nil {
				log.Fatalln(err)
			}

			if factor > 0 {
				dungeon[x] = append(
					dungeon[x],
					&dir{
						id:     i,
						to:     y,
						factor: factor,
					},
				)
				dungeon[y] = append(
					dungeon[y],
					&dir{
						id:     i,
						to:     x,
						factor: factor,
					},
				)
			}
		}

		//        for m, dirs := range dungeon {
		//           log.Println(m, dirs)
		//        }

		p := make([]int, 0, n)
		max := make([]float64, n)
		step(
			dungeon,
			0,
			1,
			-1,
			0,
			p,
			max,
			n-1,
		)
		fmt.Printf("%.4f\n", max[n-1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

}

type dir struct {
	id     int
	to     int
	factor float64
}

func (d *dir) String() string {
	return fmt.Sprint(
		"to:", d.to,
		"factor:", d.factor,
	)
}

func step(
	dungeon [][]*dir,
	k int,
	mikael float64,
	dirFrom int,
	depth int,
	p []int,
	max []float64,
	exit int,
) {
	p = append(p, k)

	if max[k] < mikael {
		max[k] = mikael
	}

	if k == exit {
		return
	}

	// worse than the best path to the exit
	if mikael < max[len(max)-1] {
		return
	}

	// too much size is lost
	if mikael < max[k] {
		return
	}

	pre := p[:depth]
LOOP:
	for _, dir := range dungeon[k] {
		// avoid circle
		for _, n := range pre {
			if n == dir.to {
				continue LOOP
			}
		}

		if dir.id != dirFrom {
			step(
				dungeon,
				dir.to,
				mikael*dir.factor,
				dir.id,
				depth+1,
				p,
				max,
				exit,
			)
		}
	}
}
