package controller

import (
	"itv/internal/dto"
	"itv/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	movieService *service.MovieService
}

func NewMovieController(movieService *service.MovieService) *MovieController {
	return &MovieController{
		movieService: movieService,
	}
}

// GetAllMovies godoc
//	@Summary		Get all movies
//	@Description	Get a list of all movies
//	@Tags			movies
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{array}		dto.MovieResponseDTO
//	@Failure		500	{object}	dto.ErrorResponseDTO
//	@Router			/movies [get]
func (c *MovieController) GetAllMovies(ctx *gin.Context) {
	movies, err := c.movieService.GetAllMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponseDTO{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, movies)
}

// GetMovieByID godoc
//	@Summary		Get a movie by ID
//	@Description	Get a movie by its ID
//	@Tags			movies
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		int	true	"Movie ID"
//	@Success		200	{object}	dto.MovieResponseDTO
//	@Failure		404	{object}	dto.ErrorResponseDTO
//	@Failure		500	{object}	dto.ErrorResponseDTO
//	@Router			/movies/{id} [get]
func (c *MovieController) GetMovieByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponseDTO{Error: "invalid ID format"})
		return
	}

	movie, err := c.movieService.GetMovieByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponseDTO{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

// CreateMovie godoc
//	@Summary		Create a new movie
//	@Description	Create a new movie with the provided data
//	@Tags			movies
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			movie	body		dto.CreateMovieDTO	true	"Movie data"
//	@Success		201		{object}	dto.MovieResponseDTO
//	@Failure		400		{object}	dto.ErrorResponseDTO
//	@Failure		500		{object}	dto.ErrorResponseDTO
//	@Router			/movies [post]
func (c *MovieController) CreateMovie(ctx *gin.Context) {
	var movieDTO dto.CreateMovieDTO
	if err := ctx.ShouldBindJSON(&movieDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponseDTO{Error: err.Error()})
		return
	}

	movie, err := c.movieService.CreateMovie(movieDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponseDTO{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, movie)
}

// UpdateMovie godoc
//	@Summary		Update a movie
//	@Description	Update a movie with the provided data
//	@Tags			movies
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		int					true	"Movie ID"
//	@Param			movie	body		dto.UpdateMovieDTO	true	"Movie data"
//	@Success		200		{object}	dto.MovieResponseDTO
//	@Failure		400		{object}	dto.ErrorResponseDTO
//	@Failure		404		{object}	dto.ErrorResponseDTO
//	@Failure		500		{object}	dto.ErrorResponseDTO
//	@Router			/movies/{id} [put]
func (c *MovieController) UpdateMovie(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponseDTO{Error: "invalid ID format"})
		return
	}

	var movieDTO dto.UpdateMovieDTO
	if err := ctx.ShouldBindJSON(&movieDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponseDTO{Error: err.Error()})
		return
	}

	movie, err := c.movieService.UpdateMovie(uint(id), movieDTO)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponseDTO{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

// DeleteMovie godoc
//	@Summary		Delete a movie
//	@Description	Delete a movie by its ID
//	@Tags			movies
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path	int	true	"Movie ID"
//	@Success		204	"No Content"
//	@Failure		404	{object}	dto.ErrorResponseDTO
//	@Failure		500	{object}	dto.ErrorResponseDTO
//	@Router			/movies/{id} [delete]
func (c *MovieController) DeleteMovie(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponseDTO{Error: "invalid ID format"})
		return
	}

	err = c.movieService.DeleteMovie(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponseDTO{Error: err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// SearchMovies godoc
//	@Summary		Search movies
//	@Description	Search movies based on a query
//	@Tags			movies
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			query	query		string	true	"Search query"
//	@Success		200		{array}		dto.MovieResponseDTO
//	@Failure		500		{object}	dto.ErrorResponseDTO
//	@Router			/movies/search [get]
func (c *MovieController) SearchMovies(ctx *gin.Context) {
	query := ctx.Query("query")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponseDTO{Error: "search query is required"})
		return
	}

	movies, err := c.movieService.SearchMovies(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponseDTO{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, movies)
}
