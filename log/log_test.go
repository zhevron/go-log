package log

import "testing"

func TestGetLogger(t *testing.T) {
	l := NewLogger("Test", Info)
	logger := GetLogger("Test")

	if logger != l {
		t.Errorf("expected %#q, got %#q", l, logger)
	}
}

func TestGetLogger_New(t *testing.T) {
	l := GetLogger("Test2")

	if l == nil {
		t.Error("expected non-nil, got nil")
	}

	if l.Name != "Test2" {
		t.Errorf("expected %#q, got %#q", "Test", l.Name)
	}

	if l.Level != Info {
		t.Errorf("expected %#q, got %#q", Info, l.Level)
	}
}

func TestLevelString(t *testing.T) {
	m := map[Level]string{
		None:    "",
		Debug:   "DEBUG",
		Info:    "INFO",
		Warning: "WARNING",
		Error:   "ERROR",
		Fatal:   "FATAL",
	}

	for l, s := range m {
		if l.String() != s {
			t.Errorf("expected %#q, got %#q", s, l.String())
		}
	}
}
