CREATE TABLE IF NOT EXISTS todos(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    body TEXT,
    detail JSON,
    status VARCHAR(50) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME
)