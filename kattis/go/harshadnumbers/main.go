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
	n, _ := nextInt(b, 0)
	if err := scanner.Err(); err != nil {
		return err
	}

	for {
		sum := sumDigits(n)
		if n%sum == 0 {
			fmt.Fprint(w, n)
			break
		}
		n++
	}

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

func sumDigits(n int) int {
	var sum int
	for _, c := range fmt.Sprint(n) {
		sum += int(c) - '0'
	}
	return sum
}
