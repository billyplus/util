package util

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

var (
	appRootPath string
)

func init() {
	appRootPath, _ = GetCurrentPath()
}

// GetAppRootPath return root path of current application
func GetAppRootPath() string {
	return appRootPath
}

// GetCurrentPath return current path of running application
func GetCurrentPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return wd, nil
}

func GetFullPath(file string) (string, error) {
	r, err := filepath.Abs(file)
	if err != nil {
		return "", errors.Wrap(err, "GetFullPath: ")
	}
	return r, nil
}

// CreateDirIfNotExist create dir is not exist, else do nothing
func CreateDirIfNotExist(path string) error {
	isExist := IsExist(path)
	if !isExist {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return errors.Wrapf(err, "CreateDirIfNotExist:(MkdirAll %v)", path)
		}
	}
	return nil
}
