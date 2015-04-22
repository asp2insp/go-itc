package ITC

import (
	"fmt"
	"strings"
)

// An Event is a recursively defined binary interval tree
// with
type Event struct {
	n      int    // Value of this node
	el, er *Event // left child, right child
}

// An Id is a recursively defined binary interval tree
// which can be divded or merged with other Ids.
type Id struct {
	n      int // Value of this node
	il, ir *Id // left child, right child
}

// A stamp is a causal event tracker. It is composed
// of an Id (the who) and an Event (the what).
// It does not encode any payload data, merely a causal
// history which may be attached to a message.
type Stamp struct {
	event *Event
	id    *Id
}

// ================= FUNCTIONS ===============

// Convert the id into a compact string for printing
func (id *Id) String() string {
	switch {
	case id.n == -1:
		l, r := "nil", "nil"
		if id.il != nil {
			l = id.il.String()
		}
		if id.ir != nil {
			r = id.ir.String()
		}
		return fmt.Sprintf("(%s, %s)", l, r)
	default:
		return fmt.Sprintf("%d", id.n)
	}
}

// Reverse the String conversion
func stringToId(s string) *Id {
	return sReaderToId(strings.NewReader(s))
}

// Recursive helper function that consumes a reader
func sReaderToId(s *strings.Reader) (id *Id) {
	id = &Id{n: -1}
	for { // Iterate
		ch, _, err := s.ReadRune()
		switch {
		// right paren or end of string ends the id
		case err != nil, ch == ')':
			return
		case ch == '(' || ch == ',':
			// left paren begins a new left tree,
			// comma begins a new right tree
			nId := sReaderToId(s)
			if ch == '(' {
				id.il = nId
			} else {
				id.ir = nId
			}
		// Atoms are always nodes of their own
		case ch == '1' || ch == '0':
			return &Id{n: int(ch) - '0'}
		default:
		}
	}
	return
}
