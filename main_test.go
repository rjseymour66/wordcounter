package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {

	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4
	res := count(b, false, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {

	b := bytes.NewBufferString("This is a line\nThis is another line\nThis is the third line")

	exp := 3
	res := count(b, true, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead", exp, res)
	}
}

func TestCountBytes(t *testing.T) {

	b := bytes.NewBufferString("0123456789\n")

	exp := 11
	result := count(b, false, true)

	if result != exp {
		t.Errorf("Expected %d, got %d instead", exp, result)
	}
}
