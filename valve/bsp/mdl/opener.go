package mdl

import "io"

func ReadFromStream(stream io.Reader) {
	reader := Reader{
		stream: stream,
	}
	reader.Read()
}
