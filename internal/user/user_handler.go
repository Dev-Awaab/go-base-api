package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService
}

func NewUserHandler(us UserService) *Handler{
	return &Handler{
		UserService: us,
	}
}

func (h *Handler) CreateUser(c *gin.Context){
	var u CreateUserReq

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err:=h.UserService.Create(c.Request.Context(), &u)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, res)
}