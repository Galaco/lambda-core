package mdl

import (
	"io"
	"unsafe"
	"bytes"
	"encoding/binary"
)

type Reader struct {
	stream io.Reader
	buf []byte
}


func (reader *Reader) Read() (*Mdl,error) {
	err := reader.getByteBuffer()
	if err != nil {
		return nil,err
	}

	header,err := reader.readHeader()
	if err != nil {
		return nil, err
	}
	if header.studioHDR2Index > 0 {
		// reader header2
	}

	// Read all properties

	return &Mdl{
		header: *header,
	}, nil
}

// Reads studiohdr header information
func (reader *Reader) readHeader() (*studiohdr,error) {
	header := studiohdr{}
	headerSize := unsafe.Sizeof(header)

	err := binary.Read(bytes.NewBuffer(reader.buf[:headerSize]), binary.LittleEndian, &header)

	return &header,err
}

// Read stream to []byte buffer
func (reader *Reader) getByteBuffer() error {
	buf := bytes.Buffer{}
	_,err := buf.ReadFrom(reader.stream)
	if err == nil {
		reader.buf = buf.Bytes()
	}

	return err
}