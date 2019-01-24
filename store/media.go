package store

import "fmt"

// MediaStatus enum
type MediaStatus int

const (
	unknown  MediaStatus = iota
	created  MediaStatus = iota
	imported MediaStatus = iota
	approved MediaStatus = iota
)

// Media aggregate
type Media struct {
	ID        string
	Name      string
	OrigPath  string
	ThumbPath string
	BigPath   string
	Type      string
	Status    MediaStatus
}

// String Stringer func
func (media *Media) String() string {
	format := `
	Media: %s
	Name: %s
	origPath: %s
	thumbPath: %s
	bigPath: %s
	type: %s
	status: %s
	`
	return fmt.Sprintf(format, media.ID, media.Name, media.OrigPath,
		media.ThumbPath, media.BigPath, media.Type, media.Status)
}
