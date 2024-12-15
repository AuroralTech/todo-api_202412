# Dockerfile
FROM golang:1.23

# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

# ワーキングディレクトリの設定
WORKDIR /go/src/app

# ModuleモードをON
ENV GO111MODULE=on

# ホストのファイルをコンテナの作業ディレクトリに移行
COPY . .

# makeとgitのインストール（aptを使用）
RUN apt-get update && apt-get install -y \
    make \
    git \
    && rm -rf /var/lib/apt/lists/*

# airとdlvのインストール
RUN go install github.com/air-verse/air@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest

# go.modを参照し、go.sumファイルの更新を行う
RUN go mod tidy

EXPOSE ${API_PORT}

# localではホットリロードを有効にしたいのでairで起動する
CMD ["air", "-c", ".air.toml"]
