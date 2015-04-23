package ITC

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// An Event is a recursively defined binary interval tree
// with
type Event struct {
	n      int    // Value of this node
	el, er *Event // left child, right child
}

// An Id is a recursively defined binary interval tree
// which can be divded or merged with other Ids. It is defined
// by recursively dividing an interval into sub intervals
// where the id is defined (n = 1) or not defined (n = 0)
// Ids can be represented by a nested tree of intervals:
// ((1, 0), (0, 1)) represents a subdivision into 4 parts,
// the outer two parts have value 1 and the middle two have
// value 0. We would represent this graphically as follows:
// #__# though an equally valid representation is ##____##
type Id struct {
	n      int // Value of this node: 1 or 0 for a leaf, -1 for internal node
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

// Convert the event into a compact string for printing
func (e *Event) String() string {
	switch {
	case e.el == nil && e.er == nil: // Atom
		return fmt.Sprintf("%d", e.n)
	default: // standard case, both subtrees should be defined
		l, r := "nil", "nil"
		if e.el != nil {
			l = e.el.String()
		}
		if e.er != nil {
			r = e.er.String()
		}
		return fmt.Sprintf("(%d, %s, %s)", e.n, l, r)
	}
}

// Reverse the String conversion
func stringToEvent(s string) *Event {
	scan := new(scanner.Scanner)
	scan.Init(strings.NewReader(s))
	scan.Mode = scanner.ScanInts
	return sReaderToEvent(scan)
}

// Recursive helper function that consumes a reader
func sReaderToEvent(s *scanner.Scanner) (e *Event) {
	e = &Event{n: 0}
	// An event node is either a leaf: "123" or a 3-tuple: "(root, left, right)"
	for s.Scan() != scanner.EOF {
		tok := s.TokenText()
		if n, err := strconv.Atoi(tok); err == nil {
			// If the next token is a number, we're done
			e.n = n
			break
		} else if tok == "(" {
			// If this is a start of a new tree, we'll recurse 3 times
			e.n = sReaderToEvent(s).n
			e.el = sReaderToEvent(s)
			e.er = sReaderToEvent(s)
			break
		}
		// Ignore all other tokens
	}
	return
}
