package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/mskelton/lorem/util"
	"github.com/spf13/cobra"
)

var isSentence bool
var rootCmd = &cobra.Command{
	Use:   "lorem",
	Short: "Generate lorem text",
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())
		count := getCount(args)

		if isSentence {
			fmt.Println(getSentences(count))
		} else {
			fmt.Println(getParagraphs(count))
		}
	},
	Args:    cobra.MaximumNArgs(1),
	Version: "1.0.0",
}

func init() {
	rootCmd.Flags().BoolVarP(&isSentence, "sentence", "s", false, "Generate a sentence of text")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func getCount(args []string) int {
	if len(args) == 1 {
		count, err := strconv.Atoi(args[0])
		cobra.CheckErr(err)

		return count
	}

	return 1
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
	return strings.TrimSpace(util.LoremText)
}
