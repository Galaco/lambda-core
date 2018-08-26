package core

import "strconv"

// A Handle is a unique identifier for any object in the
// engine that wants one.
type Handle string

var handleCounter = int64(0)

func NewHandle() Handle {
	handleCounter++
	return Handle("handle&" + strconv.FormatInt(handleCounter, 10))
}
