package jf

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"io"
)

type Cipher struct {
	hexMasterKey string
	block        cipher.Block
}

func NewCipher(hexMasterKey string) (ret *Cipher, err error) {
	ret = &Cipher{hexMasterKey: hexMasterKey}
	var key []byte
	if key, err = hex.DecodeString(hexMasterKey); err == nil {
		ret.block, err = aes.NewCipher(key)
	}
	return
}

func (o *Cipher) Decrypt(secret string) (ret []byte, err error) {
	secretType := secret[0:2]
	if secretType != "JE" {
		err = fmt.Errorf("'%v' is not a master.key secret, it has to start with JE", secret)
		return
	}

	secretData := secret[2:]
	aesBlob := base58.Decode(secretData)

	// first byte is iv size, == 16
	if ivLength := int(aesBlob[0]); ivLength != 16 {
		err = fmt.Errorf("IV is not length 16")
		return
	}

	// drop last 2 bytes of CRC
	aesSecret := aesBlob[1 : len(aesBlob)-2]
	ret = o.DecryptData(aesSecret)
	return
}

func (o *Cipher) Encrypt(data []byte) []byte {
	dataWithPadding := PKCS7AddPadding(data)
	buffer := make([]byte, aes.BlockSize+len(dataWithPadding))

	iv := buffer[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	encrypter := cipher.NewCBCEncrypter(o.block, iv)
	encrypter.CryptBlocks(buffer[len(iv):], dataWithPadding)

	return buffer
}

func (o *Cipher) DecryptData(data []byte) []byte {
	iv := data[0:aes.BlockSize]

	buffer := make([]byte, len(data)-len(iv))
	decrypter := cipher.NewCBCDecrypter(o.block, iv)
	decrypter.CryptBlocks(buffer, data[len(iv):])

	return PKCS7RemovePadding(buffer)
}

func PKCS7AddPadding(ciphertext []byte) (ret []byte) {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	ret = append(ciphertext, padText...)
	return
}

func PKCS7RemovePadding(plantText []byte) (ret []byte) {
	length := len(plantText)
	padding := int(plantText[length-1])
	ret = plantText[:(length - padding)]
	return ret
}
