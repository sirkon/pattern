package pattern

import (
	"testing"
	"unsafe"
)

func Test_isPrefixOf(t *testing.T) {
	const pattern = `a..b..c..d..e..f`
	p, err := NewPattern(pattern)
	if err != nil {
		t.Fatal(err)
	}

	patternBytes := []byte(pattern)
	res := isPrefixOf(
		uintptr(unsafe.Pointer(&p.pattern[0])),
		uintptr(unsafe.Pointer(&p.mask[0])),
		uintptr(unsafe.Pointer(&patternBytes[0])),
		p.length,
	)
	if !res {
		t.Errorf("pattern `%s` must match against itself", pattern)
	}

	const failmatch = `a..b..c..d..e..g`
	failBytes := []byte(failmatch)
	res = isPrefixOf(
		uintptr(unsafe.Pointer(&p.pattern[0])),
		uintptr(unsafe.Pointer(&p.mask[0])),
		uintptr(unsafe.Pointer(&failBytes[0])),
		p.length,
	)
	if res {
		t.Errorf("pattern `%s` must not match '%s'", pattern, failmatch)
	}
}
