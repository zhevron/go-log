package log

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestLoggerOutput(t *testing.T) {
	l := NewLogger("Test", Info)

	buf := new(bytes.Buffer)
	l.Output(buf)

	if len(l.writers) != 1 {
		t.Errorf("expected 1, got %d", len(l.writers))
	}

	if l.writers[0].writer != buf {
		t.Errorf("expected %#q, got %#q", buf, l.writers[0].writer)
	}
}

func TestLoggerOutput_Duplicate(t *testing.T) {
	l := NewLogger("Test", Info)

	buf := new(bytes.Buffer)
	l.Output(buf)
	l.Output(buf)

	if len(l.writers) != 1 {
		t.Errorf("expected 1, got %d", len(l.writers))
	}
}

func TestLoggerOutput_Multiple(t *testing.T) {
	l := NewLogger("Test", Info)

	buf1 := new(bytes.Buffer)
	buf2 := new(bytes.Buffer)
	l.Output(buf1, buf2)

	if len(l.writers) != 2 {
		t.Errorf("expected 2, got %d", len(l.writers))
	}
}

func TestLoggerWrite(t *testing.T) {
	l := NewLogger("Test", Info)

	buf := new(bytes.Buffer)
	l.Output(buf)

	msg := "Testing log message"
	l.write(msg, 0)

	str := string(buf.Bytes())
	if !strings.Contains(str, msg) {
		t.Errorf("did not find %#q in strfer %#q", msg, str)
	}
}

func TestLoggerWrite_IncludeTimeStamp(t *testing.T) {
	l := NewLogger("Test", Info)
	l.IncludeTimeStamp = true

	buf := new(bytes.Buffer)
	l.Output(buf)

	msg := "Testing log message"
	l.write(msg, 0)

	str := string(buf.Bytes())
	if !strings.Contains(str, msg) {
		t.Errorf("did not find %#q in strfer %#q", msg, str)
	}

	now := time.Now().Format("2006-01-02")
	if !strings.Contains(str, now) {
		t.Errorf("did not find %#q in strfer %#q", now, str)
	}
}

func TestLoggerWrite_ShowFileAndLineNumber(t *testing.T) {
	l := NewLogger("Test", Info)
	l.ShowFileAndLineNumber = true

	buf := new(bytes.Buffer)
	l.Output(buf)

	msg := "Testing log message"
	l.write(msg, 1)

	str := string(buf.Bytes())
	if !strings.Contains(str, msg) {
		t.Errorf("did not find %#q in strfer %#q", msg, str)
	}

	file := "logger_test.go"
	if !strings.Contains(str, file) {
		t.Errorf("did not find %#q in strfer %#q", file, str)
	}
}

func TestLoggerWrite_ShowFileAndLineNumber_Fail(t *testing.T) {
	l := NewLogger("Test", Info)
	l.ShowFileAndLineNumber = true

	buf := new(bytes.Buffer)
	l.Output(buf)

	msg := "Testing log message"
	l.write(msg, 20)

	str := string(buf.Bytes())
	if !strings.Contains(str, msg) {
		t.Errorf("did not find %#q in strfer %#q", msg, str)
	}

	file := "<unknown>"
	if !strings.Contains(str, file) {
		t.Errorf("did not find %#q in strfer %#q", file, str)
	}
}

func TestLoggerWrite_UseUTC(t *testing.T) {
	l := NewLogger("Test", Info)
	l.UseUTC = true

	buf := new(bytes.Buffer)
	l.Output(buf)

	msg := "Testing log message"
	l.write(msg, 0)

	str := string(buf.Bytes())
	if !strings.Contains(str, msg) {
		t.Errorf("did not find %#q in strfer %#q", msg, str)
	}

	now := time.Now().UTC().Format("2006-01-02")
	if !strings.Contains(str, now) {
		t.Errorf("did not find %#q in strfer %#q", now, str)
	}
}

func TestLoggerDebug(t *testing.T) {
	l := NewLogger("Test", Debug)

	buf := new(bytes.Buffer)
	l.Output(buf)

	msg := "Testing log message"
	l.Debug(msg)

	str := string(buf.Bytes())
	if !strings.Contains(str, msg) {
		t.Errorf("did not find %#q in strfer %#q", msg, str)
	}

	level := Debug.String()
	if !strings.Contains(str, level) {
		t.Errorf("did not find %#q in strfer %#q", l, str)
	}
}

