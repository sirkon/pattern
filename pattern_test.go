/*
package pattern provides match and lookup operations in slices of byte for so called patterns

*/

package pattern_test

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/sirkon/pattern/internal/ragel"

	"github.com/sirkon/pattern"
)

func TestPattern_isPrefixOf(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		source  string
		want    bool
	}{
		{
			name:    "ok",
			pattern: "..:..:..:..:..:..",
			source:  "00:11:22:33:44:55",
			want:    true,
		},
		{
			name:    "source-not-enough",
			pattern: "..:..:..:..:..:..",
			source:  "00:11:22:33:44:5",
			want:    false,
		},
		{
			name:    "source-mismatch",
			pattern: "..:..:..:..:..:..",
			source:  "00-11-22-33-44-55",
			want:    false,
		},
		{
			name:    "bad-pattern",
			pattern: `\`,
			source:  "",
			want:    false,
		},
		{
			name:    "regression-1",
			pattern: "..:..:..:..:..:..",
			source:  "0:11:22:33:44:55    ",
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := pattern.NewPattern(tt.pattern)
			if err != nil {
				if tt.want {
					t.Error(err)
				}
				return
			}
			if got := p.IsPrefixOf([]byte(tt.source)); got != tt.want {
				t.Errorf("Pattern.isPrefixOf() = %v, want %v", got, tt.want)
			}
			if tt.want {
				if !p.Match([]byte(tt.source)) {
					t.Errorf("Match must give true on %s", tt.source)
				}
			}
		})
	}
}

func TestPattern_Lookup(t *testing.T) {
	p, _ := pattern.NewPattern("..:..:..:..:..:..")

	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "ok-match",
			source: "00:11:22:33:44:55",
			want:   "00:11:22:33:44:55",
		},
		{
			name:   "ok-offset",
			source: "        00:11:22:33:44:55",
			want:   "00:11:22:33:44:55",
		},
		{
			name:   "ok-junk",
			source: "        00:11:22:33:44:55    ",
			want:   "00:11:22:33:44:55    ",
		},
		{
			name:   "fail-match",
			source: "00:11:22:33:44:5",
			want:   "",
		},
		{
			name:   "fail-offset",
			source: "            00:11:22:33:44:5",
			want:   "",
		},
		{
			name:   "fail-no-anchor-short",
			source: "       ",
			want:   "",
		},
		{
			name:   "fail-no-anchor",
			source: "                                        ",
			want:   "",
		},
		{
			name:   "fail-no-match",
			source: "0:11:22:33:44:55    ",
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := p.Lookup([]byte(tt.source))
			var want []byte
			if len(tt.want) > 0 {
				want = []byte(tt.want)
			} else {
				tt.want = "nil"
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Pattern.Lookup() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestPattern_Lookup2(t *testing.T) {
	const patternText = "......."
	p, _ := pattern.NewPattern(patternText)
	if p.Lookup([]byte("  ")) != nil {
		t.Errorf("Must be nil result here")
	}
	if p.Lookup([]byte("                      ")) == nil {
		t.Errorf("Must not be nil here")
	}
	if p.String() != patternText {
		t.Errorf("wrong name")
	}
}

func TestRagelCheck(t *testing.T) {
	var p ragel.Stuff

	failedSources := []string{
		"00:11:22:33:44:5",
		" 00:11:22:33:44:55",
	}
	for _, source := range failedSources {
		if p.Match([]byte(source)) {
			t.Errorf("source %s must not be matched against a pattern", source)
		}
	}

	lookupSources := []string{
		"00:11:22:33:44:55",
		"         00:11:22:33:44:55",
		"         00:11:22:33:44:55    ",
	}
	for _, source := range lookupSources {
		if p.Index([]byte(source)) < 0 {
			t.Errorf("failed to find valid pattern in the source")
		}
	}

	failedLookup := []string{
		"00:11:22:33:44:5",
		"     00:11:22:33:44:5",
	}
	for _, source := range failedLookup {
		if p.Index([]byte(source)) >= 0 {
			t.Errorf("source %s must not be matched against a pattern", source)
		}
	}
}

var macAddrs [][]byte
var aaSampes [][]byte
var offsetAddrs [][]byte

func init() {
	for i := 0; i < 16; i++ {
		opts := make([]interface{}, 12)
		for j := range opts {
			opts[j] = i
		}
		macAddrs = append(macAddrs, []byte(fmt.Sprintf("%x%x:%x%x:%x%x:%x%x:%x%x:%x%x", opts...)))
	}

	for i := 'a'; i <= 'z'; i++ {
		aaSampes = append(aaSampes, []byte(fmt.Sprintf("a%s%sa", string(i), string(i))))
	}

	for i := 0; i < 16; i++ {
		opts := make([]interface{}, 12)
		for j := range opts {
			opts[j] = i
		}
		for k := 0; k < 128; k++ {
			offsetAddrs = append(offsetAddrs, []byte(
				strings.Repeat(" ", k)+fmt.Sprintf("%x%x:%x%x:%x%x:%x%x:%x%x:%x%x", opts...),
			))
		}
	}
}

func BenchmarkPattern_Match(b *testing.B) {
	ptrn, _ := pattern.NewPattern("..:..:..:..:..:..")
	for i := 0; i < b.N; i++ {
		for _, macAddr := range macAddrs {
			if !ptrn.Match(macAddr) {
				b.Fatalf("the pattern %s must gives true when matched against %s", ptrn, string(macAddr))
			}
		}
	}
}

func BenchmarkRegexp_Match(b *testing.B) {
	ptrn := regexp.MustCompile("^..:..:..:..:..:..$")
	for i := 0; i < b.N; i++ {
		for _, macAddr := range macAddrs {
			if !ptrn.Match(macAddr) {
				b.Fatalf("the pattern %s must gives true when matched against %s", ptrn, string(macAddr))
			}
		}
	}
}

func BenchmarkRagel_Match(b *testing.B) {
	var ptrn ragel.Stuff
	for i := 0; i < b.N; i++ {
		for _, macAddr := range macAddrs {
			if !ptrn.Match(macAddr) {
				b.Fatalf("the pattern %s must give true when matched against %s", ptrn, string(macAddr))
			}
		}
	}
}

func BenchmarkPattern_Lookup(b *testing.B) {
	ptrn, _ := pattern.NewPattern("..:..:..:..:..:..")
	for i := 0; i < b.N; i++ {
		for _, off := range offsetAddrs {
			if ptrn.Lookup(off) == nil {
				b.Fatalf("the pattern %s must return non-nil against %s", ptrn, string(off))
			}
		}
	}
}

func BenchmarkRegexp_Lookup(b *testing.B) {
	ptrn := regexp.MustCompile("..:..:..:..:..:..")
	for i := 0; i < b.N; i++ {
		for _, off := range offsetAddrs {
			if ptrn.Find(off) == nil {
				b.Fatalf("the pattern %s must return non-nil against %s", ptrn, string(off))
			}
		}
	}
}

func BenchmarkRagel_Lookup(b *testing.B) {
	var ptrn ragel.Stuff
	for i := 0; i < b.N; i++ {
		for _, off := range offsetAddrs {
			if ptrn.Index(off) < 0 {
				b.Fatalf("the pattern %s must give true when matched against %s", ptrn, string(off))
			}
		}
	}
}
