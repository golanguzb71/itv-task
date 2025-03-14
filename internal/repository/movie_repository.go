package repository

import (
	"errors"
	"gorm.io/gorm"
	"itv/internal/model"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) FindAll() ([]model.Movie, error) {
	var movies []model.Movie
	result := r.db.Find(&movies)
	return movies, result.Error
}

func (r *MovieRepository) FindByID(id uint) (*model.Movie, error) {
	var movie model.Movie
	result := r.db.First(&movie, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("movie not found")
		}
		return nil, result.Error
	}
	return &movie, nil
}

func (r *MovieRepository) Create(movie *model.Movie) error {
	return r.db.Create(movie).Error
}

func (r *MovieRepository) Update(movie *model.Movie) error {
	return r.db.Save(movie).Error
}

func (r *MovieRepository) Delete(id uint) error {
	return r.db.Delete(&model.Movie{}, id).Error
}

func (r *MovieRepository) SearchMovies(query string) ([]model.Movie, error) {
	var movies []model.Movie
	result := r.db.Where("title LIKE ? OR director LIKE ? OR plot LIKE ?",
		"%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&movies)
	return movies, result.Error
}
