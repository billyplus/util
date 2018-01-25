package encoding

import (
	"bytes"
	"io"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GBKReaderToUTF8 将gbk的reader转换成utf的reader
func GBKReaderToUTF8(r io.Reader) io.Reader {
	return transform.NewReader(r, simplifiedchinese.GBK.NewDecoder())
}

// GBKToUTF8 将gbk的[]byte转换成utf的[]byte
func GBKToUTF8(data []byte) ([]byte, error) {
	gbkreader := bytes.NewReader(data)
	utf8reader := GBKReaderToUTF8(gbkreader)
	return ioutil.ReadAll(utf8reader)
}
