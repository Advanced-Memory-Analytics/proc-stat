package _test

import (
	"path/filepath"
	"runtime"
)

var thisDir []string

func init() {
	_, path, _, _ := runtime.Caller(0)
	thisDir = []string{filepath.Dir(path)}
}

func GetTestDir(dirName ...string) string {
	return filepath.Join(append(thisDir, dirName...)...)
}
