package ITC

import (
	"testing"

	"github.com/ca-geo/go-misc/testutils"
)

// ================= LEQ =======================

// leq(n1, n2) === n1 <= n2
func TestCompareTwoAtoms(t *testing.T) {
	e1 := stringToEvent("1")
	e2 := stringToEvent("3")
	testutils.ExpectFalse(leq(e2, e1), "3 !<= 1", t)
	testutils.ExpectTrue(leq(e1, e2), "1 <= 3", t)
}

func TestCompareAtomAndTree(t *testing.T) {
	e1 := stringToEvent("(2, 9, 8)")
	e2 := stringToEvent("3")
	testutils.ExpectFalse(leq(e2, e1), "3 !<= 2", t)
	testutils.ExpectTrue(leq(e1, e2), "2 <= 3", t)
}

// ================= NORM =======================

// normEvent(n) => n
func TestNormEventAtom(t *testing.T) {
	e := stringToEvent("1")
	normed := normEvent(e)
	testutils.CheckString("1", normed.String(), t)
}

// normEvent((3, 1, 1)) => 4
func TestNormEventTreeToAtom(t *testing.T) {
	e := stringToEvent("(3,1,1)")
	normed := normEvent(e)
	testutils.CheckString("4", normed.String(), t)
}

// normEvent((3, 1, 2)) => (4, 0, 1)
func TestNormEventTreeWithMismatchedLeaf(t *testing.T) {
	e := stringToEvent("(3,1,2)")
	normed := normEvent(e)
	testutils.CheckString("(4, 0, 1)", normed.String(), t)
}

// (2,(2,1,0),3) => (4,(0,1,0),1)
func TestNormEventTreePartialSink(t *testing.T) {
	e := stringToEvent("(2,(2,1,0),3)")
	normed := normEvent(e)
	testutils.CheckString("(4, (0, 1, 0), 1)", normed.String(), t)
}

// ================= LIFT =======================

// lift(n, m) => n+m
func TestLiftAtom(t *testing.T) {
	e := stringToEvent("1")
	eLifted := lift(e, 3)
	testutils.CheckString("4", eLifted.String(), t)
}

// lift((n, el, er), m) => (n+m, el, er)
func TestLiftRoot(t *testing.T) {
	e := stringToEvent("(1, 4, 7)")
	eLifted := lift(e, 3)
	testutils.CheckString("(4, 4, 7)", eLifted.String(), t)
}

// ================= SINK =======================

// sink(n, m) => n-m
func TestSinkAtom(t *testing.T) {
	e := stringToEvent("5")
	eSunk := sink(e, 3)
	testutils.CheckString("2", eSunk.String(), t)
}

// sink((n, el, er), m) => (n-m, el, er)
func TestSinkRoot(t *testing.T) {
	e := stringToEvent("(8, 4, 7)")
	eSunk := sink(e, 3)
	testutils.CheckString("(5, 4, 7)", eSunk.String(), t)
}
