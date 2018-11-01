package stringtable

import (
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/bsp/primitives/texinfo"
	"github.com/galaco/source-tools-common/texdatastringtable"
)

func GetTable(stringData *lumps.TexdataStringData, stringTable *lumps.TexDataStringTable) *texdatastringtable.TexDataStringTable {
	// Prepare texture lookup table
	return texdatastringtable.NewTable(stringData.GetData(), stringTable.GetData())
}

func SortUnique(stringTable *texdatastringtable.TexDataStringTable, texInfos *[]texinfo.TexInfo) []string {
	materialList := make([]string, 0)
	for _, ti := range *texInfos {
		target, _ := stringTable.GetString(int(ti.TexData))
		found := false
		for _, cur := range materialList {
			if cur == target {
				found = true
				break
			}
		}
		if found == false {
			materialList = append(materialList, target)
		}
	}

	return materialList
}

