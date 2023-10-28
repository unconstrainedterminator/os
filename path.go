package util

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func GetCurrentAbs() string {
	currentPath := GetCurrentPath()
	dir := filepath.Dir(currentPath)
	return dir
}

func GetCurrentPath() string {
	dir := GetCurrentPathByExecutable()
	tmp, err := filepath.EvalSymlinks(os.TempDir())
	if err != nil {
		glog.Error(context.Background(), err)
	}

	if strings.Contains(dir, tmp) {
		return GetCurrentPathByCaller()
	}
	return dir
}

func GetCurrentPathByExecutable() string {
	currentPath, err := os.Executable()
	if err != nil {
		glog.Error(context.Background(), err)
	}

	res, err := filepath.EvalSymlinks(filepath.Dir(currentPath))
	if err != nil {
		glog.Error(context.Background(), err)
	}
	return res
}

func GetCurrentPathByCaller() string {
	var currentPath string

	_, filename, _, ok := runtime.Caller(0)
	if ok {
		currentPath = path.Dir(filename)
	}
	return currentPath
}
