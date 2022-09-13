package crypto

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

/**
 * Generates a salt
 * @param length
 * @return salt
 */
func GenerateSalt(saltSize int) []byte {
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt[:]); err != nil {
		panic(err)
	}
	return salt
}

/**
 * Hash with sha512
 * @param plainText string
 * @return string
 */
func Hash(password string) string {
	sha512Hasher := sha512.New()
	sha512Hasher.Write([]byte(password))
	hashedPasswordBytes := sha512Hasher.Sum(nil)
	base64EncodedPasswordHash := base64.URLEncoding.EncodeToString(hashedPasswordBytes)
	return base64EncodedPasswordHash
}

/**
 * Hashs a password with a salt
 * @param password
 * @param salt
 * @return base64EncodedPasswordHash
 */
func HashWithSalt(password string, salt []byte) string {
	passwordBytes := []byte(password)
	sha512Hasher := sha512.New()
	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)
	hashedPasswordBytes := sha512Hasher.Sum(nil)
	base64EncodedPasswordHash := base64.URLEncoding.EncodeToString(hashedPasswordBytes)
	return base64EncodedPasswordHash
}

/**
 * Hashs a password with a salt and a number of iterations
 * using PBKDF2 algorithm with sha512 hash function and salt
 * @param password
 * @param salt
 * @param iterations
 * @return string
 */
func HashWithSaltAndIteration(password string, salt []byte, iterations int) string {
	passwordBytes := []byte(password)
	sha512Hasher := sha512.New()
	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)
	for i := 0; i < iterations; i++ {
		newSalt := sha512Hasher.Sum(nil)
		sha512Hasher.Reset()
		passwordBytes = append(passwordBytes, newSalt...)
		sha512Hasher.Write(passwordBytes)
	}
	hashedPasswordBytes := sha512Hasher.Sum(nil)
	base64EncodedPasswordHash := base64.URLEncoding.EncodeToString(hashedPasswordBytes)
	return base64EncodedPasswordHash
}