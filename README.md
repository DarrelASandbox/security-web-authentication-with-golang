<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#01-json-encoding">01-json-encoding</a></li>
    <li><a href="#02-authentication-basics">02-authentication-basics</a></li>
    <li><a href="#03-hmac-cookie">03-hmac-cookie</a></li>
    <li><a href="#04-jwt-cookie">04-jwt-cookie</a></li>
    <li><a href="#05-oatuh2">05-oatuh2</a></li>
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

## 01-json-encoding

- **Marshal**
  - go modules
- `02-encode-decode` - decode: `curl -XGET -H "Content-type: application/json" -d '{"First":"James"}' 'localhost:8080/decode'`

&nbsp;

---

&nbsp;

## 02-authentication-basics

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

## 03-hmac-cookie

&nbsp;

---

&nbsp;

## 04-jwt-cookie

- [Russ Cox - Our Software Dependency Problem](https://research.swtch.com/deps)
- [Go Package - jwt-go](https://github.com/dgrijalva/jwt-go)

&nbsp;

---

&nbsp;

## 05-oatuh2

- **OAuth2 package**
  - create a config struct
  - **authcodeURL**
    - state is some string, anything, some unique ID for this login attempt
    - returns string that you want to redirect your users to
  - **Exchange**
    - converts a code into a token
  - **TokenSource**
    - gives us a token source
  - **NewClient**
    - gives us an http client
    - this client is authenticated with the oauth resource provider
- [Hackernoon - Build your own OAuth2 Server in Go](https://hackernoon.com/build-your-own-oauth2-server-in-go-7d0f660732c3)
- [GitHub GraphQL API](https://docs.github.com/en/graphql)
- Github.com -> Settings -> Developer settings -> OAuth Apps
  - Application name
  - Homepage URL: https://example.com
  - Authorization callback URL: http://localhost:8080/oauth2/receive

&nbsp;

---

&nbsp;
