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

type Hashes struct {
	Md5, Sha1, Sha256, Sha384, Sha512 string
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

func scrape(element, attr, value, url string) interface{} {
	var result []string
	c.OnHTML(element, func(e *colly.HTMLElement) {
		if e.Attr(attr) == value {
			if element == "span" {
				result = append(result, e.Text)
			} else {
				result = append(result, e.ChildText("span"))
			}
		}
	})
	c.Visit(url)
	if len(result) >= 2 {
		return Hashes{
			Md5:    result[0],
			Sha1:   result[1],
			Sha256: result[2],
			Sha384: result[3],
			Sha512: result[4],
		}
	}
	return result[0]
}

func RetrieveHash(newhash, hashtype string) string {
	h := api{}
	sources := h.addSource()
	for _, source := range sources {
		if source.hashtype == hashtype {
			data := scrape("span", "title", "decrypted "+hashtype+" hash", source.url+newhash)
			if decrypted, ok := data.(string); ok {
				return decrypted
			}

		}
	}
	return ""
}

func GenerateHash(text string) Hashes {
	url := "https://hashtoolkit.com/generate-hash/?text=" + text
	data := scrape("td", "class", "res-hash", url)
	if generated, ok := data.(Hashes); ok {
		var ha Hashes
		ha = generated
		return ha
	}
	return Hashes{}
}
