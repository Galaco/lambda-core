package gameinfo

import (
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/valve/file"
	"github.com/galaco/Gource-Engine/valve/libwrapper/vpk"
	"github.com/galaco/KeyValues"
	"strings"
)

// Read game resource data paths from gameinfo.txt
// All games should ship with a gameinfo.txt, but it isn't actually mandatory
func RegisterGameResourcePaths(basePath string, gameInfo *keyvalues.KeyValue) {
	searchPaths := gameInfo.FindByKey("GameInfo").FindByKey("FileSystem").FindByKey("SearchPaths")

	for _,searchPath := range (*searchPaths.GetAllValues()) {
		kv := searchPath.(keyvalues.KeyValue)
		path := (*kv.GetAllValues())[0].(string)

		// Current directory
		if strings.Contains(path, "|gameinfo_path|") {
			path = strings.Replace(path, "|gameinfo_path|", "/" + basePath, 1)
		}
		// Executable directory
		if strings.Contains(path, "|all_source_engine_paths|") {
			path = strings.Replace(path, "|all_source_engine_paths|", basePath + "/../", 1)
		}
		if strings.Contains(*kv.GetKey(), "mod") {
			path = basePath + "/../" + path
		}

		// Strip vpk extension, then load it
		if strings.HasSuffix(path, ".vpk") {
			path = strings.Replace(path, ".vpk", "", 1)
			vpkHandle, err := vpk.OpenVPK(path)
			if err != nil {
				debug.Log(err)
				continue
			}
			file.AddVpk(vpkHandle)
			debug.Log("Registered vpk: " + path)
		} else {
			file.AddSearchDirectory(path)
			debug.Log("Registered path: " + path)
		}

	}
}