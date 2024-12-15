package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AuroralTech/todo-api_202412/config"
	"github.com/rs/cors"
)

const requestTimeout = 5 * time.Minute

func main() {

	// ポート番号を設定
	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		log.Fatalf("環境変数が不足しています. API_PORT: %s", apiPort)
	}

	// データベースの設定
	db, err := config.NewDatabase()
	if err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}

	// bun疎通確認（データベースのバージョンを取得）
	var dbVersion string
	if err = db.NewSelect().ColumnExpr("version()").Scan(context.Background(), &dbVersion); err != nil {
		log.Fatalf("bun接続エラー: %v", err)
	}

	// CORSの設定
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                            // すべてのオリジンを許可
		AllowCredentials: true,                                     // クレデンシャル情報（Cookieなど）を許可
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"}, // 許可するHTTPメソッド
	})
	httpHandler := c.Handler(http.DefaultServeMux)

	// サーバーの設定
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", apiPort),
		Handler:      httpHandler,
		ReadTimeout:  requestTimeout,
		WriteTimeout: requestTimeout,
	}

	// サーバーの起動
	log.Printf("サーバーを起動します: :%s", apiPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}
