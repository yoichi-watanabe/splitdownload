package splitdownload

import (
	"log"
	"net/http"
	"net/url"
)

// Downloader は構造体定義
type Downloader struct {
	URL        *url.URL
	HTTPClient *http.Client

	Logger *log.Logger
}

// NewDownloader はクライアント生成メソッド
func NewDownloader() *Downloader {

	// todo
	return &Downloader{}

}

