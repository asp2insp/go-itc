package ITC

import (
	"testing"

	"github.com/asp2insp/go-misc/testutils"
)

// ================ FORK ==================

func TestFork(t *testing.T) {
	base := NewStamp()
	s1, s2 := base.Fork()
	testutils.CheckString("(1, 0)", s1.id.String(), t)
	testutils.CheckString("(0, 1)", s2.id.String(), t)
}

// ================ JOIN ==================

func TestJoin(t *testing.T) {
	s1 := &Stamp{
		id:    stringToId("(1, 0)"),
		event: stringToEvent("4"),
	}
	s2 := &Stamp{
		id:    stringToId("(0, 1)"),
		event: stringToEvent("(2, 0, 3)"),
	}
	testutils.CheckString("[1, (4, 0, 1)]", Join(s1, s2).String(), t)
}
