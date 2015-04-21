package ITC

import (
	"testing"

	"github.com/ca-geo/go-misc/testutils"
)

func _TestStringToTreeAtom(t *testing.T) {
	var id *Id = stringToId("0")
	testutils.CheckString("0", id.String(), t)
	id = stringToId("1")
	testutils.CheckString("1", id.String(), t)
}

func TestStringToTreeOneLevel(t *testing.T) {
	var id *Id = stringToId("(0, 1)")
	testutils.CheckString("(0, 1)", id.String(), t)
}
