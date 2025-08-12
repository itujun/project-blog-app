-- Komentar
CREATE TABLE comments (
  id BINARY(16) PRIMARY KEY,
  post_id BINARY(16) NOT NULL,
  author_id BINARY(16) NOT NULL,
  content TEXT NOT NULL,
  is_approved BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL,
  CONSTRAINT fk_comments_post FOREIGN KEY (post_id) REFERENCES posts(id),
  CONSTRAINT fk_comments_author FOREIGN KEY (author_id) REFERENCES users(id)
);