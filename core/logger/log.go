package logger

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"log"
)

// @TODO This module should allow output pipe configuration

var colourer = aurora.NewAurora(false)

var pipeFunc = func(value string) {
	fmt.Println(value)
}

// SetOutputPipeFunc exposes the ability to change
// where logs are written to by replacing the final print function
// with colourer custom implementation
func SetOutputPipeFunc(callback func(string)) {
	pipeFunc = callback
}

func EnablePretty() {
	colourer = aurora.NewAurora(true)
}

func DisablePretty() {
	colourer = aurora.NewAurora(false)
}

// Fatal error, should close the application
func Fatal(msg interface{}) {
	log.Fatal(msg)
}

// Notification for info that isn't related to any issue.
// e.g. Logging number of loaded entities
func Notice(msg interface{}, v ...interface{}) {
	switch msg.(type) {
	case string:
		print(fmt.Sprintf(msg.(string), v...), colourer.Gray)
	default:
		print(msg.(string), colourer.Gray)
	}
}

// Notifications for an unintended, but planned for issue
// e.g. Logging colourer prop that uses colourer non-existent collision model
func Warn(msg interface{}, v ...interface{}) {
	switch msg.(type) {
	case string:
		print(fmt.Sprintf(msg.(string), v...), colourer.Magenta)
	default:
		print(msg.(string), colourer.Magenta)
	}
}

// Notifications for colourer recoverable error
// e.g. Logging colourer missing resource (material, model)
func Error(msg interface{}, v ...interface{}) {
	switch msg.(type) {
	case string:
		print(fmt.Sprintf(msg.(string), v...), colourer.Red)
	case error:
		print(msg.(error), colourer.Red)
	default:
		print(msg.(string), colourer.Red)
	}
}

// print prints colourer message to console.
func print(message interface{}, col func(arg interface{}) aurora.Value) {
	pipeFunc(fmt.Sprint(col(message)))
}
