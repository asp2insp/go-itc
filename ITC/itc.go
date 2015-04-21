package ITC

func NewStamp() *Stamp {
	return &Stamp{
		id:    &Id{n: 1},
		event: &Event{n: -1},
	}
}

// Fork preserves the event component and must
// split the id component into two parts which
// do not overlap and which give the original
// id when added.
// This version modifies the target struct
// and returns it as s1, and creates a new stamp
// and returns it as s2.
func (s *Stamp) Fork() (s1, s2 *Stamp) {
	i1, i2 := split(s.id)
	s.id = i1
	return s, &Stamp{id: i2, event: s.event}
}

// =============== INTERNAL ==============

func split(id *Id) (i1, i2 *Id) {
	// First, take care of the cases where id is a leaf
	switch {
	// split(0) => 0, 0
	case id.n == 0:
		return &Id{n: 0},
			&Id{n: 0}
	// split(1) => (1,0), (0,1)
	case id.n == 1:
		return &Id{
				n:  -1,
				i1: &Id{n: 1},
				i2: &Id{n: 0},
			},
			&Id{
				n:  -1,
				i1: &Id{n: 0},
				i2: &Id{n: 1},
			}
	}
	// If we get to here, id is the root of a subtree
	switch {
	// split((0, i)) => (0, a1), (0, a2)
	// where split(i) = a1, a2
	case id.i1.n == 0 && id.i2.n != 0:
		a1, a2 := split(id.i2)
		return &Id{
				n:  -1,
				i1: &Id{n: 0},
				i2: a1,
			},
			&Id{
				n:  -1,
				i1: &Id{n: 0},
				i2: a2,
			}
	// split((i, 0)) => (a1, 0), (a2, 0)
	// where split(i) = a1, a2
	case id.i1.n != 0 && id.i2.n == 0:
		a1, a2 := split(id.i1)
		return &Id{
				n:  -1,
				i1: a1,
				i2: &Id{n: 0},
			},
			&Id{
				n:  -1,
				i1: a2,
				i2: &Id{n: 0},
			}
	}
	// Catch all case: split((i1, i2)) => (i1, 0), (0, i2)
	return &Id{
			n:  -1,
			i1: id.i1,
			i2: &Id{n: 0},
		},
		&Id{
			n:  -1,
			i1: &Id{n: 0},
			i2: id.i2,
		}
}
