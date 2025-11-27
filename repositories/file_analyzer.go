package repositories

import (
	"database/sql"

	"github.com/ali-hassan-Codes/file_analyzer_2/models"
)

type FileAnalyzerRepoInterface interface {
	InsertFileInfo(file models.FileInfo) error
}

type FileAnalyzerRepository struct {
	DB *sql.DB
}

func NewFileAnalyzerRepository(db *sql.DB) *FileAnalyzerRepository {
	return &FileAnalyzerRepository{DB: db}
}

func (repo *FileAnalyzerRepository) InsertFileInfo(file models.FileInfo) error {
	query := `
        INSERT INTO file_stats
        (file_name, paragraphs, line_count, word_count, char_count, alphabetic, numeric_count, vowels, non_vowels)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
	_, err := repo.DB.Exec(query,
		file.FileName,
		file.Paragraphs,
		file.LineCount,
		file.WordCount,
		file.CharCount,
		file.Alphabetic,
		file.NumericCount, // updated
		file.Vowels,
		file.NonVowels,
	)
	return err
}
