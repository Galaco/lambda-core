package vvd

import "io"

func ReadFromStream(stream io.Reader) (*Vvd, error) {
	reader := Reader{
		stream: stream,
	}
	return reader.Read()
}
