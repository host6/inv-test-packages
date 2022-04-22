package hit

import (
	"os"
	"runtime"
	"sync"
	"testing"
	"unsafe"
)

type hitsType map[string]*int

var m = sync.Mutex{}
var hits hitsType = map[string]*int{}

func BumpHits(t *testing.T, str string) int {
	m.Lock()
	defer m.Unlock()
	t.Log("before", hits)
	val := new(int)
	*val = 1
	runtime.SetFinalizer(val, func(v *int) {
		t.Log("finit")
	})
	hits[str] = val
	t.Log(unsafe.Pointer(&hits), len(hits))
	t.Log("after", hits)
	t.Log("pid", os.Getpid())
	return len(hits)
}
