// ABC
// You will be given three integers A, B and C. The numbers will not be given in that exact order, but we do know that A is less than B and B less than C. In order to make for a more pleasant viewing, we want to rearrange them in a given order.
// 
// Input
// The first line contains the three positive integers A, B and C, not necessarily in that order. The three numbers will be less than or equal to 100.
// 
// The second line contains three uppercase letters ’A’, ’B’ and ’C’ (with no spaces between them) representing the desired order.
// 
// Output
// Output A, B and C in the desired order on a single line, separated by single spaces.
// 
// Sample Input 1	
// 1 5 3
// ABC
// 
// Sample Output 1
// 1 3 5
// 
// Sample Input 2	
// 6 4 2
// CAB
// 
// Sample Output 2
// 6 2 4

package main

import (
    "sort"
    "fmt"
    "log"
)

func main() {
    a := make([]int, 3)
    _, err := fmt.Scanln(&a[0], &a[1], &a[2])
    if err != nil {
        log.Fatalln(err)
    }
    sort.Ints(a)

    var order string
    _, err = fmt.Scanln(&order)
    if err != nil {
        log.Fatalln(err)
    }

    for i, c := range order {
        if i > 0 {
            fmt.Print(" ")
        }
        switch c {
        case 'A':
            fmt.Print(a[0])
        case 'B':
            fmt.Print(a[1])
        case 'C':
            fmt.Print(a[2])
        }
    }
}
