-- Pivot post_tags (many-to-many)
CREATE TABLE post_tags (
  post_id BINARY(16) NOT NULL,
  tag_id BINARY(16) NOT NULL,
  PRIMARY KEY (post_id, tag_id),
  CONSTRAINT fk_pt_post FOREIGN KEY (post_id) REFERENCES posts(id),
  CONSTRAINT fk_pt_tag FOREIGN KEY (tag_id) REFERENCES tags(id)
);