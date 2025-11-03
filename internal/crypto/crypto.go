package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql/driver"
	"encoding/base64"
	"errors"
	"io"
)

type Encrypter struct{ gcm cipher.AEAD }

func New(key []byte) (*Encrypter, error) {
	if len(key) < 32 {
		return nil, errors.New("AES key must be 32 bytes")
	}
	block, err := aes.NewCipher(key[:32])
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return &Encrypter{gcm: gcm}, nil
}

func (e *Encrypter) Encrypt(plain string) (string, error) {
	nonce := make([]byte, e.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ct := e.gcm.Seal(nonce, nonce, []byte(plain), nil)
	return base64.StdEncoding.EncodeToString(ct), nil
}

func (e *Encrypter) Decrypt(b64 string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", err
	}
	nonceSize := e.gcm.NonceSize()
	if len(raw) < nonceSize {
		return "", errors.New("ciphertext too short")
	}
	nonce, ct := raw[:nonceSize], raw[nonceSize:]
	pt, err := e.gcm.Open(nil, nonce, ct, nil)
	return string(pt), err
}

// GORM type helper
type EncryptedString struct {
	Raw string
	Enc *Encrypter
}

func (es EncryptedString) Value() (driver.Value, error) {
	if es.Enc == nil || es.Raw == "" {
		return "", nil
	}
	return es.Enc.Encrypt(es.Raw)
}
