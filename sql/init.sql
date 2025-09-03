-- booksテーブルの作成
CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- サンプルデータの挿入
INSERT IGNORE INTO books (title, author) VALUES
('Goプログラミング入門', '山田太郎'),
('データベース設計', '田中花子'),
('Web開発の基礎', '佐藤次郎');
