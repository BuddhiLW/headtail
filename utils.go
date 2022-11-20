package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter(s)")
		return
	}

	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Enter a valid interger")
	}

	out, err := Head(os.Args[1], n)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(out)
}

func Head(path string, n int) ([]string, error) {
	lines := make([]string, 0, n)

	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(f)

	for i := 0; s.Scan() && i < n; i++ {
		lines = append(lines, s.Text())
	}

	return lines, nil
}

func Tail(path string, n int) ([]string, error) {
	lines := make([]string, 0, n)

	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(f)

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	if n > len(lines) {
		n = len(lines)
	}

	return lines[len(lines)-n:], nil

}
