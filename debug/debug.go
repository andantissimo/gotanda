// +build !release

package debug

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	pwd    = ""
	gopath = ""
)

// IsEnabled is true if debug build
const IsEnabled = true

func init() {
	pwd, _ = os.Getwd()
	pwd = filepath.ToSlash(pwd) + "/"
	gopath = filepath.ToSlash(os.Getenv("GOPATH"))
	if len(gopath) > 0 && gopath[len(gopath)-1] != '/' {
		gopath += "/"
	}
	gopath += "src/"
}

// Assert tests an expression is true
func Assert(expression bool, a ...interface{}) {
	if !expression {
		pc, file, line, _ := runtime.Caller(1)
		caller := formatCaller(pc, file, line)
		desc := fmt.Sprint(a...)
		log.Fatalln(caller + ": assert(" + desc + ")")
	}
}

// Printf outputs a message to the standard logger
func Printf(format string, a ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	caller := formatCaller(pc, file, line)
	message := fmt.Sprintf(format, a...)
	log.Println(caller + ": " + message)
}

func formatCaller(pc uintptr, file string, line int) string {
	if len(gopath) < len(file) && file[:len(gopath)] == gopath {
		file = file[len(gopath):len(file)]
	} else if len(pwd) < len(file) && file[:len(pwd)] == pwd {
		file = "./" + file[len(pwd):len(file)]
	}
	return fmt.Sprintf("%s:%d", file, line)
}
