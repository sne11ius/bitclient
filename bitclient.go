package bitclient

import (
	"github.com/dghubble/sling"
	"strings"
)

type BitClient struct {
	sling    *sling.Sling
	basePath string
}

func NewBitClient(url string, username string, password string) *BitClient {
	var splitUrl = strings.SplitAfterN(url, "/", 4)
	var last = splitUrl[len(splitUrl)-1]
	var basePath = ""
	if 0 != len(last) {
		basePath = "/" + last
	}
	return &BitClient{
		sling:    sling.New().Base(url).Set("Accept", "application/json").SetBasicAuth(username, password),
		basePath: basePath,
	}
}
