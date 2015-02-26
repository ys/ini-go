package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/scanner"
)

const (
	LBRACK   = "["
	RBRACK   = "]"
	COMMENT1 = ";"
	COMMENT2 = "#"
	ASSIGN   = "="
	EOL      = "\n"
)

type Ini map[string]KeyVals

type KeyVals map[string]string

func readSection(s *scanner.Scanner) string {
	tok := s.Scan()
	var val string
	for tok != scanner.EOF {
		c := s.TokenText()
		if c == RBRACK {
			return val
		} else {
			val = val + c
		}
		tok = s.Scan()
	}
	fmt.Println("Error near: [", val)
	os.Exit(1)
	return ""
}

func readValue(s *scanner.Scanner) string {
	tok := s.Scan()
	c := s.TokenText()
	var v string
	if c != ASSIGN {
		os.Exit(1)
	}
	tok = s.Scan()
	for tok != scanner.EOF {
		c := s.TokenText()
		switch c {
		case COMMENT1, COMMENT2:
			tok = scanner.EOF
		default:
			v = v + c
			tok = s.Scan()
		}
	}
	return v

}

func parse(file *os.File) {
	var s scanner.Scanner
	var c string
	ini := Ini{}
	var section string
	linesScanner := bufio.NewScanner(file)
	for linesScanner.Scan() {
		line := linesScanner.Text()
		s.Init(strings.NewReader(line))
		s.Scan()
		c = s.TokenText()
		switch c {
		case LBRACK:
			section = readSection(&s)
			ini[section] = KeyVals{}
		case COMMENT1, COMMENT2:
		default:
			k := c
			v := readValue(&s)
			ini[section][k] = v
		}
	}
	fmt.Println(ini)
}

func main() {
	file, err := os.Open("lol.ini") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	parse(file)
}
