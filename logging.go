package main

import (
	"os"

	"github.com/op/go-logging"
)

var logFormatDefault = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)
var logFormatDebug = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{callpath} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

// Logging configuration
var log = logging.MustGetLogger("main")

// loggingInit - initial function for logging subsystem.
func loggingInit() {
	// Create backend for logs output.
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	var logFormat logging.Formatter
	if globalConfig.Debug {
		logFormat = logFormatDebug
	} else {
		logFormat = logFormatDefault
	}

	// Set log format.
	backendFormatter := logging.NewBackendFormatter(backend, logFormat)
	// Set the backends to be used and set logging level.
	backendFormatterLeveled := logging.AddModuleLevel(backendFormatter)
	logging.SetBackend(backendFormatterLeveled)
	// Set logging level.
	if globalConfig.Debug {
		backendFormatterLeveled.SetLevel(logging.DEBUG, "main")
	} else {
		backendFormatterLeveled.SetLevel(logging.INFO, "main")
	}
}
