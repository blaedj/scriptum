package store

import (
	"github.com/oklog/ulid"
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
	contents := "hello World"
	docID := newULID()
	ownerID := "myownerulid"
	fs := NewFileStore()
	err := fs.Save(contents, docID, ownerID)
	if err != nil {
		t.Errorf("FileStore.save(%v, %v) failed with error: %v", contents, docID, err)
	}
	defer os.RemoveAll("./docs")
}
