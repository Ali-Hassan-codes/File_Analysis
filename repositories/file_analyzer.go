package repositories

import (
    "database/sql"

    "github.com/ali-hassan-Codes/file_analyzer_2/models"
)

type FileAnalyzerRepository struct {
    DB *sql.DB
}

func NewFileAnalyzerRepository(db *sql.DB) *FileAnalyzerRepository {
    return &FileAnalyzerRepository{DB: db}
}



// Insert file analysis record
func (repo *FileAnalyzerRepository) InsertFileInfo(file models.FileInfo) error {
    query := `
        INSERT INTO file_info
        (file_name, paragraphs, line_count, word_count, char_count, alphabetic, digits, vowels, non_vowels)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
    _, err := repo.DB.Exec(query,
        file.FileName,
        file.Paragraphs,
        file.LineCount,
        file.WordCount,
        file.CharCount,
        file.Alphabetic,
        file.Digits,      // updated field name
        file.Vowels,
        file.NonVowels,
    )
    return err
}
