package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Color string

const (
	ColorGreen = "\u001b[32m"
	ColorReset = "\u001b[0m"
)

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func main() {
	var textBuffer []string
	printVersion := flag.Bool("v", false, "Print version")
	flag.Parse()

	if *printVersion {
		colorize(ColorGreen, "lower - text case converter to lowercase, version 0.0.2 😉")
		return
	}
	if len(os.Args) > 1 {
		textBuffer = append(textBuffer, os.Args[1])
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			textBuffer = append(textBuffer, scan.Text())
		}
		if err := scan.Err(); err != nil {
			log.Println(err)
		}
	}

	for _, sentence := range textBuffer {
		fmt.Println(strings.ToLower(sentence))
	}
}
