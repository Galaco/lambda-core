package stringtable

import (
	bsplib "github.com/galaco/bsp"
	"github.com/galaco/source-tools-common/texdatastringtable"
	"github.com/galaco/bsp/lumps"
)

func GetTable(bsp *bsplib.Bsp) *texdatastringtable.TexDataStringTable{
	// Prepare texture lookup table
	stringData := bsp.GetLump(bsplib.LUMP_TEXDATA_STRING_DATA).(*lumps.TexdataStringData).GetData()
	stringTable := bsp.GetLump(bsplib.LUMP_TEXDATA_STRING_TABLE).(*lumps.TexDataStringTable).GetData()
	return texdatastringtable.NewTable(
		stringData,
		stringTable)
}
