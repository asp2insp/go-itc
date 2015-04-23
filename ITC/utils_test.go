package ITC

import (
	"testing"

	"github.com/asp2insp/go-misc/testutils"
)

// ========== STRING <=> ID ============

func TestStringToIdAtom(t *testing.T) {
	var id *Id = stringToId("0")
	testutils.CheckString("0", id.String(), t)
	id = stringToId("1")
	testutils.CheckString("1", id.String(), t)
}

func TestStringToIdOneLevel(t *testing.T) {
	var id *Id = stringToId("(0, 1)")
	testutils.CheckString("(0, 1)", id.String(), t)
}

func TestStringToIdHalves(t *testing.T) {
	var id *Id = stringToId("((1, 0), 0)")
	testutils.CheckString("((1, 0), 0)", id.String(), t)
	id = stringToId("(0, (1, 0))")
	testutils.CheckString("(0, (1, 0))", id.String(), t)
}

func TestStringToIdTwoLevels(t *testing.T) {
	var id *Id = stringToId("((1, 0), (0, 1))")
	testutils.CheckString("((1, 0), (0, 1))", id.String(), t)
}

// ========== STRING <=> EVENT ============

func TestStringToEventAtom(t *testing.T) {
	var e *Event = stringToEvent("0")
	testutils.CheckString("0", e.String(), t)
	e = stringToEvent("1")
	testutils.CheckString("1", e.String(), t)
}

func TestStringToEventOneLevel(t *testing.T) {
	var e *Event = stringToEvent("(1, 2, 3)")
	testutils.CheckString("(1, 2, 3)", e.String(), t)
}

func TestStringToEventHalves(t *testing.T) {
	var e *Event = stringToEvent("(2,(2,1,0),0)")
	testutils.CheckString("(2, (2, 1, 0), 0)", e.String(), t)
	e = stringToEvent("(2, 0, (2, 1, 0))")
	testutils.CheckString("(2, 0, (2, 1, 0))", e.String(), t)
}

func TestStringToEventTwoLevels(t *testing.T) {
	var e *Event = stringToEvent("(2,(3,2,0),(1, 2, (0, (1, 0, 2), 0)))")
	testutils.CheckString("(2, (3, 2, 0), (1, 2, (0, (1, 0, 2), 0)))", e.String(), t)
}
