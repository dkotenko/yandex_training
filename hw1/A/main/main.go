package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rng struct {
	left  int
	right int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	p, v := parseTwoArgs(scanner)
	q, m := parseTwoArgs(scanner)
	n := calcTrees(p, v, q, m)

	fmt.Printf("%d", n)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func (r *rng) calc() int {
	return absDiffInt(r.right, r.left) + 1
}

func (r *rng) isOverlap(another *rng) bool {
	return r.left <= another.right && another.left <= r.right
}

func calcTrees(p, v, q, m int) int {
	v = absDiffInt(v, 0)
	m = absDiffInt(m, 0)
	vRange := rng{left: p - v, right: p + v}
	mRange := rng{left: q - m, right: q + m}

	leftRange := &vRange
	rightRange := &mRange
	if mRange.left < vRange.left {
		leftRange = &mRange
		rightRange = &vRange
	}

	overlap := 0
	if leftRange.isOverlap(rightRange) {
		overlapRange := rng{
			left:  max(leftRange.left, rightRange.left),
			right: min(leftRange.right, rightRange.right),
		}
		overlap = overlapRange.calc()
	}
	return leftRange.calc() + rightRange.calc() - overlap
}

func parseTwoArgs(scanner *bufio.Scanner) (int, int) {
	scanner.Scan()
	row := strings.Split(scanner.Text(), " ")
	first, _ := strconv.Atoi(row[0])
	second, _ := strconv.Atoi(row[1])
	return first, second
}
