package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

func Decode(jwt string, secret string) (interface{}, error) {
	token := strings.Split(jwt, ".")
	// check if the jwt token contains
	// header, payload and token
	if len(token) != 3 {
		splitErr := errors.New(" Invalid token: token should contain header, payload and secret")
		return nil, splitErr
	}
	// base64decode payload
	decodedPayload := base64decode(token[1])
	payload := Payload{}
	// parses payload from string to a struct
	ParseErr := json.Unmarshal([]byte(decodedPayload), &payload)
	if ParseErr != nil {
		return nil, fmt.Errorf(" Invalid payload: %s", ParseErr.Error())
	}
	// checks if the token has expired.
	if payload.Exp != 0 && time.Now().Unix() > payload.Exp {
		return nil, errors.New(" Expired token: token has expired")
	}
	signatureValue := token[0] + "." + token[1]
	// verifies if the header and signature is exactly whats in
	// the signature
	if isValidHash(signatureValue, token[2], secret) == false {
		return nil, errors.New(" Invalid token")
	}
	return payload, nil
}

func isValidHash(value string, hash string, secret string) bool {
	return hash == Hash(value, secret)
}

func base64decode(word string) string {
	length := len(word) % 4
	if length%4 > 0 {
		word += strings.Repeat("=", 4-length)
	}
	decoded, err := base64.URLEncoding.DecodeString(word)
	if err != nil {
		log.Fatalf("Decoding Error %s", err)
	}
	return string(decoded)
}
