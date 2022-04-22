package pkg2

import (
	"inv-test/hit"
	"testing"
)

func TestPkg2(t *testing.T) {
	t.Log(hit.BumpHits(t, "sdsd"))
}
