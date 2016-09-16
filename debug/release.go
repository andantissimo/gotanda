// +build release

package debug

// IsEnabled is false if release build
const IsEnabled = false

// Assert tests an expression is true
func Assert(expression bool, a ...interface{}) {
	// do nothing for release build
}

// Printf outputs a message to the standard logger
func Printf(format string, a ...interface{}) {
	// do nothing for release build
}
