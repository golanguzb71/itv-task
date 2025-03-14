package controller

import (
	"github.com/gin-gonic/gin"
	"itv/internal/dto"
	"itv/internal/service"
	"net/http"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Login godoc
//	@Summary		User login
//	@Description	Authenticate user and return JWT token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			login	body		dto.LoginDTO	true	"Login credentials"
//	@Success		200		{object}	dto.TokenResponseDTO
//	@Failure		400		{object}	dto.ErrorResponseDTO
//	@Failure		401		{object}	dto.ErrorResponseDTO
//	@Router			/auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	if err := ctx.ShouldBindJSON(&loginDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponseDTO{Error: err.Error()})
		return
	}

	token, err := c.authService.Login(loginDTO)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponseDTO{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, token)
}
