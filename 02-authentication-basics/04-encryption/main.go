package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	msg := "some message"

	encodingBase64(msg)
	encryptEmail(msg)
	sha256Encryption()
}

func encodingBase64(msg string) {
	fmt.Println("encodingBase64:")

	encoded := encode(msg)
	fmt.Println(encoded)

	s, err := decode(encoded)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Decoded message: ", s)
}

func encode(msg string) string {
	return base64.URLEncoding.EncodeToString([]byte(msg))
}

func decode(encoded string) (string, error) {
	s, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return "", fmt.Errorf("couldn't decode string %w", err)
	}

	return string(s), nil
}

func encryptEmail(msg string) {
	fmt.Println("\n\nencryptEmail:")

	password := "qwe"
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatalln("couldn't bcrypt password:", err)
	}

	bs = bs[:16]

	result, err := cryptoAES(bs, msg)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(result))

	result2, err := cryptoAES(bs, string(result))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(result2))

	fmt.Println("\nReuse code with encryptWriter:")
	wtr := &bytes.Buffer{}
	encWriter, err := encryptWriter(wtr, bs)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = io.WriteString(encWriter, msg)
	if err != nil {
		log.Fatalln(err)
	}

	encrypted := wtr.String()
	fmt.Println(encrypted)
}

func cryptoAES(key []byte, input string) ([]byte, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error %w", err)
	}

	// initialization vector
	iv := make([]byte, aes.BlockSize)
	// can improve by randomizing bytes with
	// _, err = io.ReadFull(rand.Reader, iv)

	buf := &bytes.Buffer{}
	s := cipher.NewCTR(b, iv)
	sw := cipher.StreamWriter{
		S: s,
		W: buf,
	}
	_, err = sw.Write([]byte(input))
	if err != nil {
		return nil, fmt.Errorf("couldn't sw.Write to streamwriter %w", err)
	}

	return buf.Bytes(), nil
}

func encryptWriter(wtr io.Writer, key []byte) (io.Writer, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("couldn't newCipher %w", err)
	}

	//initialization vector
	iv := make([]byte, aes.BlockSize)

	s := cipher.NewCTR(b, iv)

	return cipher.StreamWriter{
		S: s,
		W: wtr,
	}, nil
}

func sha256Encryption() {
	fmt.Println("\n\nsha256Encryption:")
	f, err := os.Open("04-encryption/sample.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	h := sha256.New()

	_, err = io.Copy(h, f)
	if err != nil {
		log.Fatalln("error %w", err)
	}

	fmt.Printf("sha256.New() is %T\n", h)
	fmt.Printf("\nBefore Sum:\n%v\n", h)
	bs := h.Sum(nil)
	fmt.Printf("\nAfter Sum:\n%x\n", bs)

	bs = h.Sum(nil)
	bs = h.Sum(bs)
	fmt.Printf("\nSum again:\n%x\n", bs)
}
