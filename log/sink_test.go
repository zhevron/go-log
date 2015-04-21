package log

import (
	"bytes"
	"testing"
)

func TestNewSink(t *testing.T) {
	s := NewSink(LevelDebug, nil)
	if s == nil {
		t.Error("expected non-nil, got nil")
	}
}

func TestSinkWrite(t *testing.T) {
	b := new(bytes.Buffer)
	s := &Sink{Level: LevelDebug, Writer: b}

	msg := "Test"
	if err := s.Write(LevelDebug, msg); err != nil {
		t.Error(err)
	}
	if string(b.Bytes()) != msg {
		t.Errorf("expected %#q, got %#q", msg, string(b.Bytes()))
	}
}

func TestSinkWrite_SkipLevel(t *testing.T) {
	b := new(bytes.Buffer)
	s := &Sink{Level: LevelInformation, Writer: b}

	if err := s.Write(LevelDebug, "Test"); err != nil {
		t.Error(err)
	}
	if len(b.Bytes()) != 0 {
		t.Errorf("expected %d, got %d", 0, len(b.Bytes()))
	}
}
