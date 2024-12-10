// データベースの設定

package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func loadDatabaseURL() (string, error) {

	// 環境変数から設定を読み込み
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := "5432"

	// 環境変数が不足している場合はエラーを返す
	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
		return "", fmt.Errorf("環境変数が不足しています. POSTGRES_USER: %s, POSTGRES_PASSWORD: %s, POSTGRES_DB: %s, POSTGRES_HOST: %s, POSTGRES_PORT: %s", dbUser, dbPassword, dbName, dbHost, dbPort)
	}

	// PostgreSQL接続文字列の構築
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	return psqlInfo, nil
}

func NewDatabase() (*sql.DB, error) {
	// データベース設定を取得
	dbURL, err := loadDatabaseURL()
	if err != nil {
		return nil, fmt.Errorf("failed to load database URL: %w", err)
	}

	// データベースに接続
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	return db, nil
}
