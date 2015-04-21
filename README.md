go-log - Simple logging library
===============================

[![wercker status](https://app.wercker.com/status/c98a6d8f01bffef5bca40c3563347dba/m "wercker status")](https://app.wercker.com/project/bykey/c98a6d8f01bffef5bca40c3563347dba)

[![Coverage Status](https://coveralls.io/repos/zhevron/go-log/badge.svg?branch=HEAD)](https://coveralls.io/r/zhevron/go-log?branch=HEAD)
[![GoDoc](https://godoc.org/gopkg.in/zhevron/go-log.v0/log?status.svg)](https://godoc.org/gopkg.in/zhevron/go-log.v0/log)

**go-log** is a simple logging library for [Go](https://golang.org/).  

For package documentation, refer to the GoDoc badge above.

## Installation

```
go get gopkg.in/zhevron/go-log.v0/log
```

## Usage

### Log to file

```go
package main

import (
  "os"

  "gopkg.in/zhevron/go-log.v0/log"
)

func main() {
  // Open the log file.
  f, err := os.Create("myLogFile.log")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  // Add it as a sink.
  log.AddSink(log.NewSink(log.LevelDebug, f))

  log.Info("This message will appear in the log file")
}
```

### Log to stdout

```go
package main

import "gopkg.in/zhevron/go-log.v0/log"

func main() {
  // Add the standard Stdout sink.
  log.AddSink(log.Stdout)

  log.Info("This message will appear in stdout")
}
```

## License

**go-log** is licensed under the [MIT license](http://opensource.org/licenses/MIT).
