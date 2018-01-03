package api

import (
	golog "log"
	"os"
)

var debug bool

type Logger interface {
	Print(...interface{})
	Printf(string, ...interface{})
}

var log Logger = golog.New(os.Stderr, "[vsz-api] ", golog.LstdFlags)

// SetLogger allows you to override the default internal logger.  "[vsz-api]" will still be prefixed to log output
func SetLogger(l Logger) {
	log = l
}

// EnableDebug enables debug-level logging
func EnableDebug() {
	debug = true
}

// DisableDebug disables debug-level logging
func DisableDebug() {
	debug = false
}
