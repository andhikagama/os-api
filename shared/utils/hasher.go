package utils

import (
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/blowfish"
)

const _CipherKey string = "$4lT3d!"

func NewEncryptedByte(value string) ([]byte, error) {
	var returnMe, valueInByteArr, paddedByteArr, keyByteArr []byte
	valueInByteArr = []byte(value)
	keyByteArr = []byte(_CipherKey)
	paddedByteArr = blowfishChecksizeAndPad(valueInByteArr)
	returnMe, err := blowfishEncrypt(paddedByteArr, keyByteArr)
	if err != nil {
		return nil, err
	}
	return returnMe, nil
}

func NewEncryptedString(value string) (string, error) {
	encryptedByteArr, err := NewEncryptedByte(value)
	if err != nil {
		return "", err
	}
	returnMe := encodeBase64(encryptedByteArr)
	return returnMe, nil
}

func NewDecryptedByte(value string) ([]byte, error) {
	var returnMe, keyByteArr []byte
	keyByteArr = []byte(_CipherKey)
	decodeB64, err1 := decodeBase64(value)
	if err1 != nil {
		return nil, err1
	}
	returnMe, err2 := blowfishDecrypt(decodeB64, keyByteArr)
	if err2 != nil {
		return nil, err2
	}
	return returnMe, nil
}

func NewDecryptedString(value string) (string, error) {
	decryptedByteArr, err := NewDecryptedByte(value)
	if decryptedByteArr == nil {
		return "", err
	}
	var returnMe string = string(decryptedByteArr[:])
	return returnMe, nil
}

func NewRandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	rndStr := strings.ToUpper(hex.EncodeToString(bytes))

	return rndStr, nil
}

func NewMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func blowfishChecksizeAndPad(value []byte) []byte {
	modulus := len(value) % blowfish.BlockSize
	if modulus != 0 {
		padnglen := blowfish.BlockSize - modulus
		for i := 0; i < padnglen; i++ {
			value = append(value, 0)
		}
	}
	return value
}

func blowfishEncrypt(value, key []byte) ([]byte, error) {
	bcipher, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}
	returnMe := make([]byte, blowfish.BlockSize+len(value))
	eiv := returnMe[:blowfish.BlockSize]
	ecbc := cipher.NewCBCEncrypter(bcipher, eiv)
	ecbc.CryptBlocks(returnMe[blowfish.BlockSize:], value)
	return returnMe, nil
}

func decodeBase64(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func blowfishDecrypt(value, key []byte) ([]byte, error) {
	dcipher, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}
	div := value[:blowfish.BlockSize]
	decrypted := value[blowfish.BlockSize:]
	if len(decrypted)%blowfish.BlockSize != 0 {
		return nil, err
	}
	dcbc := cipher.NewCBCDecrypter(dcipher, div)
	dcbc.CryptBlocks(decrypted, decrypted)
	return decrypted, nil
}
