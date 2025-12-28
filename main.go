package main

import (
	"flag"
	"net/http"
	"os"
	"strings"

	"github.com/dys2p/eco/lang"
	"github.com/dys2p/eco/ssg"
)

//go:generate gotext-update-templates -srclang=en-US -lang=de-DE,en-US -out=catalog.go -d . -d ./proxysto.re

type TemplateData struct {
	ssg.TemplateData
	Active string
	Onion  bool
}

func main() {
	listen := flag.Bool("listen", false, "run webserver at 127.0.0.1:8080")
	flag.Parse()

	langs := lang.MakeLanguages(nil, "de", "en")
	for _, site := range []struct {
		SrcDir string
		Active string
		Onion  bool
		OutDir string
	}{
		{"./dys2p.com", "dys2p", false, "/tmp/ssg-build/dys2p.com"},
		{"./dys2p.com", "dys2p", true, "/tmp/ssg-build/dys2pwwos5w5kez2chufdk3b3oyj5n4n4iiseyke2tzuqahgvmovloyd.onion"},
		{"./proxysto.re", "proxystore", false, "/tmp/ssg-build/proxysto.re"},
		{"./proxysto.re", "proxystore", true, "/tmp/ssg-build/proxyoxiemywllckvpix543gqcmvvltrnb7inbwtk2knkehqt72tyfyd.onion"},
	} {
		ws, err := ssg.MakeWebsite(os.DirFS(site.SrcDir), nil, langs, func(r *http.Request, data ssg.TemplateData) any {
			active := site.Active
			if active == "proxystore" && strings.Contains(r.URL.Path, "warum.html") {
				active = "why"
			}
			return TemplateData{data, active, site.Onion}
		})
		if err != nil {
			panic(err)
		}
		err = ws.WriteFiles(site.OutDir)
		if err != nil {
			panic(err)
		}
	}

	if *listen {
		ssg.ListenAndServe("/tmp/ssg-build/proxysto.re")
	}
}
