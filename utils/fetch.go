package utils

import (
	"errors"
	"github.com/gocolly/colly"
)

type hashInfo struct {
	hashtype string
	url      string
}

type hash struct {
	hash []hashInfo
}

func (ha *hash) AddHash() []hashInfo {
	ha.hash = append(ha.hash, hashInfo{hashtype: "md5", url: "https://hashtoolkit.com/decrypt-md5-hash/?hash="})
	ha.hash = append(ha.hash, hashInfo{hashtype: "sha1", url: "https://hashtoolkit.com/decrypt-sha1-hash/?hash="})
	ha.hash = append(ha.hash, hashInfo{hashtype: "sha256", url: "https://hashtoolkit.com/decrypt-sha256-hash/?hash="})
	return ha.hash
}

func RetrieveHash(newhash, hashType string) (string, error) {
	var response string
	c := colly.NewCollector()
	h := hash{}
	hashes := h.AddHash()
	for _, hash := range hashes {
		if hash.hashtype == hashType {
			switch hash.hashtype {
			case "md5":
				c.OnHTML("span", func(e *colly.HTMLElement) {
					if e.Attr("title") == "decrypted "+hashType+" hash" {
						response = string(e.Text)
					}
				})
				c.Visit(hash.url + newhash)
				return response, nil

			case "sha1":
				c.OnHTML("span", func(e *colly.HTMLElement) {
					if e.Attr("title") == "decrypted "+hashType+" hash" {
						response = string(e.Text)
					}
				})
				c.Visit(hash.url + newhash)
				return response, nil

			case "sha256":
				c.OnHTML("span", func(e *colly.HTMLElement) {
					if e.Attr("title") == "decrypted "+hashType+" hash" {
						response = string(e.Text)
					}
				})
				c.Visit(hash.url + newhash)
				return response, nil
			}
		}
	}
	return response, errors.New("Hash could not be deciphered")
}
