package ITC

import (
	"testing"

	"github.com/asp2insp/go-misc/testutils"
)

func TestStringToTreeAtom(t *testing.T) {
	var id *Id = stringToId("0")
	testutils.CheckString("0", id.String(), t)
	id = stringToId("1")
	testutils.CheckString("1", id.String(), t)
}

func TestStringToTreeOneLevel(t *testing.T) {
	var id *Id = stringToId("(0, 1)")
	testutils.CheckString("(0, 1)", id.String(), t)
}

func TestStringToTreeHalves(t *testing.T) {
	var id *Id = stringToId("((1, 0), 0)")
	testutils.CheckString("((1, 0), 0)", id.String(), t)
	id = stringToId("(0, (1, 0))")
	testutils.CheckString("(0, (1, 0))", id.String(), t)
}

func TestStringToTreeTwoLevels(t *testing.T) {
	var id *Id = stringToId("((1, 0), (0, 1))")
	testutils.CheckString("((1, 0), (0, 1))", id.String(), t)
}
