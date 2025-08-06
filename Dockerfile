FROM golang:1.22-alpine

# gitと必要な依存関係をインストール
RUN apk add --no-cache git

WORKDIR /app
COPY . .

# 依存関係の解決とビルド
RUN go mod tidy
RUN go build -o main .

CMD ["./main"]
