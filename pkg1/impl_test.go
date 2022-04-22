package pkg1

import (
	"inv-test/hit"
	"testing"
)

func TestPkg1(t *testing.T) {
	t.Log(hit.BumpHits(t, "dfdfdfdfdf"))
}
