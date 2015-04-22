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
