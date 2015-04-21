package ITC

import "fmt"

// An Event is a recursively defined binary interval tree
// with
type Event struct {
	n      int    // Value of this node
	e1, e2 *Event // left child, right child
}

// An Id is a recursively defined binary interval tree
// which can be divded or merged with other Ids.
type Id struct {
	n      int // Value of this node
	i1, i2 *Id // left child, right child
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
		return fmt.Sprintf("(%s, %s)", id.i1.String(), id.i2.String())
	default:
		return fmt.Sprintf("%d", id.n)
	}
}

// // Reverse the String conversion
// func stringToId(s string) *Id {
// 	return sReaderToId(s.Reader())
// }

// // Recursive helper function that consumes a reader
// func sReaderToId(s *strings.Reader) (id *Id) {
// 	ch, _, err := s.ReadRune()
// 	for {
// 		switch {
// 		case err != nil:
// 			return nil
// 		case ch == ',':
// 			break
// 		case ch == ')' || ' ':
//       continue
//     case ch == '('
//     default:
// 		}
// 	}
// 	return
// }
