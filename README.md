# Starting Point of a Linked List Cycle

> Striver SDE Sheet 035 · Day 6 Linked List — Floyd's meeting point, then two 1x pointers from head and meeting point converge on the entry. Go, stdlib only, tested.

## Problem

A linked list may contain a cycle. Return the exact node where the cycle begins (not just whether one exists), in O(1) space.

## Approach

Run tortoise-and-hare until they meet inside the cycle. Now the arithmetic: with head-to-entry distance a, and the meeting point r steps into a cycle of length c, the 2:1 speed ratio forces a ≡ c − r (mod c). That congruence means a pointer restarted at the head and a pointer left at the meeting point, both now moving one step at a time, walk exactly into each other at the cycle's entry. Two phases, three pointers, no marking of visited nodes.

**Complexity:** O(n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
