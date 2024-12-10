package config

import (
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

// loadDataSourceNameはデータベースの設定を読み込み、データベースの接続文字列（DSN）を返す
func loadDataSourceName() (string, error) {
	// 環境変数から設定を読み込み
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := "5432"
	dbName := os.Getenv("POSTGRES_DB")

	// 環境変数が不足している場合はエラーを返す
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		return "", fmt.Errorf("environment variable is missing. POSTGRES_USER: %s, POSTGRES_PASSWORD: %s, POSTGRES_HOST: %s, POSTGRES_PORT: %s, POSTGRES_DB: %s", dbUser, dbPassword, dbHost, dbPort, dbName)
	}

	// PostgreSQL接続文字列の構築（postgres://user:pass@localhost:5432/database?sslmode=disable）
	dbDsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	return dbDsn, nil
}

func NewDatabase() (*bun.DB, error) {
	// データベース接続文字列（DSN）を取得
	dbDsn, err := loadDataSourceName()
	if err != nil {
		return nil, fmt.Errorf("failed to loadDataSourceName: %w", err)
	}

	// DSNをDBドライバ（pgx）で使用するconfig形式に変換
	pgxConfig, err := pgx.ParseConfig(dbDsn)
	if err != nil {
		return nil, fmt.Errorf("failed to pgx.ParseConfig: %w", err)
	}
	pgxConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	// sql.DBを作成
	sqlDB := stdlib.OpenDB(*pgxConfig)
	// bun.DBを作成
	bunDB := bun.NewDB(sqlDB, pgdialect.New())

	// デバック用のログ出力設定を追加（https://bun.uptrace.dev/guide/debugging.html#bundebug）
	bunDB.AddQueryHook(bundebug.NewQueryHook(
		// disable the hook
		bundebug.WithEnabled(false),

		// BUNDEBUG=1 logs failed queries
		// BUNDEBUG=2 logs all queries
		bundebug.FromEnv("BUNDEBUG"),
	))
	return bunDB, nil
}
