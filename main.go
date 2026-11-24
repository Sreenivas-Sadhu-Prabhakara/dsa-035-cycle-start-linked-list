package main

import "fmt"

// Node is a singly linked list node.
type Node struct {
	Val  int
	Next *Node
}

// CycleStart returns the node where the list's cycle begins, or nil if the
// list has no cycle. Floyd's meeting point plus the head-restart step.
func CycleStart(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// Meeting point found: a pointer from the head and one from the
			// meeting point, both at speed 1, meet at the cycle entry.
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

func main() {
	a := &Node{Val: 1}
	b := &Node{Val: 2}
	c := &Node{Val: 3}
	d := &Node{Val: 4}
	a.Next, b.Next, c.Next, d.Next = b, c, d, b // cycle starts at node 2
	fmt.Println(CycleStart(a).Val)
}
