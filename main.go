package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/MoXcz/dossier-org/api"
	"github.com/MoXcz/dossier-org/helpers"
	_ "github.com/lib/pq"
)

func main() {
	listenAddr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	helpers.Logger.Info("starting server", "addr", *listenAddr)
	err := http.ListenAndServe(*listenAddr, api.Routes())
	helpers.Logger.Error(err.Error())
	os.Exit(1)
}
