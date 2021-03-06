CREATE TABLE IF NOT EXISTS device_keys (
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
    device_id INT UNSIGNED NOT NULL,
    dkey VARCHAR(255) NOT NULL,

    FOREIGN KEY (device_id) REFERENCES devices (id),

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;