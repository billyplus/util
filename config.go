package util

import (
	"encoding/json"
	"io/ioutil"
)

//ParseConfigFromFile 从文件中读取配置
func ParseConfigFromFile(cfg interface{}, file string) error {
	d, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	//var conf interface{}
	err = json.Unmarshal(d, cfg)
	if err != nil {
		return err
	}
	return nil
}

//WriteConfigToFile 写入配置文件
func WriteConfigToFile(data []byte, file string) error {
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
