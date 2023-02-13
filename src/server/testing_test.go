package main

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	got := -1
	if got != 1 {
		t.Errorf("-1 = %d; want 1", got)
		fmt.Println("Error")
	}
}
