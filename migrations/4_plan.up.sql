CREATE TABLE IF NOT EXISTS plans (
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
    uuid VARCHAR(255) UNIQUE NOT NULL,
    owner_id INT UNSIGNED NOT NULL,
    name VARCHAR(255) NOT NULL,

    FOREIGN KEY (owner_id) REFERENCES users (id),

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;