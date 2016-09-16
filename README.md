# Gotanda

Cross-platform debugging and logging library for Go.

## Installation

```bash
go get -u github.com/andantissimo/gotanda
```

## Usage

### Example Code

```go
package main

import (
    "github.com/andantissimo/gotanda/debug"
    "github.com/andantissimo/gotanda/log"
)

// open the syslog|eventlog before using log.*
var _ = log.Open(log.LOCAL0, "Gotanda")

func main() {
    var x = 42
    debug.Assert(x == 42, "x == 42")
    debug.Printf("x = %d", x)
    if debug.IsEnabled {
        log.Infof("Debug is enabled")
    } else {
        log.Infof("Debug is disabled")
    }
    log.Warningf("Warning message")
    log.Errf("Error message")
}
```

### Debug Build

`go run main.go` or press F5 on Visual Studio Code

### Release Build

`go build -tags "release" main.go`

## License

WTFPL
