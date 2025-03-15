CREATE DATABASE IF NOT EXISTS mydb;
USE mydb;

CREATE TABLE IF NOT EXISTS feed (
  id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
  owner_id SMALLINT UNSIGNED NOT NULL,
  title VARCHAR(30) NOT NULL,
  content VARCHAR(200) DEFAULT NULL,
  img_urls TEXT CHARACTER SET ascii COLLATE ascii_bin DEFAULT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(id),
  KEY IDX_ownerId(owner_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS comment (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  feed_id SMALLINT UNSIGNED NOT NULL,
  owner_id SMALLINT UNSIGNED NOT NULL,
  content VARCHAR(200) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(id),
  KEY IDX_feedId(feed_id),
  KEY IDX_ownerId(owner_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET NAMES utf8mb4;

INSERT INTO feed (id, owner_id, title, content, img_urls)
VALUES (1, 1, '첫 번째 게시글입니다.', '첫 번째 게시글의 내용입니다.', 'https://example.com/image.jpg');

INSERT INTO comment (feed_id, owner_id, content)
VALUES (1, 1, '댓글 1입니다.');

INSERT INTO comment (feed_id, owner_id, content)
VALUES (1, 1, '댓글 2입니다!!!');