package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	n := calcTrees(0, 7, 12, 5)
	assert.Equal(t, 25, n)

}

func TestUnaryNonOverlapingRanges(t *testing.T) {
	n := calcTrees(-2, 0, 2, 0)
	assert.Equal(t, 2, n)

}
func TestNotOverlap(t *testing.T) {
	n := calcTrees(-2, 1, 2, 1)
	assert.Equal(t, 6, n)
}

func TestOverlapUnary(t *testing.T) {
	n := calcTrees(-2, 2, 2, 2)
	assert.Equal(t, 9, n)
}

func TestOverlap(t *testing.T) {
	n := calcTrees(-2, 5, 2, 5)
	assert.Equal(t, 15, n)
}

func TestNotOverlapNegativeRange(t *testing.T) {
	n := calcTrees(-2, -1, 2, -1)
	assert.Equal(t, 6, n)
}

func TestOverlapNegativeRange(t *testing.T) {
	n := calcTrees(-2, -5, 2, -5)
	assert.Equal(t, 15, n)
}

func TestOverlapWithUnary(t *testing.T) {
	n := calcTrees(-2, 0, 2, -5)
	assert.Equal(t, 11, n)
}
