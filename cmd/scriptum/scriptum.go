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
		flAddr  = flag.String("addr", env.String("ADDR", ":8774"), "the server address (default :8774)")
		flDebug = flag.Bool("debug", env.Bool("DEBUG", false), "enable debug logging (default false)")
	)
	flag.Parse()
	// fstore := store.
	fstore := store.FileStore{}
	logger := logutil.NewServerLogger(*flDebug)
	scriptum.NewScriptumService(fstore, logger)
}
