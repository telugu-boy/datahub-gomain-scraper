package main

import "time"

const (
	ThreadCnt  = 1024
	DNSServer  = "1.1.1.1:53"
	ReqTimeout = 5 * time.Second
	MaxRespLen = 64_000 // 64KB
	MaxRedir   = 5
)

var RequestHeaders = map[string]string{"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0"}

var HtmlMetaBlackList = []string{
	"video:duration", "og:image:type", "og:image:height",
	"og:image:width", "twitter:card"}
