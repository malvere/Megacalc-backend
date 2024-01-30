package server

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func (s *Server) getHmacToken() (string, error) {
	signature, err := generateHMAC(s.config.Telegram.Token, "WebAppData")
	if err != nil {
		return "", err
	}
	return signature, nil
}

func generateHMAC(message, secretKey string) (string, error) {
	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	_, err := h.Write([]byte(message))
	if err != nil {
		return "", err
	}
	hashInBytes := h.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}
