package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var (
		input   io.Reader
		keyword string
	)
	if len(os.Args) == 2 { //Pipe mode "AToolLikeGrep <keyword>"
		keyword = os.Args[1]
		input = os.Stdin
	} else if len(os.Args) == 3 { //Standard mode "AToolLikeGrep <path> <keyword>"
		path := os.Args[1]
		file, err := os.Open(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file:", err)
			return
		}
		defer file.Close()
		keyword = os.Args[2]
		input = file
	} else {
		fmt.Println("Usage:")
		fmt.Println("  From file:  AToolLikeGrep <path> <keyword>")
		fmt.Println("  From pipe:  cat <path> | AToolLikeGrep <keyword>")
		return
	}
	err := runSearch(input, keyword)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error during search:", err)
	}

}

func runSearch(input io.Reader, keyword string) error {
	scanner := bufio.NewScanner(input)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			fmt.Printf("\"%s\" in line %d\n", line, lineNum)
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return err
	}
	return nil
}
