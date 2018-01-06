package crypto

import (
	"fmt"
	"testing"
)

func TestXor(t *testing.T) {
	src := "这是一场屠杀"
	sb := []byte(src)
	key := byte(0x03)

	result := make([]byte, len(sb))
	for i, b := range sb {
		result[i] = b ^ key
	}
	fmt.Println(string(result))

	back := make([]byte, len(sb))
	for i, b := range result {
		back[i] = b ^ key
	}
	fmt.Println(string(back))
	t.Log(result)
}
