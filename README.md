# Go + MySQL + Docker CRUD環境構築

このプロジェクトはGo言語、MySQL、Dockerを使用した基本的なAPIサーバーの環境構築例です。

## 構成

- Go: APIサーバー（1.22）
- MySQL: データベース（8.0）
- Docker & Docker Compose: コンテナ管理

## 環境構築手順

### 前提条件

- Docker
- Docker Compose

### セットアップ手順

1. リポジトリをクローンする
```bash
git clone <repository-url>
cd go-mysql-crud
```

2. Docker Composeでビルド・起動する
```bash
docker-compose up --build
```

3. 動作確認
   - http://localhost:8080/health にアクセスして「ok」レスポンスが返ってくれば成功

