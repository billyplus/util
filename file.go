package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
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

	err = ioutil.WriteFile(file, d, 0664)
	if err != nil {
		return err
	}
	return nil
}
