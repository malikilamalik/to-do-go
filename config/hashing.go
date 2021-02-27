package config

import (
	"crypto/subtle"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

var (
	memory     uint32 = 32 * 2048
	iterations uint32 = Uint32EnvVariable("ARGON2_ITERATIONS")
	threads    uint8  = Uint8EnvVariable("ARGON2_THREADS")
	keyLength  uint32 = Uint32EnvVariable("ARGON2_KEYLENGTH")
	salt       string = StringEnvVariable("ARGON2_SALT")
)

func GenerateFromPassword(password string) (encodedHash string, err error) {
	hash := argon2.IDKey([]byte(password), []byte(salt), iterations, memory, threads, keyLength)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	return b64Hash, nil
}

func ComparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	hash, _ := base64.RawStdEncoding.DecodeString(encodedHash)
	otherHash := argon2.IDKey([]byte(password), []byte(salt), iterations, memory, threads, keyLength)
	if subtle.ConstantTimeCompare(otherHash, hash) == 1 {
		return true, nil
	}
	return false, nil
}
