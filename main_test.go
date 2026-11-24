package main

import "testing"

func TestCycleStart(t *testing.T) {
	// 1 -> 2 -> 3 -> 4 -> back to 2
	a := &Node{Val: 1}
	b := &Node{Val: 2}
	c := &Node{Val: 3}
	d := &Node{Val: 4}
	a.Next, b.Next, c.Next, d.Next = b, c, d, b
	if got := CycleStart(a); got != b {
		t.Errorf("expected node 2, got %v", got)
	}

	// Cycle at the head.
	x := &Node{Val: 1}
	y := &Node{Val: 2}
	x.Next, y.Next = y, x
	if got := CycleStart(x); got != x {
		t.Errorf("expected head node for full-circle list")
	}

	// Self loop.
	s := &Node{Val: 9}
	s.Next = s
	if got := CycleStart(s); got != s {
		t.Errorf("expected self-loop node")
	}

	// No cycle.
	p := &Node{Val: 1, Next: &Node{Val: 2}}
	if got := CycleStart(p); got != nil {
		t.Errorf("expected nil for straight list, got %v", got)
	}
	if CycleStart(nil) != nil {
		t.Errorf("expected nil for empty list")
	}
}
