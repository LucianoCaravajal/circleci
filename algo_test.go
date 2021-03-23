package main

import (
"testing"
)

func TestSum(t *testing.T) {
    total := 5 + 5
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}

func TestMul(t *testing.T) {
    total := 5 * 5
    if total != 25 {
       t.Errorf("Mul was incorrect, got: %d, want: %d.", total, 25)
    }
}

func TestDiv(t *testing.T) {
    total := 5 / 5
    if total != 1 {
       t.Errorf("Div was incorrect, got: %d, want: %d.", total, 1)
    }
}

func TestSub(t *testing.T) {
    total := 5 - 5
    if total != 0 {
       t.Errorf("Sub was incorrect, got: %d, want: %d.", total, 0)
    }
}
