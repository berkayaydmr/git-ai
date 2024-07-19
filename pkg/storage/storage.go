package storage

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/berkayaydmr/git-ai/pkg/cryptographer"
	"github.com/berkayaydmr/git-ai/pkg/errors"
	"github.com/berkayaydmr/git-ai/pkg/storage/models"
)

type StorageInterface interface {
	GetProfiles() ([]models.Profile, error)
	GetProfileByName(name string) (string, error)
	NewProfile(profileModel models.Profile) error
	OpenStorage() (*models.Data, error)
	SaveStorage(data *models.Data) error
	RemoveProfile(name string) error
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

func (s *Storage) RemoveProfile(name string) error {
	data, err := s.OpenStorage()
	if err != nil {
		return err
	}

	if len(data.Profiles) == 0 {
		return errors.ErrNoProfileFound
	}

	if !s.IsProfileExistWithName(name) {
		return errors.ErrProfileNotFound
	}

	for i, profile := range data.Profiles {
		if profile.Name == name {
			data.Profiles = append(data.Profiles[:i], data.Profiles[i+1:]...)
			break
		}
	}

	return s.SaveStorage(data)
}

func (s *Storage) IsProfileExistWithName(name string) bool {
	data, err := s.OpenStorage()
	if err != nil {
		return false
	}

	if len(data.Profiles) == 0 {
		return false
	}

	for _, apiKey := range data.Profiles {
		if apiKey.Name == name {
			return true
		}
	}

	return false
}
func (s *Storage) GetProfiles() ([]models.Profile, error) {
	data, err := s.OpenStorage()
	if err != nil {
		return nil, err
	}

	if len(data.Profiles) == 0 {
		return []models.Profile{}, nil
	}

	return data.Profiles, nil
}

func (s *Storage) GetProfileByName(name string) (string, error) {
	data, err := s.OpenStorage()
	if err != nil {
		return "", err
	}

	if len(data.Profiles) == 0 {
		return "", errors.ErrNoProfileFound
	}

	if len(data.Profiles) == 1 {
		return data.Profiles[0].Key, nil
	}

	for _, apiKey := range data.Profiles {
		if apiKey.Name == name {
			return apiKey.Key, nil
		}
	}

	return "", errors.ErrProfileNotFound
}

func (s *Storage) NewProfile(apiKeyModel models.Profile) error {
	data, err := s.OpenStorage()
	if err != nil {
		return err
	}

	if s.IsProfileExistWithName(apiKeyModel.Name) {
		return errors.ErrApiKeyNameExists
	}

	if len(data.Profiles) == 0 {
		data.Profiles = []models.Profile{apiKeyModel}
	} else {
		data.Profiles = append(data.Profiles, apiKeyModel)
	}

	if err := s.SaveStorage(data); err != nil {
		return err
	}

	return nil
}

func (s *Storage) OpenStorage() (*models.Data, error) {
	file, err := os.OpenFile(s.StorageFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
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
