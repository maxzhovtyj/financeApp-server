package hash

import (
	"crypto/sha1"
	"fmt"
)

// PasswordHashing provides hashing logic to securely store passwords.
type PasswordHashing interface {
	Hash(password string) (string, error)
}

// SHA1Hashing uses SHA1 to hash passwords with provided salt.
type SHA1Hashing struct {
	salt string
}

func NewSHA1Hashing(salt string) *SHA1Hashing {
	return &SHA1Hashing{salt: salt}
}

// Hash creates SHA1 hash of given password.
func (h *SHA1Hashing) Hash(password string) (string, error) {
	hash := sha1.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt))), nil
}
