-- +goose Up
CREATE TABLE IF NOT EXISTS file_stats (
    id INT AUTO_INCREMENT PRIMARY KEY,
    file_name VARCHAR(255),
    paragraphs INT,
    line_count INT,
    word_count INT,
    char_count INT,
    alphabetic INT,
    numeric_count INT,   -- changed from "numeric" to "numeric_count"
    vowels INT,
    non_vowels INT
);

-- +goose Down
DROP TABLE IF EXISTS file_stats;
