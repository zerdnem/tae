package utils

import (
	"github.com/gocolly/colly"
)

var c = colly.NewCollector()

type info struct {
	hashtype, url string
}

type api struct {
	hash []info
}

func (ha *api) addSource() []info {
	ha.hash = []info{
		{hashtype: "md5", url: "https://hashtoolkit.com/decrypt-md5-hash/?hash="},
		{hashtype: "sha1", url: "https://hashtoolkit.com/decrypt-sha1-hash/?hash="},
		{hashtype: "sha256", url: "https://hashtoolkit.com/decrypt-sha256-hash/?hash="},
		{hashtype: "sha384", url: "https://hashtoolkit.com/decrypt-sha384-hash/?hash="},
		{hashtype: "sha512", url: "https://hashtoolkit.com/decrypt-sha512-hash/?hash="},
	}
	return ha.hash
}

func DecryptHash(newhash, hashtype string) string {
	h := api{}
	sources := h.addSource()
	for _, source := range sources {
		if source.hashtype == hashtype {
			var response string
			c.OnHTML("span", func(e *colly.HTMLElement) {
				if e.Attr("title") == "decrypted "+hashtype+" hash" {
					response = string(e.Text)
				}
			})
			c.Visit(source.url + newhash)
			return response

		}
	}
	return ""
}
