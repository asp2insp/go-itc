package ITC

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
	e *Event
	i *Id
}

// ================= FUNCTIONS ===============
