// Adding Words
// Psychologists at Wassamatta University believe that humans are able to more easily deal with words than with numbers. So they have devised experiments to find out if this is true. In an interesting twist, one of their experiments deals with using words which represent numbers. Rather than adding numbers, they want to add words. You are a research programmer on the project, and your job is to write a program that demonstrates this ability.
//
// Input
// Input is a sequence of up to 2000 commands, one per line, ending at end of file. Each command is a definition, a calculation, or a clear. All tokens within a command are separated by single spaces.
//
// A definition has the format def x y where x is a variable name and y is a an integer in the range [−1000,1000]. No two variables are ever defined to have the same numeric value at the same time. If x had been defined previously, defining it again erases its old definition. Variable names are up to 30 lowercase characters, each character from the range a to z.
//
// A calculation command starts with the word calc, and is followed by one to 15 more variable names separated by addition or subtraction operators. The end of a calc command is an equals sign. For example:
//
// calc foo + bar - car =
// The clear command instructs your program to forget all of its definitions.
//
// Output
// Your program should produce no output for definitions, but for calculations it should produce the value of the calculation. If there is not a word for the result, or some word in the calculation has not been defined, then the result of the calculation should be unknown. The word unknown is never used as a variable in the input.
//
// Sample Input 1
// def foo 3
// calc foo + bar =
// def bar 7
// def programming 10
// calc foo + bar =
// def is 4
// def fun 8
// calc programming - is + fun =
// def fun 1
// calc programming - is + fun =
// clear
//
// Sample Output 1
// foo + bar = unknown
// foo + bar = programming
// programming - is + fun = unknown
// programming - is + fun = bar

package main

import (
	"bufio"
	//"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)

	words := newWords()
	var s, q, res string
	var err error
	for scanner.Scan() {
		s = scanner.Text()

		if s == "clear" {
			words.Clear()
		} else if strings.HasPrefix(s, "def ") {
			err = words.Def(s[4:])
			if err != nil {
				log.Fatalln(err)
			}
		} else if strings.HasPrefix(s, "calc ") {
			q = s[5:]
			res, err = words.Calc(q)
			if err != nil {
				log.Fatalln(err)
			}
			w.WriteString(q)
			w.WriteByte(' ')
			w.WriteString(res)
			w.WriteByte('\n')
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	w.Flush()
}

type Words struct {
	vars   map[string]int
	values map[int]string
}

func newWords() *Words {
	return &Words{
		vars:   make(map[string]int),
		values: make(map[int]string),
	}
}

func (w *Words) Def(cmd string) error {
	args := strings.Split(cmd, " ")
	name := args[0]
	valNew, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	valOld, defined := w.vars[name]
	if defined {
		delete(w.values, valOld)
	}

	w.vars[name] = valNew
    w.values[valNew] = name
	return nil
}

func (w *Words) Calc(cmd string) (string, error) {
    //fmt.Println("Cmd", cmd)
	args := strings.Split(cmd, " ")

    var val, res int
	var defined, plus bool
	for i, a := range args {
		if a == "+" {
			plus = true
		} else if a == "-" {
			plus = false
		} else if a == "=" {
			name := w.values[res]
			if name == "" {
				return "unknown", nil
			}
            return name, nil
		} else {
			val, defined = w.vars[a]
			//fmt.Println(i, a, val, defined)
			if !defined {
				return "unknown", nil
			}

            if i == 0 || plus {
                res += val
            } else {
                res -= val
            }
		}
	}
	return "unknown", nil
}

func (w *Words) Clear() {
	w.vars = make(map[string]int)
	w.values = make(map[int]string)
}
