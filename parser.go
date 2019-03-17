package pattern

import (
	"fmt"
)

type parserState int

const (
	parserStateRegular parserState = iota
	parserStateExtended
)

var acceptableEscapes = map[byte]byte{
	'n':  '\n',
	'r':  '\r',
	't':  '\t',
	'\\': '\\',
	'.':  '.',
	'`':  '`',
}

// CheckPattern checks if pattern is correct
func CheckPattern(pattern string) bool {
	_, _, _, err := parsePattern(pattern)
	return err == nil
}

func parsePattern(pattern string) ([]byte, byte, int, error) {
	state := parserStateRegular
	var ptrn []byte
	var first byte
	var offset int

	firstStillUnset := true
	var count int

	for _, c := range []byte(pattern) {
		switch state {
		case parserStateRegular:
			if c == '\\' {
				state = parserStateExtended
				continue
			}

			if c == '.' {
				ptrn = append(ptrn, 0)
			} else {
				ptrn = append(ptrn, c)
				if firstStillUnset {
					first = c
					offset = count
					firstStillUnset = false
				}
			}
			count++
		case parserStateExtended:
			char, acceptableEscape := acceptableEscapes[c]
			if !acceptableEscape {
				return nil, 0, 0, fmt.Errorf(`parsing error, only \n, \r, \t, \\, \. and \`+"` are allowed, got \\%s", string(c))
			}
			state = parserStateRegular
			ptrn = append(ptrn, char)
			if firstStillUnset {
				first = char
				offset = count
				firstStillUnset = false
			}
			count++
		}
	}

	if state == parserStateExtended {
		return nil, 0, 0, fmt.Errorf("empty escape at the end of pattern string")
	}

	if len(ptrn) == 0 {
		return nil, 0, 0, fmt.Errorf("empty pattern")
	}

	return ptrn, first, offset, nil
}
