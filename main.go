package main

import (
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

const version = "1.1.1"

var isVersion bool
var isSentence bool

func init() {
	flag.BoolVar(&isSentence, "sentence", false, "")
	flag.BoolVar(&isSentence, "s", false, "")
	flag.BoolVar(&isVersion, "version", false, "")
	flag.BoolVar(&isVersion, "v", false, "")

	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), strings.TrimSpace(getHelp()))
	}
}

func main() {
	flag.Parse()

	// Print version if requested
	if isVersion {
		fmt.Println(version)
		return
	}

	count := getCount()

	// Print the lorem text
	if isSentence {
		fmt.Println(getSentences(count))
	} else {
		fmt.Println(getParagraphs(count))
	}
}

func getCount() int {
	arg := flag.Arg(0)
	count, err := strconv.Atoi(arg)
	if err != nil {
		return 1
	}

	return count
}

func getParagraphs(count int) string {
	elems := strings.Split(getText(), "\n")
	start := rand.Intn(len(elems) - count + 1)

	return strings.Join(elems[start:start+count], "\n\n")
}

func getSentences(count int) string {
	elems := regexp.MustCompile(`\.\s`).Split(getText(), -1)
	start := rand.Intn(len(elems) - count + 1)

	return strings.Join(elems[start:start+count], ". ") + "."
}

func getText() string {
	return strings.TrimSpace(LoremText)
}

func getHelp() string {
	return `
Generate lorem text

Usage:
  lorem [count]

Flags:
  -s, --sentence   generate a sentence of text
  -h, --help       help for lorem
  -v, --version    version for lorem
`
}
