# todo-api_202412

## ディレクトリ構成

```
.
├── cmd # エントリーポイント
├── db # データベース関連
│   ├── config # データベースの設定
│   └── migrations # マイグレーション用のSQLファイル
└── pkg # パッケージ
    ├── domain # ドメイン
    │   ├── entity # エンティティ（ドメインモデル）
    │   └── interface # レポジトリのインターフェース
    ├── handler # ハンドラー
    ├── repository # レポジトリ
    └── usecase # ユースケース
```
