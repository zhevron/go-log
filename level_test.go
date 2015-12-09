package log

import (
	"testing"

	"github.com/zhevron/match"
)

var levelTests = []struct {
	in  level
	out string
}{
	{Debug, "DEBUG"},
	{Info, "INFO"},
	{Warning, "WARNING"},
	{Error, "ERROR"},
	{Fatal, "FATAL"},
	{level(255), "UNKNOWN"},
}

func TestLevelString(t *testing.T) {
	for _, tt := range levelTests {
		match.Equals(t, tt.in.String(), tt.out)
	}
}
