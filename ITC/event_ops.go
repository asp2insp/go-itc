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

func min(e1, e2 *Event) int {
	if e1.n < e2.n {
		return e1.n
	}
	return e2.n
}

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
