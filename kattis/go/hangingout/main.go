// Hanging Out on the Terrace
// The HiQ office in Stockholm has a pretty awesome rooftop terrace, often used in company parties and events such as programming competitions.
//
// Unfortunately, fire safety rules limit the number of people who can be on the terrace at any one point in time – at most L people. During a party, people come and go to the terrace, but it is pretty annoying to keep track of the number of people who are currently on the terrace. Furthermore, people often enter the terrace in groups. If a group of people wish to enter the terrace, but their addition would exceed the fire safety limit, the group will instead go and play ping pong inside.
//
// Your task is to write a program that determines, given the sizes of the groups which attempted to enter the terrace during a party and when people left the terrace, how many times a group was denied entry to the terrace.
//
// Input
// The first line of input contains the fire safety limit 1≤L≤200 and the number of events 0≤x≤100.
//
// The next x lines contains the events. An event starts with one of the words “enter” or “leave”, depending on whether the event describes a group attempting to enter the terrace or some set of people leaving it. This is followed by an integer 1≤p≤200 – the number of people entering/leaving at this time.
//
// The number of people who leave the terrace will never exceed the number of people currently on the terrace.
//
// Output
// Output the number of groups who were not allowed to enter the terrace during the party.
//
// Explanation of Sample Input 1
// There may be at most 4 people on the terrace at the same time. The first thing that happens is 3 people entering the terrace. Then, a group of 2 people attempt to enter. This would bring the total number up to 3+2=5. Since this is larger than 4, this group may not enter. A single person then leaves, meaning 2 people remain on the terrace. That person then comes back, bringing the total number up to three again. Finally, a pair of people try to enter the terrace, but is again denied since this would bring the total number up to 5.
//
// Thus, a total of 2 groups were not allowed to enter the terrace.
//
// Sample Input 1
// 4 5
// enter 3
// enter 2
// leave 1
// enter 1
// enter 2
//
// Sample Output 1
// 2
//
//
// Design
//
// We need a counter which keeps track of the number of people who are currently on he terrace.
//
// Another counter will be used to count the number of times when a group of people was not allowed to enter the terrace.
//
// Examining only the first character of an event is enough to decide if that event is an enter or exit event.

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
	defer w.Flush()

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	b := scanner.Bytes()
	limit, _ := nextInt(b, 0)

	var people, denied int
	for scanner.Scan() {
		b = scanner.Bytes()
		p, _ := nextInt(b, 6)
		if b[0] == 'e' { // enter
			if people+p <= limit {
				people += p
			} else {
				denied++
			}
		} else {
			people -= p
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Fprint(w, denied)
	return nil
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
