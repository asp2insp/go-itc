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
	id := &Id{n: -1, i1: &Id{n: 1}, i2: &Id{n: 0}}
	i1, i2 := split(id)
	testutils.CheckString("((1, 0), 0)", i1.String(), t)
	testutils.CheckString("((0, 1), 0)", i2.String(), t)
}

// __## => __#_, ___#
func TestSplitRightHalf(t *testing.T) {
	id := &Id{n: -1, i1: &Id{n: 0}, i2: &Id{n: 1}}
	i1, i2 := split(id)
	testutils.CheckString("(0, (1, 0))", i1.String(), t)
	testutils.CheckString("(0, (0, 1))", i2.String(), t)
}

// #### => ##__, __##
func TestSplitOneEquiv(t *testing.T) {
	id := &Id{n: -1, i1: &Id{n: 1}, i2: &Id{n: 1}}
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

// norm(i) => i
func TestNormAtoms(t *testing.T) {
	id := stringToId("0")
	testutils.CheckString("0", norm(id).String(), t)
	id = stringToId("1")
	testutils.CheckString("1", norm(id).String(), t)
}

// norm(0, 0) => 0
func TestNormZero(t *testing.T) {
	id := stringToId("(0, 0)")
	testutils.CheckString("0", norm(id).String(), t)
}

// norm(1, 1) => 1
func TestNormOne(t *testing.T) {
	id := stringToId("(1, 1)")
	testutils.CheckString("1", norm(id).String(), t)
}

// norm(i) => i
func TestNormDeepIsNoOp(t *testing.T) {
	id := stringToId("((0, 1), 0)")
	testutils.CheckString("((0, 1), 0)", norm(id).String(), t)
	id = stringToId("((0, 1), (1, 1))")
	testutils.CheckString("((0, 1), (1, 1))", norm(id).String(), t)
}

// ================= SUM =======================

// _# => #
func TestSumRightHalf(t *testing.T) {

}
