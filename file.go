package util

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
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

//GetFileListByExt 根据文件后缀名获取文件路径列表，包括子目录
func GetFileListByExt(path string, ext string) (pathlist []string, err error) {
	var plist []string
	err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if ext == "" || filepath.Ext(path) == ext {
			plist = append(plist, path)
		}
		return nil
	})
	if err != nil {
		return
	}
	pathlist = plist
	return
}

// ReadDirByExt 根据文件后缀名返回某个目录下的全部文件名，不包括子目录
func ReadDirByExt(path string, ext string) (filelist []string, err error) {
	var files []os.FileInfo
	files, err = ioutil.ReadDir(path)
	if err != nil {
		err = errors.Wrap(err, "ReadDirByExt")
		return
	}
	for _, f := range files {
		name := f.Name()
		if ext == "" || strings.HasSuffix(name, ext) {
			filelist = append(filelist, name)
		}
	}
	return
}

// LineCounter 从io.Reader里面读取，并返回总行数。
// 根据wc的定义：A line is defined as a string of characters delimited by a <newline> character
func LineCounts(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
