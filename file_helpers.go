package poci

import (
	"os"
	"testing"
)

func CreateTempFile(t *testing.T, prefix string) *os.File {
	t.Helper()
	f, err := CreateTempFileBase(prefix)
	RequireNoError(t, err)
	return f
}

func CreateTempFileBase(prefix string) (*os.File, error) {
	f, err := os.CreateTemp(os.TempDir(), prefix)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func CreateConfFile(t *testing.T, content []byte) string {
	t.Helper()
	fName, err := CreateConfFileBase(content)
	if err != nil {
		t.Fatalf("Error writing conf file: %v", err)
	}
	return fName
}

func CreateConfFileBase(content []byte) (string, error) {
	conf, err := CreateTempFileBase(_EMPTY_)
	if err != nil {
		return _EMPTY_, err
	}
	fName := conf.Name()
	_ = conf.Close()
	if err := os.WriteFile(fName, content, 0666); err != nil {
		return _EMPTY_, err
	}
	return fName, nil
}

func RemoveContents(dir string) error {
	return os.RemoveAll(dir)
}

var (
	_ = CreateTempFile
	_ = CreateConfFile
	_ = RemoveContents
)
