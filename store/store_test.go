package store

import (
	"fmt"
	"github.com/oklog/ulid"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"
)

func newULID() string {
	t := time.Unix(100000, 0)
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	docID := ulid.MustNew(ulid.Timestamp(t), entropy)
	return docID.String()
}

func TestSuccessfulSave(t *testing.T) {
	rootDir := "./test"
	contents := "hello World"
	docID := newULID()
	ownerID := "myownerulid"
	fs := NewFileStore(rootDir)

	err := fs.Save(contents, docID, ownerID)
	if err != nil {
		t.Errorf("FileStore.save(%v, %v, %v) failed with error: %v", contents, docID, ownerID, err)
	}

	fileContents, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/%s", rootDir, ownerID, docID))
	if err != nil {
		t.Errorf("FileStore.save(%v, %v, %v) => Read failed with error: %v", contents, docID, ownerID, err)
	}
	if string(fileContents) != contents {
		t.Errorf("FileStore.save(%v, %v, %v) => Read expected %s, got %s",
			contents, docID, ownerID, contents, fileContents)
	}
	defer os.RemoveAll(rootDir)
}
