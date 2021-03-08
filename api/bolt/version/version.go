package version

import (
	"strconv"

	"github.com/boltdb/bolt"
	portainer "github.com/portainer/portainer/api"
	"github.com/portainer/portainer/api/bolt/errors"
	"github.com/portainer/portainer/api/bolt/internal"
)

const (
	// BucketName represents the name of the bucket where this service stores data.
	BucketName  = "version"
	versionKey  = "DB_VERSION"
	instanceKey = "INSTANCE_ID"
	editionKey  = "EDITION"
)

// Service represents a service to manage stored versions.
type Service struct {
	db *bolt.DB
}

// NewService creates a new instance of a service.
func NewService(db *bolt.DB) (*Service, error) {
	err := internal.CreateBucket(db, BucketName)
	if err != nil {
		return nil, err
	}

	return &Service{
		db: db,
	}, nil
}

// DBVersion retrieves the stored database version.
func (service *Service) DBVersion() (int, error) {
	var data []byte

	err := service.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))

		value := bucket.Get([]byte(versionKey))
		if value == nil {
			return errors.ErrObjectNotFound
		}

		data = make([]byte, len(value))
		copy(data, value)

		return nil
	})
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(string(data))
}

// Edition retrieves the stored portainer edition.
func (service *Service) Edition() (portainer.SoftwareEdition, error) {
	editionData, err := service.getKey(editionKey)
	if err != nil {
		return 0, err
	}

	edition, err := strconv.Atoi(string(editionData))
	if err != nil {
		return 0, err
	}

	return portainer.SoftwareEdition(edition), nil
}

// StoreDBVersion store the database version.
func (service *Service) StoreDBVersion(version int) error {
	return service.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))

		data := []byte(strconv.Itoa(version))
		return bucket.Put([]byte(versionKey), data)
	})
}

// InstanceID retrieves the stored instance ID.
func (service *Service) InstanceID() (string, error) {
	var data []byte

	err := service.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))

		value := bucket.Get([]byte(instanceKey))
		if value == nil {
			return errors.ErrObjectNotFound
		}

		data = make([]byte, len(value))
		copy(data, value)

		return nil
	})
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// StoreInstanceID store the instance ID.
func (service *Service) StoreInstanceID(ID string) error {
	return service.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))

		data := []byte(ID)
		return bucket.Put([]byte(instanceKey), data)
	})
}

func (service *Service) getKey(key string) ([]byte, error) {
	var data []byte

	err := service.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))

		value := bucket.Get([]byte(key))
		if value == nil {
			return errors.ErrObjectNotFound
		}

		data = make([]byte, len(value))
		copy(data, value)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *Service) setKey(key string, value string) error {
	return service.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))

		data := []byte(value)
		return bucket.Put([]byte(key), data)
	})
}
