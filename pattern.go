/*
package pattern provides match and lookup operations in slices of byte for so called patterns

*/

package pattern

import (
	"bytes"
	"unsafe"
)

// NewPattern pattern constructor
func NewPattern(pattern string) (*Pattern, error) {
	var res Pattern

	ptrn, mask, first, offset, err := parsePattern(pattern)
	if err != nil {
		return nil, err
	}
	res.text = pattern
	res.pattern = ptrn
	res.mask = mask
	res.first = first
	res.firstOffset = offset
	return &res, nil
}

// Pattern represents a pattern.
type Pattern struct {
	text        string
	pattern     []byte
	mask        []byte
	first       byte
	firstOffset int
}

func (p *Pattern) String() string {
	return p.text
}

// IsPrefixOf checks if a pattern matches first symbols of source
func (p *Pattern) IsPrefixOf(source []byte) bool {
	return p.length() <= len(source) && p.isPrefixOf(source)
}

// Match matches against a source
func (p *Pattern) Match(source []byte) bool {
	return len(source) == p.length() && p.isPrefixOf(source)
}

// Lookup returns a rest of source with a head matching against current pattern. Returns nil if no matches weren't found
func (p *Pattern) Lookup(source []byte) []byte {
	if len(source) < p.length() {
		return nil
	}
	if p.first == 0 {
		return source
	}

	var realPos int
	origSource := source
	source = source[:len(source)-p.length()+1+p.firstOffset] // we are not going to search a pattern outside of source
	for {
		// Look for first non-wildcard of a pattern
		pos := bytes.IndexByte(source, p.first)
		if pos < 0 {
			return nil
		}

		// Passing it as this symbol is before there could be a match
		if pos < p.firstOffset {
			source = source[pos+1:]
			realPos += pos + 1
			continue
		}

		realPos += pos
		if p.isPrefixOf(origSource[realPos-p.firstOffset:]) {
			return origSource[realPos-p.firstOffset:]
		}
		source = source[pos+1:]
		realPos += 1
	}
}

// ShortLookup the result is the same Lookup what do. But this method is optimized for shorter distances.
func (p *Pattern) ShortLookup(source []byte) []byte {
	if len(source) < p.length() {
		return nil
	}
	if p.first == 0 {
		return source
	}

	var realPos int
	origSource := source
	source = source[:len(source)-p.length()+1+p.firstOffset] // we are not going to search a pattern outside of source
	for {
		// Look for first non-wildcard of a pattern
		pos := -1
		for i, c := range source {
			if c == p.first {
				pos = i
				break
			}
		}
		if pos < 0 {
			return nil
		}

		// Passing it as this symbol is before there could be a match
		if pos < p.firstOffset {
			source = source[pos+1:]
			realPos += pos + 1
			continue
		}

		realPos += pos
		if p.isPrefixOf(origSource[realPos-p.firstOffset:]) {
			return origSource[realPos-p.firstOffset:]
		}
		source = source[pos+1:]
		realPos += 1
	}
}

func (p *Pattern) length() int {
	return len(p.pattern)
}

func (p *Pattern) isPrefixOf(source []byte) bool {
	u64len := len(p.pattern) >> 3
	pmask := unsafe.Pointer(&p.mask[0])
	ppattern := unsafe.Pointer(&p.pattern[0])
	psource := unsafe.Pointer(&source[0])
	for k := 0; k <= u64len; k++ {
		j := k << 3
		mask := *(*uint64)(unsafe.Pointer(uintptr(pmask) + uintptr(j)))
		pattern := *(*uint64)(unsafe.Pointer(uintptr(ppattern) + uintptr(j)))
		value := *(*uint64)(unsafe.Pointer(uintptr(psource) + uintptr(j)))

		if mask&value != pattern {
			return false
		}
	}
	for i := u64len * 8; i < len(p.pattern); i++ {
		v := *(*byte)(unsafe.Pointer(uintptr(pmask) + uintptr(i)))
		if v != 0 && source[i] != v {
			return false
		}
	}

	return true
}
