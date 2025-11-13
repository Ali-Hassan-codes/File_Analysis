package models



type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}
type FileInfo struct {
    ID         int    `db:"id"`
    FileName   string `db:"file_name"`
    Paragraphs int    `db:"paragraphs"`
    LineCount  int    `db:"line_count"`
    WordCount  int    `db:"word_count"`
    CharCount  int    `db:"char_count"`
    Alphabetic int    `db:"alphabetic"`
    Digits     int    `db:"digits"`      // renamed from Numeric
    Vowels     int    `db:"vowels"`
    NonVowels  int    `db:"non_vowels"`
}
