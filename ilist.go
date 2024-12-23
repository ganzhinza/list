package list

import (
	"fmt"
)

type List struct {
	root *Elem
}

type Elem struct {
	Val        interface{}
	next, prev *Elem
}

func New() *List {
	var l List
	l.root = &Elem{}
	l.root.next = l.root
	l.root.prev = l.root
	return &l
}

func (l *List) Push(e Elem) *Elem {
	e.prev = l.root
	e.next = l.root.next
	l.root.next = &e
	if e.next != l.root {
		e.next.prev = &e
	}
	return &e
}

func (l *List) String() string {
	el := l.root.next
	var s string
	for el != l.root {
		s += fmt.Sprintf("%v ", el.Val)
		el = el.next
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

func (l *List) Pop() *List {
	if l.root.next == l.root {
		return nil
	}
	l.root.next = l.root.next.next
	l.root.next.prev = l.root
	return l
}

func (l *List) Reverse() *List {
	p := l.root.next
	for p.next != l.root {
		p.next, p.prev = p.prev, p.next
		p = p.prev
	}
	l.root.next = p
	p.next, p.prev = p.prev, p.next
	return l
}
