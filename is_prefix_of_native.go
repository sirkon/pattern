// +build !amd64

package pattern

import (
	"unsafe"
)

func isPrefixOf(pattern, mask, source uintptr, length int) bool {
	u64len := length >> 3
	for k := 0; k <= u64len; k++ {
		j := k << 3
		mask := *(*uint64)(unsafe.Pointer(mask + uintptr(j)))
		pattern := *(*uint64)(unsafe.Pointer(pattern + uintptr(j)))
		value := *(*uint64)(unsafe.Pointer(source + uintptr(j)))

		if mask&value != pattern {
			return false
		}
	}

	return true
}
