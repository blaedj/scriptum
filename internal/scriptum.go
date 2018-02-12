package service

import (
	"context"
	"fmt"
	"github.com/oklog/ulid"
	"log"
	"math/rand"
	"time"

	pb "github.com/blaedj/scriptum/rpc"
	"github.com/blaedj/scriptum/store"
)

type ScriptumService struct {
	store  store.FileStore
	logger log.Logger
}

func (svc *ScriptumService) NewDocument(ctx context.Context, doc pb.Document) (pb.NewDocumentResponse, error) {
	t := time.Unix(100000, 0)
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	docID := ulid.MustNew(ulid.Timestamp(t), entropy)

	contents := fmt.Sprintf("%s\n%s\n", doc.Title, doc.Body)
	err := svc.store.Save(contents, docID.String())
	if err != nil {
		return pb.NewDocumentResponse{
			Err:        fmt.Sprintf("%v", err),
			DocumentId: docID.String(),
		}, err
	}

	return pb.NewDocumentResponse{
		Err:        "",
		DocumentId: docID.String(),
	}, nil
}

// func (svc *ScriptumService) DeleteDocument(ctx context.Context, req pb.DeleteDocumentRequest) (pb.DeleteDocumentResponse, error) {

// }
