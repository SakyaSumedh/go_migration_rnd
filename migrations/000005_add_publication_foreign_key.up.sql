ALTER TABLE books
    ADD COLUMN publication_id INT UNSIGNED DEFAULT NULL,
    ADD FOREIGN KEY (publication_id) REFERENCES publications(id) ON DELETE CASCADE; 