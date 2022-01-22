# 使うDokerイメージ
FROM golang:latest
# ポートを開ける
EXPOSE 10001/TCP
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/
# ファイルコピー
ADD . /go/src/
# サーバーを起動する
CMD go run server.go