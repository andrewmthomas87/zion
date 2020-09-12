package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register registers authentication API handlers.
func Register(r *gin.RouterGroup, api *API) {
	r.POST("/sign-up", api.SignUp)
	r.POST("/sign-in", api.SignIn)
}

// NewAPI creates an api.
func NewAPI(service *Service) *API {
	return &API{s: service}
}

// API implements handler functions for authentication.
type API struct {
	s *Service
}

type signUpRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// SignUp handles sign up requests.
func (a *API) SignUp(c *gin.Context) {
	var json signUpRequest
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := a.s.SignUp(c.Request.Context(), json.Username, json.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

type signInRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// SignIn handles sign in requests.
func (a *API) SignIn(c *gin.Context) {
	var json signInRequest
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := a.s.SignIn(c.Request.Context(), json.Username, json.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
