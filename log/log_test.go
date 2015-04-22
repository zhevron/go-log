package log

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestAddSink(t *testing.T) {
	defer func() { sinks = nil }()
	s := &Sink{Level: LevelDebug, Writer: nil}

	AddSink(s)
	if len(sinks) != 1 {
		t.Errorf("expected %d, got %d", 1, len(sinks))
	}
}

func TestAddSink_Duplicate(t *testing.T) {
	defer func() { sinks = nil }()
	s := &Sink{Level: LevelDebug, Writer: nil}

	AddSink(s)
	if len(sinks) != 1 {
		t.Errorf("expected %d, got %d", 1, len(sinks))
	}

	AddSink(s)
	if len(sinks) != 1 {
		t.Errorf("expected %d, got %d", 1, len(sinks))
	}
}

func TestLog(t *testing.T) {
	defer func() { sinks = nil }()
	b := createBufferSink(LevelDebug)

	IncludeTimeStamp = false
	ShowFileAndLineNumber = false
	UseUTC = false

	msg := "Testing log message"
	Log(LevelDebug, msg)

	buf := string(b.Bytes())
	if !strings.Contains(buf, msg) {
		t.Errorf("did not find %#q in buffer %#q", msg, buf)
	}
}

func TestLog_IncludeTimeStamp(t *testing.T) {
	defer func() { sinks = nil }()
	b := createBufferSink(LevelDebug)

	IncludeTimeStamp = true
	ShowFileAndLineNumber = false
	UseUTC = false

	msg := "Testing log message"
	Log(LevelDebug, msg)

	buf := string(b.Bytes())
	if !strings.Contains(buf, msg) {
		t.Errorf("did not find %#q in buffer %#q", msg, buf)
	}

	now := time.Now().Format("2006-01-02")
	if !strings.Contains(buf, now) {
		t.Errorf("did not find %#q in buffer %#q", now, buf)
	}
}

func TestLog_ShowFileAndLineNumber(t *testing.T) {
	defer func() { sinks = nil }()
	b := createBufferSink(LevelDebug)

	IncludeTimeStamp = false
	ShowFileAndLineNumber = true
	UseUTC = false

	msg := "Testing log message"
	Log(LevelDebug, msg)

	buf := string(b.Bytes())
	if !strings.Contains(buf, msg) {
		t.Errorf("did not find %#q in buffer %#q", msg, buf)
	}

	file := "log_test.go"
	if !strings.Contains(buf, file) {
		t.Errorf("did not find %#q in buffer %#q", file, buf)
	}
}

func TestLog_UseUTC(t *testing.T) {
	defer func() { sinks = nil }()
	b := createBufferSink(LevelDebug)

	IncludeTimeStamp = true
	ShowFileAndLineNumber = false
	UseUTC = true

	msg := "Testing log message"
	Log(LevelDebug, msg)

	buf := string(b.Bytes())
	if !strings.Contains(buf, msg) {
		t.Errorf("did not find %#q in buffer %#q", msg, buf)
	}

	now := time.Now().UTC().Format("2006-01-02")
	if !strings.Contains(buf, now) {
		t.Errorf("did not find %#q in buffer %#q", now, buf)
	}
}

func TestLogf(t *testing.T) {
	defer func() { sinks = nil }()
	b := createBufferSink(LevelDebug)

	IncludeTimeStamp = false
	ShowFileAndLineNumber = false
	UseUTC = false

	msg := "Testing log message"
	Logf(LevelDebug, "%s", msg)

	buf := string(b.Bytes())
	if !strings.Contains(buf, msg) {
		t.Errorf("did not find %#q in buffer %#q", msg, buf)
	}
}

func TestDebug(t *testing.T) {
	defer func() { sinks = nil }()
	b := createBufferSink(LevelDebug)

	IncludeTimeStamp = false
	ShowFileAndLineNumber = false
	UseUTC = false

	msg := "Testing log message"
	Debug(msg)

	buf := string(b.Bytes())
	if !strings.Contains(buf, msg) {
		t.Errorf("did not find %#q in buffer %#q", msg, buf)
	}

	l := LevelDebug.String()
	if !strings.Contains(buf, l) {
		t.Errorf("did not find %#q in buffer %#q", l, buf)
	}
}

func TestDebugf(t *testing.T) {
	defer func() { sinks = nil }()
	b := createBufferSink(LevelDebug)

	IncludeTimeStamp = false
	ShowFileAndLineNumber = false
	UseUTC = false

	msg := "Testing log message"
	Debugf("%s", msg)

	buf := string(b.Bytes())
	if !strings.Contains(buf, msg) {
		t.Errorf("did not find %#q in buffer %#q", msg, buf)
	}

	l := LevelDebug.String()
	if !strings.Contains(buf, l) {
		t.Errorf("did not find %#q in buffer %#q", l, buf)
	}
}

func TestInfo(t *testing.T) {
	defer func() { sinks = nil }()
	b := createBufferSink(LevelDebug)

	IncludeTimeStamp = false
	ShowFileAndLineNumber = false
	UseUTC = false

	msg := "Testing log message"
	Info(msg)

	buf := string(b.Bytes())
	if !strings.Contains(buf, msg) {
		t.Errorf("did not find %#q in buffer %#q", msg, buf)
	}

	l := LevelInformation.String()
	if !strings.Contains(buf, l) {
		t.Errorf("did not find %#q in buffer %#q", l, buf)
	}
}

