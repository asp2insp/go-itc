package ITC

type Event struct {
	n      int
	e1, e2 *Event
}

type Id struct {
	n      int
	i1, i2 *Id
}

type Stamp struct {
	e      *Event
	i      *Id
	serial int
}
