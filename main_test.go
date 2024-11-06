package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4
	res := count(b)

	if res != exp {
		t.Errorf("expected %d, got %d instead.", exp, res)
	}
}