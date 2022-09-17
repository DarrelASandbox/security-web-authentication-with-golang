package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))

	password := "qwe"
	hashedPassword, err := hashPassword(password)
	if err != nil {
		panic(err)
	}

	err = comparePassword(password, hashedPassword)
	if err != nil {
		log.Fatalln("Invalid credentials")
	}

	log.Println("Logged in successfully!")
}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}
	return nil
}
