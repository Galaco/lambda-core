package stringtable

import (
	"github.com/galaco/bsp/primitives/texinfo"
	"github.com/galaco/source-tools-common/texdatastringtable"
)

func SortUnique(stringTable *texdatastringtable.TexDataStringTable, texInfos []*texinfo.TexInfo) []string {
	materialList := make([]string, 0)
	for _, ti := range texInfos {
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
