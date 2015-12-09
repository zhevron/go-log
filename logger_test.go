package log

import (
	"testing"
	"time"

	"github.com/zhevron/match"
)

var loggerWriteTests = []struct {
	in        string
	out       string
	level     level
	calldepth int
}{
	{"Test Message", "Test Message", Info, 1},
	{"Test Message", "", Debug, 1},
	{"Test Message", time.Now().UTC().Format("2006-01-02"), Info, 1},
	{"Test Message", "logger_test.go", Info, 1},
	{"Test Message", "<unknown>", Info, 10},
}

func TestGetLogger(t *testing.T) {
	l := NewLogger("Test", Info)
	match.Equals(t, GetLogger("Test"), l)
	l = GetLogger("Test2")
	match.IsNotNil(t, l)
	match.Equals(t, l.Name, "Test2")
	match.Equals(t, l.MinimumLevel, Info)
}

func TestNewLogger(t *testing.T) {
	l := NewLogger("Test3", Info)
	match.IsNotNil(t, l)
}

func TestLoggerWrite(t *testing.T) {
	logger := NewLogger("TestLoggerWrite", Info)
	logger.UseUTC = true
	logger.ShowFileAndLineNumber = true
	for _, tt := range loggerWriteTests {
		buffer.Reset()
		logger.write(tt.in, tt.level, tt.calldepth)
		match.Contains(t, string(buffer.Bytes()), tt.out)
	}
}

func TestLoggerLevels(t *testing.T) {
	logger := NewLogger("TestLoggerLevels", Debug)
	buffer.Reset()
	logger.Debug("STR_" + Debug.String())
	logger.Debugf("FORMAT_" + Debug.String())
	match.Contains(t, string(buffer.Bytes()), "STR_"+Debug.String()).Contains("FORMAT_" + Debug.String())
	buffer.Reset()
	logger.Info("STR_" + Info.String())
	logger.Infof("FORMAT_" + Info.String())
	match.Contains(t, string(buffer.Bytes()), "STR_"+Info.String()).Contains("FORMAT_" + Info.String())
	buffer.Reset()
	logger.Warning("STR_" + Warning.String())
	logger.Warningf("FORMAT_" + Warning.String())
	match.Contains(t, string(buffer.Bytes()), "STR_"+Warning.String()).Contains("FORMAT_" + Warning.String())
	buffer.Reset()
	logger.Error("STR_" + Error.String())
	logger.Errorf("FORMAT_" + Error.String())
	match.Contains(t, string(buffer.Bytes()), "STR_"+Error.String()).Contains("FORMAT_" + Error.String())
}

func TestLoggerFatal(t *testing.T) {
	logger := NewLogger("TestLoggerFatal", Info)
	defer func() {
		r := recover()
		match.IsNotNil(t, r)
		match.Contains(t, string(buffer.Bytes()), "STR_"+Fatal.String())
	}()
	buffer.Reset()
	logger.Fatal("STR_" + Fatal.String())
}

func TestLoggerFatalf(t *testing.T) {
	logger := NewLogger("TestLoggerFatalf", Info)
	defer func() {
		r := recover()
		match.IsNotNil(t, r)
		match.Contains(t, string(buffer.Bytes()), "FORMAT_"+Fatal.String())
	}()
	buffer.Reset()
	logger.Fatalf("FORMAT_" + Fatal.String())
}
