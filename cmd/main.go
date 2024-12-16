package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AuroralTech/todo-api_202412/config"
	userHandler "github.com/AuroralTech/todo-api_202412/pkg/handler/user"
	repository "github.com/AuroralTech/todo-api_202412/pkg/repository"
	userUsecase "github.com/AuroralTech/todo-api_202412/pkg/usecase/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	// echoサーバーの設定
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// CORSの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	}))

	// 依存性の注入
	userRepository := repository.NewUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepository)
	userHandler := userHandler.NewUserHandler(userUsecase)

	// ルーティングの設定
	e.PUT("/users", userHandler.UpdateUser)

	// サーバーの起動
	log.Printf("サーバーを起動します: :%s", apiPort)
	if err := e.Start(fmt.Sprintf(":%s", apiPort)); err != nil {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}
