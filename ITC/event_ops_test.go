package ITC

import (
	"testing"

	"github.com/asp2insp/go-misc/testutils"
)

// ================= FILL =======================

// fill(0, e) = e
func TestFillWithOutId(t *testing.T) {
	i := stringToId("0")
	e := stringToEvent("(3, 0, 1)")
	testutils.CheckString("(3, 0, 1)", fill(i, e).String(), t)
}

// fill(1, e) = max(e)
func TestFillWithFullId(t *testing.T) {
	i := stringToId("1")
	e := stringToEvent("(3, 0, 1)")
	testutils.CheckString("(3, 0, 1)", fill(i, e).String(), t)
}

// fill(i, n) = n
func TestFillWithAtomicEvent(t *testing.T) {
	i := stringToId("0")
	i2 := stringToId("1")
	e := stringToEvent("7")
	testutils.CheckString("7", fill(i, e).String(), t)
	testutils.CheckString("7", fill(i2, e).String(), t)
}

// ================= JOIN =======================

// join(n1, n2) === max(n1, n2)
func TestJoinTwoAtoms(t *testing.T) {
	n1 := stringToEvent("1")
	n2 := stringToEvent("3")
	testutils.ExpectTrue(n2 == join(n1, n2), "Join of atoms should take max", t)
	testutils.ExpectTrue(n2 == join(n2, n1), "Join order shouldn't matter", t)
}

// join((n1, l1, r1), (n2, l2, r2)) = norm((n1,
//                                         join(l1, lift(l2, n2- n1)),
//                                         join(r1, lift(r2, n2- n1))))
func TestJoinShallow(t *testing.T) {
	e1 := stringToEvent("(2, 0, 5)")
	e2 := stringToEvent("(3, 7, 0)")
	testutils.CheckString("(7, 3, 0)", join(e1, e2).String(), t)
	testutils.CheckString("(7, 3, 0)", join(e2, e1).String(), t)
}

// join(n1, (n2, l2, r2)) === join((n1, 0, 0), (n2, l2, r2))
func TestJoinHalves(t *testing.T) {
	e1 := stringToEvent("5")
	e2 := stringToEvent("(3, 7, 0)")
	testutils.CheckString("(5, 5, 0)", join(e1, e2).String(), t)
	testutils.CheckString("(5, 5, 0)", join(e2, e1).String(), t)
}

// ================= LEQ =======================

// leq(n1, n2) === n1 <= n2
func TestCompareTwoAtoms(t *testing.T) {
	e1 := stringToEvent("1")
	e2 := stringToEvent("3")
	testutils.ExpectFalse(leq(e2, e1), "3 !<= 1", t)
	testutils.ExpectTrue(leq(e1, e2), "1 <= 3", t)
}

// leq(n1, (n2, l2, r2)) === n1 <= n2
func TestCompareAtomLessThanTree(t *testing.T) {
	base := stringToEvent("4")
	less := stringToEvent("(2, 10, 0)")
	more := stringToEvent("(8, 0, 1)")
	testutils.ExpectFalse(leq(more, base), "8 !<= 4", t)
	testutils.ExpectFalse(leq(less, base), "2 <= 4", t)
}

// leq((n1, l1 r1), n2) === n1 <= n2
//                           and leq(lift(l1, n1), n2)
//                           and leq(lift(r1, n1), n2)
func TestCompareTreeLessThanAtom(t *testing.T) {
	base := stringToEvent("(3, 9, 0)")
	less := stringToEvent("3")
	more := stringToEvent("20")
	testutils.ExpectFalse(leq(more, base), "20 !<= 2+1 or 2+9", t)
	testutils.ExpectTrue(leq(less, base), "3 <= 3+9 and 3+0", t)
}

// leq((n1, l1, r1), (n2, l2, r2)) === n1 <= n2
//                                     and leq(lift(l1, n1), lift(l2, n2))
//                                     and leq(lift(r1, n1), lift(r2, n2))
func TestCompareTwoTreesByRoot(t *testing.T) {
	base := stringToEvent("(2, 9, 8)")
	less := stringToEvent("(1, 0, 5)")
	more := stringToEvent("(10, 0, 9)")
	testutils.ExpectFalse(leq(more, base), "10 !<= 2", t)
	testutils.ExpectTrue(leq(less, base), "1 <= 2", t)
}

// leq((n1, l1, r1), (n2, l2, r2)) === n1 <= n2
//                                     and leq(lift(l1, n1), lift(l2, n2))
//                                     and leq(lift(r1, n1), lift(r2, n2))
func TestCompareTwoTreesByLeaves(t *testing.T) {
	base := stringToEvent("(2, 5, 6)")
	less := stringToEvent("(2, 5, 5)")
	more := stringToEvent("(2, 6, 6)")
	testutils.ExpectFalse(leq(more, base), "6 !<= 5", t)
	testutils.ExpectTrue(leq(less, base), "5 <= 6", t)
}

// leq((n1, l1, r1), (n2, l2, r2)) === n1 <= n2
//                                     and leq(lift(l1, n1), lift(l2, n2))
//                                     and leq(lift(r1, n1), lift(r2, n2))
func TestCompareTwoTreesRecursive(t *testing.T) {
	base := stringToEvent("(2, 0, (1, 2, 0))")
	less := stringToEvent("(2, 0, (1, 1, 0))")
	more := stringToEvent("(2, 0, (1, 5, 0))")
	testutils.ExpectFalse(leq(more, base), "5 !<= 2", t)
	testutils.ExpectTrue(leq(less, base), "1 <= 2", t)
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
	testutils.ExpectTrue(e != eLifted, "lift should return a new event", t)
	testutils.CheckString("4", eLifted.String(), t)
}

// lift((n, el, er), m) => (n+m, el, er)
func TestLiftRoot(t *testing.T) {
	e := stringToEvent("(1, 4, 7)")
	eLifted := lift(e, 3)
	testutils.ExpectTrue(e != eLifted, "lift should return a new event", t)
	testutils.CheckString("(4, 4, 7)", eLifted.String(), t)
}

// ================= SINK =======================

// sink(n, m) => n-m
func TestSinkAtom(t *testing.T) {
	e := stringToEvent("5")
	eSunk := sink(e, 3)
	testutils.ExpectTrue(e != eSunk, "sink should return a new event", t)
	testutils.CheckString("2", eSunk.String(), t)
}

// sink((n, el, er), m) => (n-m, el, er)
func TestSinkRoot(t *testing.T) {
	e := stringToEvent("(8, 4, 7)")
	eSunk := sink(e, 3)
	testutils.ExpectTrue(e != eSunk, "sink should return a new event", t)
	testutils.CheckString("(5, 4, 7)", eSunk.String(), t)
}
