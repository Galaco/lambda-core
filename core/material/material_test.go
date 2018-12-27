package material

import "testing"

func TestMaterial_Bind(t *testing.T) {
	t.Skip()
}

func TestMaterial_Destroy(t *testing.T) {
	t.Skip()
}

func TestMaterial_GetFilePath(t *testing.T) {
	sut := Material{
		FilePath: "foo/bar.vmt",
	}

	if sut.GetFilePath() != "foo/bar.vmt" {
		t.Errorf("incorrect filepath returned. Expected %s, but received: %s", "foo/bar.vmt", sut.GetFilePath())
	}
}

func TestMaterial_Height(t *testing.T) {
	t.Skip()
}

func TestMaterial_Width(t *testing.T) {
	t.Skip()
}
