package ITC

// Split the given id into two parts which, when
// added together, form the given id, and do not
// overlap
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
				il: &Id{n: 1},
				ir: &Id{n: 0},
			},
			&Id{
				n:  -1,
				il: &Id{n: 0},
				ir: &Id{n: 1},
			}
	}
	// If we get to here, id is the root of a subtree
	// We first check to see if it's a stick
	switch {
	// split((0, i)) => (0, a1), (0, a2)
	// where split(i) = a1, a2
	case id.il.n == 0 && id.ir.n != 0:
		a1, a2 := split(id.ir)
		return &Id{
				n:  -1,
				il: &Id{n: 0},
				ir: a1,
			},
			&Id{
				n:  -1,
				il: &Id{n: 0},
				ir: a2,
			}
	// split((i, 0)) => (a1, 0), (a2, 0)
	// where split(i) = a1, a2
	case id.il.n != 0 && id.ir.n == 0:
		a1, a2 := split(id.il)
		return &Id{
				n:  -1,
				il: a1,
				ir: &Id{n: 0},
			},
			&Id{
				n:  -1,
				il: a2,
				ir: &Id{n: 0},
			}
	}
	// Catch all case: split((i1, ir)) => (i1, 0), (0, ir)
	return &Id{
			n:  -1,
			il: id.il,
			ir: &Id{n: 0},
		},
		&Id{
			n:  -1,
			il: &Id{n: 0},
			ir: id.ir,
		}
}

// Normalize a single Id. Used to clean up
// nodes in the id tree that can be joined
// to reduce the size of the id tree. Usually
// used as part of a bigger operation, like sum.
func norm(id *Id) *Id {
	switch {
	case id.n != -1: // norm(i) = i
		return id
	case id.il.n == 0 && id.ir.n == 0: // norm((0,0)) => 0
		return &Id{n: 0}
	case id.il.n == 1 && id.ir.n == 1: // norm((1,1)) => 0
		return &Id{n: 1}
	default: // norm(i) = i
		return id
	}
}

// Sum recursively combines two Id trees and produces
// a normalized result
func sum(i1, i2 *Id) *Id {
	switch {
	case i1.n == 0:
		return i2
	case i2.n == 0:
		return i1
	default:
		return norm(&Id{n: -1, il: sum(i1.il, i2.il), ir: sum(i1.ir, i2.ir)})
	}
}
