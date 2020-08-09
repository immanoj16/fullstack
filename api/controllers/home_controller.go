package controllers

import (
	"net/http"

	"github.com/immanoj16/fullstack/api/responses"
)

// Home route
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this awesome API")
}
