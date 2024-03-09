package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	home  = 1
	guest = 2
)

type total struct {
	homeG  int
	guestG int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t1h, t2g := parseTwoArgs(scanner)
	t1g, t2h := parseTwoArgs(scanner)
	where := parseArg(scanner)
	fmt.Print(calc(t1h, t2g, t1g, t2h, where))
}

func calc(t1h, t2g, t1g, t2h, whereFirst int) int {
	t1 := total{homeG: t1h, guestG: t1g}
	t2 := total{homeG: t2h, guestG: t2g}

	if whereFirst == guest {
		t1.homeG, t1.guestG = t1.guestG, t1.homeG
		t2.homeG, t2.guestG = t2.guestG, t2.homeG
	}

	toScore := 0
	switch {
	case t1.homeG+t1.guestG > t2.homeG+t2.guestG:
		return 0

	case t1.homeG+t1.guestG < t2.homeG+t2.guestG:
		toScore = t2.homeG + t2.guestG - t1.homeG - t1.guestG
	}
	if t1.guestG <= t2.guestG {
		toScore++
	}

	return toScore
}

func parseTwoArgs(scanner *bufio.Scanner) (int, int) {
	scanner.Scan()
	row := strings.Split(scanner.Text(), ":")
	first, _ := strconv.Atoi(row[0])
	second, _ := strconv.Atoi(row[1])
	return first, second
}

func parseArg(scanner *bufio.Scanner) int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
