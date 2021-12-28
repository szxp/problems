package main

import (
	"bufio"
	"bytes"
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

	dict := make(map[string]string)
	var b []byte
	var i int
	for scanner.Scan() {
		b = scanner.Bytes()
		if len(b) == 0 {
			break
		}
		i = bytes.IndexByte(b, ' ')
		dict[string(b[i+1:])] = string(b[:i])
	}

	var en string
	var ok bool
	for scanner.Scan() {
		b = scanner.Bytes()
		en, ok = dict[string(b)]
		if ok {
			w.WriteString(en)
		} else {
			w.WriteString("eh")
		}
		w.WriteByte('\n')
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}
