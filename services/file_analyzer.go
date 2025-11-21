package services

import (
	"bufio"
	"os"
	"strings"
	"unicode"

	"github.com/ali-hassan-Codes/file_analyzer_2/models"
	"github.com/ali-hassan-Codes/file_analyzer_2/repositories"
)

type IFileAnalyzerService interface {
	AnalyzeFile(filePath string) (models.FileInfo, error)
}

// Service
type FileAnalyzerService struct {
	repo *repositories.FileAnalyzerRepository
}

func NewFileAnalyzerService(repo *repositories.FileAnalyzerRepository) *FileAnalyzerService {
	return &FileAnalyzerService{repo: repo}
}

// AnalyzeFile reads the file, calculates stats, and saves to DB
func (s *FileAnalyzerService) AnalyzeFile(filePath string) (models.FileInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return models.FileInfo{}, err
	}
	defer file.Close()

	paraCount, lineCount, wordCount := 0, 0, 0
	charCount, alphaCount, digitCount := 0, 0, 0
	vowelCount, nonVowelCount := 0, 0
	isParaCounted := false
	vowels := "aeiouAEIOU"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			if !isParaCounted {
				paraCount++
				isParaCounted = true
			}

			words := strings.Fields(trimmed)
			wordCount += len(words)

			for _, char := range trimmed {
				charCount++
				if unicode.IsLetter(char) {
					alphaCount++
					if strings.ContainsRune(vowels, char) {
						vowelCount++
					} else {
						nonVowelCount++
					}
				} else if unicode.IsDigit(char) {
					digitCount++
				}
			}
		} else {
			isParaCounted = false
		}
	}

	if err := scanner.Err(); err != nil {
		return models.FileInfo{}, err
	}

	fileInfo := models.FileInfo{
		FileName:   filePath,
		Paragraphs: paraCount,
		LineCount:  lineCount,
		WordCount:  wordCount,
		CharCount:  charCount,
		Alphabetic: alphaCount,
		Vowels:     vowelCount,
		NonVowels:  nonVowelCount,
	}

	// Save to DB
	err = s.repo.InsertFileInfo(fileInfo)
	if err != nil {
		return models.FileInfo{}, err
	}

	return fileInfo, nil
}
