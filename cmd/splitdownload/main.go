// Package main はimgconvコマンドのエントリーポイント
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"net/http"
	"net/url"

	// 自作パッケージ
	"github.com/yoichi-watanabe/splitdownload"
)

// エンドポイント
func main() {
	os.Exit(Run(os.Args))
}

// Run はmainパッケージのexecute関数
func Run(args []string) int {

	// コマンドライン引数からダウンロードファイルのURLを取得
	flag.Parse()
	argURL := flag.Arg(0)

	downloader := splitdownload.NewDownloader()
	downloader.URL = argURL


	//todo 以下処理をdownloaderに委譲
	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	// クエリを組み立て
	values := url.Values{}
	values.Add("key", "value")

	// Request を生成
	req, err := http.NewRequest("GET", argURL, nil)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	// 組み立てたクエリを生クエリ文字列に変換して設定
	req.URL.RawQuery = values.Encode()

	// Doメソッドでリクエストを投げる
	// http.Response型のポインタ（とerror）が返ってくる
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	// 関数を抜ける際に必ずresponseをcloseするようにdeferでcloseを呼ぶ
	defer resp.Body.Close()

	// dubug用
	fmt.Println(req)

	return 0
}
