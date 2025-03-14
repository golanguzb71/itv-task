package service

import (
	"gorm.io/gorm"
	"itv/internal/dto"
	"itv/internal/model"
	"itv/internal/repository"
)

type MovieService struct {
	movieRepo *repository.MovieRepository
	db        *gorm.DB
}

func NewMovieService(movieRepo *repository.MovieRepository, db *gorm.DB) *MovieService {
	return &MovieService{
		movieRepo: movieRepo,
		db:        db,
	}
}

func (s *MovieService) GetAllMovies() ([]dto.MovieResponseDTO, error) {
	movies, err := s.movieRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var moviesDTO []dto.MovieResponseDTO
	for _, movie := range movies {
		moviesDTO = append(moviesDTO, mapMovieToDTO(&movie))
	}

	return moviesDTO, nil
}

func (s *MovieService) GetMovieByID(id uint) (*dto.MovieResponseDTO, error) {
	movie, err := s.movieRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	movieDTO := mapMovieToDTO(movie)
	return &movieDTO, nil
}

func (s *MovieService) CreateMovie(movieDTO dto.CreateMovieDTO) (*dto.MovieResponseDTO, error) {
	movie := model.Movie{
		Title:       movieDTO.Title,
		Director:    movieDTO.Director,
		Year:        movieDTO.Year,
		Plot:        movieDTO.Plot,
		Rating:      movieDTO.Rating,
		Duration:    movieDTO.Duration,
		ReleaseDate: movieDTO.ReleaseDate,
	}

	err := s.movieRepo.Create(&movie)
	if err != nil {
		return nil, err
	}

	responseDTO := mapMovieToDTO(&movie)
	return &responseDTO, nil
}

func (s *MovieService) UpdateMovie(id uint, movieDTO dto.UpdateMovieDTO) (*dto.MovieResponseDTO, error) {
	movie, err := s.movieRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if movieDTO.Title != "" {
		movie.Title = movieDTO.Title
	}
	if movieDTO.Director != "" {
		movie.Director = movieDTO.Director
	}
	if movieDTO.Year != 0 {
		movie.Year = movieDTO.Year
	}
	if movieDTO.Plot != "" {
		movie.Plot = movieDTO.Plot
	}
	if movieDTO.Rating != 0 {
		movie.Rating = movieDTO.Rating
	}
	if movieDTO.Duration != 0 {
		movie.Duration = movieDTO.Duration
	}
	if movieDTO.ReleaseDate != "" {
		movie.ReleaseDate = movieDTO.ReleaseDate
	}

	err = s.movieRepo.Update(movie)
	if err != nil {
		return nil, err
	}

	responseDTO := mapMovieToDTO(movie)
	return &responseDTO, nil
}

func (s *MovieService) DeleteMovie(id uint) error {
	_, err := s.movieRepo.FindByID(id)
	if err != nil {
		return err
	}

	return s.movieRepo.Delete(id)
}

func (s *MovieService) SearchMovies(query string) ([]dto.MovieResponseDTO, error) {
	movies, err := s.movieRepo.SearchMovies(query)
	if err != nil {
		return nil, err
	}

	var moviesDTO []dto.MovieResponseDTO
	for _, movie := range movies {
		moviesDTO = append(moviesDTO, mapMovieToDTO(&movie))
	}

	return moviesDTO, nil
}

func mapMovieToDTO(movie *model.Movie) dto.MovieResponseDTO {
	return dto.MovieResponseDTO{
		ID:          movie.ID,
		Title:       movie.Title,
		Director:    movie.Director,
		Year:        movie.Year,
		Plot:        movie.Plot,
		Rating:      movie.Rating,
		Duration:    movie.Duration,
		ReleaseDate: movie.ReleaseDate,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
	}
}
