package gopuzzlegame

import "github.com/google/uuid"

type Image struct {
	ID      uuid.UUID `json:"id"`
	FileURL string    `json:"file_url"`
}
