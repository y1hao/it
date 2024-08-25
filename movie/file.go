package main

import (
	"bufio"
	"iter"
	"os"
)

func ReadLines(filename string) (seq iter.Seq[string], cleanUp func(), err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	cleanUp = func() {
		file.Close()
	}
	scanner := bufio.NewScanner(file)
	seq = func(yield func(string) bool) {
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				break
			}
		}
	}
	return
}
