package service

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/oklog/ulid"
	"math/rand"
	"time"

	pb "github.com/blaedj/scriptum/rpc"
	"github.com/blaedj/scriptum/store"
)

type ScriptumService struct {
	store  store.FileStore
	logger log.Logger
}

func NewScriptumService(storage store.FileStore, logger log.Logger) (*ScriptumService, error) {
	return &ScriptumService{store: storage, logger: logger}, nil
}

func (svc ScriptumService) NewDocument(ctx context.Context, doc *pb.Document, ownerID string) (*pb.NewDocumentResponse, error) {
	t := time.Unix(100000, 0)
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	docID := ulid.MustNew(ulid.Timestamp(t), entropy)

	err := svc.store.Save(doc.Contents, docID.String(), ownerID)
	if err != nil {
		return &pb.NewDocumentResponse{
			Err:        fmt.Sprintf("%v", err),
			DocumentId: "",
		}, err
	}

	return &pb.NewDocumentResponse{
		Err:        "",
		DocumentId: docID.String(),
	}, nil
}

func (svc ScriptumService) DeleteDocument(ctx context.Context, req *pb.DeleteDocumentRequest) (*pb.DeleteDocumentResponse, error) {
	// noop for now
	return &pb.DeleteDocumentResponse{Err: ""}, nil
}

//SaveDocument(context.Context, *Document) (*SaveDocumentResponse, error)
func (svc ScriptumService) SaveDocument(ctx context.Context, doc *pb.Document) (*pb.SaveDocumentResponse, error) {
	// noop for now
	return &pb.SaveDocumentResponse{}, nil
}
