package crypto

//XorWithKey 使用指定的key进行异或运算
func XorWithKey(src []byte, key byte) (dst []byte) {
	buff := make([]byte, len(src))
	for i, data := range src {
		buff[i] = data ^ key
	}
	return buff
}

//XorStringWithKey 使用指定的key对字符串进行异或运算
func XorStringWithKey(src string, key byte) (dst string) {
	srcbytes := []byte(src)
	buff := XorWithKey(srcbytes, key)
	dst = string(buff)
	return

}
