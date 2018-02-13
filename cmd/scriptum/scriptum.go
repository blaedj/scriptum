package main

import (
	"flag"
	scriptum "github.com/blaedj/scriptum/internal"
	pb "github.com/blaedj/scriptum/rpc"
	"github.com/blaedj/scriptum/store"
	"github.com/kolide/kit/env"
	"github.com/kolide/kit/logutil"
	"net/http"
)

func main() {
	var (
		flAddr      = flag.String("addr", env.String("ADDR", ":8774"), "the server address (default :8774)")
		flDebug     = flag.Bool("debug", env.Bool("DEBUG", false), "enable debug logging (default false)")
		flStorePath = flag.String("storage_dir", env.String("STORAGE_DIR", "/usr/local/var/scriptum"), "Root directory for file store")
	)
	flag.Parse()
	//  TODO: properly create the FileStore, pass in the url
	fstore := store.NewFileStore(flStorePath)
	logger := logutil.NewServerLogger(*flDebug)
	svc, err := scriptum.NewScriptumService(fstore, logger)
	twirpHandler := pb.NewScriptumServiceServer(*svc, nil)
	http.ListenAndServe(flAddr, twirpHandler)
}
