package hasher

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

const saltSize = 16

type Hasher struct {
	parameter string
}

func New(p string) *Hasher {
	return &Hasher{
		parameter: p,
	}
}

func (h *Hasher) GenerateSalt() (string, error) {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(salt), nil
}

func (h *Hasher) HashPassword(password, salt string) string {
	var passwordBytes = []byte(h.parameter + password + salt)

	var sha256Hasher = sha256.New()

	sha256Hasher.Write(passwordBytes)

	var hashedPasswordBytes = sha256Hasher.Sum(nil)

	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

func (h *Hasher) DoPasswordsMatch(hashedPassword, currPassword, salt string) bool {
	var currPasswordHash = h.HashPassword(currPassword, salt)

	return hashedPassword == currPasswordHash
}
