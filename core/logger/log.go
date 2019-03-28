package logger

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"io"
)

var colourer = aurora.NewAurora(false)

var internalWriter io.Writer

// SetWriter exposes the ability to change
// where logs are written to by replacing the final print function
// with colourer custom implementation
func SetWriter(writer io.Writer) {
	internalWriter = writer
}

// EnablePretty
func EnablePretty() {
	colourer = aurora.NewAurora(true)
}

// DisablePretty
func DisablePretty() {
	colourer = aurora.NewAurora(false)
}

// Panic error, should close the application
func Panic(msg interface{}) {
	Error(msg)
	panic(msg)
}

// Notice Notification for info that isn't related to any issue.
// e.g. Logging number of loaded entities
func Notice(msg interface{}, v ...interface{}) {
	switch t := msg.(type) {
	case string:
		print(fmt.Sprintf(t, v...), colourer.Gray)
	default:
		print(msg.(string), colourer.Gray)
	}
}

// Warn Notifications for an unintended, but planned for issue
// e.g. Logging colourer prop that uses colourer non-existent collision model
func Warn(msg interface{}, v ...interface{}) {
	switch t := msg.(type) {
	case string:
		print(fmt.Sprintf(t, v...), colourer.Magenta)
	default:
		print(msg.(string), colourer.Magenta)
	}
}

// Error Notifications for colourer recoverable error
// e.g. Logging colourer missing resource (material, model)
func Error(msg interface{}, v ...interface{}) {
	switch t := msg.(type) {
	case string:
		print(fmt.Sprintf(t, v...), colourer.Red)
	case error:
		print(t, colourer.Red)
	default:
		print(msg.(string), colourer.Red)
	}
}

// print prints colourer message to console.
func print(message interface{}, col func(arg interface{}) aurora.Value) {
	if _,err := internalWriter.Write([]byte(fmt.Sprint(col(message)))); err != nil {
		panic(err)
	}
}
