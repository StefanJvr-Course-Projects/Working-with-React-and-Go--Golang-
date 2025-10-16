package repository

import "backend/internal/models"

type DataabaseRepo interface {
	AllMovies() ([]*models.Movie, error)
}
