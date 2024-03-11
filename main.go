package main

import (
	"os"

	"github.com/dys2p/eco/ssg"
)

//go:generate gotext-update-templates -srclang=en-US -lang=de-DE,en-US -out=catalog.go -d . -d ./proxysto.re

func main() {
	ssg.Must(ssg.MakeWebsite(os.DirFS("./dys2p.com"), nil, "de", "en")).StaticHTML("/tmp/ssg-build/dys2p.com")
	ssg.Must(ssg.MakeWebsite(os.DirFS("./proxysto.re"), nil, "de", "en")).StaticHTML("/tmp/ssg-build/proxysto.re")

	ssg.ListenAndServe("/tmp/ssg-build/proxysto.re")
}
