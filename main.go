package main

import (
	"github.com/dys2p/eco/ssg"
)

//go:generate gotext-update-templates -srclang=en-US -lang=de-DE,en-US -out=catalog.go -d . -d ./proxysto.re

func main() {
	ssg.StaticHTML("./dys2p.com", "/tmp/ssg-build/dys2p.com")
	ssg.StaticHTML("./proxysto.re", "/tmp/ssg-build/proxysto.re")

	ssg.ListenAndServe("/tmp/ssg-build/proxysto.re")
}
