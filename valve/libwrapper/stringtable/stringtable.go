package stringtable

import (
	bsplib "github.com/galaco/bsp"
	"github.com/galaco/source-tools-common/texdatastringtable"
	"github.com/galaco/bsp/lumps"
)

func GetTable(bsp *bsplib.Bsp) *texdatastringtable.TexDataStringTable{
	// Prepare texture lookup table
	stringDataLump := *bsp.GetLump(bsplib.LUMP_TEXDATA_STRING_DATA).GetContents()
	stringTableLump := *bsp.GetLump(bsplib.LUMP_TEXDATA_STRING_TABLE).GetContents()
	return texdatastringtable.NewTable(
		*stringDataLump.GetData().(*string),
		*stringTableLump.(lumps.TexDataStringTable).GetData().(*[]int32))
}
