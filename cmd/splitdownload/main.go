// Package main はimgconvコマンドのエントリーポイント
package main

import (
	"fmt"
	"os"
	"time"

	"net/http"
	"net/url"
)

// エンドポイント
func main() {
	os.Exit(Run(os.Args))
}

// Run はmainパッケージのexecute関数
func Run(args []string) int {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	// クエリを組み立て
	values := url.Values{}
	values.Add("key", "value")

	// Request を生成
	req, err := http.NewRequest("GET", "[アクセス先URL]", nil)
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

	return 0
}
