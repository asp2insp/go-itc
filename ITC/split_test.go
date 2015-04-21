package ITC

import (
	"testing"

	"github.com/ca-geo/go-misc/testutils"
)

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
	id := &Id{
		n: -1,
		i1: &Id{
			n:  -1,
			i1: &Id{n: 1},
			i2: &Id{n: 0},
		},
		i2: &Id{
			n:  -1,
			i1: &Id{n: 0},
			i2: &Id{n: 1},
		},
	}
	i1, i2 := split(id)
	testutils.CheckString("((1, 0), 0)", i1.String(), t)
	testutils.CheckString("(0, (0, 1))", i2.String(), t)
}
