package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", emailForm)
	http.HandleFunc("/submit", submitButton)
	http.ListenAndServe(":8080", nil)
}

func getCode(msg string) string {
	h := hmac.New(sha256.New, []byte("some random key"))
	h.Write([]byte(msg))
	return fmt.Sprintf("%x", h.Sum(nil))
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

	code := getCode(email)

	// hash / message digest / digest / hash value | what we stored
	c := http.Cookie{
		Name:  "session",
		Value: code + " | " + email,
	}

	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func emailForm(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{}
	}

	isEqual := true
	// slice of string, client code & client email
	ss := strings.SplitN(c.Value, " | ", 2)
	if len(ss) == 2 {
		cCode := ss[0]
		cEmail := ss[1]

		code := getCode(cEmail)
		isEqual = hmac.Equal([]byte(cCode), []byte(code))
	}

	message := "Not logged in"
	if isEqual {
		message = "Logged in"
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
