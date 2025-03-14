package dto

import (
	"time"
)

type CreateMovieDTO struct {
	Title       string    `json:"title" binding:"required"`
	Director    string    `json:"director" binding:"required"`
	Year        int       `json:"year" binding:"required,min=1800,max=2100"`
	Plot        string    `json:"plot"`
	Rating      float32   `json:"rating" binding:"min=0,max=10"`
	Duration    int       `json:"duration" binding:"required,min=1"`
	ReleaseDate time.Time `json:"release_date" binding:"required"`
}

type UpdateMovieDTO struct {
	Title       string    `json:"title"`
	Director    string    `json:"director"`
	Year        int       `json:"year" binding:"omitempty,min=1800,max=2100"`
	Plot        string    `json:"plot"`
	Rating      float32   `json:"rating" binding:"omitempty,min=0,max=10"`
	Duration    int       `json:"duration" binding:"omitempty,min=1"`
	ReleaseDate time.Time `json:"release_date"`
}
type MovieResponseDTO struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Director    string    `json:"director"`
	Year        int       `json:"year"`
	Plot        string    `json:"plot"`
	Rating      float32   `json:"rating"`
	Duration    int       `json:"duration"`
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type LoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenResponseDTO struct {
	Token string `json:"token"`
}

type ErrorResponseDTO struct {
	Error string `json:"error"`
}
