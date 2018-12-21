package material

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

// @TODO This implementation is not great. It'll do until more parameters are supported,
// but its fragile and awkward to use.

// Vmt file format properties
// @TODO This could probably be replaced with the KeyValues implementation
// @TODO This does not support recursive 'include' parameter
type Vmt struct {
	Filename    string
	ShaderName  string
	properties  map[string]VmtProperty
	BaseTexture string
}

// GetFilePath returns Vmt filepath on disk
func (vmt *Vmt) GetFilePath() string {
	return vmt.Filename
}

// GetProperty returns a vmt property
func (vmt *Vmt) GetProperty(name string) VmtProperty {
	if _, ok := vmt.properties[strings.ToLower(name)]; ok {
		return vmt.properties[strings.ToLower(name)]
	}

	return VmtProperty{
		value: "",
	}
}

func (vmt *Vmt) Unload() {

}

// VmtProperty is a single vmt property.
// It allows for reading a property as any supported type, but its up to
// a calling function to know the expected type, as it should already know
// what the property is.
type VmtProperty struct {
	value string
}

// AsInt returns property as int64
func (property VmtProperty) AsInt() (int64, error) {
	return strconv.ParseInt(property.value, 10, 32)
}

// AsBool returns property as boolean
func (property VmtProperty) AsBool() (bool, error) {
	return strconv.ParseBool(property.value)
}

// AsFloat returns property as float64
func (property VmtProperty) AsFloat() (float64, error) {
	return strconv.ParseFloat(property.value, 32)
}

// AsString returns property as it was read
func (property VmtProperty) AsString() string {
	return property.value
}

// ParseVmt reads a vmt file stream and returns a Vmt
// struct with the parsed properties
func ParseVmt(filename string, stream io.Reader) (*Vmt, error) {
	reader := bufio.NewReader(stream)
	vmt := &Vmt{
		Filename: filename,
		properties: map[string]VmtProperty{
			"basetexture": {value: ""},
		},
	}

	shaderName, err := reader.ReadString([]byte("{")[0])
	if err != nil {
		return nil, err
	}
	vmt.ShaderName = sanitise(shaderName)

	depth := 1

	for depth > 0 {
		l, err := reader.ReadString([]byte("\n")[0])
		if err != nil {
			if err != io.EOF {
				return nil, errors.New("invalid vmt filesystem")
			} else {
				line := string(l)
				// Remove any comments
				line = sanitise(strings.Split(line, "//")[0])
				if isEndOfScope(line) || len(line) == 0 {
					depth = 0
					break
				}
			}
		}
		line := string(l)

		// Remove any comments
		line = sanitise(strings.Split(line, "//")[0])

		// Are we changing scope (end of vmt/new complex property)
		if isNewScope(line) {
			depth++
		} else if isEndOfScope(line) {
			depth--
		}

		if len(line) == 0 {
			continue
		}

		// Read the key value
		splitSet := strings.Split(line, " ")
		kv := [2]string{}
		for _, s := range splitSet {
			s := sanitise(s)
			if len(s) < 1 || s == " " {
				continue
			}
			if isPropertyName(s) {
				kv[0] = trimPropertyName(s)
			} else {
				kv[1] = s
			}
		}
		if len(kv[0]) > 1 {
			vmt.properties[strings.ToLower(kv[0])] = VmtProperty{value: kv[1]}
		}
	}

	return vmt, nil
}

func isNewScope(line string) bool {
	return strings.Contains(line, "{")
}

func isEndOfScope(line string) bool {
	return strings.Contains(line, "}")
}

func isPropertyName(property string) bool {
	return strings.Contains(property, "$")
}

func trimPropertyName(property string) string {
	return strings.TrimLeft(property, "$")
}

func sanitise(property string) string {
	property = strings.Replace(property, "\t", " ", -1)

	// Remove tabs
	//if strings.Contains(property, "\t") {
	//	set := strings.Split(property, "\t")
	//	for _,s := range set {
	//		if len(s) > 1 {
	//			property = s
	//		}
	//	}
	//}

	if strings.Contains(property, "\r") {
		property = strings.Replace(property, "\r", " ", -1)
	}

	if strings.Contains(property, "\n") {
		property = strings.Replace(property, "\n", " ", -1)
	}

	// Remove " escapes
	if strings.Contains(property, "\"") {
		property = strings.Replace(property, "\"", " ", -1)
	}
	// Remove ' escapes
	if strings.Contains(property, "'") {
		property = strings.Replace(property, "'", " ", -1)
	}

	property = strings.Replace(property, "\\", "/", -1)

	return strings.Trim(property, " ")
}