func TestLoggerDebugf(t *testing.T) {
	l := NewLogger("Test", Debug)

	buf := new(bytes.Buffer)
	l.Output(buf)

	msg := "Testing log message"
	l.Debugf("%s", msg)

	str := string(buf.Bytes())
	if !strings.Contains(str, msg) {
		t.Errorf("did not find %#q in strfer %#q", msg, str)
	}

	level := Debug.String()
	if !strings.Contains(str, level) {
		t.Errorf("did not find %#q in strfer %#q", l, str)
	}
}

func TestLoggerInfo(t *testing.T) {
	l := NewLogger("Test", Info)

	buf := new(bytes.Buffer)
	l.Output(buf)

	msg := "Testing log message"
	l.Info(msg)

	str := string(buf.Bytes())
	if !strings.Contains(str, msg) {
		t.Errorf("did not find %#q in strfer %#q", msg, str)
	}

	level := Info.String()
	if !strings.Contains(str, level) {
		t.Errorf("did not find %#q in strfer %#q", l, str)
	}
}

func TestLoggerInfof(t *testing.T) {
	l := NewLogger("Test", Info)

	buf := new(bytes.Buffer)
	l.Output(buf)

	msg := "Testing log message"
	l.Infof("%s", msg)

	str := string(buf.Bytes())
	if !strings.Contains(str, msg) {
		t.Errorf("did not find %#q in strfer %#q", msg, str)
	}

	level := Info.String()
	if !strings.Contains(str, level) {
		t.Errorf("did not find %#q in strfer %#q", l, str)
	}
}

func TestLoggerWarning(t *testing.T) {
	l := NewLogger("Test", Warning)

	buf := new(bytes.Buffer)
	l.Output(buf)

	msg := "Testing log message"
	l.Warning(msg)

	str := string(buf.Bytes())
	if !strings.Contains(str, msg) {
		t.Errorf("did not find %#q in strfer %#q", msg, str)
	}

	level := Warning.String()
	if !strings.Contains(str, level) {
		t.Errorf("did not find %#q in strfer %#q", l, str)
	}
}

func TestLoggerWarningf(t *testing.T) {
	l := NewLogger("Test", Warning)

	buf := new(bytes.Buffer)
	l.Output(buf)

	msg := "Testing log message"
	l.Warningf("%s", msg)

	str := string(buf.Bytes())
	if !strings.Contains(str, msg) {
		t.Errorf("did not find %#q in strfer %#q", msg, str)
	}

	level := Warning.String()
	if !strings.Contains(str, level) {
		t.Errorf("did not find %#q in strfer %#q", l, str)
	}
}

func TestLoggerError(t *testing.T) {
	l := NewLogger("Test", Error)

	buf := new(bytes.Buffer)
	l.Output(buf)

	msg := "Testing log message"
	l.Error(msg)

	str := string(buf.Bytes())
	if !strings.Contains(str, msg) {
		t.Errorf("did not find %#q in strfer %#q", msg, str)
	}

	level := Error.String()
	if !strings.Contains(str, level) {
		t.Errorf("did not find %#q in strfer %#q", l, str)
	}
}

func TestLoggerErrorf(t *testing.T) {
	l := NewLogger("Test", Error)

	buf := new(bytes.Buffer)
	l.Output(buf)

	msg := "Testing log message"
	l.Errorf("%s", msg)

	str := string(buf.Bytes())
	if !strings.Contains(str, msg) {
		t.Errorf("did not find %#q in strfer %#q", msg, str)
	}

	level := Error.String()
	if !strings.Contains(str, level) {
		t.Errorf("did not find %#q in strfer %#q", l, str)
	}
}

func TestLoggerFatal(t *testing.T) {
	l := NewLogger("Test", Fatal)

	buf := new(bytes.Buffer)
	l.Output(buf)

	msg := "Testing log message"
	defer func() {
		if r := recover(); r == nil {
			t.Error("did not recover from panic")
		}

		str := string(buf.Bytes())
		if !strings.Contains(str, msg) {
			t.Fatal("did not find msg in strfer")
		}

		level := Fatal.String()
		if !strings.Contains(str, level) {
			t.Errorf("did not find %#q in strfer %#q", l, str)
		}
	}()

	l.Fatal(msg)
}

func TestLoggerFatalf(t *testing.T) {
	l := NewLogger("Test", Fatal)

	buf := new(bytes.Buffer)
	l.Output(buf)

	msg := "Testing log message"
	defer func() {
		if r := recover(); r == nil {
			t.Error("did not recover from panic")
		}

		str := string(buf.Bytes())
		if !strings.Contains(str, msg) {
			t.Fatal("did not find msg in strfer")
		}

		level := Fatal.String()
		if !strings.Contains(str, level) {
			t.Errorf("did not find %#q in strfer %#q", l, str)
		}
	}()

	l.Fatalf("%s", msg)
}
