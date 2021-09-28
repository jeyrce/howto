package lib

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"testing"
)

var (
	message = "woqutech"
	header  = map[string]string{
		"Author": "jeyrce@gmail.com",
	}
	private = "private.pem"
	public  = "public.pem"
)

// 密钥对导出到文件
func TestSecretFile(t *testing.T) {
	pKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatal(err)
	}
	pBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: header,
		Bytes:   x509.MarshalPKCS1PrivateKey(pKey),
	}
	pFile, err := os.Create(private)
	if err != nil {
		t.Fatal(err)
	}
	err = pem.Encode(pFile, &pBlock)
	if err != nil {
		t.Fatal(err)
	}
	pubBlock := pem.Block{
		Type:    "RSA PUBLIC KEY",
		Headers: header,
		Bytes:   x509.MarshalPKCS1PublicKey((pKey.Public()).(*rsa.PublicKey)),
	}
	pub, err := os.Create(public)
	if err != nil {
		t.Fatal(err)
	}
	err = pem.Encode(pub, &pubBlock)
	if err != nil {
		t.Fatal(err)
	}
}

// 从文件读取公钥和私钥
func ReadFile() (*rsa.PublicKey, *rsa.PrivateKey, error) {
	pubFile, err := os.ReadFile(public)
	if err != nil {
		return nil, nil, err
	}
	pubBlock, _ := pem.Decode(pubFile)
	if pubBlock == nil {
		return nil, nil, errors.New("empty")
	}
	key, err := x509.ParsePKCS1PublicKey(pubBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}
	privFile, err := os.ReadFile(private)
	if err != nil {
		return nil, nil, err
	}
	priBlock, _ := pem.Decode(privFile)
	if priBlock == nil {
		return nil, nil, errors.New("empty")
	}
	priKey, _ := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
	return key, priKey, nil
}

// 公钥加密-私钥解密
func TestSecret(t *testing.T) {
	key, priKey, err := ReadFile()
	if err != nil {
		t.Fatal(err)
	}
	// 加密过程
	secret, err := rsa.EncryptOAEP(md5.New(), rand.Reader, key, []byte(message), nil)
	if err != nil {
		t.Fatal(err)
	}
	// 解密过程
	raw, err := rsa.DecryptOAEP(md5.New(), rand.Reader, priKey, secret, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(raw))
}

// 私钥加签-公钥验签
func TestSign(t *testing.T) {
	pubKey, priKey, err := ReadFile()
	if err != nil {
		t.Fatal(err)
	}
	// 加签: 能签名的消息很短, 所以一般不直接对消息加签, 而是对消息的hash值进行签名
	hash := sha256.Sum256([]byte(message))
	code, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hash[:])
	if err != nil {
		t.Fatal(err)
	}
	// 解签
	fake := sha256.Sum256([]byte("假消息"))
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash[:], code)
	if err != nil {
		t.Fatal(err)
	}
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, fake[:], code)
	if err != nil {
		t.Fatal(err)
	}
}
