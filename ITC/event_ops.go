package ITC

// lift returns a new event with the
// same children as the given event,
// but with the root node value increased
// by the given amount
func lift(e *Event, m int) *Event {
	return &Event{
		n:  e.n + m,
		el: e.el,
		er: e.er,
	}
}

// sink returns a new event with the
// same children as the given event,
// but with the root node value decreased
// by the given amount
func sink(e *Event, m int) *Event {
	return &Event{
		n:  e.n - m,
		el: e.el,
		er: e.er,
	}
}

// normEvent will normalize the given event
// in a non-recursive fashion.
func normEvent(e *Event) *Event {
	switch {
	case isLeaf(e): // norm(n) => n
		return e
	// norm((n, m, m) => n+m
	case isLeaf(e.el) && isLeaf(e.er) && e.el.n == e.er.n:
		e = lift(e, e.el.n)
		e.el = nil
		e.er = nil
		return e
	default:
		m := min(e.el, e.er)
		e = lift(e, m)
		e.el = sink(e.el, m)
		e.er = sink(e.er, m)
		return e
	}
}

// min is a utility function which compares the value of two events directly.
// Returns the node value of the smaller event
func min(e1, e2 *Event) int {
	if e1.n < e2.n {
		return e1.n
	}
	return e2.n
}

// max is a utility function which compares the value of two events directly.
// Returns the event directly.
func max(e1, e2 *Event) *Event {
	if e1.n > e2.n {
		return e1
	}
	return e2
}

// isLeaf returns true if the given event has no children
func isLeaf(e *Event) bool {
	return e.el == nil && e.er == nil
}

// Compares normalized event trees
// where one subtree is always 0. We exploit this
// in the comparison case where a single node
// is compared to a tree. When comparing two trees,
// we take the deepest one across all branches.
// In other words, an event tree must completely dominate
// another in order to be compared greater
func leq(e1, e2 *Event) bool {
	switch {
	// leq(n1, n2) === n1 <= n2
	// leq(n1, (n2, l2, r2)) === n1 <= n2
	case isLeaf(e1):
		return e1.n <= e2.n
	// leq((n1, l1 r1), n2) === n1 <= n2
	//                           and leq(lift(l1, n1), n2)
	//                           and leq(lift(r1, n1), n2)
	case isLeaf(e2):
		return e1.n <= e2.n &&
			leq(lift(e1.el, e1.n), e2) &&
			leq(lift(e1.er, e1.n), e2)
	// leq((n1, l1, r1), (n2, l2, r2)) === n1 <= n2
	//                                     and leq(lift(l1, n1), lift(l2, n2))
	//                                     and leq(lift(r1, n1), lift(r2, n2))
	default:
		return e1.n <= e2.n &&
			leq(lift(e1.el, e1.n), lift(e2.el, e2.n)) &&
			leq(lift(e1.er, e1.n), lift(e2.er, e2.n))
	}
}

func fillNils(e *Event) *Event {
	return &Event{
		n:  e.n,
		el: new(Event),
		er: new(Event),
	}
}

// join recursively stiches together two event trees,
// creating a new event tree that dominates
// both input trees. Produces a normalized tree.
func join(e1, e2 *Event) *Event {
	switch {
	case isLeaf(e1) && isLeaf(e2): // join(n1, n2) === max(n1, n2)
		return max(e1, e2)
	case isLeaf(e1): // join(n1, (n2, l2, r2)) === join((n1, 0, 0), (n2, l2, r2))
		return join(fillNils(e1), e2)
	case isLeaf(e2): // join((n1, l1, r1)) === join((n1, l1, r1), (n2, 0, 0))
		return join(e1, fillNils(e2))
	// join((n1, l1, r1), (n2, l2, r2)) === join((n2, l2, r2), (n1, l1, r1))
	case e1.n > e2.n:
		return join(e2, e1)
	// join((n1, l1, r1), (n2, l2, r2)) = norm((n1,
	//                                         join(l1, lift(l2, n2- n1)),
	//                                         join(r1, lift(r2, n2- n1))))
	default:
		return normEvent(&Event{
			n:  e1.n,
			el: join(e1.el, lift(e2.el, e2.n-e1.n)),
			er: join(e1.er, lift(e2.er, e2.n-e1.n)),
		})
	}
}
