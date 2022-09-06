package service

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/crypto/scrypt"
)

func scrypt_options() (int, int, int, int) {
	N := 32768
	r := 8
	p := 1
	key_length := 32

	return N, r, p, key_length
}

func HashPass(pass string) (string, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)

	if err != nil {
		return "", err
	}

	N, r, p, key_length := scrypt_options()

	shash, err := scrypt.Key([]byte(pass), salt, N, r, p, key_length)

	if err != nil {
		return "", err
	}

	hassed_pass := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))

	return hassed_pass, nil
}

func comparePass(suplied string, stored string) (bool, error) {
	pass_salt := strings.Split(stored, ".")

	salt, err := hex.DecodeString(pass_salt[1])

	if err != nil {
		return false, err
	}

	N, r, p, key_length := scrypt_options()
	suplied_hash, err := scrypt.Key([]byte(suplied), salt, N, r, p, key_length)

	if err != nil {
		return false, err
	}

	return hex.EncodeToString(suplied_hash) == pass_salt[0], nil
}
