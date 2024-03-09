package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask1(t *testing.T) {
	n := calc(0, 0, 0, 0, 1)
	assert.Equal(t, 1, n)
}

func TestTask2(t *testing.T) {
	n := calc(0, 2, 0, 3, 1)
	assert.Equal(t, 5, n)
}

func TestTask3(t *testing.T) {
	n := calc(0, 2, 0, 3, 2)
	assert.Equal(t, 6, n)
}

func TestHomeDraw(t *testing.T) {
	n := calc(3, 0, 0, 3, 1)
	assert.Equal(t, 1, n)
}

func TestHomeWin(t *testing.T) {
	n := calc(5, 3, 2, 1, 2)
	assert.Equal(t, 0, n)
}

func TestGuestDraw(t *testing.T) {
	n := calc(3, 0, 0, 3, 1)
	assert.Equal(t, 1, n)
}

func TestGuestLoose(t *testing.T) {
	n := calc(0, 0, 0, 4, 1)
	assert.Equal(t, 4, n)
}

func TestHomeLoose(t *testing.T) {
	n := calc(0, 4, 0, 0, 1)
	assert.Equal(t, 5, n)
}
