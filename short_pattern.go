package pattern

import (
	"fmt"
	"unsafe"
)

// ShortPattern efficient structure for short patterns (1…8 bytes long)
type ShortPattern struct {
	mask    uint64
	pattern uint64
	length  int
	first   byte
	offset  int
}

// NewShortPattern constructor of Short Patterns
func NewShortPattern(pattern string) (*ShortPattern, error) {
	length, ptrn, mask, first, offset, err := parsePattern(pattern)
	if err != nil {
		return nil, err
	}
	if len(ptrn) > 8 {
		return nil, fmt.Errorf("short pattern cannot be longer than 8 characters, got %d", len(ptrn))
	}
	return &ShortPattern{
		mask:    *(*uint64)(unsafe.Pointer(&mask[0])),
		pattern: *(*uint64)(unsafe.Pointer(&ptrn[0])),
		length:  length,
		first:   first,
		offset:  offset,
	}, nil
}
