package material

import (
	"github.com/galaco/Gource-Engine/lib/stringtable"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/bsp/primitives/texinfo"
	"github.com/galaco/source-tools-common/texdatastringtable"
)

func LoadMaterials(stringData *lumps.TexdataStringData, stringTable *lumps.TexDataStringTable, texInfos *[]texinfo.TexInfo) *texdatastringtable.TexDataStringTable {
	materialStringTable := stringtable.GetTable(stringData, stringTable)

	LoadMaterialList(stringtable.SortUnique(materialStringTable, texInfos))

	return materialStringTable
}