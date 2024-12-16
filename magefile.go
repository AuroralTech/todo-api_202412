//go:build mage
// +build mage

package main

import (
	"fmt"
	"os/exec"

	"github.com/AuroralTech/todo-api_202412/config"
	"github.com/joho/godotenv"
)

var Aliases = map[string]interface{}{
	"mup":      MigrationUp,
	"mdown":    MigrationDown,
	"createmg": CreateMigrationFile,
}

// MigrationUp (arg string) 実行するとマイグレーションがされる 例: "all", "1", "2", "3"
func MigrationUp(arg string) error {
	fmt.Println("マイグレーションを開始します...")

	// .envファイルを読み込む
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(".envファイルを読み込めませんでした")
		return err
	}
	// データベース接続文字列（DSN）を取得
	dbDsn, err := config.LoadDataSourceName()
	if err != nil {
		fmt.Println("データベース接続文字列を取得できませんでした")
		return err
	}

	var cmd []byte

	if arg == "all" {
		// 全てのマイグレーションを実行
		cmd, err = exec.Command(
			"docker", "compose", "run", "--rm", "migrate",
			"-database", dbDsn,
			"-path", "db/migrations", "up",
		).CombinedOutput()
	} else {
		// 引数がある場合は指定された数のマイグレーションのみ実行
		cmd, err = exec.Command(
			"docker", "compose", "run", "--rm", "migrate",
			"-database", dbDsn,
			"-path", "db/migrations", "up", arg,
		).CombinedOutput()
	}

	fmt.Println("Output Command:", string(cmd))
	if err != nil {
		fmt.Println("エラーによりマイグレーションが完了できませんでした")
		return err
	}

	fmt.Println("マイグレーションが完了しました!!")
	return nil
}

// MigrationDown (arg string) 実行するとマイグレーションがされる 例: "all", "1", "2", "3"
func MigrationDown(arg string) error {
	fmt.Println("マイグレーションを開始します...")

	// .envファイルを読み込む
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(".envファイルを読み込めませんでした")
		return err
	}
	// データベース接続文字列（DSN）を取得
	dbDsn, err := config.LoadDataSourceName()
	if err != nil {
		fmt.Println("データベース接続文字列を取得できませんでした")
		return err
	}

	var cmd []byte

	if arg == "all" {
		// 全てのマイグレーションを実行
		cmd, err = exec.Command(
			"docker", "compose", "run", "--rm", "migrate",
			"-database", dbDsn,
			"-path", "db/migrations", "down", "-all",
		).CombinedOutput()
	} else {
		// 引数がある場合は指定された数のマイグレーションのみ実行
		cmd, err = exec.Command(
			"docker", "compose", "run", "--rm", "migrate",
			"-database", dbDsn,
			"-path", "db/migrations", "down", arg,
		).CombinedOutput()
	}

	fmt.Println("Output Command:", string(cmd))
	if err != nil {
		fmt.Println("エラーによりマイグレーションが完了できませんでした")
		return err
	}

	fmt.Println("マイグレーションが完了しました!!")
	return nil
}

// CreateMigrationFile (name string) 引数に値を渡し、マイグレーションファイルを作成する 例: create_users、add_uuid_to_users
func CreateMigrationFile(name string) error {
	fmt.Println("マイグレーションファイルを作成します...")
	cmd, err := exec.Command("docker", "compose", "run", "--rm", "migrate", "create", "-ext", "sql", "-dir", "db/migrations", name).Output()
	fmt.Println("Output Command:", string(cmd))
	if err != nil {
		fmt.Println("エラーによりマイグレーションファイルを作成できませんでした")
		return err
	}
	fmt.Println("マイグレーションファイルを作成しました!!")
	return nil
}

func run(cmd *exec.Cmd) error {
	b, err := cmd.CombinedOutput()
	if len(b) > 0 {
		fmt.Println(string(b))
	}
	return err
}
