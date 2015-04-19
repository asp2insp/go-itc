package itc

func NewStamp() *Stamp {
	return &Stamp{
		id:    newId(1),
		event: newEvent(-1),
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
	switch {
	case id.n == 0: // split(0) => (0, 0)
		return &Id{n: 0}, &Id{n: 0}
	case id.n == 1: // split(1) => ((1,0), (0,1))
		return &Id{}
	}
}
