package main

import (
	"log"
	"net/http"
	"os"

	ipfs "github.com/ipfs/go-ipfs-api"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("expected 1 argument")
	}

	ipfs := ipfs.NewLocalShell()
	if ipfs == nil {
		log.Fatalln("local IPFS node not found")
	}

	if !ipfs.IsUp() {
		log.Fatalln("local IPFS node not up")
	}

	err := http.ListenAndServe(":3000", &ipfsHandler{
		ipfs: ipfs,
		path: os.Args[1],
	})
	if err != nil {
		log.Fatalln(err)
	}
}
