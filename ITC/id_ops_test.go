package ITC

import (
	"testing"

	"github.com/asp2insp/go-misc/testutils"
)

// ================= SPLIT =======================

// ____ => __, __
func TestSplitZero(t *testing.T) {
	id := &Id{n: 0}
	i1, i2 := split(id)
	testutils.CheckString("0", i1.String(), t)
	testutils.CheckString("0", i2.String(), t)
}

// #### => ##, ##
func TestSplitOne(t *testing.T) {
	id := &Id{n: 1}
	i1, i2 := split(id)
	testutils.CheckString("(1, 0)", i1.String(), t)
	testutils.CheckString("(0, 1)", i2.String(), t)
}

// ##__ => #___, _#__
func TestSplitLeftHalf(t *testing.T) {
	id := &Id{n: -1, il: &Id{n: 1}, ir: &Id{n: 0}}
	i1, i2 := split(id)
	testutils.CheckString("((1, 0), 0)", i1.String(), t)
	testutils.CheckString("((0, 1), 0)", i2.String(), t)
}

// __## => __#_, ___#
func TestSplitRightHalf(t *testing.T) {
	id := &Id{n: -1, il: &Id{n: 0}, ir: &Id{n: 1}}
	i1, i2 := split(id)
	testutils.CheckString("(0, (1, 0))", i1.String(), t)
	testutils.CheckString("(0, (0, 1))", i2.String(), t)
}

// #### => ##__, __##
func TestSplitOneEquiv(t *testing.T) {
	id := &Id{n: -1, il: &Id{n: 1}, ir: &Id{n: 1}}
	i1, i2 := split(id)
	testutils.CheckString("(1, 0)", i1.String(), t)
	testutils.CheckString("(0, 1)", i2.String(), t)
}

// #__# => #___, ___#
func TestSplitGeneral(t *testing.T) {
	id := stringToId("((1,0), (0,1))")
	i1, i2 := split(id)
	testutils.CheckString("((1, 0), 0)", i1.String(), t)
	testutils.CheckString("(0, (0, 1))", i2.String(), t)
}

// ================= NORM =======================

// normId(i) => i
func TestNormAtoms(t *testing.T) {
	id := stringToId("0")
	testutils.CheckString("0", normId(id).String(), t)
	id = stringToId("1")
	testutils.CheckString("1", normId(id).String(), t)
}

// normId(0, 0) => 0
func TestNormZero(t *testing.T) {
	id := stringToId("(0, 0)")
	testutils.CheckString("0", normId(id).String(), t)
}

// normId(1, 1) => 1
func TestNormOne(t *testing.T) {
	id := stringToId("(1, 1)")
	testutils.CheckString("1", normId(id).String(), t)
}

// normId(i) => i
func TestNormDeepIsNoOp(t *testing.T) {
	id := stringToId("((0, 1), 0)")
	testutils.CheckString("((0, 1), 0)", normId(id).String(), t)
	id = stringToId("((0, 1), (1, 1))")
	testutils.CheckString("((0, 1), (1, 1))", normId(id).String(), t)
}

// ================= SUM =======================

// sum(1, 0) => 1
// sum(0, 1) => 1
func TestSumAtoms(t *testing.T) {
	id1 := stringToId("0")
	id2 := stringToId("1")
	testutils.CheckString("1", sum(id1, id2).String(), t)
	testutils.CheckString("1", sum(id2, id1).String(), t)
}

// sum((i, j), 0) => (i, j)
func TestSumHalves(t *testing.T) {
	id1 := stringToId("(1, 0)")
	id2 := stringToId("0")
	testutils.CheckString("(1, 0)", sum(id1, id2).String(), t)
	testutils.CheckString("(1, 0)", sum(id2, id1).String(), t)
}

// sum((l1, r1), (l2, r2)) => normId(sum(l1, l2), sum(r1, r2))
// ##__ + __## => ####
func TestSumRecursiveOneLevel(t *testing.T) {
	id1 := stringToId("(1, 0)")
	id2 := stringToId("(0, 1)")
	testutils.CheckString("1", sum(id1, id2).String(), t)
}

// sum((l1, r1), (l2, r2)) => normId(sum(l1, l2), sum(r1, r2))
// ##______ + ___#____ => ##_#____
func TestSumRecursiveStick(t *testing.T) {
	id1 := stringToId("((1, 0), 0)")
	id2 := stringToId("((0, (0, 1)), 0)")
	testutils.CheckString("((1, (0, 1)), 0)", sum(id1, id2).String(), t)
}
