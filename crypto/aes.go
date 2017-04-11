package crypto

import (
	"crypto/aes"
	log "github.com/cihub/seelog"
)

//EncryptData 加密
func EncryptData(src []byte, key []byte) (dst []byte, err error) {
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		log.Error("AES加密出错！", err)
		return
	}
	log.Debug("blocksize is :", aesCipher.BlockSize())
	return
}

//DecryptData 解密
func DecryptData(src []byte, key []byte) (dst []byte, err error) {
	return
}
