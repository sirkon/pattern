package pattern

import (
	"bytes"
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
	text    string
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
		text:    pattern,
	}, nil
}

func (p *ShortPattern) String() string {
	return p.text
}

// Match checks if text matches the pattern
func (p *ShortPattern) Match(text []byte) bool {
	if p.length != len(text) {
		return false
	}
	return *(*uint64)(unsafe.Pointer(&text[0]))&p.mask == p.pattern
}

// IsPrefix check if text has a head that matches a pattern
func (p *ShortPattern) IsPrefix(text []byte) bool {
	if p.length > len(text) {
		return false
	}
	return *(*uint64)(unsafe.Pointer(&text[0]))&p.mask == p.pattern
}

// Lookup look for pattern in a given text
func (p *ShortPattern) Lookup(text []byte) []byte {
	if len(text) < p.length {
		return nil
	}
	if p.first == 0 {
		return text
	}

	var realPos int
	origSource := text
	text = text[:len(text)-p.length+1+p.offset] // we are not going to search a pattern outside of text
	// ppatern := uintptr(unsafe.Pointer(&p.pattern[0]))
	// pmask := uintptr(unsafe.Pointer(&p.mask[0]))
	for {
		// Look for first non-wildcard of a pattern
		pos := bytes.IndexByte(text, p.first)
		if pos < 0 {
			return nil
		}

		// Passing it as this symbol is before there could be a match
		if pos < p.offset {
			text = text[pos+1:]
			realPos += pos + 1
			continue
		}

		realPos += pos
		// if isPrefixOf(ppatern, pmask, uintptr(unsafe.Pointer(&origSource[realPos-p.firstOffset])), len(p.pattern)) {
		// 	return origSource[realPos-p.firstOffset:]
		// }
		if *(*uint64)(unsafe.Pointer(&origSource[realPos-p.offset]))&p.mask == p.pattern {
			return origSource[realPos-p.offset:]
		}
		text = text[pos+1:]
		realPos += 1
	}
}
