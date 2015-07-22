package goutils

import (
	"crypto"
	//"crypto/md5"
	//"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io"
)

func VerifyRSASignWithSHA1(publicKey string, content, sign string) bool {
	pubKeyData, err := base64.StdEncoding.DecodeString(publicKey)
	pk, err := x509.ParsePKIXPublicKey(pubKeyData)
	if err != nil {
		fmt.Println("read_public_key_error", err)
		return false
	}

	signData, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		fmt.Println("decode_base64_sign_error", err)

	}
	// var h crypto.Hash
	hash := sha1.New()
	io.WriteString(hash, string(content))
	hashed := hash.Sum(nil)

	err = rsa.VerifyPKCS1v15(pk.(*rsa.PublicKey), crypto.SHA1, hashed, signData)
	if err == nil {
		return true
	}
	fmt.Println("verify_sign_error", err)
	return false

}
