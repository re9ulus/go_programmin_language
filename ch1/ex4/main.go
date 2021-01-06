// Modify dup2 to print files with duplications
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)

			if len(files) > 0 {
				// solution without arrays/func return values
				for _, arg := range files {
					f, err := os.Open(arg)
					if err != nil {
						continue
					}
					input := bufio.NewScanner(f)
					for input.Scan() {
						if input.Text() == line {
							fmt.Println(arg)
							break
						}
					}
					f.Close()
				}
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
