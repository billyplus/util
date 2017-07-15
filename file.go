package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//IsExist 判断一个文件或者文件夹是否存在
func IsExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

//WriteJSONToFile 写入配置文件
func WriteJSONToFile(file string, data interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, d, 0664)

}

//WritePrettyJSONToFile 写入配置文件
func WritePrettyJSONToFile(file string, data interface{}) error {
	d, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, d, 0664)

}

//GetFileListByExt 根据文件后缀名获取文件列表
func GetFileListByExt(path string, ext string) ([]string, error) {
	var flist []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		fname := f.Name()
		if strings.HasSuffix(fname, ext) {
			flist = append(flist, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return flist, nil
}
