package logger

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"log"
)

// @TODO This module should allow output pipe configuration

var a = aurora.NewAurora(false)

// Fatal error, should close the application
func Fatal(msg interface{}) {
	log.Fatal(msg)
}

// Notification for info that isn't related to any issue.
// e.g. Logging number of loaded entities
func Notice(msg interface{}, v ...interface{}) {
	switch msg.(type) {
	case string:
		print(fmt.Sprintf(msg.(string), v...), a.Gray)
	default:
		print(msg.(string), a.Gray)
	}
}

// Notifications for an unintended, but planned for issue
// e.g. Logging a prop that uses a non-existent collision model
func Warn(msg interface{}, v ...interface{}) {
	switch msg.(type) {
	case string:
		print(fmt.Sprintf(msg.(string), v...), a.Magenta)
	default:
		print(msg.(string), a.Magenta)
	}
}

// Notifications for a recoverable error
// e.g. Logging a missing resource (material, model)
func Error(msg interface{}, v ...interface{}) {
	switch msg.(type) {
	case string:
		print(fmt.Sprintf(msg.(string), v...), a.Red)
	case error:
		print(msg.(error), a.Red)
	default:
		print(msg.(string), a.Red)
	}
}

// print prints a message to console.
func print(message interface{}, col func(arg interface{}) aurora.Value) {
	fmt.Println(col(message))
}
