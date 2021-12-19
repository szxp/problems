// A New Alphabet
// /problems/anewalphabet/file/statement/en/img-0001.jpg
// Photo by r. nial bradshaw
// A New Alphabet has been developed for Internet communications. While the glyphs of the new alphabet don’t necessarily improve communications in any meaningful way, they certainly make us feel cooler.
//
// You are tasked with creating a translation program to speed up the switch to our more elite New Alphabet by automatically translating ASCII plaintext symbols to our new symbol set.
//
// The new alphabet is a one-to-many translation (one character of the English alphabet translates to anywhere between 1 and 6 other characters), with each character translation as follows:
//
// Original
//
// New
//
// English Description
//
// Original
//
// New
//
// English Description
//
// a
//
// @
//
// at symbol
//
// n
//
// []\[]
//
// brackets, backslash, brackets
//
// b
//
// 8
//
// digit eight
//
// o
//
// 0
//
// digit zero
//
// c
//
// (
//
// open parenthesis
//
// p
//
// |D
//
// bar, capital D
//
// d
//
// |)
//
// bar, close parenthesis
//
// q
//
// (,)
//
// parenthesis, comma, parenthesis
//
// e
//
// 3
//
// digit three
//
// r
//
// |Z
//
// bar, capital Z
//
// f
//
// #
//
// number sign (hash)
//
// s
//
// $
//
// dollar sign
//
// g
//
// 6
//
// digit six
//
// t
//
// ']['
//
// quote, brackets, quote
//
// h
//
// [-]
//
// bracket, hyphen, bracket
//
// u
//
// |_|
//
// bar, underscore, bar
//
// i
//
// |
//
// bar
//
// v
//
// \/
//
// backslash, forward slash
//
// j
//
// _|
//
// underscore, bar
//
// w
//
// \/\/
//
// four slashes
//
// k
//
// |<
//
// bar, less than
//
// x
//
// }{
//
// curly braces
//
// l
//
// 1
//
// digit one
//
// y
//
// `/
//
// backtick, forward slash
//
// m
//
// []\/[]
//
// brackets, slashes, brackets
//
// z
//
// 2
//
// digit two
//
// For instance, translating the string “Hello World!” would result in:
//
// [-]3110 \/\/0|Z1|)!
// Note that uppercase and lowercase letters are both converted, and any other characters remain the same (the exclamation point and space in this example).
//
// Input
// Input contains one line of text, terminated by a newline. The text may contain any characters in the ASCII range 32–126 (space through tilde), as well as 9 (tab). Only characters listed in the above table (A–Z, a–z) should be translated; any non-alphabet characters should be printed (and not modified). Input has at most 10000 characters.
//
// Output
// Output the input text with each letter (lowercase and uppercase) translated into its New Alphabet counterpart.
//
// Sample Input 1
// All your base are belong to us.
//
// Sample Output 1
// @11 `/0|_||Z 8@$3 @|Z3 8310[]\[]6 ']['0 |_|$.
//
// Sample Input 2
// What's the Frequency, Kenneth?
//
// Sample Output 2
// \/\/[-]@'][''$ ']['[-]3 #|Z3(,)|_|3[]\[](`/, |<3[]\[][]\[]3']['[-]?
//
// Sample Input 3
// A new alphabet!
//
// Sample Output 3
// @ []\[]3\/\/ @1|D[-]@83']['!

package main

import (
	"bufio"
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

	var s string
	for scanner.Scan() {
		s = scanner.Text()

		for _, c := range s {
			if (65 <= c && c <= 90) || (97 <= c && c <= 122) {
				w.WriteString(tr(c))
			} else {
				w.WriteRune(c)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}

func tr(char rune) string {
	switch char {
	case 65, 97:
		return a
	case 66, 98:
		return b
	case 67, 99:
		return c
	case 68, 100:
		return d
	case 69, 101:
		return e
	case 70, 102:
		return f
	case 71, 103:
		return g
	case 72, 104:
		return h
	case 73, 105:
		return i
	case 74, 106:
		return j
	case 75, 107:
		return k
	case 76, 108:
		return l
	case 77, 109:
		return m
	case 78, 110:
		return n
	case 79, 111:
		return o
	case 80, 112:
		return p
	case 81, 113:
		return q
	case 82, 114:
		return r
	case 83, 115:
		return s
	case 84, 116:
		return t
	case 85, 117:
		return u
	case 86, 118:
		return v
	case 87, 119:
		return w
	case 88, 120:
		return x
	case 89, 121:
		return y
	case 90, 122:
		return z
	}
	return ""
}

const a = "@"
const b = "8"
const c = "("
const d = "|)"
const e = "3"
const f = "#"
const g = "6"
const h = "[-]"
const i = "|"
const j = "_|"
const k = "|<"
const l = "1"
const m = "[]\\/[]"
const n = "[]\\[]"
const o = "0"
const p = "|D"
const q = "(,)"
const r = "|Z"
const s = "$"
const t = "']['"
const u = "|_|"
const v = "\\/"
const w = "\\/\\/"
const x = "}{"
const y = "`/"
const z = "2"
