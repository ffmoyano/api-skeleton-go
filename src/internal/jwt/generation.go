package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub string `json:"sub"`
	Exp int64  `json:"exp"`
}

func base64encode(word []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(word), "=")
}

func Hash(src string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(src))
	return base64encode(h.Sum(nil))
}

func Generate(payload Payload, secret string) (string, error) {
	jsonHeader, err := json.Marshal(Header{"HS256", "JWT"})
	if err != nil {
		return fmt.Sprintf("Error while encoding to json the gopherjwt header: %s", err), err
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Sprintf("Error while encoding to json the gopherjwt payload: %s", err), err
	}
	signatureValue := base64encode(jsonHeader) + "." + base64encode(jsonPayload)
	return signatureValue + "." + Hash(signatureValue, secret), nil
}