func TestInfof(t *testing.T) {
	defer func() { sinks = nil }()
	b := createBufferSink(LevelDebug)

	IncludeTimeStamp = false
	ShowFileAndLineNumber = false
	UseUTC = false

	msg := "Testing log message"
	Infof("%s", msg)

	buf := string(b.Bytes())
	if !strings.Contains(buf, msg) {
		t.Errorf("did not find %#q in buffer %#q", msg, buf)
	}

	l := LevelInformation.String()
	if !strings.Contains(buf, l) {
		t.Errorf("did not find %#q in buffer %#q", l, buf)
	}
}

func TestWarning(t *testing.T) {
	defer func() { sinks = nil }()
	b := createBufferSink(LevelDebug)

	IncludeTimeStamp = false
	ShowFileAndLineNumber = false
	UseUTC = false

	msg := "Testing log message"
	Warning(msg)

	buf := string(b.Bytes())
	if !strings.Contains(buf, msg) {
		t.Errorf("did not find %#q in buffer %#q", msg, buf)
	}

	l := LevelWarning.String()
	if !strings.Contains(buf, l) {
		t.Errorf("did not find %#q in buffer %#q", l, buf)
	}
}

func TestWarningf(t *testing.T) {
	defer func() { sinks = nil }()
	b := createBufferSink(LevelDebug)

	IncludeTimeStamp = false
	ShowFileAndLineNumber = false
	UseUTC = false

	msg := "Testing log message"
	Warningf("%s", msg)

	buf := string(b.Bytes())
	if !strings.Contains(buf, msg) {
		t.Errorf("did not find %#q in buffer %#q", msg, buf)
	}

	l := LevelWarning.String()
	if !strings.Contains(buf, l) {
		t.Errorf("did not find %#q in buffer %#q", l, buf)
	}
}

func TestError(t *testing.T) {
	defer func() { sinks = nil }()
	b := createBufferSink(LevelDebug)

	IncludeTimeStamp = false
	ShowFileAndLineNumber = false
	UseUTC = false

	msg := "Testing log message"
	Error(msg)

	buf := string(b.Bytes())
	if !strings.Contains(buf, msg) {
		t.Errorf("did not find %#q in buffer %#q", msg, buf)
	}

	l := LevelError.String()
	if !strings.Contains(buf, l) {
		t.Errorf("did not find %#q in buffer %#q", l, buf)
	}
}

func TestErrorf(t *testing.T) {
	defer func() { sinks = nil }()
	b := createBufferSink(LevelDebug)

	IncludeTimeStamp = false
	ShowFileAndLineNumber = false
	UseUTC = false

	msg := "Testing log message"
	Errorf("%s", msg)

	buf := string(b.Bytes())
	if !strings.Contains(buf, msg) {
		t.Errorf("did not find %#q in buffer %#q", msg, buf)
	}

	l := LevelError.String()
	if !strings.Contains(buf, l) {
		t.Errorf("did not find %#q in buffer %#q", l, buf)
	}
}

func TestFatal(t *testing.T) {
	defer func() { sinks = nil }()
	b := createBufferSink(LevelDebug)

	IncludeTimeStamp = false
	ShowFileAndLineNumber = false
	UseUTC = false

	msg := "Testing log message"
	defer func() {
		if r := recover(); r == nil {
			t.Error("did not recover from panic")
		}

		buf := string(b.Bytes())
		if !strings.Contains(buf, msg) {
			t.Fatal("did not find msg in buffer")
		}

		l := LevelFatal.String()
		if !strings.Contains(buf, l) {
			t.Errorf("did not find %#q in buffer %#q", l, buf)
		}
	}()
	Fatal(msg)
}

func TestFatalf(t *testing.T) {
	defer func() { sinks = nil }()
	b := createBufferSink(LevelDebug)

	IncludeTimeStamp = false
	ShowFileAndLineNumber = false
	UseUTC = false

	msg := "Testing log message"
	defer func() {
		if r := recover(); r == nil {
			t.Error("did not recover from panic")
		}

		buf := string(b.Bytes())
		if !strings.Contains(buf, msg) {
			t.Fatal("did not find msg in buffer")
		}

		l := LevelFatal.String()
		if !strings.Contains(buf, l) {
			t.Errorf("did not find %#q in buffer %#q", l, buf)
		}
	}()
	Fatalf("%s", msg)
}

func TestLevelString(t *testing.T) {
	m := map[Level]string{
		LevelNone:        "",
		LevelDebug:       "DEBUG",
		LevelInformation: "INFO",
		LevelWarning:     "WARNING",
		LevelError:       "ERROR",
		LevelFatal:       "FATAL",
	}

	for l, s := range m {
		if l.String() != s {
			t.Errorf("expected %#q, got %#q", s, l.String())
		}
	}
}

func createBufferSink(level Level) *bytes.Buffer {
	buf := new(bytes.Buffer)
	AddSink(&Sink{Level: level, Writer: buf})
	return buf
}
