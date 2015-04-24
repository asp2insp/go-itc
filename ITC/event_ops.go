package ITC

func lift(e *Event, m int) *Event {
	e.n += m
	return e
}

func sink(e *Event, m int) *Event {
	e.n -= m
	if isLeaf(e) && e.n == 0 {
		return nil
	} else {
		return e
	}
}

func normEvent(e *Event) *Event {
	switch {
	case isLeaf(e): // norm(n) => n
		return e
	case isLeaf(e.el) && isLeaf(e.er) && e.el.n == e.er.n:
		e = lift(e, e.el.n)
		e.el = nil
		e.er = nil
		return e
	default:
		m := min(e.el, e.er)
		lift(e, m)
		sink(e.el, m)
		sink(e.er, m)
		return e
	}
}

func min(e1, e2 *Event) int {
	if e1.n < e2.n {
		return e1.n
	} else {
		return e2.n
	}
}

func isLeaf(e *Event) bool {
	return e.el == nil && e.er == nil
}
