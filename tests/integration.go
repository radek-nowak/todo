package tests

import (
	"os"
	"path"
)

type TestCase struct {
	Name     string
	FileName string
	FilePath string
	Data     string
	Err      *error
}

func NewIntegrationTest(test TestCase) func() {
	os.WriteFile(path.Join(test.FilePath, test.FileName), []byte(test.Data), 0644)

	teardown := func() {
		os.RemoveAll(test.FilePath)
	}

	return teardown
}
