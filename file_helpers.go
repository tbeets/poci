package poci

import (
	"os"
	"path/filepath"
	"testing"
)

func CreateTempFile(t testing.TB, prefix string) *os.File {
	t.Helper()
	tempDir := t.TempDir()
	f, err := os.CreateTemp(tempDir, prefix)
	Require_NoError(t, err)
	return f
}

func CreateConfFile(t testing.TB, content []byte) string {
	t.Helper()
	conf := CreateTempFile(t, _EMPTY_)
	fName := conf.Name()
	conf.Close()
	if err := os.WriteFile(fName, content, 0666); err != nil {
		t.Fatalf("Error writing conf file: %v", err)
	}
	return fName
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
