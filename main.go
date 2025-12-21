package main

import (
	"os"

	"github.com/dys2p/eco/lang"
	"github.com/dys2p/eco/ssg"
)

//go:generate gotext-update-templates -srclang=en-US -lang=de-DE,en-US -out=catalog.go -d . -d ./proxysto.re

func main() {
	langs := lang.MakeLanguages(nil, "de", "en")
	must(ssg.Must(ssg.MakeWebsite(os.DirFS("./dys2p.com"), nil, langs)).WriteFiles("/tmp/ssg-build/dys2p.com", false))
	must(ssg.Must(ssg.MakeWebsite(os.DirFS("./dys2p.com"), nil, langs)).WriteFiles("/tmp/ssg-build/dys2pwwos5w5kez2chufdk3b3oyj5n4n4iiseyke2tzuqahgvmovloyd.onion", true))
	must(ssg.Must(ssg.MakeWebsite(os.DirFS("./proxysto.re"), nil, langs)).WriteFiles("/tmp/ssg-build/proxysto.re", false))
	must(ssg.Must(ssg.MakeWebsite(os.DirFS("./proxysto.re"), nil, langs)).WriteFiles("/tmp/ssg-build/proxyoxiemywllckvpix543gqcmvvltrnb7inbwtk2knkehqt72tyfyd.onion", true))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
