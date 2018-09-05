package vtx

import (
	"bytes"
	"encoding/binary"
	"io"
	"unsafe"
)

type Reader struct {
	stream io.Reader
	buf    []byte
}

func (reader *Reader) Read() (*Vtx, error) {
	err := reader.getByteBuffer()
	if err != nil {
		return nil, err
	}

	// Read header
	header, err := reader.readHeader()
	if err != nil {
		return nil, err
	}

	offset := int32(0)

	//bodyparts
	offset += header.BodyPartOffset
	bodyParts := reader.readBodyParts(offset, header.NumBodyParts)
	offset += (header.NumBodyParts * 8)

	//models
	models := make([]modelHeader, 0)
	for _, part := range bodyParts {
		models = append(models, reader.readModels(offset + part.ModelOffset, part.NumModels)...)
	}
	offset += int32(len(models) * 8)

	//modellods
	modelLods := make([]modelLODHeader, 0)
	for _, model := range models {
		modelLods = append(modelLods, reader.readModelLODs(offset + model.LODOffset, model.NumLODs)...)
	}
	offset += int32(models[len(models)-1].LODOffset + 12)

	//meshes
	meshes := make([]meshHeader, 0)
	for _, modelLod := range modelLods {
		meshes = append(meshes, reader.readMeshes(offset + modelLod.MeshOffset, modelLod.NumMeshes)...)
	}
	offset += int32(modelLods[len(modelLods)-1].MeshOffset + 12)

	//stripgroups
	stripGroups := make([]stripGroupHeader, 0)
	for _, mesh := range meshes {
		stripGroups = append(stripGroups, reader.readStripGroups(offset + mesh.StripGroupHeaderOffset, mesh.NumStripGroups)...)
	}
	offset += int32(len(stripGroups) * 25)

	//indices

	//vertices

	//strips

	//vertexes



	return &Vtx{
		Header: header,
		BodyParts: bodyParts,
		Models: models,
	}, nil
}

// Reads studiohdr header information
func (reader *Reader) readHeader() (header, error) {
	header := header{}
	headerSize := unsafe.Sizeof(header)

	err := binary.Read(bytes.NewBuffer(reader.buf[:headerSize]), binary.LittleEndian, &header)

	return header, err
}

func (reader *Reader) readBodyParts(offset int32, num int32) []bodyPartHeader {
	ret := make([]bodyPartHeader, num)
	structSize := int32(unsafe.Sizeof(bodyPartHeader{}))
	binary.Read(bytes.NewBuffer(reader.buf[offset:offset + (structSize * num)]), binary.LittleEndian, &ret)
	return ret
}

func (reader *Reader) readModels(offset int32, num int32) []modelHeader {
	ret := make([]modelHeader, num)
	structSize := int32(unsafe.Sizeof(modelHeader{}))
	binary.Read(bytes.NewBuffer(reader.buf[offset:offset + (structSize * num)]), binary.LittleEndian, &ret)
	return ret
}

func (reader *Reader) readModelLODs(offset int32, num int32) []modelLODHeader {
	ret := make([]modelLODHeader, num)
	structSize := int32(unsafe.Sizeof(modelLODHeader{}))
	binary.Read(bytes.NewBuffer(reader.buf[offset:offset + (structSize * num)]), binary.LittleEndian, &ret)
	return ret
}

func (reader *Reader) readMeshes(offset int32, num int32) []meshHeader {
	ret := make([]meshHeader, num)
	structSize := int32(unsafe.Sizeof(meshHeader{}))
	binary.Read(bytes.NewBuffer(reader.buf[offset:offset + (structSize * num)]), binary.LittleEndian, &ret)
	return ret
}

func (reader *Reader) readStripGroups(offset int32, num int32) []stripGroupHeader {
	ret := make([]stripGroupHeader, num)
	structSize := int32(unsafe.Sizeof(stripGroupHeader{}))
	binary.Read(bytes.NewBuffer(reader.buf[offset:offset + (structSize * num)]), binary.LittleEndian, &ret)
	return ret
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
