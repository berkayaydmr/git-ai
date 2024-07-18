package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/berkayaydmr/git-ai/pkg/cryptographer"
	"github.com/berkayaydmr/git-ai/pkg/errors"
	"github.com/berkayaydmr/git-ai/pkg/storage/models"
)

type StorageInterface interface {
	GetApiKeys() ([]models.ApiKey, error)
	GetByName(name string) (string, error)
	GetByGptVersion(version string) (string, error)
	NewApiKey(apiKeyModel models.ApiKey) error
	OpenStorage() (*models.Data, error)
	SaveStorage(data *models.Data) error
	RemoveApiKey(name string) error
}

type Storage struct {
	Cryptographer cryptographer.Cryptographer
	StorageFile   string
}

func New(cipher cryptographer.Cryptographer, fileName string) StorageInterface {
	return &Storage{
		Cryptographer: cipher,
		StorageFile:   fileName,
	}
}

func (s *Storage) RemoveApiKey(name string) error {
	data, err := s.OpenStorage()
	if err != nil {
		return err
	}

	if len(data.ApiKeys) == 0 {
		return errors.ErrNoneOfApiKeysFound
	}

	if !s.IsApiKeyExistWithName(name) {
		return errors.ErrApiKeyNotFound
	}

	for i, apiKey := range data.ApiKeys {
		if apiKey.Name == name {
			data.ApiKeys = append(data.ApiKeys[:i], data.ApiKeys[i+1:]...)
			break
		}
	}

	return s.SaveStorage(data)
}

func (s *Storage) IsApiKeyExistWithName(name string) bool {
	data, err := s.OpenStorage()
	if err != nil {
		return false
	}

	if len(data.ApiKeys) == 0 {
		return false
	}

	for _, apiKey := range data.ApiKeys {
		if apiKey.Name == name {
			return true
		}
	}

	return false
}
func (s *Storage) GetApiKeys() ([]models.ApiKey, error) {
	data, err := s.OpenStorage()
	if err != nil {
		return nil, err
	}

	if len(data.ApiKeys) == 0 {
		return []models.ApiKey{}, nil
	}

	return data.ApiKeys, nil
}

func (s *Storage) GetByName(name string) (string, error) {
	data, err := s.OpenStorage()
	if err != nil {
		return "", err
	}

	if len(data.ApiKeys) == 0 {
		return "", errors.ErrNoneOfApiKeysFound
	}

	if len(data.ApiKeys) == 1 {
		return data.ApiKeys[0].Key, nil
	}

	for _, apiKey := range data.ApiKeys {
		if apiKey.Name == name {
			fmt.Println("name", name)
			return apiKey.Key, nil
		}
	}

	return "", errors.ErrApiKeyNotFound
}

func (s *Storage) GetByGptVersion(version string) (string, error) {
	data, err := s.OpenStorage()
	if err != nil {
		return "", err
	}

	if len(data.ApiKeys) == 0 {
		return "", errors.ErrNoneOfApiKeysFound
	}

	for _, apiKey := range data.ApiKeys {
		if apiKey.GptVersion.String() == version {
			return apiKey.Key, nil
		}
	}

	return "", errors.ErrApiKeyNotFound
}

func (s *Storage) NewApiKey(apiKeyModel models.ApiKey) error {
	data, err := s.OpenStorage()
	if err != nil {
		return err
	}

	if s.IsApiKeyExistWithName(apiKeyModel.Name) {
		return errors.ErrApiKeyNameExists
	}

	if len(data.ApiKeys) == 0 {
		data.ApiKeys = []models.ApiKey{apiKeyModel}
	} else {
		data.ApiKeys = append(data.ApiKeys, apiKeyModel)
	}

	if err := s.SaveStorage(data); err != nil {
		return err
	}

	return nil
}

func (s *Storage) OpenStorage() (*models.Data, error) {
	file, err := os.OpenFile(s.StorageFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("98")
		return nil, err
	}

	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if info.Size() == 0 {
		return &models.Data{}, nil
	}

	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		fmt.Println("110")
		return nil, err
	}

	decrypted, err := s.Cryptographer.Decrypt(buf[:n])
	if err != nil {
		return nil, err
	}

	var data *models.Data
	if err := json.Unmarshal(decrypted, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Storage) SaveStorage(data *models.Data) error {
	file, err := os.Create(s.StorageFile)
	if err != nil {
		return err
	}

	defer file.Close()

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return err
	}

	encrypted, err := s.Cryptographer.Encrypt(buf.Bytes())
	if err != nil {
		return err
	}

	if _, err := file.Write(encrypted); err != nil {
		return err
	}

	return nil
}
