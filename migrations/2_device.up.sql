CREATE TABLE IF NOT EXISTS devices (
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
    owner_id INT UNSIGNED NOT NULL,
    uuid VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255),
    address VARCHAR(255) NOT NULL,

    FOREIGN KEY (owner_id) REFERENCES users (id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;