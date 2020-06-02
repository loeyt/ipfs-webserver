package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"time"

	ipfs "github.com/ipfs/go-ipfs-api"
)

type ipfsHandler struct {
	ipfs *ipfs.Shell
	path string
}

func (h *ipfsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const indexPage = "index.html"
	rPath := r.URL.Path
	if !path.IsAbs(rPath) {
		rPath = "/" + rPath
	}
	rPath = path.Join(h.path, path.Clean(rPath))

	lsObject, err := h.ipfs.FileList(rPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var buf io.ReadSeeker

	if lsObject.Type == "Directory" {
		rPath = path.Join(rPath, indexPage)
	}
	rc, err := h.ipfs.Cat(rPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := ioutil.ReadAll(rc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	buf = bytes.NewReader(b)
	http.ServeContent(w, r, rPath, time.Time{}, buf)
}
