package core

import "strconv"

type Handle string

var handleCounter = int64(0)

func NewHandle() Handle {
	handleCounter++
	return Handle("handle&" + strconv.FormatInt(handleCounter, 10))
}
