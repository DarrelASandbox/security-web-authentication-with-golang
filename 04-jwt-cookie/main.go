package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	http.HandleFunc("/", emailForm)
	http.HandleFunc("/submit", submitButton)
	http.ListenAndServe(":8080", nil)
}

type claims struct {
	jwt.StandardClaims
	Email string
}

const key = "some random key"

func getJWT(msg string) (string, error) {

	c := claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},

		Email: msg,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &c)
	ss, err := token.SignedString([]byte(key))

	if err != nil {
		return "", fmt.Errorf("error: %w", err)
	}

	fmt.Printf("%v %v", ss, err)
	return ss, nil
}

func submitButton(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("email")
	if email == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	ss, err := getJWT(email)
	if err != nil {
		http.Error(w, "couldn't getJWT", http.StatusInternalServerError)
		return
	}

	// hash / message digest / digest / hash value | what we stored
	c := http.Cookie{
		Name:  "session",
		Value: ss,
	}

	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func emailForm(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{}
	}

	ss := c.Value
	// t is token before verification
	token, err := jwt.ParseWithClaims(ss, &claims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("signing method is different: %w", err)
		}

		return []byte(key), nil
	})

	/*
		`StandardClaims` has the `Valid() error` method
		which means it implements the `Claims` interface

		type Claims interface {
			Valid() error
		}

		when you `ParseWithClaims``
		the `Valid()` method gets run and
		if all is well, then returns no "error" and
		type `TOKEN` which has a field `VALID` will be true
	*/

	isEqual := err == nil && token.Valid

	message := "Not logged in"
	if isEqual {
		message = "Logged in"
		newClaims := token.Claims.(*claims)
		fmt.Println("Email:", newClaims.Email)
		fmt.Println("ExpiresAt:", newClaims.ExpiresAt) // embedded struct
	}

	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<p>Cookie value: ` + c.Value + `</p>
		<p>` + message + `</p>
		<form action="/submit" method="POST">
			<input type="email" name="email"/>
			<input type="submit"/>
		</form>
	</body>
	</html>`

	io.WriteString(w, html)
}
