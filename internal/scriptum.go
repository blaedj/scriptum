package service

import (
	"context"
	pb "github.com/blaedj/scriptum/rpc"
	"github.com/oklog/ulid"
	"log"
	"math/rand"
	"time"
)

type ScriptumService struct {
	logger log.Logger
}

func (svc *ScriptumService) NewDocument(ctx context.Context, doc pb.Document) (pb.NewDocumentResponse, error) {
	t := time.Unix(100000, 0)
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	docID := ulid.MustNew(ulid.Timestamp(t), entropy)
	return pb.NewDocumentResponse{
		Err:        "",
		DocumentId: docID.String(),
	}, nil
}
