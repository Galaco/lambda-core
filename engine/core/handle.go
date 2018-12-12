package core

import "strconv"

// Handle A Handle is a unique identifier for any object in the
// engine that wants one.
type Handle string

var handleCounter = int64(0)

// NewHandle returns a new handle
func NewHandle() Handle {
	handleCounter++
	return Handle("handle&" + strconv.FormatInt(handleCounter, 10))
}
