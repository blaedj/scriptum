package store

import (
	"errors"
	"fmt"
	"os"
)

type FileStore struct {
	ownerUUID string
}

func NewFileStore(ownerID string) *FileStore {
	return &FileStore{ownerUUID: ownerID}
}

//  TODO: what are the error conditions here?
//  TODO: the './docs' should be a config variable or flag value
func (fs *FileStore) Save(contents, docUUID string) error {
	dirpath := fmt.Sprintf("docs/%s", fs.ownerUUID)
	err := os.MkdirAll(dirpath, 0700)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to mkdir: %v", err))
	}
	f, err := os.Create(fmt.Sprintf("%s/%s", dirpath, docUUID))
	if err != nil {
		return errors.New(fmt.Sprintf("couldn't create file: %v", err))
	}

	_, err = f.WriteString(contents)
	if err != nil {
		return errors.New(fmt.Sprintf("couldn't write the file: %v", err))
	}
	err = f.Sync()

	if err != nil {
		return errors.New(fmt.Sprintf("Unable to sync file: %v", err))
	}
	defer f.Close()

	return nil
}
