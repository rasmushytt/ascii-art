package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := ""
	if len(os.Args) == 3 {
		args = string(os.Args[2])
		args = args + ".txt"
	}
	if _, err := os.Stat(args); os.IsNotExist(err) {
		fmt.Println("Usage: go run . [STRING] [OPTION]\nEX: go run . 'something' font\nPlease use only ascii characters.")
		os.Exit(0)
	}

	font, _ := readLines(args)
	text := os.Args[1]

	textLines := strings.Split(text, "\\n")

	fmt.Print(writeText(textLines, font))
}

// This function prints the lines
func writeText(textLines []string, font []string) string {
	s := ""
	for i := 0; i < len(textLines); i++ {
		x := 0
		for j := 0; j < 8; j++ {
			for k := 0; k < len(textLines[i]); k++ {
				// Find the right line for the letter
				pos := 1 + (int(textLines[i][k])-32)*9 + x
				s += font[pos]
			}
			x++
			s += "\n"
		}

	}
	return s
}

// This function reads the font file
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
