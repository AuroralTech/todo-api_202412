package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// 環境変数から設定を読み込み
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := "5432"
	apiPort := os.Getenv("API_PORT")

	// PostgreSQL接続文字列の構築
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// データベースに接続
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}
	defer db.Close()

	// 接続テスト
	err = db.Ping()
	if err != nil {
		log.Fatalf("データベースPingエラー: %v", err)
	}
	log.Println("データベースに正常に接続されました")

	// ヘルスチェックエンドポイント
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "データベース接続エラー")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Healthy")
	})

	// サーバー起動
	log.Printf("サーバーを起動します: :%s", apiPort)
	if err := http.ListenAndServe(":"+apiPort, nil); err != nil {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}
