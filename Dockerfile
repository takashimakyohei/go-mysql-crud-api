FROM golang:1.22

WORKDIR /app
COPY . .

# 依存関係の解決とairのインストール
RUN go mod tidy
RUN go install github.com/cosmtrek/air@v1.44.0

# 起動時に依存関係を解決してからairを実行
CMD sh -c "go get github.com/go-sql-driver/mysql && go mod tidy && air"
