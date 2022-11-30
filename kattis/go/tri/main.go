// Tri
// Little Mirko wrote into his math notebook an equation containing three positive integers, the equals sign and one of the four basic arithmetic operations (addition, subtraction, multiplication and division).
//
// During another class, his friend Slavko erased the equals sign and the operations from Mirko’s notebook. Help Mirko by reconstructing the equation from the three integers
//
// Input
// The first line of input contains three integers less than 100, separated by spaces. The input data will guarantee that a solution, although not necessarily unique, will always exist.
//
// Output
// On a single line, output a valid equation containing the three integers (in the same order), an equals sign and one of the four operations. If there are multiple solutions, output any of them.
//
// Sample Input 1
// 5 3 8
//
// Sample Output 1
// 5+3=8
//
// Sample Input 2
// 5 15 3
//
// Sample Output 2
// 5=15/3

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
	buf := scanner.Bytes()
	if err := scanner.Err(); err != nil {
		return err
	}

	a, i := nextInt(buf, 0)
	b, i := nextInt(buf, i)
	c, i := nextInt(buf, i)

	if a+b == c {
		fmt.Fprint(w, a, "+", b, "=", c)
	} else if a-b == c {
		fmt.Fprint(w, a, "-", b, "=", c)
	} else if a*b == c {
		fmt.Fprint(w, a, "*", b, "=", c)
	} else if a/b == c {
		fmt.Fprint(w, a, "/", b, "=", c)
	} else if a == b+c {
		fmt.Fprint(w, a, "=", b, "+", c)
	} else if a == b-c {
		fmt.Fprint(w, a, "=", b, "-", c)
	} else if a == b*c {
		fmt.Fprint(w, a, "=", b, "*", c)
	} else if a == b/c {
		fmt.Fprint(w, a, "=", b, "/", c)
	}

	return w.Flush()
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
