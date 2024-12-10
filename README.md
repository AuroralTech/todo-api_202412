# todo-api_202412

## ディレクトリ構成

```
.
├── cmd # エントリーポイント
├── config # 設定ファイル
├── db # データベース関連
│   └── migrations # マイグレーション用のSQLファイル
└── pkg # パッケージ
    ├── entity # エンティティ（ドメインモデル）
    ├── handler # ハンドラー
    ├── repository # レポジトリ
    │   └── interface # レポジトリのインターフェース
    └── usecase # ユースケース
```
