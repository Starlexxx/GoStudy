// Doubly circular linked list implementation with basic operations.
package list

import (
	"fmt"
)

// List - double linked list.
type List struct {
	head *Elem
	tail *Elem
}

// Elem - list element.
type Elem struct {
	Val        interface{}
	next, prev *Elem
}

// New creates a list and returns a pointer to it.
func New() *List {
	var l List
	l.head = nil
	l.tail = nil

	return &l
}

// Push inserts an element at the beginning of the list.
func (l *List) Push(e Elem) *Elem {
	el := &Elem{Val: e.Val}
	if l.head == nil {
		l.head = el
		l.head.next = l.head
		l.head.prev = l.head
		l.tail = l.head
	} else {
		el.next = l.head
		el.prev = l.head.prev
		l.head.prev.next = el
		l.head.prev = el
		l.head = el
	}

	return el
}

// String implements fmt.Stringer interface representing the list as a string.
func (l *List) String() string {
	if l.head == nil {
		return ""
	}

	el := l.head
	s := fmt.Sprintf("%v", el.Val)
	el = el.next
	for el != l.head {
		s += fmt.Sprintf(" %v", el.Val)
		el = el.next
	}

	return s
}

// Pop removes the first element of the list.
func (l *List) Pop() *List {
	if l.head == nil {
		return l
	}

	if l.head.next == l.head {
		l.head = nil
		l.tail = nil
		return l
	}

	l.head.prev.next = l.head.next
	l.head.next.prev = l.head.prev
	l.head = l.head.next

	return l
}

// Reverse reverses the list.
func (l *List) Reverse() *List {
	if l.head == nil || l.head.next == l.head {
		return l
	}

	el := l.head
	for el != l.tail {
		el.next, el.prev = el.prev, el.next
		el = el.prev
	}
	el.next, el.prev = el.prev, el.next
	l.head, l.tail = l.tail, l.head

	return l
}
