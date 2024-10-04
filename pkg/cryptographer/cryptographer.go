// Package cryptographer provides an Fx module to encrypt and decrypt data.
//
// Configuration:
//   - CRYPTOGRAPHER_KEY
//
// Note: The keys should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
package cryptographer

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"

	"go.uber.org/fx"
)

const moduleName = "cryptographer"

// Module is the Fx module.
var Module = fx.Module(
	moduleName,
	fx.Provide(
		New,
	),
	fx.Provide(
		fx.Private,

		parseConfig,
	),
)

// Cryptographer defines the behavior of a cryptographer.
type Cryptographer interface {
	// Encrypt encrypts the data.
	Encrypt(data []byte) ([]byte, error)
	// Decrypt decrypts the data.
	Decrypt(data []byte) ([]byte, error)
}

func ParseConfig() (Config, error) {
	key, ok := os.LookupEnv(keyEnv)
	if !ok {
		return Config{}, errKeyNotSet
	}

	config := Config{
		Key: []byte(key),
	}

	return config, nil
}

// Config is the configuration for the cryptographer.
type Config struct {
	Key []byte
}

// Parameter is the input for the New function that creates a new cryptographer.
type Parameter struct {
	fx.In

	Config Config
}

// Result is the output for the New function that creates a new cryptographer.
type Result struct {
	fx.Out

	Cryptographer Cryptographer
}

type cryptographer struct {
	key []byte
}

const keyEnv = "CRYPTOGRAPHER_KEY"

var errKeyNotSet = fmt.Errorf("%s is not set", keyEnv)

// New creates a new cryptographer.
func New(parameter Parameter) Result {
	return Result{
		Cryptographer: cryptographer{
			key: parameter.Config.Key,
		},
	}
}

func parseConfig() (Config, error) {
	key, ok := os.LookupEnv(keyEnv)
	if !ok {
		return Config{}, errKeyNotSet
	}

	config := Config{
		Key: []byte(key),
	}

	return config, nil
}

func (c cryptographer) Encrypt(decrypted []byte) ([]byte, error) {
	gcm, err := createGCM(c.key)
	if err != nil {
		return nil, err
	}

	nonce, err := generateNonce(gcm.NonceSize())
	if err != nil {
		return nil, err
	}

	encrypted := gcm.Seal(nonce, nonce, decrypted, nil)

	return encrypted, nil
}

func (c cryptographer) Decrypt(encrypted []byte) ([]byte, error) {
	gcm, err := createGCM(c.key)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, encrypted := encrypted[:nonceSize], encrypted[nonceSize:]

	decrypted, err := gcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

func createGCM(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return gcm, nil
}

func generateNonce(size int) ([]byte, error) {
	nonce := make([]byte, size)

	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	return nonce, nil
}
