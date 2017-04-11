package io

import (
	"errors"
	//log "github.com/gogap/logrus"
	"net"
)

var (
	//BlockSize 复制块大小
	BlockSize = 32 * 1024
)

//Copy 复制两个端口
func Copy(inConn net.Conn, outConn net.Conn) (n int, err error) {
	buff := make([]byte, BlockSize)
	//log.Println("start copy")
	for {
		if inConn == nil {
			n = 0
			err = errors.New("没有初始化连接！")
			return
		}
		nr, er := inConn.Read(buff)
		if er != nil {
			err = er
			break
		}
		if nr > 0 {
			//log.Debugln("buff copied: ", string(buff))
			nw, ew := outConn.Write(buff[0:nr])
			if nw > 0 {
				n += int(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = errors.New("short writting buff")
				break
			}
		}

	}
	return
}

//CopyXorWithKey 复制两个连接，利用xor加密
func CopyXorWithKey(inConn net.Conn, outConn net.Conn, key byte) (n int, err error) {
	buff := make([]byte, BlockSize)
	//log.Debugln("start copy")
	for {
		if inConn == nil {
			n = 0
			err = errors.New("没有初始化连接！")
			return
		}
		nr, er := inConn.Read(buff)
		if er != nil {
			err = er
			break
		}
		if nr > 0 {
			//log.Debugln("buff copied: ", string(buff))
			encryptedBuff := make([]byte, BlockSize)
			for i, data := range buff {
				encryptedBuff[i] = data ^ key
			}
			nw, ew := outConn.Write(encryptedBuff[0:nr])
			if nw > 0 {
				n += int(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = errors.New("short writting buff")
				break
			}
		}

	}
	return
}
