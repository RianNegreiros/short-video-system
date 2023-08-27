package models

type Video struct {
	VideoID     uint
	UserID      uint
	Title       string
	Description string
	Category    string
	Tags        []string
	FileName    string
	FilePath    string
}
