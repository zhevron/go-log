package log

import (
	"bytes"
	"os"
	"testing"

	"github.com/zhevron/match"
)

var buffer = new(bytes.Buffer)

func TestMain(m *testing.M) {
	AddOutput(buffer)
	retCode := m.Run()
	os.Exit(retCode)
}

func TestAddOutput(t *testing.T) {
	buf := new(bytes.Buffer)
	AddOutput(buf)
	match.Equals(t, len(writers), 2).Fatal()
	match.Equals(t, writers[1].writer, buf).Fatal()
	writers = writers[:len(writers)-1]
}

func TestLoggerAddOutput_Duplicate(t *testing.T) {
	buf := new(bytes.Buffer)
	AddOutput(buf)
	AddOutput(buf)
	match.Equals(t, len(writers), 2).Fatal()
	writers = writers[:len(writers)-1]
}

func TestLoggerAddOutput_Multiple(t *testing.T) {
	buf1 := new(bytes.Buffer)
	buf2 := new(bytes.Buffer)
	AddOutput(buf1, buf2)
	match.Equals(t, len(writers), 3).Fatal()
	writers = writers[:len(writers)-2]
}
