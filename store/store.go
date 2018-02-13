package store

import (
	"errors"
	"fmt"
	"os"
)

type FileStore struct {
	rootDir string // root path to store files at.
}

func NewFileStore(root string) *FileStore {
	return &FileStore{rootDir: root}
}

func dirExists(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

//  TODO: what are the error conditions here?
//  TODO: the 'docs' should be a config variable or flag value
func (fs *FileStore) Save(contents, docUUID, ownerID string) error {
	dirpath := fmt.Sprintf("%s/%s", fs.rootDir, ownerID)
	haveDir, _ := dirExists(dirpath)
	if !haveDir {
		err2 := os.MkdirAll(dirpath, 0700)
		if err2 != nil {
			return errors.New(fmt.Sprintf("failed to mkdir: %v", err2))
		}
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
