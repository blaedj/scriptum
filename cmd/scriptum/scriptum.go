package main

import (
	"flag"
	scriptum "github.com/blaedj/scriptum/internal"
	pb "github.com/blaedj/scriptum/rpc"
	"github.com/blaedj/scriptum/store"
	"github.com/kolide/kit/env"
	"github.com/kolide/kit/logutil"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		flAddr      = flag.String("addr", env.String("ADDR", ":8774"), "the server address (default :8774)")
		flDebug     = flag.Bool("debug", env.Bool("DEBUG", false), "enable debug logging (default false)")
		flStorePath = flag.String("storage_dir", env.String("STORAGE_DIR", "/usr/local/var/scriptum"), "Root directory for file store")
		flStoreType = flag.String("storage_type", env.String("STORAGE_TYPE", "file"), "Type of storage to use, can be one of: [file] (default file)")
	)

	flag.Parse()
	var documentStore store.Store

	switch *flStoreType {
	case "file":
		documentStore = store.NewFileStore(*flStorePath)
	case "mongo":
		// TODO: replace with a go-kit logger
		log.Fatalf("mongo store not yet implemented")
	default:
		// TODO: replace with a go-kit logger
		log.Fatalf("unrecognized file store type: '%s'", *flStoreType)
	}
	logger := logutil.NewServerLogger(*flDebug)
	svc, err := scriptum.NewScriptumService(documentStore, logger)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}
	twirpHandler := pb.NewScriptumServiceServer(*svc, nil)
	logger.Log("event", "server_started", "addr", *flAddr)
	http.ListenAndServe(*flAddr, twirpHandler)
}
