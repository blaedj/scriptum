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

func (svc ScriptumService) NewDocument(ctx context.Context, doc *pb.Document) (*pb.NewDocumentResponse, error) {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	docID := ulid.MustNew(ulid.Timestamp(t), entropy)

	ownerID := doc.OwnerId
	err := svc.store.Save(doc.Contents, docID.String(), ownerID)
	if err != nil {
		svc.logger.Log("event", "new_document_failure", "err", err)
		return &pb.NewDocumentResponse{
			Err:        fmt.Sprintf("%v", err),
			DocumentId: "",
		}, err
	}

	svc.logger.Log("event", "new_document_success", "documentID", docID.String())
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
