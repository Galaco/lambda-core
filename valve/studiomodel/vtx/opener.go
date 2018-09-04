package vtx

import "io"

func ReadFromStream(stream io.Reader) (*Vtx, error) {
	reader := Reader{
		stream: stream,
	}
	return reader.Read()
}
