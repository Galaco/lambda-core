package vvd

import (
	"bytes"
	"encoding/binary"
	"github.com/go-gl/mathgl/mgl32"
	"io"
	"unsafe"
)

type Reader struct {
	stream io.Reader
	buf    []byte
}

func (reader *Reader) Read() (*Vvd, error) {
	err := reader.getByteBuffer()
	if err != nil {
		return nil, err
	}

	offset := 0
	// Read header
	header, offset, err := reader.readHeader(offset)
	if err != nil {
		return nil, err
	}

	// Read fixups
	fixups, offset := reader.readFixups(int(header.FixupTableStart), int(header.NumFixups))

	//Read vertices
	vertices, offset := reader.readVertices(int(header.VertexDataStart), &header)

	//Read tangents
	tangents, offset := reader.readTangents(int(header.TangentDataStart), &header)

	return &Vvd{
		Header:   header,
		Fixups:   fixups,
		Vertices: vertices,
		Tangents: tangents,
	}, nil
}

// Reads studiohdr header information
func (reader *Reader) readHeader(offset int) (header, int, error) {
	header := header{}
	headerSize := unsafe.Sizeof(header)

	err := binary.Read(bytes.NewBuffer(reader.buf[offset:headerSize]), binary.LittleEndian, &header)

	return header, int(headerSize), err
}

func (reader *Reader) readFixups(offset int, numFixups int) ([]fixup, int) {
	fixupSize := int(unsafe.Sizeof(fixup{}))
	fixups := make([]fixup, numFixups)
	if numFixups > 0 {
		binary.Read(bytes.NewBuffer(reader.buf[offset:offset+(fixupSize*numFixups)]), binary.LittleEndian, &fixups)
	}

	return fixups, offset + (fixupSize * numFixups)
}

// read vertex data
func (reader *Reader) readVertices(offset int, header *header) ([]vertex, int) {
	vertexSize := int(unsafe.Sizeof(vertex{}))
	// Compute number of vertices to read
	numVertices := 0
	for i := int32(0); i < header.NumLODs; i++ {
		numVertices += int(header.NumLODVertexes[i])
	}
	numVertices = int(header.NumLODVertexes[0])
	vertexes := make([]vertex, numVertices)
	binary.Read(bytes.NewBuffer(reader.buf[offset:offset+(vertexSize*numVertices)]), binary.LittleEndian, &vertexes)

	return vertexes, offset + (vertexSize * numVertices)
}

// read tangent data
// NOTE: There is 1 tangent for every vertex
func (reader *Reader) readTangents(offset int, header *header) ([]mgl32.Vec4, int) {
	tangentSize := int(unsafe.Sizeof(mgl32.Vec4{}))
	// Compute number of tangents to read
	numTangents := 0
	for i := int32(0); i < header.NumLODs; i++ {
		numTangents += int(header.NumLODVertexes[i])
	}
	numTangents = int(header.NumLODVertexes[0])
	tangents := make([]mgl32.Vec4, tangentSize)
	binary.Read(bytes.NewBuffer(reader.buf[offset:offset+(tangentSize*numTangents)]), binary.LittleEndian, &tangents)

	return tangents, offset + (tangentSize * numTangents)
}

// Read stream to []byte buffer
func (reader *Reader) getByteBuffer() error {
	buf := bytes.Buffer{}
	_, err := buf.ReadFrom(reader.stream)
	if err == nil {
		reader.buf = buf.Bytes()
	}

	return err
}
