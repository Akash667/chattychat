package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {

	return &Handler{
		Service: s,
	}

}

func (h *Handler) CreateUser(c *gin.Context) {
	var u UserCreateReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": c.Params.ByName("username")})
		log.Println("Error occured during binding create request")
		return
	}
	log.Println(u.Email)

	res, err := h.Service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Error occured during executing create request")
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) LoginUser(c *gin.Context) {

	var u LoginUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		log.Println("Error occured during binding login info", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": c.Params.ByName("email")})
		return
	}

	user, err := h.Service.LoginUser(c.Request.Context(), &u)

	if err != nil {
		log.Println("Error occured during loggin in user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("token", user.accessToken, 3600, "/", "localhost", false, true)

	res := &LoginUserRes{
		Username: user.Username,
		ID:       user.ID,
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
