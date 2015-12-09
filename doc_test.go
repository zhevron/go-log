package log

import "os"

var logger = NewLogger("ExampleLogger", Info)

func ExampleAddOutput() {
	f, err := os.Create("myLogFile.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	AddOutput(f)
}

func ExampleLoggerDebug() {
	logger.Debug("This is a debug message")
}

func ExampleLoggerDebugf() {
	msg := "This is a debug message"
	logger.Debugf("%s", msg)
}

func ExampleLoggerInfo() {
	logger.Info("This is an info message")
}

func ExampleLoggerInfof() {
	msg := "This is an info message"
	logger.Infof("%s", msg)
}

func ExampleLoggerWarning() {
	logger.Warning("This is a warning message")
}

func ExampleLoggerWarningf() {
	msg := "This is a warning message"
	logger.Warningf("%s", msg)
}

func ExampleLoggerError() {
	logger.Error("This is an error message")
}

func ExampleLoggerErrorf() {
	msg := "This is an error message"
	logger.Errorf("%s", msg)
}

func ExampleLoggerFatal() {
	logger.Fatal("This is a fatal message")
}

func ExampleLoggerFatalf() {
	msg := "This is a fatal message"
	logger.Fatalf("%s", msg)
}
