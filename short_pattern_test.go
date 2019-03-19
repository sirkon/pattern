package pattern_test

import (
	"testing"

	"github.com/sirkon/pattern"
)

func TestShortPattern_Match(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name           string
		pattern        string
		text           string
		want           bool
		wantParseError bool
	}{
		{
			name:           "simple-ok",
			pattern:        "a.",
			text:           "ab",
			want:           true,
			wantParseError: false,
		},
		{
			name:           "simple-failure",
			pattern:        "a.",
			text:           "ba",
			want:           false,
			wantParseError: false,
		},
		{
			name:           "failed-parsing",
			pattern:        "........a",
			text:           "",
			want:           false,
			wantParseError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := pattern.NewShortPattern(tt.pattern)
			if tt.wantParseError && err == nil {
				t.Errorf("pattern '%s' should not be parsed", tt.pattern)
				return
			} else if err != nil {
				if tt.wantParseError {
					return
				}
				t.Errorf("unexpected pattern '%s' parsing error: %s", tt.pattern, err)
			}
			if got := p.Match([]byte(tt.text)); got != tt.want {
				t.Errorf("ShortPattern.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
