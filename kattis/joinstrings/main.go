// Join Strings
// You are given a collection of N non-empty strings, denoted by S1,S2,…,Sn. Then you are given N-1 operations which you execute in the order they are given. The ith operation is has the following format: ‘a b’ (1-based indexing, without the quotes), which means that you have to make the following changes:
// 
// Sa=Sa+Sb, i.e. concatenate ath string and bth string and store the result in ath string,
// 
// Sb = "", i.e. make the bth string empty, after doing the previous step.
// 
// You are ensured that after the ith operation, there will be no future operation that will be accessing Sb. Given these operations to join strings, print the last string that will remain at the end of this process.
// 
// Input
// The first line contains an integer N (1≤N≤105) denoting the number of strings given. Each of the next N lines contains a string denoting the Si. All the characters in the string Si are lowercase alphabets from ‘a’ to ‘z’. The total number of characters over all the strings is at most 106, i.e ∑Ni=1|Si|≤106, where |Si| denotes the length of the ith string. After these N strings, each of the next N-1 lines contain two integers a and b, such that a≠b and 1≤a,b≤N denoting the ith operation.
// 
// Output
// Print the last string which remains at the end of the N-1 operations.
// 
// Warning
// The I/O files are large. Please use fast I/O methods.
// 
// Sample Input 1	
// 4
// cute
// cat
// kattis
// is
// 3 2
// 4 1
// 3 4
// 
// Sample Output 1
// kattiscatiscute
// 
// Sample Input 2	
// 3
// howis
// this
// practicalexam
// 1 2
// 1 3
// 
// Sample Output 2
// howisthispracticalexam

package main

import (
    "bufio"
    "log"
    "strconv"
    "strings"
    "os"
    "io"
)

func main() {
    const maxCapacity = 2*1024*1024
    w := bufio.NewWriterSize(os.Stdout, maxCapacity)

    scanner := bufio.NewScanner(os.Stdin)
    buf := make([]byte, maxCapacity)
    scanner.Buffer(buf, maxCapacity)

    scanner.Scan()
    n, err := strconv.Atoi(scanner.Text())
    if err != nil {
        log.Fatalln(err)
    }

    v := make([]string, n)
    for i:=0; i<n; i++ {
        scanner.Scan()
        v[i] = scanner.Text()
    }

    tails := make([][]int, n)
    a := 1
    b := 0
    for scanner.Scan() {
        s := scanner.Text()
        ind := strings.IndexByte(s, ' ')

        a, err = strconv.Atoi(s[:ind])
        if err != nil {
            log.Fatalln(err)
        }

        b, err = strconv.Atoi(s[ind+1:])
        if err != nil {
            log.Fatalln(err)
        }

        tails[a-1] = append(tails[a-1], b-1)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalln(err)
    }

    //for i, id := range tails {
    //    log.Println(i, id)
    //}

    w.Write([]byte(v[a-1]))
    write(w, a-1, tails, v)
    w.Flush()
}

func write(
    w io.Writer,
    i int,
    tails [][]int,
    v []string,
) {
    for _, id := range tails[i] {
        w.Write([]byte(v[id]))
        if len(tails[id]) > 0 {
            write(w, id, tails, v)
        }
    }
}
