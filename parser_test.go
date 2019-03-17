package pattern

import (
	"reflect"
	"testing"
)

func Test_patternParser(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		ptrn    []byte
		mask    []byte
		first   byte
		offset  int
		wantErr bool
	}{
		{
			name:    "all-pattern",
			pattern: "......",
			ptrn:    []byte{0, 0, 0, 0, 0, 0},
			mask:    []byte{0, 0, 0, 0, 0, 0},
			first:   0,
			offset:  0,
			wantErr: false,
		},
		{
			name:    "simple-pattern",
			pattern: "...ab..",
			ptrn:    []byte{0, 0, 0, 'a', 'b', 0, 0},
			mask:    []byte{0, 0, 0, 0xff, 0xff, 0, 0},
			first:   'a',
			offset:  3,
			wantErr: false,
		},
		{
			name:    "all-escapes",
			pattern: `..\n\r\t\.\` + "`...",
			ptrn:    []byte{0, 0, '\n', '\r', '\t', '.', '`', 0, 0, 0},
			mask:    []byte{0, 0, 0xff, 0xff, 0xff, 0xff, 0xff, 0, 0, 0},
			first:   '\n',
			offset:  2,
			wantErr: false,
		},
		{
			name:    "empty-error",
			pattern: "",
			ptrn:    nil,
			first:   0,
			offset:  0,
			wantErr: true,
		},
		{
			name:    "unknown-escape-error",
			pattern: `..\a`,
			ptrn:    nil,
			first:   0,
			offset:  0,
			wantErr: true,
		},
		{
			name:    "unclosed-escape-error",
			pattern: `....\`,
			ptrn:    nil,
			first:   0,
			offset:  0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ptrn, mask, first, offset, err := parsePattern(tt.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePattern() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(ptrn, tt.ptrn) {
				t.Errorf("parsePattern() ptrn = %v, want %v", ptrn, tt.ptrn)
			}
			if !reflect.DeepEqual(mask, tt.mask) {
				t.Errorf("parsePattern() mask = %v, want %v", mask, tt.mask)
			}
			if first != tt.first {
				t.Errorf("parsePattern() first = %v, want %v", first, tt.first)
			}
			if offset != tt.offset {
				t.Errorf("parsePattern() offset = %v, want %v", offset, tt.offset)
			}
			if tt.wantErr == CheckPattern(tt.pattern) {
				t.Errorf("CheckPattern() return true if pattern is correct")
			}
		})
	}
}
