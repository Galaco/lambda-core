package material

import (
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/bsp/primitives/texinfo"
	"github.com/galaco/lambda-core/lib/stringtable"
	"github.com/golang-source-engine/filesystem"
	stringtableLib "github.com/golang-source-engine/stringtable"
)

// LoadMaterials is the base bsp material loader function.
// All bsp materials should be loaded by this function.
// Note that this covers bsp referenced materials only, model & entity
// materials are loaded mostly ad-hoc.
func LoadMaterials(fs *filesystem.FileSystem, stringData *lumps.TexDataStringData, stringTable *lumps.TexDataStringTable, texInfos *[]texinfo.TexInfo) *stringtableLib.StringTable {
	materialStringTable := stringtable.NewTable(stringData, stringTable)
	LoadErrorMaterial()
	LoadMaterialList(fs, stringtable.SortUnique(materialStringTable, texInfos))

	return materialStringTable
}
