log - Simple logging library
============================

[![wercker status](https://app.wercker.com/status/054f2a1e2351df9bc41d07889f2bcf36/s/master "wercker status")](https://app.wercker.com/project/bykey/054f2a1e2351df9bc41d07889f2bcf36)
[![Coverage Status](https://coveralls.io/repos/zhevron/log/badge.svg?branch=master&service=github)](https://coveralls.io/github/zhevron/log?branch=master)
[![GoDoc](https://godoc.org/github.com/zhevron/log?status.svg)](https://godoc.org/github.com/zhevron/log)

**log** is a simple logging library for [Google Go](https://golang.org/).  
For full package documentation, see the GoDoc link above.

## Usage

### Log to file

```go
// Create and open a file.
f, err := os.Create("myLogFile.log")
if err != nil {
  panic(err)
}
defer f.Close()

// Add the file as an output.
log.AddOutput(f)

// Create a new logger.
logger := log.NewLogger("MyLogger", log.Info)

logger.Info("This message will appear in the log file")
```

### Log to stdout

```go
// Add stdout as an output.
log.AddOutput(os.Stdout)

// Create a new logger.
logger := log.NewLogger("MyLogger", log.Info)

logger.Info("This message will appear in stdout")
```

## License

**log** is licensed under the [MIT license](http://opensource.org/licenses/MIT).
