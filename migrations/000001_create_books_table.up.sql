CREATE TABLE IF NOT EXISTS books (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title varchar(255),
    author varchar(255),
    publication varchar(255) NOT NULL
) ENGINE=INNODB;

COMMIT;