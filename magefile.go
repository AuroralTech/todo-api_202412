//go:build mage
// +build mage

package main

import (
	"fmt"
	"os/exec"
)

var Aliases = map[string]interface{}{
	"mup":      MigrationUp,
	"mdown":    MigrationDown,
	"createmg": CreateMigrationFile,
}

// MigrationUp (arg string) 実行するとマイグレーションがされる 例: "all", "1", "2", "3"
func MigrationUp(arg string) error {
	fmt.Println("マイグレーションを開始します...")

	var cmd []byte
	var err error

	if arg == "all" {
		// 全てのマイグレーションを実行
		cmd, err = exec.Command("docker", "compose", "run", "--rm", "migrate",
			"-database", "postgres://root:password@db/local_db?sslmode=disable",
			"-path", "postgres", "up").CombinedOutput()
	} else {
		// 引数がある場合は指定された数のマイグレーションのみ実行
		cmd, err = exec.Command("docker", "compose", "run", "--rm", "migrate",
			"-database", "postgres://root:password@db/local_db?sslmode=disable",
			"-path", "postgres", "up", arg).CombinedOutput()
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

	var cmd []byte
	var err error

	if arg == "all" {
		// 全てのマイグレーションを実行
		cmd, err = exec.Command("docker", "compose", "run", "--rm", "migrate",
			"-database", "postgres://root:password@db/local_db?sslmode=disable",
			"-path", "postgres", "down").CombinedOutput()
	} else {
		// 引数がある場合は指定された数のマイグレーションのみ実行
		cmd, err = exec.Command("docker", "compose", "run", "--rm", "migrate",
			"-database", "postgres://root:password@db/local_db?sslmode=disable",
			"-path", "postgres", "down", arg).CombinedOutput()
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
	cmd, err := exec.Command("docker", "compose", "run", "--rm", "migrate", "create", "-ext", "sql", "-dir", "postgres", name).Output()
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
