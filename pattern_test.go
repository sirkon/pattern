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
		{
			name:    "regression-aa",
			pattern: "a..a",
			source:  "aaaa",
			want:    true,
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
			if !reflect.DeepEqual(got, p.ShortLookup([]byte(tt.source))) {
				t.Errorf("Different result caught for Lookup and ShortLookup")
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
	if p.ShortLookup([]byte("                      ")) == nil {
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
var aaSamples []byte
var offsetAddrs [][]byte
var uuids [][]byte
var offsetUUIDs [][]byte
var realWorldMacAddr [][]byte
var falsePositivesMacAddrs [][]byte

func init() {
	for i := 0; i < 16; i++ {
		opts := make([]interface{}, 12)
		for j := range opts {
			opts[j] = i
		}
		macAddrs = append(macAddrs, []byte(fmt.Sprintf("%x%x:%x%x:%x%x:%x%x:%x%x:%x%x", opts...)))
	}

	for i := 'a'; i <= 'z'; i++ {
		aaSamples = append(aaSamples, []byte(fmt.Sprintf("a%s%sa", string(i), string(i)))...)
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
		for k := 0; k < 16; k++ {
			falsePositivesMacAddrs = append(falsePositivesMacAddrs,
				[]byte(strings.Repeat("   :   ", k)+fmt.Sprintf("%x%x:%x%x:%x%x:%x%x:%x%x:%x%x", opts...)),
			)
		}
	}

	uuidSample := "00000000-0000-0000-0000-000000000000"
	for i := 0; i < 16; i++ {
		uuid := strings.ReplaceAll(uuidSample, "0", fmt.Sprintf("%x", i))
		uuids = append(uuids, []byte(uuid))
	}

	for i := 0; i < 16; i++ {
		uuid := strings.ReplaceAll(uuidSample, "0", fmt.Sprintf("%x", i))
		for k := 0; k < 8; k++ {
			offsetUUIDs = append(offsetUUIDs, []byte(strings.Repeat(" ", k)+uuid))
		}
	}

	line1 := `Dec  3 19:37:13 localhost dhcpd: DHCPREQUEST for 192.168.100.10 (192.168.100.1) from 11:22:33:44:55:66 (MYHOST) via eth0`
	line2 := `11:22:33:44:55:66`
	line3 := `2019:15:00:01-00Z localhost myapp: REQUEST FROM MAC: 11:22:33:44:55:66 made`
	for i := 0; i < 100000; i++ {
		realWorldMacAddr = append(realWorldMacAddr, []byte(line1), []byte(line2), []byte(line3))
	}
}

func BenchmarkShortPattern_AAMatch(b *testing.B) {
	ptrn, _ := pattern.NewShortPattern("a..a")
	for i := 0; i < b.N; i++ {
		samples := aaSamples
		for len(samples) > 0 {
			if !ptrn.Match(samples[:4]) {
				b.Fatalf("the pattern %s must give true when matched against %s", ptrn, string(samples[:4]))
			}
			samples = samples[4:]
		}
	}
}

func BenchmarkRegexp_AAMatch(b *testing.B) {
	ptrn := regexp.MustCompile(`^a..a$`)
	for i := 0; i < b.N; i++ {
		samples := aaSamples
		for len(samples) > 0 {
			if !ptrn.Match(samples[:4]) {
				b.Fatalf("the pattern %s must gives true when matched against %s", ptrn, string(samples[:4]))
			}
			samples = samples[4:]
		}
	}
}

func BenchmarkRagel_AAMatch(b *testing.B) {
	var ptrn ragel.Stuff
	for i := 0; i < b.N; i++ {
		samples := aaSamples
		for len(samples) > 0 {
			if !ptrn.AAMatch(samples[:4]) {
				b.Fatalf("the pattern %s must gives true when matched against %s", ptrn, string(samples[:4]))
			}
			samples = samples[4:]
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

func BenchmarkPattern_UUIDMatch(b *testing.B) {
	ptrn, _ := pattern.NewPattern("........-....-....-....-............")
	for i := 0; i < b.N; i++ {
		for _, uuid := range uuids {
			if !ptrn.Match(uuid) {
				b.Fatalf("the pattern %s must gives true when matched against %s", ptrn, string(uuid))
			}
		}
	}
}

func BenchmarkRegexp_UUIDMatch(b *testing.B) {
	ptrn := regexp.MustCompile("^........-....-....-....-............$")
	for i := 0; i < b.N; i++ {
		for _, uuid := range uuids {
			if !ptrn.Match(uuid) {
				b.Fatalf("the pattern %s must gives true when matched against %s", ptrn, string(uuid))
			}
		}
	}
}

func BenchmarkRagel_UUIDMatch(b *testing.B) {
	var ptrn ragel.Stuff
	for i := 0; i < b.N; i++ {
		for _, uuid := range uuids {
			if !ptrn.UUIDMatch(uuid) {
				b.Fatalf("the pattern %s must give true when matched against %s", ptrn, string(uuid))
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

func BenchmarkPattern_UUIDLookup(b *testing.B) {
	ptrn, _ := pattern.NewPattern("........-....-....-....-............")
	for i := 0; i < b.N; i++ {
		for _, off := range offsetUUIDs {
			if ptrn.Lookup(off) == nil {
				b.Fatalf("the pattern %s must return non-nil against %s", ptrn, string(off))
			}
		}
	}
}

func BenchmarkRegexp_UUIDLookup(b *testing.B) {
	ptrn := regexp.MustCompile("........-....-....-....-............")
	for i := 0; i < b.N; i++ {
		for _, off := range offsetUUIDs {
			if ptrn.Find(off) == nil {
				b.Fatalf("the pattern %s must return non-nil against %s", ptrn, string(off))
			}
		}
	}
}

func BenchmarkRagel_UUIDLookup(b *testing.B) {
	var ptrn ragel.Stuff
	for i := 0; i < b.N; i++ {
		for _, off := range offsetUUIDs {
			if ptrn.UUIDIndex(off) < 0 {
				b.Fatalf("the pattern %s must give true when matched against %s", ptrn, string(off))
			}
		}
	}
}

func BenchmarkPattern_RealWorldMacAddrLookup(b *testing.B) {
	ptrn, _ := pattern.NewPattern(`..:..:..:..:..:..`)
	for i := 0; i < b.N; i++ {
		for _, line := range realWorldMacAddr {
			if ptrn.Lookup(line) == nil {
				b.Fatal("didn't locate existing pattern in a line")
			}
		}
	}
}

func BenchmarkRegexp_RealWorldMacAddrLookup(b *testing.B) {
	ptrn := regexp.MustCompile(`..:..:..:..:..:...*`)
	for i := 0; i < b.N; i++ {
		for _, line := range realWorldMacAddr {
			if ptrn.Find(line) == nil {
				b.Fatal("didn't locate existing pattern in a line")
			}
		}
	}
}

func BenchmarkRagel_RealWorldMacAddrLookup(b *testing.B) {
	var ptrn ragel.Stuff
	for i := 0; i < b.N; i++ {
		for _, line := range realWorldMacAddr {
			if ptrn.Index(line) < 0 {
				b.Fatal("didn't locate existing pattern in a line")
			}
		}
	}
}

func BenchmarkPattern_LotsOfFalsePositivesOnMacAddrLookup(b *testing.B) {
	ptrn, _ := pattern.NewPattern(`..:..:..:..:..:..`)
	for i := 0; i < b.N; i++ {
		for _, line := range falsePositivesMacAddrs {
			if ptrn.Lookup(line) == nil {
				b.Fatal("didn't locate existing pattern in a line")
			}
		}
	}
}

func BenchmarkRegexp_LotsOfFalsePositivesOnMacAddrLookup(b *testing.B) {
	ptrn := regexp.MustCompile(`..:..:..:..:..:..`)
	for i := 0; i < b.N; i++ {
		for _, line := range falsePositivesMacAddrs {
			if ptrn.Find(line) == nil {
				b.Fatal("didn't locate existing pattern in a line")
			}
		}
	}
}

func BenchmarkRagel_LotsOfFalsePositivesOnMacAddrLookup(b *testing.B) {
	var ptrn ragel.Stuff
	for i := 0; i < b.N; i++ {
		for _, line := range falsePositivesMacAddrs {
			if ptrn.Index(line) < 0 {
				b.Fatal("didn't locate existing pattern in a line")
			}
		}
	}
}
