package main

import (
	"testing"
)

func TestReversePrintByStack(t *testing.T) {
	reverseByStack(makeList(10))
}

func TestReverseByRecursion(t *testing.T) {
	reverseByRecursion(makeList(10))
}

func TestReverseLinkedList(t *testing.T) {
	list := reverseLinkedList(makeList(10))
	PrintList(list)
}