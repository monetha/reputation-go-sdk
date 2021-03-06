package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/monetha/go-verifiable-data/cmd"
	data "github.com/monetha/go-verifiable-data/cmd/passport-scanner/assets-data"
	"github.com/monetha/go-verifiable-data/cmd/passport-scanner/middleware"
	"github.com/shurcooL/httpgzip"
)

var (
	listen  = flag.String("listen", ":8080", "listen address")
	noCache = flag.Bool("nocache", true, "prevent browser caching")
)

func main() {
	flag.Parse()
	if cmd.HasPrintedVersion() {
		return
	}

	log.Printf("listening on %q...", *listen)

	fileServer := httpgzip.FileServer(
		data.Assets,
		httpgzip.FileServerOptions{
			IndexHTML: true,
		},
	)

	if *noCache {
		fileServer = middleware.NoCache(fileServer)
	}

	err := http.ListenAndServe(*listen, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}

		fileServer.ServeHTTP(resp, req)
	}))
	if err != nil {
		log.Fatalln(err)
	}
}
