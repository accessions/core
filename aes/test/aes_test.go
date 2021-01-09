package test

import (
	"fmt"
	"github.com/accessions/core/aes"
	"testing"
)



func TestAes_AesECBDecrypt(t *testing.T)  {
	data := []byte("this is native data.")
	key := []byte("eros")
	encrypt, err := aes.AesECBEncrypt(data, key)
	if err == nil {
		fmt.Println(1, err)
	} else {
		fmt.Println(2, string(encrypt))
	}
}
