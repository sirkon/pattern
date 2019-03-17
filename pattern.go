/*
package pattern provides match and lookup operations in slices of byte for so called patterns

*/

package pattern

import (
	"bytes"
)

// NewPattern pattern constructor
func NewPattern(pattern string) (*Pattern, error) {
	var res Pattern

	ptrn, first, offset, err := parsePattern(pattern)
	if err != nil {
		return nil, err
	}
	res.text = pattern
	res.pattern = ptrn
	res.first = first
	res.firstOffset = offset
	return &res, nil
}

// Pattern represents a pattern.
type Pattern struct {
	text        string
	pattern     []byte
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
	if p.first == 0 {
		if len(source) < p.length() {
			return nil
		}
		return source
	}

	var realPos int
	origSource := source
	for {
		if len(origSource)-realPos < p.length() {
			return nil
		}

		// Look for first non-wildcard of a pattern
		pos := bytes.IndexByte(source, p.first)
		if pos < 0 {
			return nil
		}
		if pos < p.firstOffset {
			source = source[pos+1:]
			realPos += pos + 1
			continue
		}

		realPos += pos
		if len(origSource)-realPos+p.firstOffset < p.length() {
			return nil
		}
		if p.isPrefixOf(origSource[realPos-p.firstOffset:]) {
			return origSource[realPos-p.firstOffset:]
		}
		source = source[realPos+1:]
		realPos += 1
	}
}

func (p *Pattern) length() int {
	return len(p.pattern)
}

func (p *Pattern) isPrefixOf(source []byte) bool {
	for i, mask := range p.pattern {
		if mask != 0 && source[i] != mask {
			return false
		}
	}
	return true
}
