<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#basics">Basics</a>
      <ol>
        <li><a href="#01-json-encoding">01-json-encoding</a></li>
        <li><a href="#02-authentication-basics">02-authentication-basics</a></li>
      </ol>
    </li>
  </ol>
</details>

&nbsp;

## About The Project

- Web Authentication With Golang - Google's Go Language
- Learn Golang Web Authentication, Encryption, JWT, HMAC, & OAuth with the Go Language
- [Todd McLeod](https://github.com/GoesToEleven)
- [Original Repo: golang-arch](https://github.com/GoesToEleven/golang-arch)

&nbsp;

---

&nbsp;

## Basics

### 01-json-encoding

- **Marshal**
  - go modules
- `02-encode-decode` - decode: `curl -XGET -H "Content-type: application/json" -d '{"First":"James"}' 'localhost:8080/decode'`

### 02-authentication-basics

- **base64**
  - reversible
  - never use with http; only https
- **Password storage**
  - [Go Package - bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
    - For password
  - [Go Package - hmac](https://pkg.go.dev/crypto/hmac)
    - For verifying that some message has not change
  - [Go Package - jwt](https://github.com/golang-jwt/jwt)
    - `go list -m -versions github.com/golang-jwt/jwt`
  - [Go Package - uuid](https://github.com/gofrs/uuid)

```sh
# In 03-jwt folder
go get github.com/golang-jwt/jwt
```

- **Hashing**
  - MD5 - don’t use
  - SHA
  - Bcrypt
  - Scrypt
- **Signing**
  - **Symmetric Key**
    - HMAC
    - same key to sign (encrypt) / verify (decrypt)
  - **Asymmetric Key**
    - RSA
    - ECDSA - better than RSA; faster; smaller keys
    - private key to sign (encrypt) / public key to verify (decrypt)
  - **JWT**
- **Encryption**
  - **Symmetric key**
    - AES
  - **Asymmetric Key**
    - RSA

&nbsp;

---

&nbsp;
