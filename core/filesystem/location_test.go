package filesystem

import "testing"

func TestGetFile(t *testing.T) {
	t.Skip()
}

func TestRegisterPakfile(t *testing.T) {
	t.Skip()
}

func TestRegisterVpk(t *testing.T) {
	t.Skip()
}

func TestUnregisterVpk(t *testing.T) {
	t.Skip()
}

func TestRegisterLocalDirectory(t *testing.T) {
	dir := "foo/bar/baz"
	RegisterLocalDirectory(dir)
	found := false
	for _,path := range localDirectories {
		if path == dir {
			found = true
			break
		}
	}
	if found == false {
		t.Error("local filepath was not found in registered paths")
	}
}

func TestUnregisterLocalDirectory(t *testing.T) {
	dir := "foo/bar/baz"
	RegisterLocalDirectory(dir)
	UnregisterLocalDirectory(dir)
	found := false
	for _,path := range localDirectories {
		if path == dir {
			found = true
			break
		}
	}
	if found == true {
		t.Error("local filepath was not found in registered paths")
	}
}

func TestUnregisterPakfile(t *testing.T) {
	t.Skip()
}
